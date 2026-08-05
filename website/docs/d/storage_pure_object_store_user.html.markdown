---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_pure_object_store_user"
description: |-
        The PureObjectStoreUser object defines a user entity in the PureStorage FlashBlade object store context. It serves as a central point for managing access and user identities within object store accounts.
        #### Purpose
        PureObjectStoreUser is designed to facilitate secure user management in object store environments, allowing administrators to track user creation, association, and access within storage accounts.
        #### Key Concepts
        - **User Identity Management**: Provides a standardized model for user creation and association with object store accounts.
        - **Access Integration**: Links users to their accounts and arrays, supporting secure, multi-tenant environments.
        - **Privilege Enforcement**: Ensures only authorized personnel can view or modify user details.
        - **Audit Support**: Supports audit trails for user actions and associations within object stores.

---

# Data Source: intersight_storage_pure_object_store_user
The PureObjectStoreUser object defines a user entity in the PureStorage FlashBlade object store context. It serves as a central point for managing access and user identities within object store accounts.
#### Purpose
PureObjectStoreUser is designed to facilitate secure user management in object store environments, allowing administrators to track user creation, association, and access within storage accounts.
#### Key Concepts
- **User Identity Management**: Provides a standardized model for user creation and association with object store accounts.
- **Access Integration**: Links users to their accounts and arrays, supporting secure, multi-tenant environments.
- **Privilege Enforcement**: Ensures only authorized personnel can view or modify user details.
- **Audit Support**: Supports audit trails for user actions and associations within object stores.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_pure_object_store_user.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `created`:(string) Creation time of the user. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the object store user. 
* `object_store_account_name`:(string) Name of the object store account that owns this user. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `uuid`:(string) A globally unique identifier for the object store user. 
 
