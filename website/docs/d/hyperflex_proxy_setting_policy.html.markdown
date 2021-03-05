---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_proxy_setting_policy"
description: |-
  A policy specifying the HTTP proxy settings to be used by the HyperFlex installation process and HyperFlex storage controller VMs. This policy is required when the internet access of your servers including CIMC and HyperFlex storage controller VMs is secured by a HTTP proxy.
---

# Data Source: intersight_hyperflex_proxy_setting_policy
A policy specifying the HTTP proxy settings to be used by the HyperFlex installation process and HyperFlex storage controller VMs. This policy is required when the internet access of your servers including CIMC and HyperFlex storage controller VMs is secured by a HTTP proxy.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_proxy_setting_policy.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string) Description of the policy. 
* `hostname`:(string) HTTP Proxy server FQDN or IP. 
* `is_password_set`:(bool) Indicates whether the value of the 'password' property has been set. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `password`:(string) The password for the HTTP Proxy. 
* `port`:(int) The HTTP Proxy port number.The port number of the HTTP proxy must be between 1 and 65535, inclusive. 
* `username`:(string) The username for the HTTP Proxy. 
 