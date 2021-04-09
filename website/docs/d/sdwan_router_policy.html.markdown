---
subcategory: "sdwan"
layout: "intersight"
page_title: "Intersight: intersight_sdwan_router_policy"
description: |-
  A policy specifying SD-WAN router details.
---

# Data Source: intersight_sdwan_router_policy
A policy specifying SD-WAN router details.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_sdwan_router_policy.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `deployment_size`:(string) Scale of the SD-WAN router virtual machine deployment.* `Typical` - Typical deployment configuration with 4 vCPUs and 4GB RAM.* `Minimal` - Minimal deployment configuration with 2 vCPUs and 4GB RAM. 
* `description`:(string) Description of the policy. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `wan_count`:(int) Number of WAN connections across the SD-WAN site. 
* `wan_termination_type`:(string) Defines if the WAN networks are singly or dually terminated. Dually terminated WANs are configured on all the SD-WAN routers. Singly terminated WANs are configured only on one of the SD-WAN routers.* `Single` - Singly terminated WANs ar evenly distributed across SD-WAN router nodes. A given WAN connection is available only on one of the router nodes.* `Dual` - Dually terminated WANs are configured on all the SD-WAN routers. A given WAN connection is available on multiple router nodes. 
 