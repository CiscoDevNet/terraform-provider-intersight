---
subcategory: "recovery"
layout: "intersight"
page_title: "Intersight: intersight_recovery_on_demand_backup"
description: |-
  Handles requests for on demand backup for a given endpoint.
---

# Data Source: intersight_recovery_on_demand_backup
Handles requests for on demand backup for a given endpoint.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_recovery_on_demand_backup.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string) Description of the policy. 
* `file_name_prefix`:(string) The file name for the backup image. This name is added as a prefix in the name for the backup image. A unique file name for the backup image is created along with a timestamp. For example: prefix-1572431305418. 
* `is_password_set`:(bool) Indicates whether the value of the 'password' property has been set. 
* `location_type`:(string) Specifies whether the backup will be stored locally or remotely.* `Network Share` - The backup is stored remotely on a separate server.* `Local Storage` - The backup is stored locally on the endpoint. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `password`:(string) Password of Backup server. 
* `path`:(string) The file system path where the backup images must be stored. Include the IP address/hostname of the network share location and the complete file system path. For example: 172.29.109.234/var/backups/. 
* `protocol`:(string) Protocol for transferring the backup image to the network share location.* `SCP` - Secure Copy Protocol (SCP) to access the file server.* `SFTP` - SSH File Transfer Protocol (SFTP) to access file server.* `FTP` - File Transfer Protocol (FTP) to access file server. 
* `retention_count`:(int) Number of backup copies maintained on the local or remote server. When the created backup files exceed this number, the initial backup files are overwritten in a sequential manner. 
* `user_name`:(string) Username for the backup server. 
 