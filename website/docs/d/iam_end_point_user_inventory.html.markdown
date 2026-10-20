---
subcategory: "iam"
layout: "intersight"
page_title: "Intersight: intersight_iam_end_point_user_inventory"
description: |-
        The IAM EndPointUser object represents a single local user account. It is a building block used within the EndPointUserPolicy to define the users that should be created on an endpoint.
        #### Purpose
        The purpose is to define the identity of a local user, primarily their username. It acts as a referenceable entity that can be associated with roles and a password within the broader EndPointUserPolicy.
        #### Key Concepts
        - **User Identity:** The core function of this is to specify the name (username) of the local user.
        - **Component of a Policy:** It does not function on its own but is created as part of an EndPointUserPolicy.
        - **Reusable within an Organization:** As an organization-owned object, a single EndPointUser can be referenced in multiple EndPointUserRole mappings across different policies, though it typically represents a unique user identity.

---

# Data Source: intersight_iam_end_point_user_inventory
The IAM EndPointUser object represents a single local user account. It is a building block used within the EndPointUserPolicy to define the users that should be created on an endpoint.
#### Purpose
The purpose is to define the identity of a local user, primarily their username. It acts as a referenceable entity that can be associated with roles and a password within the broader EndPointUserPolicy.
#### Key Concepts
- **User Identity:** The core function of this is to specify the name (username) of the local user.
- **Component of a Policy:** It does not function on its own but is created as part of an EndPointUserPolicy.
- **Reusable within an Organization:** As an organization-owned object, a single EndPointUser can be referenced in multiple EndPointUserRole mappings across different policies, though it typically represents a unique user identity.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iam_end_point_user_inventory.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) Device ID of the entity from where inventory is reported. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the user to be created on the endpoint. It can be any string that adheres to the following constraints. It can have alphanumeric characters, dots, underscores and hyphen. It cannot be more than 16 characters. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `user_id`:(string) UserId for the end point user. 
 
