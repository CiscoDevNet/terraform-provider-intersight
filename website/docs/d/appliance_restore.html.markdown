---
subcategory: "appliance"
layout: "intersight"
page_title: "Intersight: intersight_appliance_restore"
description: |-
        The Restore object is crucial for handling system restoration requests, enabling recovery operations to maintain system stability and data accuracy.
        #### Purpose
        The Restore object facilitates the process of restoring system data from backups, ensuring timely and accurate recovery following disruptions or data loss.
        #### Key Concepts
        - **Process Tracking:** Monitors restore processes from start to finish, providing insights into operation progress and status.
        - **Single Active Instance:** Ensures only one active restoration operation occurs, preventing overlap and conflicts in recovery efforts.
        - **Secure Operations:** Utilizes secure communication protocols for data transfer, promoting reliability and safety during restoration.
        - **Relationship Management:** Maintains connections to account structures, supporting coherent and systematic restore operations.

---

# Data Source: intersight_appliance_restore
The Restore object is crucial for handling system restoration requests, enabling recovery operations to maintain system stability and data accuracy.
#### Purpose
The Restore object facilitates the process of restoring system data from backups, ensuring timely and accurate recovery following disruptions or data loss.
#### Key Concepts
- **Process Tracking:** Monitors restore processes from start to finish, providing insights into operation progress and status.
- **Single Active Instance:** Ensures only one active restoration operation occurs, preventing overlap and conflicts in recovery efforts.
- **Secure Operations:** Utilizes secure communication protocols for data transfer, promoting reliability and safety during restoration.
- **Relationship Management:** Maintains connections to account structures, supporting coherent and systematic restore operations.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_appliance_restore.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `elapsed_time`:(int) Elapsed time in seconds since the restore process has started. 
* `end_time`:(string) End date and time of the restore process. 
* `filename`:(string) Backup filename to backup or restore. 
* `is_password_set`:(bool) Indicates whether the value of the 'password' property has been set. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `password`:(string) Password for authenticating with the file server. 
* `protocol`:(string) Communication protocol used by backup and restore workflow (e.g. scp, sftp, cifs, or local).* `scp` - Secure Copy Protocol (SCP) to access the file server.* `sftp` - SSH File Transfer Protocol (SFTP) to access file server.* `cifs` - Common Internet File System (CIFS) Protocol to access file server.* `local` - Backup file is stored in Intersight Appliance. 
* `remote_host`:(string) Hostname of the remote file server. Not required when protocol is local. 
* `remote_path`:(string) File server directory or share name to copy the file. Not required when protocol is local. 
* `remote_port`:(int) Remote TCP port on the file server (e.g. 22 for scp). Not required when protocol is local. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `start_time`:(string) Start date and time of the restore process. 
* `status`:(string) Status of the restore managed object.* `Started` - Backup or restore process has started.* `Created` - Backup or restore is in created state.* `Failed` - Backup or restore process has failed.* `Completed` - Backup or restore process has completed.* `Copied` - Backup file has been copied.* `Cleanup Failed` - Cleanup of the old backup has failed. 
* `username`:(string) Username to authenticate the fileserver. Not required when protocol is local. 
 
