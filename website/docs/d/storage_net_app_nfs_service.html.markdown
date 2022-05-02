---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_net_app_nfs_service"
description: |-
        An NFS service retrieves the NFS configuration of a storage virtual machine.

---

# Data Source: intersight_storage_net_app_nfs_service
An NFS service retrieves the NFS configuration of a storage virtual machine.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_net_app_nfs_service.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `nfs_v3_enabled`:(bool) Specifies whether NFSv3 protocol is enabled. 
* `nfs_v41_enabled`:(bool) Specifies whether NFSv4.1 protocol is enabled. 
* `nfs_v4_enabled`:(bool) Specifies whether NFSv4.0 protocol is enabled. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `svm_uuid`:(string) Unique identifier for the NetApp Storage Virtual Machine. 
 
