---
subcategory: "vnic"
layout: "intersight"
page_title: "Intersight: intersight_vnic_eth_vnic_inventory"
description: |-
        QoS settings of server’s vNIC.

---

# Data Source: intersight_vnic_eth_vnic_inventory
QoS settings of server’s vNIC.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_vnic_eth_vnic_inventory.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `cos`:(int) Class of Service to be associated to the traffic on the virtual interface. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) Device ID of the entity from where inventory is reported. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `mtu`:(int) The Maximum Transmission Unit (MTU) or packet size that the virtual interface accepts. 
* `name`:(string) Name of the virtual ethernet interface. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `trust_host_cos`:(bool) Enables usage of the Class of Service provided by the operating system. 
 
