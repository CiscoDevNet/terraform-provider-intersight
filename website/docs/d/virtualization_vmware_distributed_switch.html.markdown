---
subcategory: "virtualization"
layout: "intersight"
page_title: "Intersight: intersight_virtualization_vmware_distributed_switch"
description: |-
        The VMware Distributed Virtual Switch object is represented here.

---

# Data Source: intersight_virtualization_vmware_distributed_switch
The VMware Distributed Virtual Switch object is represented here.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_virtualization_vmware_distributed_switch.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Switch description (user provided), if any. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `identity`:(string) The internally generated identity of this switch. This entity is not manipulated by users. It aids in uniquely identifying the switch object. For VMware, this is MOR (managed object reference). 
* `max_port`:(int) Maximum number of ports allowed on this distributed virtual switch. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `mtu`:(int) Maximum transmission unit configured on a distributed virtual switch. 
* `name`:(string) User-provided name to identify the switch. 
* `network_io_control`:(bool) If network io control is enabled, will set the value as true. 
* `num_hosts`:(int) The total number of hosts attached to the distributed virtual switch. 
* `num_networks`:(int) The total number of distributed networks in the distributed virtual switch. 
* `num_stand_alone_ports`:(int) Number of stand-alone ports in use. 
* `num_uplinks`:(int) Number of uplinks configured in this distributed virtual switch. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `uuid`:(string) Universally Unique Id of this distributed virtual switch. 
* `nr_version`:(string) The running config's version details are represented. 
 
