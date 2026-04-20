---
subcategory: "apic"
layout: "intersight"
page_title: "Intersight: intersight_apic_bridge_domain"
description: |-
        BridgeDomain object has been deprecated. The BridgeDomain object is a fundamental component within the Cisco APIC framework, focusing on Layer 2 network management.
        #### Purpose
        Bridge domains define the Layer 2 boundaries within the network, facilitating efficient traffic management and segmentation. They serve as building blocks for network isolation and policy application.
        #### Key Concepts
        - **Layer 2 Segmentation:** Provides mechanisms for dividing network traffic into isolated segments for better control and security.
        - **Traffic Management:** Supports efficient handling and routing of Layer 2 traffic within the network domain.
        - **Policy Application:** Ensures consistent application of network policies within the bridge domain.

---

# Data Source: intersight_apic_bridge_domain
BridgeDomain object has been deprecated. The BridgeDomain object is a fundamental component within the Cisco APIC framework, focusing on Layer 2 network management.
#### Purpose
Bridge domains define the Layer 2 boundaries within the network, facilitating efficient traffic management and segmentation. They serve as building blocks for network isolation and policy application.
#### Key Concepts
- **Layer 2 Segmentation:** Provides mechanisms for dividing network traffic into isolated segments for better control and security.
- **Traffic Management:** Supports efficient handling and routing of Layer 2 traffic within the network domain.
- **Policy Application:** Ensures consistent application of network policies within the bridge domain.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_apic_bridge_domain.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `dn`:(string) Distinguished Name (DN) of an object in Cisco Application Policy Infrastructure Controller (APIC) GUI. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Object name in Cisco Application Policy Infrastructure Controller (APIC) GUI. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
