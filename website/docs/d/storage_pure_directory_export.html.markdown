---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_pure_directory_export"
description: |-
        The PureDirectoryExport object facilitates the creation of managed directory exports within a PureStorage FlashArray, enhancing data accessibility.
        #### Purpose
        PureDirectoryExport provides a mechanism for exporting directories, enabling efficient data access and integration with supported protocols.
        #### Key Concepts
        - **Export Management:** Facilitates the creation and management of directory exports, optimizing data accessibility and integration.
        - **Protocol Support:** Supports integration with NFS and SMB protocols, enhancing storage management capabilities.
        - **Secure Access:** Utilizes privilege sets to secure and manage directory exports within the storage network.

---

# Data Source: intersight_storage_pure_directory_export
The PureDirectoryExport object facilitates the creation of managed directory exports within a PureStorage FlashArray, enhancing data accessibility.  
#### Purpose  
PureDirectoryExport provides a mechanism for exporting directories, enabling efficient data access and integration with supported protocols.
#### Key Concepts  
- **Export Management:** Facilitates the creation and management of directory exports, optimizing data accessibility and integration. 
- **Protocol Support:** Supports integration with NFS and SMB protocols, enhancing storage management capabilities. 
- **Secure Access:** Utilizes privilege sets to secure and manage directory exports within the storage network.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_pure_directory_export.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `destroyed`:(bool) Returns a value of true if the managed directory of the export has been destroyed and is pending eradication. The export can be recovered by recovering the destroyed managed directory. 
* `directory_name`:(string) Absolute path of the managed directory in the file system. 
* `directory_resource_type`:(string) Absolute path of the managed directory in the file system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `enabled`:(bool) Returns a value of true if the export policy that manages this export is enabled. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The export name for accessing this export. 
* `path`:(string) The path of the exported managed directory. 
* `policy_name`:(string) The export policy that manages this export. An export can be managed by at most one export policy. 
* `policy_resource_type`:(string) The export policy that manages this export. An export can be managed by at most one export policy. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
