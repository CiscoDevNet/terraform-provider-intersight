---
subcategory: "resource"
layout: "intersight"
page_title: "Intersight: intersight_resource_membership_holder"
description: |-
        The MembershipHolder object acts as a container for REST resources and their membership within the system. It plays a crucial role in managing and organizing resources by associating them with specific memberships.
        #### Purpose
        A MembershipHolder serves as the central entity for resource membership management, enabling efficient organization and retrieval of resources based on their membership status. This provides a structured approach to handling resource memberships, supporting multi-tenancy and resource classification.
        #### Key Concepts
        - **Resource Organization:** Facilitates the grouping and classification of resources according to membership criteria.
        - **Access Control:** Ensures that membership data is managed securely, with access restricted to authorized users.
        - **Relationship Management:** Maintains associations between resources and their respective accounts, ensuring consistent and organized membership handling.

---

# Data Source: intersight_resource_membership_holder
The MembershipHolder object acts as a container for REST resources and their membership within the system. It plays a crucial role in managing and organizing resources by associating them with specific memberships.
#### Purpose
A MembershipHolder serves as the central entity for resource membership management, enabling efficient organization and retrieval of resources based on their membership status. This provides a structured approach to handling resource memberships, supporting multi-tenancy and resource classification.
#### Key Concepts
- **Resource Organization:** Facilitates the grouping and classification of resources according to membership criteria.
- **Access Control:** Ensures that membership data is managed securely, with access restricted to authorized users.
- **Relationship Management:** Maintains associations between resources and their respective accounts, ensuring consistent and organized membership handling.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_resource_membership_holder.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of this resource membership holder. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
