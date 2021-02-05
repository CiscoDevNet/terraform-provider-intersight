module github.com/CiscoDevNet/terraform-provider-intersight

go 1.15

require (
	github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk v0.0.0
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.4.2
	golang.org/x/oauth2 v0.0.0-20210201163806-010130855d6c
)

replace github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk v0.0.0 => ./intersight_gosdk
