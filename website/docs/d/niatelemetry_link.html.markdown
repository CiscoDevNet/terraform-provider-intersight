---
subcategory: "niatelemetry"
layout: "intersight"
page_title: "Intersight: intersight_niatelemetry_link"
description: |-
        A link managed by a nexus dashboard instance.

---

# Data Source: intersight_niatelemetry_link
A link managed by a nexus dashboard instance.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_niatelemetry_link.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `dst_fabric_name`:(string) Destination fabric name for the link. 
* `dst_interface_name`:(string) Destination interface name for the link. 
* `dst_switch_id`:(string) Destination switch identifier for the link. 
* `dst_switch_model_name`:(string) Destination switch model for the link. 
* `dst_switch_name`:(string) Destination switch name for the link. 
* `dst_switch_role`:(string) Destination switch role for the link. 
* `link_discovered`:(bool) Indicates whether the link is discovered. 
* `link_id`:(string) Unique identifier for the link. 
* `link_planned`:(bool) Indicates whether the link is planned. 
* `link_present`:(bool) Indicates whether the link is present. 
* `link_type`:(string) A description of the type of link. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `port_channel`:(bool) Indicates whether the link is a port channel. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `src_fabric_name`:(string) Source fabric name for the link. 
* `src_interface_admin_status`:(string) Administrative status of the source interface. 
* `src_interface_name`:(string) Source interface name for the link. 
* `src_interface_oper_status`:(string) Operational status of the source interface. 
* `src_switch_id`:(string) Source switch identifier for the link. 
* `src_switch_model_name`:(string) Source switch model for the link. 
* `src_switch_name`:(string) Source switch name for the link. 
* `src_switch_role`:(string) Source switch role for the link. 
 
