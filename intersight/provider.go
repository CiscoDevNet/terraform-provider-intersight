package intersight

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"apikey": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("INTERSIGHT_API_KEY", nil),
				Description: "API Key Id",
			},
			"secretkey": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("INTERSIGHT_SECRET_KEY", nil),
				Description: "Secret Key String or File Path",
			},
			"endpoint": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("INTERSIGHT_ENDPOINT_URL", nil),
				Description: "Endpoint URL",
			},
		},
		ResourcesMap:         GetResourceMapping(),
		DataSourcesMap:       GetDataSourceMapping(),
		ConfigureContextFunc: configureProvider,
	}
}

func configureProvider(c context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var de diag.Diagnostics
	var err error
	endpoint := d.Get("endpoint").(string)
	if strings.HasPrefix(endpoint, "https://") || strings.HasPrefix(endpoint, "http://") {
		endpoint = strings.TrimPrefix(endpoint, "https://")
		endpoint = strings.TrimPrefix(endpoint, "http://")
	}
	config := Config{
		ApiKey:    d.Get("apikey").(string),
		SecretKey: d.Get("secretkey").(string),
		Endpoint:  endpoint,
	}

	client := Client{}
	config.ctx, err = client.SetInputs(config.ApiKey, config.SecretKey, config.Endpoint, true)
	if err != nil {
		return nil, diag.Errorf("Failed to initialize Intersight client: %s", err)
	}
	config.ApiClient = client.GetApiClient(config.ctx, true)
	return &config, de
}
