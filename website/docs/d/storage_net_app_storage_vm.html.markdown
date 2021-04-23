---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_net_app_storage_vm"
description: |-
  NetApp Storage Virtual Machines contain data volumes and one or more Logical Interfaces ( LIFs ) through which they serve data to the clients.
---

# Data Source: intersight_storage_net_app_storage_vm
NetApp Storage Virtual Machines contain data volumes and one or more Logical Interfaces ( LIFs ) through which they serve data to the clients.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_net_app_storage_vm.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `cifs_enabled`:(bool) Status for Common Internet File System protocol ( CIFS ) allowed to run on Vservers. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `fcp_enabled`:(bool) Status for Fibre Channel Protocol ( FCP ) allowed to run on Vservers. 
* `iscsi_enabled`:(bool) Status for iSCSI protocol allowed to run on Vservers. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the tenant in storage array. 
* `nfs_enabled`:(bool) Status for Network File System Protocol ( NFS ) allowed to run on  Vservers. 
* `nvme_enabled`:(bool) Status for NVME protocol allowed to run on Vservers. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `state`:(string) The state of this tenant.* `Unknown` - Component state is not available.* `Starting` - Component is being started.* `Running` - Component is currently running.* `Stopping` - Component is being stopped.* `Stopped` - Component has been stopped.* `Deleting` - Component deletion is in progress. 
* `uuid`:(string) The uuid of this tenant in storage array. 
 