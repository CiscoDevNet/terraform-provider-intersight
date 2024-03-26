---
subcategory: "adapter"
layout: "intersight"
page_title: "Intersight: intersight_adapter_host_eth_interface"
description: |-
        Physical / Virtual port of an adapter as seen by the host.

---

# Data Source: intersight_adapter_host_eth_interface
Physical / Virtual port of an adapter as seen by the host.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_adapter_host_eth_interface.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `admin_state`:(string) Admin state of the Host Ethernet Interface. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `ep_dn`:(string) The Endpoint Config Dn of the Host Ethernet Interface. 
* `host_eth_interface_id`:(int) Unique Identifier for an Host Ethernet Interface within the adapter object. 
* `interface_type`:(string) Type of External Ethernet Interface. 
* `mac_address`:(string) Mac address of the Host Ethernet Interface. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of Host Ethernet Interface. 
* `oper_state`:(string) Operational state of an Interface. 
* `operability`:(string) Operability status of Host Ethernet Channel Interface. 
* `original_mac_address`:(string) The factory default Mac address of the Host Ethernet Interface. 
* `pci_addr`:(string) The PCI address of the Host Ethernet Interface. 
* `peer_dn`:(string) The distinguished name of the peer endpoint connected to the Host Ethernet interface. 
* `pin_group_name`:(string) Name given for Lan PinGroup. 
* `qinq_enabled`:(bool) Setting qinqEnabled to true if we have QinQ tagging enabled on the vNIC. 
* `qinq_vlan`:(int) The VLAN ID for VIC QinQ (802.1Q) Tunneling. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `standby_oper_state`:(string) Standby Operational state of an Interface. 
* `standby_vif_id`:(int) Identifier of the Standby virtual ethernet interface (Vethernet) on the networking component (e.g., Fabric Interconnect) for the corresponding Host Ethernet Interface (vNIC). 
* `veth_action`:(string) The action to be performed on the vethernet corresponding to the vNIC.* `None` - Default value for vif operation.* `ResetConnectivity` - Resets connectivity on both active and passive vif.* `ResetConnectivityActive` - Resets connectivity on the active vif.* `ResetConnectivityPassive` - Resets connectivity on the passive vif.* `Enable` - Enables the vif on both the FIs.* `Disable` - Disables the vif on both the FIs.* `EnableActive` - Enables the corresponding active vif.* `EnablePassive` - Enables the corresponding standby vif.* `DisableActive` - Disables the corresponding active vif.* `DisablePassive` - Disables the corresponding standby vif. 
* `vif_id`:(int) Identifier of the virtual ethernet interface (Vethernet) on the networking component (e.g., Fabric Interconnect) for the corresponding Host Ethernet Interface (vNIC). 
* `virtualization_preference`:(string) Virtualization Preference of the Host Ethernet Interface indicating if virtualization is enabled or not. 
* `vnic_dn`:(string) The Virtual Ethernet Interface DN connected to the Host Ethernet Interface. 
 
