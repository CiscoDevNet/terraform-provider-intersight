---
subcategory: "appliance"
layout: "intersight"
page_title: "Intersight: intersight_appliance_backup_policy"
description: |-
        BackupPolicy stores the Intersight Appliance's backup policy. There will be only
        one BackupPolicy managed object in the Intersight Appliance. Default backup policy
        managed object is created during the Intersight Appliance setup, and it is configured
        in the manual backup mode.

---

# Data Source: intersight_appliance_backup_policy
BackupPolicy stores the Intersight Appliance's backup policy. There will be only
one BackupPolicy managed object in the Intersight Appliance. Default backup policy
managed object is created during the Intersight Appliance setup, and it is configured
in the manual backup mode.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_appliance_backup_policy.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `backup_time`:(string) The next backup time set by the backup scheduler. Backup scheduler calculates the next backup time with the user-defined schedule set in the Schedule field. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `filename`:(string) Backup filename to backup or restore. 
* `is_password_set`:(bool) Indicates whether the value of the 'password' property has been set. 
* `manual_backup`:(bool) Backup mode of the appliance. Automatic backups of the appliance are not initiated if this property is set to 'true' and the backup schedule field is ignored. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `password`:(string) Password to authenticate the file server. 
* `protocol`:(string) Communication protocol used by the file server (e.g. scp or sftp).* `scp` - Secure Copy Protocol (SCP) to access the file server.* `sftp` - SSH File Transfer Protocol (SFTP) to access file server. 
* `remote_host`:(string) Hostname of the remote file server. 
* `remote_path`:(string) File server directory to copy the file. 
* `remote_port`:(int) Remote TCP port on the file server (e.g. 22 for scp). 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `username`:(string) Username to authenticate the fileserver. 
 
