---
subcategory: "bios"
layout: "intersight"
page_title: "Intersight: intersight_bios_token_settings"
description: |-
        The bios.TokenSettings object represents the status of a specific BIOS token, particularly for Memory Reliability, Availability, and Serviceability (RAS) configurations. It indicates whether a given RAS setting is currently active on the server.
        #### Purpose
        The primary function of this object is to provide a granular view of which memory RAS feature is currently enabled in the BIOS. By querying this object, administrators can confirm if a setting like Sparing or Mirroring is the active configuration, which is crucial for verifying server resilience and performance settings.
        #### Key Concepts
        - **Granular Configuration State:** Provides the status (isAssigned) for a single, specific BIOS setting identified by settingsMoRn.
        - **Memory RAS Focus:** Specifically designed to report on memory-related RAS configurations, such as ADDDC-Sparing or Mirror-Mode-1LM.
        - **Verification Tool:** Allows for programmatic verification that a desired memory RAS policy has been successfully applied to the server.
        - **Server-Specific Inventory:** Tied directly to a physical server, providing an accurate snapshot of its active BIOS token settings.

---

# Data Source: intersight_bios_token_settings
The bios.TokenSettings object represents the status of a specific BIOS token, particularly for Memory Reliability, Availability, and Serviceability (RAS) configurations. It indicates whether a given RAS setting is currently active on the server.
#### Purpose
The primary function of this object is to provide a granular view of which memory RAS feature is currently enabled in the BIOS. By querying this object, administrators can confirm if a setting like "Sparing" or "Mirroring" is the active configuration, which is crucial for verifying server resilience and performance settings.
#### Key Concepts
- **Granular Configuration State:** Provides the status (isAssigned) for a single, specific BIOS setting identified by settingsMoRn.
- **Memory RAS Focus:** Specifically designed to report on memory-related RAS configurations, such as "ADDDC-Sparing" or "Mirror-Mode-1LM".
- **Verification Tool:** Allows for programmatic verification that a desired memory RAS policy has been successfully applied to the server.
- **Server-Specific Inventory:** Tied directly to a physical server, providing an accurate snapshot of its active BIOS token settings.
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
 
