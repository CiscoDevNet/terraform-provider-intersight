module github.com/CiscoDevNet/terraform-provider-intersight

go 1.15

require (
	github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk v0.0.0
	github.com/hashicorp/terraform-plugin-sdk v1.7.0
	github.com/stretchr/testify v1.4.0 // indirect
	golang.org/x/crypto v0.0.0-20191011191535-87dc89f01550 // indirect
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
	gopkg.in/yaml.v2 v2.2.4 // indirect
)

replace github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk v0.0.0 => ./intersight_gosdk
