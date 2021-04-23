---
subcategory: "fabric"
layout: "intersight"
page_title: "Intersight: intersight_fabric_link_aggregation_policy"
description: |-
  A policy to configure the link settings for all the port channels (including LACP).
---

# Data Source: intersight_fabric_link_aggregation_policy
A policy to configure the link settings for all the port channels (including LACP).
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_fabric_link_aggregation_policy.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the policy. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `lacp_rate`:(string) Flag used to indicate whether LACP PDUs are to be sent 'fast', i.e., every 1 second.* `normal` - The expanded 4th generation UCS Fabric Interconnect with 108 ports.* `fast` - The standard 4th generation UCS Fabric Interconnect with 54 ports. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `suspend_individual`:(bool) Flag tells the switch whether to suspend the port if it didn’t receive LACP PDU. 
 