---
subcategory: "niatelemetry"
layout: "intersight"
page_title: "Intersight: intersight_niatelemetry_apic_vision"
description: |-
        Object to capture ApicVision App properties.

---

# Data Source: intersight_niatelemetry_apic_vision
Object to capture ApicVision App properties.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_niatelemetry_apic_vision.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `apic_vision_state`:(string) ApicVision App state. apicVisionState checks the current operational state of ApicVision app on APIC. 
* `apic_vision_state_last_update_ts`:(string) ApicVision App state last updated timestamp. It indicates the last updated timestamp for operational state of ApicVision app. 
* `apic_vision_version`:(string) ApicVision App version. apicVisionVersion is used to check compatibility with Nexus Cloud features. 
* `config_issues`:(string) Configuration issues depicts the failures for ApicVision managed package upgrade on APIC. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
