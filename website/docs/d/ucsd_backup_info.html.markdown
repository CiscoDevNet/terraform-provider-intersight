
---
layout: "intersight"
page_title: "Intersight: intersight_ucsd_backup_info"
sidebar_current: "docs-intersight-data-source-ucsd-backup-info"
description: |-
List of backup images available for target end device for restore operation.
---

# Data Source: intersight_ucsd._backup_info
List of backup images available for target end device for restore operation.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `backup_file_name`:(string) Auto generated backup File Name with combination of file prefix given an user input and the timestamp. 
* `backup_location`:(string) Backup location that contains the backup images for end device which can be used for restore operation. 
* `backup_server_ip`:(string) Backup server where backup images are maintained. 
* `backup_size`:(int) Size of the backup image in bytes. 
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `duration`:(int) Time taken for the backup to be completed. 
* `encryption_key`:(string) The key used for encrypting the backup file. 
* `failure_reason`:(string) Reason for backup failure. 
* `is_purged`:(bool) Backup image got purged or not. The backup images get purged based on the retention count set by the user in the backup config policy. 
* `last_modified`:(string) Last modified time when this backup record got updated. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `percentage_completion`:(int) Backup current precentage completion status information. 
* `product_version`:(string) The end device product version when the backup image was taken. 
* `protocol`:(string) Protocol used for the remote backup. possible values are FTP, SCP and SFTP. Not applicable for the localhost (127.0.0.1). 
* `stage_completion`:(string) Backup current status stage information. 
* `start_time`:(string) Start time of backup when it got initiated. 
* `status`:(string) Current status of Backup current. 
