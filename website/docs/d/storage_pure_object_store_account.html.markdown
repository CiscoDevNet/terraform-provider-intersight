---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_pure_object_store_account"
description: |-
        The PureObjectStoreAccount object models an account within the PureStorage FlashBlade object store. It is fundamental to the organization and management of buckets and users in object store environments.
        #### Purpose
        PureObjectStoreAccount provides a centralized abstraction for managing object store accounts, enabling quota control, utilization tracking, and relationship mapping to buckets and users.
        #### Key Concepts
        - **Account Management**: Offers robust capabilities for account creation, tracking, and quota enforcement.
        - **Utilization Monitoring**: Integrates storage utilization metrics for optimal resource planning.
        - **Relationship Mapping**: Links accounts to their buckets, users, and associated storage arrays.
        - **Access Control**: Privilege sets and licensing ensure secure and compliant account management.

---

# Data Source: intersight_storage_pure_object_store_account
The PureObjectStoreAccount object models an account within the PureStorage FlashBlade object store. It is fundamental to the organization and management of buckets and users in object store environments.
#### Purpose
PureObjectStoreAccount provides a centralized abstraction for managing object store accounts, enabling quota control, utilization tracking, and relationship mapping to buckets and users.
#### Key Concepts
- **Account Management**: Offers robust capabilities for account creation, tracking, and quota enforcement.
- **Utilization Monitoring**: Integrates storage utilization metrics for optimal resource planning.
- **Relationship Mapping**: Links accounts to their buckets, users, and associated storage arrays.
- **Access Control**: Privilege sets and licensing ensure secure and compliant account management.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_pure_object_store_account.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `created`:(string) Creation timestamp of the account. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `hard_limit_enabled`:(bool) If true, the account size defined by quotaLimit is used as a hard limit quota. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the object store account. 
* `object_count`:(int) The count of objects within the account. 
* `quota_limit`:(int) The effective quota limit applied against the size of the account, displayed in bytes. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `uuid`:(string) A non-modifiable, globally unique ID chosen by the system. 
 
