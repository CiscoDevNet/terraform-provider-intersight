---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_pure_file_systems"
description: |-
        The PureFileSystems object represents a filesystem within a PureStorage FlashArray, serving as a fundamental unit for organizing storage resources.
        #### Purpose
        PureFileSystems provides a mechanism for managing and organizing storage resources, ensuring efficient and secure data operations within the FlashArray.
        #### Key Concepts
        - **Filesystem Management:** It offers structured management of filesystems, optimizing storage operations and resource allocation.
        - **Directory Integration:** Supports integration with directories and export policies, enhancing storage management capabilities.
        - **Secure Access:** Utilizes privilege sets to secure and manage filesystem operations within the storage network.

---

# Data Source: intersight_storage_pure_file_systems
The PureFileSystems object represents a filesystem within a PureStorage FlashArray, serving as a fundamental unit for organizing storage resources.  
#### Purpose  
PureFileSystems provides a mechanism for managing and organizing storage resources, ensuring efficient and secure data operations within the FlashArray. 
#### Key Concepts  
- **Filesystem Management:** It offers structured management of filesystems, optimizing storage operations and resource allocation. 
- **Directory Integration:** Supports integration with directories and export policies, enhancing storage management capabilities. 
- **Secure Access:** Utilizes privilege sets to secure and manage filesystem operations within the storage network.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_pure_file_systems.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `created_time`:(int) The file system creation time, measured in milliseconds since the UNIX epoch. 
* `destroyed`:(bool) A boolean value, if set to true, lists only destroyed objects that are in the eradication pending state. If set to false, lists only objects that are not destroyed. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) A globally unique, system-generated ID. The ID cannot be modified and cannot refer to another resource. 
* `pod`:(string) The reference to the pod the file systems belongs to, or null if it is not in a pod. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
