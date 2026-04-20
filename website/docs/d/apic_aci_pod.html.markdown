---
subcategory: "apic"
layout: "intersight"
page_title: "Intersight: intersight_apic_aci_pod"
description: |-
        AciPod object has been deprecated. The AciPod object in Cisco APIC represents the network pods, focusing on the logical grouping of infrastructure components.
        #### Purpose
        ACIPods provide a structured way to manage and organize the network's physical infrastructure, facilitating efficient resource allocation and policy application.
        #### Key Concepts
        - **Logical Grouping:** Organizes network infrastructure components into logical units for streamlined management.
        - **Resource Allocation:** Supports efficient allocation and management of network resources within pods.
        - **Policy Application:** Ensures consistent application of network policies across infrastructure components.

---

# Data Source: intersight_apic_aci_pod
AciPod object has been deprecated. The AciPod object in Cisco APIC represents the network pods, focusing on the logical grouping of infrastructure components.
#### Purpose
ACIPods provide a structured way to manage and organize the network's physical infrastructure, facilitating efficient resource allocation and policy application.
#### Key Concepts
- **Logical Grouping:** Organizes network infrastructure components into logical units for streamlined management.
- **Resource Allocation:** Supports efficient allocation and management of network resources within pods.
- **Policy Application:** Ensures consistent application of network policies across infrastructure components.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_apic_aci_pod.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Pod name extracted from DN in Cisco Application Policy Infrastructure Controller (APIC). 
* `pod_type`:(string) Object pod type in Cisco Application Policy Infrastructure Controller (APIC). 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
