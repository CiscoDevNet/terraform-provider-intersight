---
subcategory: "recovery"
layout: "intersight"
page_title: "Intersight: intersight_recovery_on_demand_backup"
description: |-
        On-demand backup request handler for endpoint.

---

# Data Source: intersight_recovery_on_demand_backup
On-demand backup request handler for endpoint.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_recovery_on_demand_backup.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the policy. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `file_name_prefix`:(string) Backup image file name prefix with timestamp (e.g., prefix-1572431305418). 
* `is_password_set`:(bool) Indicates whether the value of the 'password' property has been set. 
* `location_type`:(string) Backup storage location (local or remote).* `Network Share` - Backup is stored remotely on a separate server.* `Local Storage` - Backup is stored locally on endpoint. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `password`:(string) The backup server password. 
* `path`:(string) File system path for storing backup images, including IP address/hostname (e.g., 172.29.109.234/var/backups/). 
* `protocol`:(string) Protocol for backup transfer to network share.* `SCP` - Secure Copy Protocol (SCP) to access file server.* `SFTP` - SSH File Transfer Protocol (SFTP) to access file server.* `FTP` - File Transfer Protocol (FTP) to access file server. 
* `retention_count`:(int) Number of backup copies maintained on local or remote server (older backups overwritten). 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `user_name`:(string) The backup server username. 
 
