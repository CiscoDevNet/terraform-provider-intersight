---
subcategory: "niatelemetry"
layout: "intersight"
page_title: "Intersight: intersight_niatelemetry_lc"
description: |-
  Object is available at Line Card scope.
---

# Data Source: intersight_niatelemetry_lc
Object is available at Line Card scope.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_niatelemetry_lc.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the line cards present. 
* `dn`:(string) Dn value for the line cards present. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `hardware_version`:(string) Hardware version of the line cards present. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) Model of the line cards present. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `node_id`:(int) Node Id of the line card present. 
* `operational_state`:(string) Opretaional state of the line cards present. 
* `power_state`:(string) Power state of the line cards present. 
* `record_type`:(string) Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. 
* `record_version`:(string) Version of record being pushed. This determines what was the API version for data available from the device. 
* `redundancy_state`:(string) Redundancy state of the line cards present. 
* `serial_number`:(string) Serial number of the line card present. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `site_name`:(string) The Site name represents an APIC cluster. Service Engine can onboard multiple APIC clusters / sites. 
* `vid`:(string) VID for the line card in the inventory. 
 