---
subcategory: "iam"
layout: "intersight"
page_title: "Intersight: intersight_iam_resource_roles"
description: |-
        ResourceRoles maps roles to a specific resource (e.g., an organization) within a permission, shaping the effective access model.
        #### Purpose
        It provides fine-grained control over which roles apply to which resources for a given permission.
        #### Key Concepts
        - **Resource-scoped Roles:** Tailors access at the resource boundary.
        - **Role Composition:** Works with roles, privilege sets, and endpoint roles.
        - **Permission-centric:** Operates within a permission’s assignment context.
        - **Flexible Governance:** Enables nuanced, resource-level authorization design.

---

# Data Source: intersight_iam_resource_roles
ResourceRoles maps roles to a specific resource (e.g., an organization) within a permission, shaping the effective access model.
#### Purpose
It provides fine-grained control over which roles apply to which resources for a given permission.
#### Key Concepts
- **Resource-scoped Roles:** Tailors access at the resource boundary.
- **Role Composition:** Works with roles, privilege sets, and endpoint roles.
- **Permission-centric:** Operates within a permission’s assignment context.
- **Flexible Governance:** Enables nuanced, resource-level authorization design.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iam_resource_roles.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
