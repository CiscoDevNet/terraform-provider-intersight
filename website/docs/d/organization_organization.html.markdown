---
subcategory: "organization"
layout: "intersight"
page_title: "Intersight: intersight_organization_organization"
description: |-
        Organizations provide multi-tenancy within a single Intersight account. An account can contain multiple organizations, and resources are associated to organizations using resource groups. Organizations can contain heterogeneous resources, support controlled sharing of resources across organizations, and integrate with the IAM permission/privilege model so users can have different access levels per organization.
        #### Purpose
        Partition resources within an account into separate administrative and access-control domains, enabling delegation, separation of duties, and controlled sharing across teams or business units.
        #### Key Concepts
        - **Account-level multi-tenancy**: Multiple organizations can exist under one account, acting as logical tenants within that account.
        - **Resource association via resource groups**: Resources are linked to an organization through resource groups rather than being inherently single-tenant.
        - **Heterogeneous resource containment**: An organization can include different kinds of resources (compute, network, storage, policies, etc.).
        - **Cross-organization sharing**: Resources can be shared among multiple organizations while preserving governance (often via explicit sharing constructs/rules).
        - **RBAC integration**: Organizations are tied to permissions and privilege sets; a user can access multiple organizations within the same permission, potentially with different privileges per organization.
        - **Admin-governed lifecycle**: CREATE/UPDATE/DELETE require organization management privileges; READ is broadly available to admin, view, and several functional roles.

---

# Data Source: intersight_organization_organization
Organizations provide multi-tenancy within a single Intersight account. An account can contain multiple organizations, and resources are associated to organizations using resource groups. Organizations can contain heterogeneous resources, support controlled sharing of resources across organizations, and integrate with the IAM permission/privilege model so users can have different access levels per organization.
#### Purpose
Partition resources within an account into separate administrative and access-control domains, enabling delegation, separation of duties, and controlled sharing across teams or business units.
#### Key Concepts
- **Account-level multi-tenancy**: Multiple organizations can exist under one account, acting as logical tenants within that account.
- **Resource association via resource groups**: Resources are linked to an organization through resource groups rather than being inherently single-tenant.
- **Heterogeneous resource containment**: An organization can include different kinds of resources (compute, network, storage, policies, etc.).
- **Cross-organization sharing**: Resources can be shared among multiple organizations while preserving governance (often via explicit sharing constructs/rules).
- **RBAC integration**: Organizations are tied to permissions and privilege sets; a user can access multiple organizations within the same permission, potentially with different privileges per organization.
- **Admin-governed lifecycle**: CREATE/UPDATE/DELETE require organization management privileges; READ is broadly available to admin, view, and several functional roles.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_organization_organization.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) The informative description about the usage of this organization. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the organization. There can be multiple organizations under an account. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
