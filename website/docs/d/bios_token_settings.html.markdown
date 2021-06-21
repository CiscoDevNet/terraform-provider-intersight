---
subcategory: "bios"
layout: "intersight"
page_title: "Intersight: intersight_bios_token_settings"
description: |-
  Token settings for Memory Reliability, availability and serviceability (RAS) configuration.
---

# Data Source: intersight_bios_token_settings
Token settings for Memory Reliability, availability and serviceability (RAS) configuration.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_bios_token_settings.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `is_assigned`:(string) Value that represents if the BIOS configuration is active. Possible values are \ yes\  and \ no\ . 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) Parent server serial number. 
* `settings_mo_rn`:(string) BIOS configuration name as found in dn. Possible values are \ ADDDC-Sparing\ , \ Maximum-Performance\ , \ Partial-Mirror-mode-1LM\ , \ Mirror-Mode-1LM\ , \ Mirroring\ , \ Lockstep\ , \ Sparing\ , \ Platform-Default\ . 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 