---
subcategory: "apic"
layout: "intersight"
page_title: "Intersight: intersight_apic_vrfs"
description: |-
        Vrfs object has been deprecated. The Vrfs object is a crucial element in the Cisco APIC framework, focusing on the management and configuration of Virtual Routing and Forwarding (VRF) instances.
        #### Purpose
        VRFs enable the creation of multiple virtual networks within a single physical network infrastructure. The Vrfs object provides the means to configure and manage these instances, ensuring efficient network segmentation and resource utilization.
        #### Key Concepts
        - **Network Segmentation:** Supports the division of network traffic into isolated paths, enhancing security and performance.
        - **Configuration Management:** Provides tools for setting up and maintaining VRF instances, including routing policies and network resources.
        - **Policy Enforcement:** Ensures that routing policies are applied consistently across the network segments.

---

# Data Source: intersight_apic_vrfs
Vrfs object has been deprecated. The Vrfs object is a crucial element in the Cisco APIC framework, focusing on the management and configuration of Virtual Routing and Forwarding (VRF) instances.
#### Purpose
VRFs enable the creation of multiple virtual networks within a single physical network infrastructure. The Vrfs object provides the means to configure and manage these instances, ensuring efficient network segmentation and resource utilization.
#### Key Concepts
- **Network Segmentation:** Supports the division of network traffic into isolated paths, enhancing security and performance.
- **Configuration Management:** Provides tools for setting up and maintaining VRF instances, including routing policies and network resources.
- **Policy Enforcement:** Ensures that routing policies are applied consistently across the network segments.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_apic_vrfs.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `dn`:(string) Distinguished name generated from URL parameters. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) VRF name generated from URL parameters. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
