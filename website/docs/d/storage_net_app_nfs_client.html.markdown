---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_net_app_nfs_client"
description: |-
        A currently connected NFS client.

---

# Data Source: intersight_storage_net_app_nfs_client
A currently connected NFS client.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_net_app_nfs_client.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `client_ip`:(string) Specifies IP address of the client. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `idle_duration`:(string) Specifies an ISO-8601 format of date and time to retrieve the idle time duration in hours, minutes, and seconds format. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `protocol`:(string) The NFS protocol version over which client is accessing the volume. 
* `server_ip`:(string) Specifies IP address of the server. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `svm_name`:(string) The storage virtual machine name for the NFS client. 
* `svm_uuid`:(string) Unique identifier for the NetApp Storage Virtual Machine. 
* `volume_name`:(string) The parent volume name for the NFS client. 
* `volume_uuid`:(string) Unique identifier for the NetApp Volume. 
 
