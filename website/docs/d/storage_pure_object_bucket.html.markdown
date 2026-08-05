---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_pure_object_bucket"
description: |-
        The **PureObjectBucket** object represents a volume entity within PureStorage FlashBlade, specifically modeling object buckets used for storing and organizing data in object store environments.
        #### Purpose
        PureObjectBucket provides a structured abstraction for managing object buckets, enabling customers to efficiently handle storage, quotas, and data retention policies. It supports integration with object store accounts and arrays for comprehensive storage management.
        #### Key Concepts
        - **Object Store Management**: Supports the creation and organization of object buckets, facilitating scalable storage solutions.
        - **Quota and Retention**: Integrates quota limits and retention policies for robust data governance.
        - **Relationship Mapping**: Links buckets to their respective accounts and storage arrays, enhancing data traceability.
        - **Access Control**: Utilizes privilege sets to secure bucket operations and management.

---

# Data Source: intersight_storage_pure_object_bucket
The **PureObjectBucket** object represents a volume entity within PureStorage FlashBlade, specifically modeling object buckets used for storing and organizing data in object store environments.
#### Purpose
PureObjectBucket provides a structured abstraction for managing object buckets, enabling customers to efficiently handle storage, quotas, and data retention policies. It supports integration with object store accounts and arrays for comprehensive storage management.
#### Key Concepts
- **Object Store Management**: Supports the creation and organization of object buckets, facilitating scalable storage solutions.
- **Quota and Retention**: Integrates quota limits and retention policies for robust data governance.
- **Relationship Mapping**: Links buckets to their respective accounts and storage arrays, enhancing data traceability.
- **Access Control**: Utilizes privilege sets to secure bucket operations and management.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_pure_object_bucket.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `bucket_name`:(string) Name of the bucket within the object store account. 
* `bucket_type`:(string) Type of the bucket (e.g., multi-site-writable). 
* `create_time`:(string) The time when this managed object was created. 
* `created`:(string) Creation time of the bucket. 
* `destroyed`:(bool) Whether the bucket has been destroyed. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `hard_limit_enabled`:(bool) Whether the hard limit is enabled for the bucket. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Unique name of the bucket (combination of account name and bucket name). 
* `object_count`:(int) Number of objects in the bucket. 
* `object_store_account_name`:(string) Name of the object store account that owns this bucket. 
* `quota_limit`:(int) Quota limit for the bucket in bytes. 
* `retention_lock`:(string) Retention lock status for the bucket. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `uuid`:(string) A globally unique identifier for the bucket. 
* `vendor`:(string) Vendor of the storage array. 
* `nr_version`:(string) Version of the FlashBlade array. 
* `versioning`:(string) Versioning mode for the bucket. 
 
