---
subcategory: "virtualization"
layout: "intersight"
page_title: "Intersight: intersight_virtualization_vmware_virtual_switch"
description: |-
  The VMware Virtual Switch object is represented here.
---

# Data Source: intersight_virtualization_vmware_virtual_switch
The VMware Virtual Switch object is represented here.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_virtualization_vmware_virtual_switch.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `forged_transmits`:(string) If forgedTransmits property value is set to reject, outbound frames with a source MAC address different from the one set on the adapter are dropped. If property value is set to accept, the switch does not perform filtering and permits all outbound frames.* `Reject` - Indicates that the security policy is rejected.* `Accept` - Indicates that the security policy is accepted. 
* `identity`:(string) The internally generated identity of this switch. This entity is not manipulated by users. It aids in uniquely identifying the switch object. For VMware, this is MOR (managed object reference). 
* `mac_address_changes`:(string) If macAddressChanges property value is set to reject and the MAC address of the adapter is changed to a value other than the one specified in .vmx configuration file, the switch drops all inbound frames to the adapter. If property value is set to accept and the MAC address is changed, the switch allows frames to the new MAC address to pass.* `Reject` - Indicates that the security policy is rejected.* `Accept` - Indicates that the security policy is accepted. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `mtu`:(int) Maximum transmission unit configured on a virtual switch. 
* `name`:(string) User-provided name to identify the switch. 
* `num_networks`:(int) Number of networks available on this virtual switch. 
* `num_physical_network_interfaces`:(int) Number of physical network interfaces connected with this virtual switch. 
* `promiscuous_mode`:(string) If promiscuousMode property value is set to reject, the virtual switch forwards only frames that are addressed to the adapter. If property value is set to accept, the virtual switch forwards all frames to the adapter in compliance with the active VLAN policy for the port to which it is connected.* `Reject` - Indicates that the security policy is rejected.* `Accept` - Indicates that the security policy is accepted. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 