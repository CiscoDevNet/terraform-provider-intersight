---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_pure_directory_quota"
description: |-
        The PureDirectoryQuota object represents quota policies within a PureStorage FlashArray, ensuring organized and efficient storage operations.
        #### Purpose
        PureDirectoryQuota facilitates the management of quota policies, providing structured control and allocation of storage resources within the FlashArray.
        #### Key Concepts
        - **Quota Management:** Offers structured management of quotas, optimizing storage operations and resource allocation.
        - **Integration:** Supports integration with directory objects, enhancing storage management capabilities.
        - **Secure Access:** Utilizes privilege sets to secure and manage quota policies within the storage network.

---

# Data Source: intersight_storage_pure_directory_quota
The PureDirectoryQuota object represents quota policies within a PureStorage FlashArray, ensuring organized and efficient storage operations.  
#### Purpose  
PureDirectoryQuota facilitates the management of quota policies, providing structured control and allocation of storage resources within the FlashArray. 
#### Key Concepts 
- **Quota Management:** Offers structured management of quotas, optimizing storage operations and resource allocation. 
- **Integration:** Supports integration with directory objects, enhancing storage management capabilities. 
- **Secure Access:** Utilizes privilege sets to secure and manage quota policies within the storage network.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_pure_directory_quota.<custom_name>.results[i].<propertyname>`.
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
* `policy_name`:(string) The export policy that manages this export. An export can be managed by at most one export policy. 
* `policy_resource_type`:(string) The export policy that manages this export. An export can be managed by at most one export policy. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
