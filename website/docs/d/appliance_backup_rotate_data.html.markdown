---
subcategory: "appliance"
layout: "intersight"
page_title: "Intersight: intersight_appliance_backup_rotate_data"
description: |-
        Keeps track of all current backups being manage by the backup rotate policy.
        It is used to actively manage backups as part of the rotation.

---

# Data Source: intersight_appliance_backup_rotate_data
Keeps track of all current backups being manage by the backup rotate policy.
It is used to actively manage backups as part of the rotation.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_appliance_backup_rotate_data.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `backup_time`:(string) The time at which the backup was taken. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `filename`:(string) Backup filename to backup or restore. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `protocol`:(string) Communication protocol used by the file server (e.g. scp, sftp, or CIFS).* `scp` - Secure Copy Protocol (SCP) to access the file server.* `sftp` - SSH File Transfer Protocol (SFTP) to access file server.* `cifs` - Common Internet File System (CIFS) Protocol to access file server. 
* `remote_host`:(string) Hostname of the remote file server. 
* `remote_path`:(string) File server directory or share name to copy the file. 
* `remote_port`:(int) Remote TCP port on the file server (e.g. 22 for scp). 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `username`:(string) Username to authenticate the fileserver. 
 
