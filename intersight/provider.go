package intersight

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"apikey": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("INTERSIGHT_API_KEY", nil),
				Description: "API Key Id",
			},
			"secretkeyfile": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("INTERSIGHT_SECRET_KEY_FILE_PATH", nil),
				Description: "Secret Key File Path",
			},
			"endpoint": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("INTERSIGHT_ENDPOINT_URL", nil),
				Description: "Endpoint URL",
			},
		},
		ResourcesMap:   GetResourceMapping(),
		DataSourcesMap: GetDataSourceMapping(),
		ConfigureFunc:  configureProvider,
	}
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
	endpoint := d.Get("endpoint").(string)
	if strings.HasPrefix(endpoint, "https://") || strings.HasPrefix(endpoint, "http://") {
		endpoint = strings.TrimPrefix(endpoint, "https://")
		endpoint = strings.TrimPrefix(endpoint, "http://")
	}
	config := Config{
		ApiKey:        d.Get("apikey").(string),
		SecretKeyFile: d.Get("secretkeyfile").(string),
		Endpoint:      endpoint,
	}

	client := Client{}
	config.ctx = client.SetInputs(config.ApiKey, config.SecretKeyFile, config.Endpoint, true)
	config.ApiClient = client.GetApiClient(config.ctx, true)
	return &config, nil
}
