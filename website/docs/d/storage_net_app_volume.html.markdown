---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_net_app_volume"
description: |-
        ### Overview
        The NetAppVolume object is a versatile data container that enables partitioning and management of storage resources within a NetAppCluster. It is instrumental in organizing data efficiently and securely, supporting a wide range of storage use cases.
        #### Purpose
        A NetAppVolume serves as a logical unit for data storage, providing isolation and control over data access and management. It allows administrators to define storage characteristics, implement data protection strategies, and optimize resource usage within the cluster.
        #### Key Concepts
        - **Data Isolation:** - Ensures that data stored within a volume is segregated from other volumes, enabling tailored access controls and management policies.
        - **Flexibility:** - Supports various volume types and configurations, catering to diverse application requirements and storage needs.
        - **Protection Mechanisms:** - Implements advanced data protection features, such as snapshots and replication, to safeguard data integrity and availability.
        - **Performance Management:** - Utilizes sophisticated algorithms to optimize data access speeds and responsiveness, enhancing overall system efficiency.

---

# Data Source: intersight_storage_net_app_volume
### Overview
The NetAppVolume object is a versatile data container that enables partitioning and management of storage resources within a NetAppCluster. It is instrumental in organizing data efficiently and securely, supporting a wide range of storage use cases.
#### Purpose
A NetAppVolume serves as a logical unit for data storage, providing isolation and control over data access and management. It allows administrators to define storage characteristics, implement data protection strategies, and optimize resource usage within the cluster.
#### Key Concepts
- **Data Isolation:** - Ensures that data stored within a volume is segregated from other volumes, enabling tailored access controls and management policies.
- **Flexibility:** - Supports various volume types and configurations, catering to diverse application requirements and storage needs.
- **Protection Mechanisms:** - Implements advanced data protection features, such as snapshots and replication, to safeguard data integrity and availability.
- **Performance Management:** - Utilizes sophisticated algorithms to optimize data access speeds and responsiveness, enhancing overall system efficiency.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_net_app_volume.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `autosize_mode`:(string) The autosize mode for NetApp Volume. Modes can be off or grow or grow_shrink.* `off` - The volume will not grow or shrink in size in response to the amount of used space.* `grow` - The volume will automatically grow when used space in the volume is above the grow threshold.* `grow_shrink` - The volume will grow or shrink in size in response to the amount of used space. 
* `create_time`:(string) The time when this managed object was created. 
* `created_time`:(string) Storage container's creation time. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `export_policy_name`:(string) The name of the Export Policy. 
* `flex_cache_endpoint_type`:(string) FlexCache endpoint type. The endpoint type can be the origin of a FlexCache volume, a FlexCache volume, or neither. 
* `is_object_store`:(bool) Specifies whether the volume is provisioned for an object store server. 
* `key`:(string) Unique identifier of a NetApp Volume across data center. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the storage container. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `snapshot_autodelete_enabled`:(bool) Specifies whether Snaphot copy autodelete is currently enabled on this volume. 
* `snapshot_policy_name`:(string) The name of the Snapshot Policy. 
* `snapshot_policy_uuid`:(string) The UUID of the Snapshot Policy. 
* `snapshot_reserve_percent`:(int) The space that has been set aside as a reserve for Snapshot copy usage represented as a percent. 
* `snapshot_used`:(float) The total space used by Snapshot copies in the volume represented in bytes. 
* `state`:(string) The current state of a NetApp volume.* `offline` - Read and write access to the volume is not allowed.* `online` - Read and write access to the volume is allowed.* `error` - Storage volume state of error type.* `mixed` - The constituents of a FlexGroup volume are not all in the same state. 
* `style`:(string) The style of the volume (FlexGroup or FlexVol). 
* `svm_name`:(string) The storage virtual machine name for the volume. 
* `type`:(string) NetApp volume type. The volume type can be Read-write, Data-protection, or Load-sharing.* `data-protection` - Prevents modification of the data on the Volume.* `read-write` - Data on the Volume can be modified.* `load-sharing` - The volume type is Load Sharing DP. 
* `uuid`:(string) Universally unique identifier of a NetApp Volume. 
 
