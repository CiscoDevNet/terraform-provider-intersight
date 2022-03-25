---
subcategory: "virtualization"
layout: "intersight"
page_title: "Intersight: intersight_virtualization_vmware_distributed_network"
description: |-
        The VMware Distributed Virtual PortGroup object is represented here.

---

# Data Source: intersight_virtualization_vmware_distributed_network
The VMware Distributed Virtual PortGroup object is represented here.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_virtualization_vmware_distributed_network.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `forged_transmits`:(string) If forgedTransmits property value is set to reject, outbound frames with a source MAC address different from the one set on the adapter are dropped. If property value is set to accept, no filtering is performed and all outbound frames are passed.* `Reject` - Indicates that the security policy is rejected.* `Accept` - Indicates that the security policy is accepted. 
* `identity`:(string) The internally generated identity of network. This entity cannot manipulated by users. It aids in uniquely identifying the network object. For VMware, this is MOR (managed object reference). 
* `mac_address_changes`:(string) If macAddressChanges property value is set to reject and the MAC address of the adapter is changed to a value other than the one specified in .vmx configuration file, all inbound frames are dropped. If property value is set to accept and the MAC address is changed, inbound frames to the new MAC address are received.* `Reject` - Indicates that the security policy is rejected.* `Accept` - Indicates that the security policy is accepted. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) User-provided name to identify the portgroup. 
* `num_hosts`:(int) The total number of hosts connected to this distributed virtual network. 
* `num_ports`:(int) The total number of ports in the distributed virtual network. 
* `promiscuous_mode`:(string) If promiscuousMode property value is set to reject, incoming traffic only targeted to that network will be visible. If property value is set to accept, objects defined within the network can see all incoming traffic on the virtual switch based on the VLAN policy.* `Reject` - Indicates that the security policy is rejected.* `Accept` - Indicates that the security policy is accepted. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `up_link`:(bool) Indicates if the distributed virtual network is a uplink. 
* `vlan_type`:(string) VLAN type of the distributed virtual network. It can be None, VLAN, VLAN Trunking or Private VLAN.* `None` - Do not tag traffic with any VLAN Id.* `VLAN` - Tag traffic with the Id from the VLAN Id field.* `VLAN trunking` - Pass VLAN traffic with Id within the VLAN trunk range to guest operating system.* `Private VLAN` - Associate the traffic with a private VLAN created on the distributed switch. 
 
