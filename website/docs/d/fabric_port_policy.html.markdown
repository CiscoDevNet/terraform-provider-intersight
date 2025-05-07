---
subcategory: "fabric"
layout: "intersight"
page_title: "Intersight: intersight_fabric_port_policy"
description: |-
        A policy for all the physical ports of the Fabric Interconnect.

---

# Data Source: intersight_fabric_port_policy
A policy for all the physical ports of the Fabric Interconnect.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_fabric_port_policy.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the policy. 
* `device_model`:(string) This field specifies the device model that this Port Policy is being configured for.* `UCS-FI-6454` - The standard 4th generation UCS Fabric Interconnect with 54 ports.* `UCS-FI-64108` - The expanded 4th generation UCS Fabric Interconnect with 108 ports.* `UCS-FI-6536` - The standard 5th generation UCS Fabric Interconnect with 36 ports.* `UCSX-S9108-100G` - Cisco UCS Fabric Interconnect 9108 100G with 8 ports.* `UCS-FI-6664` - The standard 6th generation UCS Fabric Interconnect with 64 ports.* `unknown` - Unknown device type, usage is TBD. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
