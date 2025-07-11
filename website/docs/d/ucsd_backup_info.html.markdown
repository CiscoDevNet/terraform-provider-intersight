---
subcategory: "ucsd"
layout: "intersight"
page_title: "Intersight: intersight_ucsd_backup_info"
description: |-
        Available backup images for target device restoration.

---

# Data Source: intersight_ucsd_backup_info
Available backup images for target device restoration.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_ucsd_backup_info.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `backup_file_name`:(string) Auto generated backup file name with user defined prefix and timestamp. 
* `backup_location`:(string) Backup location for restore operation. 
* `backup_server_ip`:(string) Backup server for storing backup images. 
* `backup_size`:(int) Backup image size in bytes. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `duration`:(int) Time taken to complete the backup. 
* `encryption_key`:(string) Encryption key for backup file. 
* `failure_reason`:(string) The cause of the backup failure. 
* `is_purged`:(bool) Backup image purge status based on retention policy. 
* `last_modified`:(string) Backup record last modified time. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `percentage_completion`:(int) Backup completion percentage status. 
* `product_version`:(string) End device product version at the backup time. 
* `protocol`:(string) Supported remote backup protocol (FTP, SCP and SFTP); not applicable for localhost (127.0.0.1). 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `stage_completion`:(string) Backup status stage information. 
* `start_time`:(string) Backup initiation start time. 
* `status`:(string) The current backup status. 
 
