---
subcategory: "fabric"
layout: "intersight"
page_title: "Intersight: intersight_fabric_net_flow_monitor"
description: |-
        A NetFlow monitor combines a NetFlow record and a NetFlow exporter, and is used to monitor traffic on the interface where it is applied.

---

# Data Source: intersight_fabric_net_flow_monitor
A NetFlow monitor combines a NetFlow record and a NetFlow exporter, and is used to monitor traffic on the interface where it is applied.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_fabric_net_flow_monitor.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Netflow Monitor name, must be a maximum of 63 characters, without spacing. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `vnic_usage_count`:(int) The count of the NetFlow monitor usage on vNICs. 
 
