module github.com/CiscoDevNet/terraform-provider-intersight

go 1.19

require (
	github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk v0.0.0
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.4.2
	golang.org/x/net v0.0.0-20220722155237-a158d28d115b // indirect
	golang.org/x/oauth2 v0.0.0-20210323180902-22b0adad7558
	golang.org/x/sys v0.0.0-20220722155257-8c9f86f7a55f // indirect
)

replace github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk v0.0.0 => ./intersight_gosdk
