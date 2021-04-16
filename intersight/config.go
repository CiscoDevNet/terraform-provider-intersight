package intersight

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"crypto/tls"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httptrace"
	"os"
	"strings"
	"time"

	gosdk "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

type Config struct {
	ApiKey    string
	SecretKey string
	Endpoint  string
	ApiClient *gosdk.APIClient
	ctx       context.Context
}

type AuthCodePKCEKeys struct {
	Verifier  string //Random URL-safe string with a minimum length of 43 characters
	Challenge string
}

type Client struct {
	hostname            string
	skipTlsVerification bool
}

func (c *Client) SetInputs(apiKeyId, secretKey, hostName string, ignoreTls bool) (context.Context, error) {
	c.hostname = hostName
	c.skipTlsVerification = ignoreTls

	var err error
	ctx := context.Background()
	if apiKeyId != "" {
		ctx, err = c.getHttpSignatureContext(ctx, apiKeyId, secretKey)
		if err != nil {
			return nil, fmt.Errorf("failed to get signature context: %v", err.Error())
		}
	} else {
		return nil, fmt.Errorf("missing API key and ID, required for authentication")
	}
	return ctx, nil
}

func newTransport(t http.RoundTripper) *transport {
	return &transport{t: t}
}

func (c *Client) getHttpSignatureContext(ctx context.Context, flagKeyId string, flagSecretKey string) (context.Context, error) {
	// Example keyId: 596cc79e5d91b400010d15ad/59b856cb16267c0001286496/5dbb09b27564612d30b68668
	// Create Authentication context.
	if flagKeyId == "" {
		return nil, errors.New("KeyId must be set")
	}
	if flagSecretKey == "" {
		return nil, errors.New("private key or its path must be set")
	}
	// The following headers are required by Intersight:
	//   (request-target)
	//   Date
	//   Digest
	//   Host
	// It is also recommended to add the following headers:
	//   (created)
	//
	//SigningScheme:    intersight.HttpSigningSchemeHs2019,
	//SigningAlgorithm: intersight.HttpSigningAlgorithmRsaPSS,

	authCfg := gosdk.HttpSignatureAuth{
		KeyId:            flagKeyId,
		Passphrase:       "",
		SigningScheme:    gosdk.HttpSigningSchemeRsaSha256,
		SigningAlgorithm: gosdk.HttpSigningAlgorithmRsaPKCS1v15,
		SignedHeaders: []string{
			gosdk.HttpSignatureParameterRequestTarget,
			gosdk.HttpSignatureParameterCreated,
			gosdk.HttpSignatureParameterExpires,
			gosdk.HttpHeaderHost,
			gosdk.HttpHeaderDate,
			gosdk.HttpHeaderDigest,
			"Content-Type",
		},
		SignatureMaxValidity: 10 * time.Minute,
	}
	//Set Private Key or Private Key Path
	if _, err := os.Stat(flagSecretKey); err != nil {
		// not a file, therefore is string
		err = authCfg.SetPrivateKey(flagSecretKey)
		if err != nil {
			return nil, err
		}
	} else {
		authCfg.PrivateKeyPath = flagSecretKey
	}

	log.Printf("Loading Intersight API private key. key-id: %v", flagKeyId)
	// Create a context that will load the Intersight API private key.
	// To create an Intersight private key:
	// - Go to Settings, API keys
	// - Generate a key
	// - Copy the key Id

	// Attach the HTTP signature attribute to the context.
	ctx, err := authCfg.ContextWithValue(ctx)
	_, err = authCfg.GetPublicKey()
	if err != nil {
		return nil, err
	}
	return ctx, err
}

func (c *Client) getHttpClientContext(ctx context.Context) context.Context {
	// Use a custom HTTP Client when requesting a token.
	t := newTransport(&http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: c.skipTlsVerification,
		},
		Proxy: http.ProxyFromEnvironment,
	})
	trace := &httptrace.ClientTrace{
		GotConn: t.GotConn,
		WroteHeaderField: func(key string, value []string) {
			log.Printf("Header %v: %v", key, value)
		},
	}
	ctx = httptrace.WithClientTrace(ctx, trace)
	httpClient := &http.Client{
		Transport: t,
	}
	return context.WithValue(ctx, oauth2.HTTPClient, httpClient)
}

// getClientCredentialsContext returns a context suitable for OAuth2 Client-credentials.
func (c *Client) getClientCredentialsContext(ctx context.Context, flagClientId string, flagClientSecret string) (context.Context, error) {
	ctx = c.getHttpClientContext(ctx)
	conf := &clientcredentials.Config{
		ClientID:     flagClientId,
		ClientSecret: flagClientSecret,
		TokenURL:     fmt.Sprintf("https://%s/iam/token", c.hostname),
		Scopes:       []string{},
	}
	return context.WithValue(ctx, gosdk.ContextOAuth2, conf.TokenSource(ctx)), nil
}

// getAuthorizationCodeContext returns a context suitable for OAuth2 authorization-code.
func (c *Client) getAuthorizationCodeContext(ctx context.Context, flagClientId string, flagClientSecret string) (context.Context, error) {
	ctx = c.getHttpClientContext(ctx)
	conf := &oauth2.Config{
		ClientID:     flagClientId,
		ClientSecret: flagClientSecret,
		Scopes:       []string{},
		Endpoint: oauth2.Endpoint{
			AuthURL:  fmt.Sprintf("https://%s/iam/token", c.hostname),
			TokenURL: fmt.Sprintf("https://%s/iam/token", c.hostname),
		},
	}

	// Redirect user to consent page to ask for permission
	// for the scopes specified above.
	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	log.Printf("Login to Intersight at %s, then visit the URL for the auth dialog: %v", c.hostname, url)

	// Use the authorization code that is pushed to the redirect
	// URL. Exchange will do the handshake to retrieve the
	// initial access token. The HTTP Client returned by
	// conf.Client will refresh the token as necessary.
	var code string
	if _, err := fmt.Scan(&code); err != nil {
		return nil, fmt.Errorf("Failed to get OAuth2 code: %v", err)
	}
	tok, err := conf.Exchange(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("Failed to perform exchange OAuth2: %v", err)
	}

	return context.WithValue(ctx, gosdk.ContextOAuth2, conf.TokenSource(ctx, tok)), nil
}

func (c *Client) getAuthorizationCodePKCEContext(ctx context.Context, flagClientId string, flagRedirectUrl string) (context.Context, error) {
	ctx = c.getHttpClientContext(ctx)
	conf := &oauth2.Config{
		ClientID:    flagClientId,
		Scopes:      []string{},
		RedirectURL: flagRedirectUrl,
		Endpoint: oauth2.Endpoint{
			AuthURL:  fmt.Sprintf("https://%s/iam/token", c.hostname),
			TokenURL: fmt.Sprintf("https://%s/iam/token", c.hostname),
		},
	}

	keys, err := generateKeys()
	if err != nil {
		return nil, fmt.Errorf("Failed to generate OAuth2 key: %v", err)
	}

	// Redirect user to consent page to ask for permission
	// for the scopes specified above.
	var csrfToken []byte
	csrfToken, err = generateRandomBytes(32)
	if err != nil {
		return nil, fmt.Errorf("Failed to generate OAuth2 key: %v", err)
	}
	url := conf.AuthCodeURL(base64.URLEncoding.EncodeToString(csrfToken),
		oauth2.SetAuthURLParam("domainname", "cisco.com"),
		oauth2.SetAuthURLParam("code_challenge_method", "S256"),
		oauth2.SetAuthURLParam("code_challenge", keys.Challenge),
	)
	log.Printf("Login to Intersight at %s, then visit the URL for the auth dialog: %v", c.hostname, url)

	// Use the authorization code that is pushed to the redirect
	// URL. Exchange will do the handshake to retrieve the
	// initial access token. The HTTP Client returned by
	// conf.Client will refresh the token as necessary.
	var code string
	if _, err = fmt.Scan(&code); err != nil {
		return nil, fmt.Errorf("Failed to get OAuth2 code: %v", err)
	}
	tok, err := conf.Exchange(ctx, code, oauth2.SetAuthURLParam("code_verifier", keys.Verifier))
	if err != nil {
		return nil, fmt.Errorf("Failed to perform exchange OAuth2: %v", err)
	}

	return context.WithValue(ctx, gosdk.ContextOAuth2, conf.TokenSource(ctx, tok)), nil
}

func (c *Client) GetApiClient(ctx context.Context, enableDebug bool) *gosdk.APIClient {
	cfg := gosdk.NewConfiguration()
	cfg.UserAgent = strings.Replace(cfg.UserAgent, "go", "terraform", -1)
	cfg.Host = c.hostname
	cfg.AddDefaultHeader("Content-Type", "application/json")
	cfg.HTTPClient = &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:       10,
			IdleConnTimeout:    30 * time.Second,
			DisableCompression: true,
			TLSClientConfig:    &tls.Config{InsecureSkipVerify: c.skipTlsVerification},
			Proxy:              http.ProxyFromEnvironment,
		},
	}
	cfg.Debug = enableDebug
	return gosdk.NewAPIClient(cfg)
}

// transport is an http.RoundTripper that keeps track of the in-flight
// request and implements hooks to report HTTP tracing events.
type transport struct {
	t       http.RoundTripper
	current *http.Request
}

// RoundTrip wraps http.DefaultTransport.RoundTrip to keep track
// of the current request.
func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	t.current = req
	log.Printf("RoundTrip URL: %v %v", t.current.Method, t.current.URL)
	/*
		b, err1 := ioutil.ReadAll(t.current.Body)
		if err1 == nil {
			log.Printf("RoundTrip Body: %v", string(b))
		}
	*/
	resp, err := t.t.RoundTrip(req)
	if err == nil {
		log.Printf("Response status: %s", resp.Status)
	} else {
		log.Printf("Response error: %v", err)
	}
	return resp, err
}

// GotConn prints whether the connection has been used previously
// for the current request.
func (t *transport) GotConn(info httptrace.GotConnInfo) {
	log.Printf("Connected to %v", t.current.URL)
}

func generateKeys() (*AuthCodePKCEKeys, error) {
	var keys AuthCodePKCEKeys

	bytes, err := generateRandomBytes(32)
	if err != nil {
		return nil, err
	}
	keys.Verifier = hex.EncodeToString(bytes)

	lCodeSum := sha256.Sum256([]byte(keys.Verifier))
	lCodeSumString := hex.EncodeToString(lCodeSum[:])
	keys.Challenge = base64.URLEncoding.EncodeToString([]byte(lCodeSumString))

	return &keys, nil
}

func generateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}
