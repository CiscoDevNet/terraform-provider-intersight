---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_pure_flash_blade_file_system"
description: |-
        The PureFlashBladeFileSystem object represents a file system entity within the PureStorage FlashBlade array. It is a foundational model for organizing and managing file-based storage resources.
        #### Purpose
        PureFlashBladeFileSystem is designed to facilitate the creation, monitoring, and management of file systems, supporting quota limits, replication status, and utilization tracking within the FlashBlade environment.
        #### Key Concepts
        - **File System Management**: Enables structured organization and oversight of file-based storage, including quotas and writable status.
        - **Replication Support**: Integrates promotion status for replicated file systems, supporting disaster recovery and high availability.
        - **Utilization Metrics**: Tracks space usage for resource planning and optimization.
        - **Relationship Mapping**: Connects file systems to their storage arrays and device registrations for seamless integration.

---

# Data Source: intersight_storage_pure_flash_blade_file_system
The PureFlashBladeFileSystem object represents a file system entity within the PureStorage FlashBlade array. It is a foundational model for organizing and managing file-based storage resources.
#### Purpose
PureFlashBladeFileSystem is designed to facilitate the creation, monitoring, and management of file systems, supporting quota limits, replication status, and utilization tracking within the FlashBlade environment.
#### Key Concepts
- **File System Management**: Enables structured organization and oversight of file-based storage, including quotas and writable status.
- **Replication Support**: Integrates promotion status for replicated file systems, supporting disaster recovery and high availability.
- **Utilization Metrics**: Tracks space usage for resource planning and optimization.
- **Relationship Mapping**: Connects file systems to their storage arrays and device registrations for seamless integration.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_pure_flash_blade_file_system.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `created_time`:(string) Creation timestamp of the file system. 
* `destroyed`:(bool) Returns a value of true if the file system has been destroyed and is pending eradication. The file system cannot be modified while it is in the destroyed state. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `file_system_id`:(string) A non-modifiable, globally unique ID chosen by the system. 
* `hard_limit_enabled`:(bool) If set to true, the file system's size, as defined by provisioned, is used as a hard limit quota. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the object (e.g., a file system or snapshot). 
* `promotion_status`:(string) The current status of the file system with respect to replication. Possible values are promoted and demoted. 
* `provisioned`:(int) The provisioned size of the file system, displayed in bytes. If set to an empty string, the file system is unlimited in size. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `writable`:(bool) Whether the file system is writable or not. If false, this overrides any protocol or file permission settings and prevents changes. 
 
