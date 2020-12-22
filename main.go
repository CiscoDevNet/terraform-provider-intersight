package main

import (
	"github.com/CiscoDevNet/terraform-provider-intersight/intersight"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: intersight.Provider,
	})
}
