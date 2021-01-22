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
The following arguments can be used to get data of already created objects in Intersight appliance:
* `cifs_enabled`:(bool) Status for Common Internet File System protocol ( CIFS ) allowed to run on Vservers. 
* `fcp_enabled`:(bool) Status for Fibre Channel Protocol ( FCP ) allowed to run on Vservers. 
* `iscsi_enabled`:(bool) Status for iSCSI protocol allowed to run on Vservers. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the tenant in storage array. 
* `nfs_enabled`:(bool) Status for Network File System Protocol ( NFS ) allowed to run on  Vservers. 
* `nvme_enabled`:(bool) Status for NVME protocol allowed to run on Vservers. 
* `state`:(string) The state of this tenant.* `Unknown` - Component state is not available.* `Starting` - Component is being started.* `Running` - Component is currently running.* `Stopping` - Component is being stopped.* `Stopped` - Component has been stopped.* `Deleting` - Component deletion is in progress. 
* `uuid`:(string) The uuid of this tenant in storage array. 
