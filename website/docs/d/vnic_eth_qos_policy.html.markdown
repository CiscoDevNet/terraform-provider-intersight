---
subcategory: "vnic"
layout: "intersight"
page_title: "Intersight: intersight_vnic_eth_qos_policy"
description: |-
        An Ethernet Quality of Service (QoS) policy assigns a system class to the outgoing traffic for a vNIC. This system class determines the quality of service for the outgoing traffic. For certain adapters additional controls can be specified like burst and rate on the outgoing traffic.

---

# Data Source: intersight_vnic_eth_qos_policy
An Ethernet Quality of Service (QoS) policy assigns a system class to the outgoing traffic for a vNIC. This system class determines the quality of service for the outgoing traffic. For certain adapters additional controls can be specified like burst and rate on the outgoing traffic.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_vnic_eth_qos_policy.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `burst`:(int) The burst traffic, in bytes, allowed on the vNIC. 
* `cos`:(int) Class of Service to be associated to the traffic on the virtual interface. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the policy. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `mtu`:(int) The Maximum Transmission Unit (MTU) or packet size that the virtual interface accepts. 
* `name`:(string) Name of the concrete policy. 
* `priority`:(string) The priority matching the System QoS specified in the fabric profile.* `Best Effort` - QoS Priority for Best-effort traffic.* `FC` - QoS Priority for FC traffic.* `Platinum` - QoS Priority for Platinum traffic.* `Gold` - QoS Priority for Gold traffic.* `Silver` - QoS Priority for Silver traffic.* `Bronze` - QoS Priority for Bronze traffic. 
* `rate_limit`:(int) The value in Mbps (0-10G/40G/100G depending on Adapter Model) to use for limiting the data rate on the virtual interface. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `trust_host_cos`:(bool) Enables usage of the Class of Service provided by the operating system. 
 
