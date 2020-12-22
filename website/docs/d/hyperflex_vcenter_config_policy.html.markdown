---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_vcenter_config_policy"
description: |-
  A policy specifying vCenter configuration.
---

# Data Source: intersight_hyperflex_vcenter_config_policy
A policy specifying vCenter configuration.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `data_center`:(string) The vCenter datacenter name. 
* `description`:(string) Description of the policy. 
* `hostname`:(string) The vCenter server FQDN or IP. 
* `is_password_set`:(bool) Indicates whether the value of the 'password' property has been set. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `password`:(string) The password for authenticating with vCenter. Follow the corresponding password policy governed by vCenter. 
* `sso_url`:(string) Overrides the default vCenter Single Sign-On URL. Do not specify unless instructed by Cisco TAC. 
* `username`:(string) The vCenter username (e.g. administrator@vsphere.local). 
