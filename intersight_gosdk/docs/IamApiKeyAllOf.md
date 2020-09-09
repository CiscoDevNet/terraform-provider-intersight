# IamApiKeyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HashAlgorithm** | Pointer to **string** | The cryptographic hash algorithm to calculate the message digest. * &#x60;SHA256&#x60; - The SHA-256 cryptographic hash, as defined by NIST in FIPS 180-4. * &#x60;SHA384&#x60; - The SHA-384 cryptographic hash, as defined by NIST in FIPS 180-4. * &#x60;SHA512&#x60; - The SHA-512 cryptographic hash, as defined by NIST in FIPS 180-4. * &#x60;SHA512_224&#x60; - The SHA-512/224 cryptographic hash, as defined by NIST in FIPS 180-4. * &#x60;SHA512_256&#x60; - The SHA-512/256 cryptographic hash, as defined by NIST in FIPS 180-4. | [optional] [default to "SHA256"]
**KeySpec** | Pointer to [**PkixKeyGenerationSpec**](pkix.KeyGenerationSpec.md) |  | [optional] 
**PrivateKey** | Pointer to **string** | Holds the private key for the API key. | [optional] 
**Purpose** | Pointer to **string** | The purpose of the API Key. | [optional] 
**SigningAlgorithm** | Pointer to **string** | The signing algorithm used by the client to authenticate API requests to Intersight. The signing algorithm must be compatible with the key generation specification. * &#x60;RSASSA-PKCS1-v1_5&#x60; - RSASSA-PKCS1-v1_5 is a RSA signature scheme specified in [RFC 8017](https://tools.ietf.org/html/rfc8017).RSASSA-PKCS1-v1_5 is included only for compatibility with existing applications. * &#x60;RSASSA-PSS&#x60; - RSASSA-PSS is a RSA signature scheme specified in [RFC 8017](https://tools.ietf.org/html/rfc8017).It combines the RSASP1 and RSAVP1 primitives with the EMSA-PSS encoding method.In the interest of increased robustness, RSASSA-PSS is required in new applications. * &#x60;Ed25519&#x60; - The Ed25519 signature algorithm, as specified in [RFC 8032](https://tools.ietf.org/html/rfc8032).Ed25519 is a public-key signature system with several attractive features, includingfast single-signature verification, very fast signing, fast key generation and high security level. * &#x60;Ecdsa&#x60; - The Elliptic Curve Digital Signature Standard (ECDSA), as defined by NIST in FIPS 186-4 and ANS X9.62. | [optional] [default to "RSASSA-PKCS1-v1_5"]
**Permission** | Pointer to [**IamPermissionRelationship**](iam.Permission.Relationship.md) |  | [optional] 
**User** | Pointer to [**IamUserRelationship**](iam.User.Relationship.md) |  | [optional] 

## Methods

### NewIamApiKeyAllOf

`func NewIamApiKeyAllOf() *IamApiKeyAllOf`

NewIamApiKeyAllOf instantiates a new IamApiKeyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamApiKeyAllOfWithDefaults

`func NewIamApiKeyAllOfWithDefaults() *IamApiKeyAllOf`

NewIamApiKeyAllOfWithDefaults instantiates a new IamApiKeyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHashAlgorithm

`func (o *IamApiKeyAllOf) GetHashAlgorithm() string`

GetHashAlgorithm returns the HashAlgorithm field if non-nil, zero value otherwise.

### GetHashAlgorithmOk

`func (o *IamApiKeyAllOf) GetHashAlgorithmOk() (*string, bool)`

GetHashAlgorithmOk returns a tuple with the HashAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashAlgorithm

`func (o *IamApiKeyAllOf) SetHashAlgorithm(v string)`

SetHashAlgorithm sets HashAlgorithm field to given value.

### HasHashAlgorithm

`func (o *IamApiKeyAllOf) HasHashAlgorithm() bool`

HasHashAlgorithm returns a boolean if a field has been set.

### GetKeySpec

`func (o *IamApiKeyAllOf) GetKeySpec() PkixKeyGenerationSpec`

GetKeySpec returns the KeySpec field if non-nil, zero value otherwise.

### GetKeySpecOk

`func (o *IamApiKeyAllOf) GetKeySpecOk() (*PkixKeyGenerationSpec, bool)`

GetKeySpecOk returns a tuple with the KeySpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySpec

`func (o *IamApiKeyAllOf) SetKeySpec(v PkixKeyGenerationSpec)`

SetKeySpec sets KeySpec field to given value.

### HasKeySpec

`func (o *IamApiKeyAllOf) HasKeySpec() bool`

HasKeySpec returns a boolean if a field has been set.

### GetPrivateKey

`func (o *IamApiKeyAllOf) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *IamApiKeyAllOf) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *IamApiKeyAllOf) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *IamApiKeyAllOf) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetPurpose

`func (o *IamApiKeyAllOf) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *IamApiKeyAllOf) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *IamApiKeyAllOf) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *IamApiKeyAllOf) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### GetSigningAlgorithm

`func (o *IamApiKeyAllOf) GetSigningAlgorithm() string`

GetSigningAlgorithm returns the SigningAlgorithm field if non-nil, zero value otherwise.

### GetSigningAlgorithmOk

`func (o *IamApiKeyAllOf) GetSigningAlgorithmOk() (*string, bool)`

GetSigningAlgorithmOk returns a tuple with the SigningAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningAlgorithm

`func (o *IamApiKeyAllOf) SetSigningAlgorithm(v string)`

SetSigningAlgorithm sets SigningAlgorithm field to given value.

### HasSigningAlgorithm

`func (o *IamApiKeyAllOf) HasSigningAlgorithm() bool`

HasSigningAlgorithm returns a boolean if a field has been set.

### GetPermission

`func (o *IamApiKeyAllOf) GetPermission() IamPermissionRelationship`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *IamApiKeyAllOf) GetPermissionOk() (*IamPermissionRelationship, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *IamApiKeyAllOf) SetPermission(v IamPermissionRelationship)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *IamApiKeyAllOf) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetUser

`func (o *IamApiKeyAllOf) GetUser() IamUserRelationship`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *IamApiKeyAllOf) GetUserOk() (*IamUserRelationship, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *IamApiKeyAllOf) SetUser(v IamUserRelationship)`

SetUser sets User field to given value.

### HasUser

`func (o *IamApiKeyAllOf) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


