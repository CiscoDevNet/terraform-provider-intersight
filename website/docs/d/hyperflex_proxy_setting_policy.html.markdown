
---
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_proxy_setting_policy"
sidebar_current: "docs-intersight-data-source-hyperflex-proxy-setting-policy"
description: |-
A policy specifying the HTTP proxy settings to be used by the HyperFlex installation process and HyperFlex storage controller VMs. This policy is required when the internet access of your servers including CIMC and HyperFlex storage controller VMs is secured by a HTTP proxy.
---

# Data Source: intersight_hyperflex._proxy_setting_policy
A policy specifying the HTTP proxy settings to be used by the HyperFlex installation process and HyperFlex storage controller VMs. This policy is required when the internet access of your servers including CIMC and HyperFlex storage controller VMs is secured by a HTTP proxy.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `description`:(string) Description of the policy. 
* `hostname`:(string) HTTP Proxy server FQDN or IP. 
* `is_password_set`:(bool) Indicates whether the value of the 'password' property has been set. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `password`:(string) The password for the HTTP Proxy. 
* `port`:(int) The HTTP Proxy port number.The port number of the HTTP proxy must be between 1 and 65535, inclusive. 
* `username`:(string) The username for the HTTP Proxy. 
