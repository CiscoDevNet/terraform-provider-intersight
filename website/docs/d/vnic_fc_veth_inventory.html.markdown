---
subcategory: "vnic"
layout: "intersight"
page_title: "Intersight: intersight_vnic_fc_veth_inventory"
description: |-
        The FcVethInventories object captures the QoS settings on the Fabric Interconnect corresponding to a server's virtual Fibre Channel interface (vHBA).
        #### Purpose
        It monitors the fabric-side configuration of virtual storage interfaces, ensuring that traffic shaping and performance settings are correctly enforced.
        #### Key Concepts
        - **Fabric-Level QoS:** Manages rate limiting and burst settings for FC traffic.
        - **Operational Monitoring:** Tracks the current configuration of virtual FC ports on the switch.

---

# Data Source: intersight_vnic_fc_veth_inventory
The FcVethInventories object captures the QoS settings on the Fabric Interconnect corresponding to a server's virtual Fibre Channel interface (vHBA).
#### Purpose
It monitors the fabric-side configuration of virtual storage interfaces, ensuring that traffic shaping and performance settings are correctly enforced.
#### Key Concepts
- **Fabric-Level QoS:** Manages rate limiting and burst settings for FC traffic.
- **Operational Monitoring:** Tracks the current configuration of virtual FC ports on the switch.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_vnic_fc_veth_inventory.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `burst`:(int) The burst traffic, in bytes, allowed on the vNIC. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the virtual FC interface. 
* `device_mo_id`:(string) Device ID of the entity from where inventory is reported. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the virtual FC interface. 
* `rate_limit`:(int) The value in Mbps to use for limiting the data rate on the virtual interface. Setting this to zero will turn rate limiting off. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
