---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_ucsm_config_policy"
description: |-
  A policy specifying UCS Manager settings such as service profile org, KVM IP addresses, and MAC prefix for server configuration in Fabric Interconnect environment.
---

# Data Source: intersight_hyperflex_ucsm_config_policy
A policy specifying UCS Manager settings such as service profile org, KVM IP addresses, and MAC prefix for server configuration in Fabric Interconnect environment.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_ucsm_config_policy.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string) Description of the policy. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `server_firmware_version`:(string) The server firmware bundle version used for server components such as CIMC, adapters, BIOS, etc. 
 