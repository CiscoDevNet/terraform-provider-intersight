---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_pure_directory"
description: |-
        The PureDirectory object represents a managed directory within a PureStorage FlashArray, offering space reporting and metrics.
        #### Purpose
        PureDirectory facilitates the organization and management of storage directories, providing structured space reporting and policy integration within the FlashArray.
        #### Key Concepts
        - **Directory Management:** Offers structured management of directories, optimizing space reporting and storage operations.
        - **Policy Integration:** Supports attachment of policies to directories, enhancing storage management capabilities.
        - **Secure Access:** Utilizes privilege sets to secure and manage directory operations within the storage network.

---

# Data Source: intersight_storage_pure_directory
The PureDirectory object represents a managed directory within a PureStorage FlashArray, offering space reporting and metrics.
#### Purpose  
PureDirectory facilitates the organization and management of storage directories, providing structured space reporting and policy integration within the FlashArray.  
#### Key Concepts  
- **Directory Management:** Offers structured management of directories, optimizing space reporting and storage operations. 
- **Policy Integration:** Supports attachment of policies to directories, enhancing storage management capabilities. 
- **Secure Access:** Utilizes privilege sets to secure and manage directory operations within the storage network.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_pure_directory.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `created_time`:(string) The managed directory creation time, measured in milliseconds since the UNIX epoch. 
* `data_reduction`:(int) The ratio of mapped sectors within a volume versus the amount of physical space the data occupies after data compression and deduplication. The data reduction ratio does not include thin provisioning savings. For example, a data reduction ratio of 5:1 means that for every 5 MB the host writes to the array, 1 MB is stored on the array's flash modules. 
* `destroyed`:(bool) A boolean value, if set to true, lists only destroyed objects that are in the eradication pending state. If set to false, lists only objects that are not destroyed. 
* `directory_name`:(string) The managed directory name without the file system name prefix. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `file_system_name`:(string) Name of file syetem associated with the directory. 
* `member_name`:(string) Absolute path of the managed directory in the file system. 
* `member_resource_type`:(string) Absolute path of the managed directory in the file system. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) A user-specified name. The name must be locally unique and can be changed. 
* `path`:(string) Absolute path of the managed directory in the file system. 
* `policy_resource_type`:(string) Absolute path of the managed directory in the file system. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `snapshots`:(int) The physical space occupied by data unique to one or more snapshots. Measured in bytes. 
* `thin_provisioning`:(int) The percentage of volume sectors that do not contain host-written data because the hosts have not written data to them or the sectors have been explicitly trimmed. 
* `total_physical`:(int) The total physical space occupied by system, shared space, volume, and snapshot data. Measured in bytes. 
* `total_reduction`:(int) The ratio of provisioned sectors within a volume versus the amount of physical space the data occupies after reduction via data compression and deduplication and with thin provisioning savings. Total reduction is data reduction with thin provisioning savings. For example, a total reduction ratio of 10:1 means that for every 10 MB of provisioned space, 1 MB is stored on the array's flash modules. 
* `total_used`:(int) The total space contributed by customer data, measured in bytes. 
* `unique`:(int) The unique physical space occupied by customer data. Unique physical space does not include shared space, snapshots, and internal array metadata. Measured in bytes. On Evergreen//One arrays, this is the effective space contributed by unique customer data, measured in bytes. Unique data does not include shared space, snapshots, and internal array metadata. 
 
