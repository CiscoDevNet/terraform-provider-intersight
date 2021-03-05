---
subcategory: "ucsd"
layout: "intersight"
page_title: "Intersight: intersight_ucsd_backup_info"
description: |-
  List of backup images available for target end device for restore operation.
---

# Data Source: intersight_ucsd_backup_info
List of backup images available for target end device for restore operation.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_ucsd_backup_info.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `backup_file_name`:(string) Auto generated backup File Name with combination of file prefix given an user input and the timestamp. 
* `backup_location`:(string) Backup location that contains the backup images for end device which can be used for restore operation. 
* `backup_server_ip`:(string) Backup server where backup images are maintained. 
* `backup_size`:(int) Size of the backup image in bytes. 
* `duration`:(int) Time taken for the backup to be completed. 
* `encryption_key`:(string) The key used for encrypting the backup file. 
* `failure_reason`:(string) Reason for backup failure. 
* `is_purged`:(bool) Backup image got purged or not. The backup images get purged based on the retention count set by the user in the backup config policy. 
* `last_modified`:(string) Last modified time when this backup record got updated. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `percentage_completion`:(int) Backup current precentage completion status information. 
* `product_version`:(string) The end device product version when the backup image was taken. 
* `protocol`:(string) Protocol used for the remote backup. possible values are FTP, SCP and SFTP. Not applicable for the localhost (127.0.0.1). 
* `stage_completion`:(string) Backup current status stage information. 
* `start_time`:(string) Start time of backup when it got initiated. 
* `status`:(string) Current status of Backup current. 
 