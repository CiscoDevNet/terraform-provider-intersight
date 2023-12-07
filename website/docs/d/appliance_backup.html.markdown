---
subcategory: "appliance"
layout: "intersight"
page_title: "Intersight: intersight_appliance_backup"
description: |-
        Backup tracks all backup requests to create a full system backup of the Intersight
        Appliance. There will be only one Backup managed object with a 'Started' state at
        any time. All other Backup managed objects will be in terminal states.

---

# Data Source: intersight_appliance_backup
Backup tracks all backup requests to create a full system backup of the Intersight
Appliance. There will be only one Backup managed object with a 'Started' state at
any time. All other Backup managed objects will be in terminal states.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_appliance_backup.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `elapsed_time`:(int) Elapsed time in seconds since the backup process has started. 
* `end_time`:(string) End date and time of the backup process. 
* `filename`:(string) Backup filename to backup or restore. 
* `is_manual`:(bool) If true, represents a manual backup. Else represents a scheduled backup. 
* `is_password_set`:(bool) Indicates whether the value of the 'password' property has been set. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `password`:(string) Password to authenticate the fileserver. 
* `protocol`:(string) Communication protocol used by the file server (e.g. scp, sftp, or CIFS).* `scp` - Secure Copy Protocol (SCP) to access the file server.* `sftp` - SSH File Transfer Protocol (SFTP) to access file server.* `cifs` - Common Internet File System (CIFS) Protocol to access file server. 
* `remote_host`:(string) Hostname of the remote file server. 
* `remote_path`:(string) File server directory or share name to copy the file. 
* `remote_port`:(int) Remote TCP port on the file server (e.g. 22 for scp). 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `start_time`:(string) Start date and time of the backup process. 
* `status`:(string) Status of the backup managed object.* `Started` - Backup or restore process has started.* `Created` - Backup or restore is in created state.* `Failed` - Backup or restore process has failed.* `Completed` - Backup or restore process has completed.* `Copied` - Backup file has been copied.* `Cleanup Failed` - Cleanup of the old backup has failed. 
* `username`:(string) Username to authenticate the fileserver. 
 
