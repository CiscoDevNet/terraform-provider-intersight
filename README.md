# Terraform provider for Cisco Intersight

- Website: https://www.terraform.io

<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="600px">

## Requirements
- [Terraform](https://www.terraform.io/downloads.html) 0.14+
- [Go](https://golang.org/doc/install) (to build the provider plugin)

## Using the Provider
The Intersight Terraform provider can be downloaded from the 
[release](https://github.com/CiscoDevNet/terraform-provider-intersight/releases) section of this repository. 
The providers are available for Darwin, FreeBSD(386, amd64, arm, arm64), 
Linux(386, amd64, arm, arm64) and Windows(386 and amd64) systems.
Documentation of the provider resources, data sources and configuration options can be found
[here](https://github.com/CiscoDevNet/terraform-provider-intersight/tree/master/website/docs).
For a detailed explanation on usage of Intersight Terraform Provider, please read the
[user guide](https://github.com/CiscoDevNet/terraform-provider-intersight/blob/master/USERGUIDE.md).
Refer to [examples](https://github.com/CiscoDevNet/terraform-provider-intersight/tree/master/examples)
for terraform configuration files of server configuration, deployment and OS installation.

## Developing the Provider
Clone repository to: `$GOPATH/src/github.com/CiscoDevNet/`
```shell
$ mkdir -p $GOPATH/src/github.com/CiscoDevNet/; cd $GOPATH/src/github.com/CiscoDevNet/
$ git clone git@github.com:CiscoDevNet/terraform-provider-intersight.git
$ cd terraform-provider-intersight
```
To compile the provider, run `make`. This will build the providers for Linux, Darwin and Windows 64-bit architecture in the 
`.build`. To use locally build providers, run the following commands in the directory where the configurations file reside.
```shell
mkdir -p terraform.d/plugins/registry.terraform.io/ciscodevnet/intersight/<PROVIDER_VERSION>/<OS_ARCH>
cp $GOPATH/src/github.com/CiscoDevNet/terraform-provider-intersight/.build/<OS_ARCH>/terraform-provider-intersight_v<PROVIDER_VERSION>
```
`PROVIDER_VERSION` should be the same version as mentioned in the `makefile` of the provider repository.
`OS_ARCH` can be one of `linux_amd64`, `darwin_amd64` and `windows`.

## Frequency of update
The provider is updated with each production push on Intersight, i.e., every Friday. Exceptions are made for high priority bug fixes.

## Contribution
The community can contribute to expanding the example base for the provider. This includes HCL configurations for resources and data sources.
A large part of the provider is generated using a custom generator. Hence, pull requests only on the `examples` directory are welcome at this time. 
For all other change requests, raise an issue.

## Community
We are on Slack - slack requires registration, but the terraform-provider-intersight team is open invitation to anyone
to register
[here](https://join.slack.com/t/cisco-intersight/shared_invite/enQtNzYzODk5MzMzNDE1LTAxNzA5YmIwYzEwN2JiODMwMmEzODYyNzg1MDQ4MGY0NmFmNTNiNGYxMTZhNjE4MWQzMTRiMmFlZGFhY2QyMWQ)
