---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_net_app_node_cdp_neighbor"
description: |-
        Information about the CDP neighbor connected to a given NetApp node port.

---

# Data Source: intersight_storage_net_app_node_cdp_neighbor
Information about the CDP neighbor connected to a given NetApp node port.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_net_app_node_cdp_neighbor.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `device_ip`:(string) The IP address of the CDP neighbor. 
* `discovered_device`:(string) The name of the CDP neighbor. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `hold_time_remaining`:(int) The period of time for which CDP advertisements are cached. 
* `interface`:(string) The interface of the CDP neighbor. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `node_name`:(string) The node that owns the port for this CDP neighbor. 
* `platform`:(string) The platform of the CDP neighbor. 
* `port`:(string) The port for this CDP neighbor. 
* `protocol`:(string) The protocol used for CDP. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `nr_version`:(string) The version of the operating system running on the CDP neighbor. 
 
