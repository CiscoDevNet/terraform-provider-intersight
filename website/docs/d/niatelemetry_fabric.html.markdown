---
subcategory: "niatelemetry"
layout: "intersight"
page_title: "Intersight: intersight_niatelemetry_fabric"
description: |-
        A fabric managed by a nexus dashboard instance.

---

# Data Source: intersight_niatelemetry_fabric
A fabric managed by a nexus dashboard instance.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_niatelemetry_fabric.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `license_tier`:(string) The license tier of the fabric. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Returns the name of the fabric. 
* `network_count`:(int) The network count of the fabric. 
* `owner_cluster_name`:(string) The name of the cluster this fabric belongs to. 
* `power_status`:(string) The power status of the fabric. 
* `power_unit`:(string) The power unit of the fabric. 
* `power_value`:(string) The power value of the fabric. 
* `security_domain`:(string) The security domain of the fabric. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `telemetry_enabled`:(bool) Indicates whether telemetry is enabled for the fabric. 
* `vrf_count`:(int) The VRF count of the fabric. 
 
