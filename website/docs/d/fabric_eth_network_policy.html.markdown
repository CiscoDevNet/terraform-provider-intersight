---
subcategory: "fabric"
layout: "intersight"
page_title: "Intersight: intersight_fabric_eth_network_policy"
description: |-
        The EthNetworkPolicy object defines the VLAN settings applied to system interfaces, ensuring consistent network segmentation and management. In simple terms, it specifies the set of VLANs assigned to those interfaces.
        #### Purpose
        EthNetworkPolicy is designed to manage VLAN configurations for interfaces, providing a structured approach to network segmentation and traffic management. It helps maintain organized and efficient network operations, reducing complexity and enhancing control.
        #### Key Concepts
        - **VLAN Management:** Enables the definition and application of VLAN settings, including native and allowed VLANs, ensuring consistent network segmentation.
        - **Policy-Based Approach:** Offers a centralized policy framework for managing VLAN configurations, streamlining network operations.
        - **Scalability:** Supports scalable network management by allowing efficient configuration of multiple interfaces, facilitating growth and adaptation.
        - **Integration:** Works seamlessly with other network policies and configurations, ensuring cohesive and integrated network management.

---

# Data Source: intersight_fabric_eth_network_policy
The EthNetworkPolicy object defines the VLAN settings applied to system interfaces, ensuring consistent network segmentation and management. In simple terms, it specifies the set of VLANs assigned to those interfaces.
#### Purpose
EthNetworkPolicy is designed to manage VLAN configurations for interfaces, providing a structured approach to network segmentation and traffic management. It helps maintain organized and efficient network operations, reducing complexity and enhancing control.
#### Key Concepts
- **VLAN Management:** Enables the definition and application of VLAN settings, including native and allowed VLANs, ensuring consistent network segmentation.
- **Policy-Based Approach:** Offers a centralized policy framework for managing VLAN configurations, streamlining network operations.
- **Scalability:** Supports scalable network management by allowing efficient configuration of multiple interfaces, facilitating growth and adaptation.
- **Integration:** Works seamlessly with other network policies and configurations, ensuring cohesive and integrated network management.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_fabric_eth_network_policy.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the policy. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `target_platform`:(string) The target platform type of the Ethernet Network policy.* `UCS Domain` - Profile/policy type for network and management configuration on UCS Fabric Interconnect.* `Unified Edge` - Profile/policy type for network, management and chassis configuration on Unified Edge. 
 
