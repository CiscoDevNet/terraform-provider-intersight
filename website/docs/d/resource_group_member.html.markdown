---
subcategory: "resource"
layout: "intersight"
page_title: "Intersight: intersight_resource_group_member"
description: |-
        The GroupMember object identifies resolved members within a resource group, providing clarity and organization to resource membership management.
        #### Purpose
        A GroupMember serves as the link between individual resources and their associated groups, ensuring that resources are accurately represented within the system. It simplifies membership management by providing a clear representation of resource-group associations.
        #### Key Concepts
        - **Membership Resolution:** Ensures that resource memberships are accurately resolved and maintained within the system.
        - **Resource References:** Facilitates easy access and management of resources within groups, supporting efficient resource handling.

---

# Data Source: intersight_resource_group_member
The GroupMember object identifies resolved members within a resource group, providing clarity and organization to resource membership management.
#### Purpose
A GroupMember serves as the link between individual resources and their associated groups, ensuring that resources are accurately represented within the system. It simplifies membership management by providing a clear representation of resource-group associations.
#### Key Concepts
- **Membership Resolution:** Ensures that resource memberships are accurately resolved and maintained within the system.
- **Resource References:** Facilitates easy access and management of resources within groups, supporting efficient resource handling.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_resource_group_member.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
