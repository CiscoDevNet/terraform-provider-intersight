---
subcategory: "apic"
layout: "intersight"
page_title: "Intersight: intersight_apic_vpc_group"
description: |-
        VpcGroup object has been deprecated. The VpcGroup object in Cisco APIC is designed to manage Virtual Port Channel (VPC) groups, focusing on link aggregation and redundancy.
        #### Purpose
        VPCGroups facilitate the aggregation of multiple network links into a single logical channel, enhancing redundancy and bandwidth utilization.
        #### Key Concepts
        - **Link Aggregation:** Supports the combination of multiple physical links into a single logical channel for improved performance.
        - **Redundancy:** Provides mechanisms for ensuring network link redundancy and reliability.
        - **Resource Management:** Facilitates the management of link resources within the VPC group.

---

# Data Source: intersight_apic_vpc_group
VpcGroup object has been deprecated. The VpcGroup object in Cisco APIC is designed to manage Virtual Port Channel (VPC) groups, focusing on link aggregation and redundancy.
#### Purpose
VPCGroups facilitate the aggregation of multiple network links into a single logical channel, enhancing redundancy and bandwidth utilization.
#### Key Concepts
- **Link Aggregation:** Supports the combination of multiple physical links into a single logical channel for improved performance.
- **Redundancy:** Provides mechanisms for ensuring network link redundancy and reliability.
- **Resource Management:** Facilitates the management of link resources within the VPC group.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_apic_vpc_group.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `dn`:(string) Distinguished Name (DN) of an object in Cisco Application Policy Infrastructure Controller (APIC) GUI. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Object name in Cisco Application Policy Infrastructure Controller (APIC) GUI. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
