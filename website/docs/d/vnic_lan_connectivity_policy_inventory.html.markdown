---
subcategory: "vnic"
layout: "intersight"
page_title: "Intersight: intersight_vnic_lan_connectivity_policy_inventory"
description: |-
        The LanConnectivityPolicy object defines the network resources and connection topology between servers and the LAN within Cisco systems. It abstracts the configuration of virtual network interfaces (vNICs), ensuring consistent device naming, optimized placement, and advanced features such as usNIC and VMQ for improved network performance.
        #### Purpose
        A LanConnectivityPolicy centralizes the management of LAN connectivity for servers by specifying how virtual interfaces are created and associated with network resources. It enables both automatic and custom placement of vNICs, supports the assignment of MAC addresses, and integrates with relevant network and adapter policies.
        #### Key Concepts
        - **Policy-Based Networking:** Facilitates consistent and scalable LAN connectivity across multiple servers.
        - **Integration with Network Policies:** Links to Ethernet network, QoS, and adapter policies for comprehensive configuration.
        - **Device Naming and Placement:** Supports consistent device naming and controls vNIC placement for optimal performance.

---

# Data Source: intersight_vnic_lan_connectivity_policy_inventory
The LanConnectivityPolicy object defines the network resources and connection topology between servers and the LAN within Cisco systems. It abstracts the configuration of virtual network interfaces (vNICs), ensuring consistent device naming, optimized placement, and advanced features such as usNIC and VMQ for improved network performance.
#### Purpose
A LanConnectivityPolicy centralizes the management of LAN connectivity for servers by specifying how virtual interfaces are created and associated with network resources. It enables both automatic and custom placement of vNICs, supports the assignment of MAC addresses, and integrates with relevant network and adapter policies.
#### Key Concepts
- **Policy-Based Networking:** Facilitates consistent and scalable LAN connectivity across multiple servers.
- **Integration with Network Policies:** Links to Ethernet network, QoS, and adapter policies for comprehensive configuration.
- **Device Naming and Placement:** Supports consistent device naming and controls vNIC placement for optimal performance.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_vnic_lan_connectivity_policy_inventory.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `azure_qos_enabled`:(bool) Enabling AzureStack-Host QoS on an adapter allows the user to carve out traffic classes for RDMA traffic which ensures that a desired portion of the bandwidth is allocated to it. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the policy. 
* `device_mo_id`:(string) Device ID of the entity from where inventory is reported. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `iqn_allocation_type`:(string) Allocation Type of iSCSI Qualified Name - Static/Pool/None.* `None` - Type indicates that there is no IQN associated to an interface.* `Static` - Type represents that static IQN is associated to an interface.* `Pool` - Type indicates that IQN value is sourced from an associated pool. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the inventoried policy object. 
* `placement_mode`:(string) The mode used for placement of vNICs on network adapters. It can either be Auto or Custom.* `custom` - The placement of the vNICs / vHBAs on network adapters is manually chosen by the user.* `auto` - The placement of the vNICs / vHBAs on network adapters is automatically determined by the system. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `static_iqn_name`:(string) User provided static iSCSI Qualified Name (IQN) for use as initiator identifiers by iSCSI vNICs in a Fabric Interconnect domain. 
* `target_platform`:(string) The platform for which the server profile is applicable. It can either be a server that is operating in standalone mode or which is attached to a Fabric Interconnect managed by Intersight.* `Standalone` - Servers which are operating in standalone mode i.e. not connected to a Fabric Interconnected.* `FIAttached` - Servers which are connected to a Fabric Interconnect that is managed by Intersight.* `UnifiedEdgeServer` - Unified Edge sleds that is managed by Intersight. 
 
