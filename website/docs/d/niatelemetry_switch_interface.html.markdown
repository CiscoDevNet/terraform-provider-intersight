---
subcategory: "niatelemetry"
layout: "intersight"
page_title: "Intersight: intersight_niatelemetry_switch_interface"
description: |-
        An interface managed by a nexus dashboard instance.

---

# Data Source: intersight_niatelemetry_switch_interface
An interface managed by a nexus dashboard instance.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_niatelemetry_switch_interface.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `admin_status`:(string) Administrative status of the interface. 
* `create_time`:(string) The time when this managed object was created. 
* `discovered`:(bool) Indicates whether the interface was discovered. 
* `discovered_config_mode`:(string) Discovered configuration mode of the interface. 
* `display_name`:(string) Display name of the interface. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `external_fabric`:(bool) Indicates whether the interface belongs to an external fabric. 
* `fabric_name`:(string) Fabric name for the interface. 
* `interface_name`:(string) The name of the interface. 
* `interface_type`:(string) The type of the interface. 
* `ip`:(string) IP address configured on the interface. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `oper_mode`:(string) Operational mode of the interface. 
* `operational_status`:(string) Operational status of the interface. 
* `platform`:(string) Platform of the interface. 
* `port_channel_id`:(int) The portchannel ID of the associated interface. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `speed`:(string) The speed of the interface. 
* `switch_id`:(string) Identifier of the switch that owns the interface. 
* `switch_name`:(string) Name of the switch that owns the interface. 
* `vlan_id`:(int) VLAN identifier associated with the interface. 
* `vpc_id`:(int) VPC identifier for the interface. 
 
