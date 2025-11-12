---
subcategory: "coremanagement"
layout: "intersight"
page_title: "Intersight: intersight_coremanagement_core_file"
description: |-
        Core file meta data for individual core files on device.

---

# Data Source: intersight_coremanagement_core_file
Core file meta data for individual core files on device.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_coremanagement_core_file.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `admin_state`:(string) Admin state prop to trigger file upload.* `None` - Admin configured None State.* `Upload` - Admin configured Upload State. 
* `core_file_download_url`:(string) The Url to download the core file. It will be set only after successful completion of core file upload to storage service. 
* `create_time`:(string) The time when this managed object was created. 
* `device_type`:(string) The device object type for the end point. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `endpoint_identifier`:(string) Endpoint device identifier. In IMM devices, it will be Fabric Interconnect hostname. 
* `file_name`:(string) The name of core file from endpoint device. 
* `file_size`:(int) File size of core file in bytes. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `node_id`:(string) Node id within cluster where core file is present. 
* `pid`:(string) Product identification of the device. 
* `reason`:(string) Reason for upload failure, if any. In successful upload case, it will be empty. 
* `serial`:(string) Serial number of the device. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `status`:(string) Status of core file upload. Valid values are InventoryComplete (default), UploadInProgress, Completed, UploadFailed, FileDownloadUrlCreationFailed, CoreRemovedDownloadOnly.* `InventoryComplete` - Default status for all mos before file upload is requested.* `UploadInProgress` - File upload is in progress.* `UploadFailed` - File upload to storage service failed.* `Completed` - File upload to storage service completed successfully.* `FileDownloadUrlCreationFailed` - File upload to storage service completed successfully but download url creation failed.* `CoreRemovedDownloadOnly` - File upload to storage service completed successfully but file removed from endpoint device. 
 
