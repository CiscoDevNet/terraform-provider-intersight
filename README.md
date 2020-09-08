# Terraform provider for Cisco Intersight

- Website: https://www.terraform.io

<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="600px">

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) 0.12+
- [Go](https://golang.org/doc/install) (to build the provider plugin)

## Using the Provider

The latest provider can be downloaded as a zip from the [release](https://github.com/CiscoDevNet/terraform-provider-intersight/releases)
section in this repository. The downloaded file has providers for Windows, Darwin (Mac OS) and Linux 64 bit architecture
systems.
Documentation about the provider resources and configuration options can be found
[here](https://github.com/CiscoDevNet/terraform-provider-intersight/tree/master/website/docs).
For a detailed explanation on usage of Intersight Terraform Provider, please read the
[user guide](https://github.com/CiscoDevNet/terraform-provider-intersight/blob/master/USERGUIDE.md).
Refer to [examples](https://github.com/CiscoDevNet/terraform-provider-intersight/tree/master/examples)
for terraform configuration files of server configuration and OS installation.

## Developing the Provider

Clone repository to: `$GOPATH/src/github.com/CiscoDevNet/`
```sh
$ mkdir -p $GOPATH/src/github.com/CiscoDevNet/; cd $GOPATH/src/github.com/CiscoDevNet/
$ git clone git@github.com:CiscoDevNet/terraform-provider-intersight.git
$ cd terraform-provider-intersight
```

To compile the provider, run `make`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.
Move the contents of ```.build``` to `%APPDATA%\terraform.d\plugins` in case of Windows and
`~/.terraform.d/plugins` for most other operating systems.

## Community
We are on Slack - slack requires registration, but the terraform-provider-intersight team is open invitation to anyone
to register
[here](https://join.slack.com/t/cisco-intersight/shared_invite/enQtNzYzODk5MzMzNDE1LTAxNzA5YmIwYzEwN2JiODMwMmEzODYyNzg1MDQ4MGY0NmFmNTNiNGYxMTZhNjE4MWQzMTRiMmFlZGFhY2QyMWQ)
