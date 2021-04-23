---
subcategory: "appliance"
layout: "intersight"
page_title: "Intersight: intersight_appliance_remote_file_import"
description: |-
  Trigger a remote file import request by configuring this mo.
---

# Data Source: intersight_appliance_remote_file_import
Trigger a remote file import request by configuring this mo.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_appliance_remote_file_import.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `filename`:(string) The name of the file to be imported. 
* `hostname`:(string) The hostname of the machine where the file is located. 
* `is_password_set`:(bool) Indicates whether the value of the 'password' property has been set. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `password`:(string) Password for remote requiest. 
* `path`:(string) The port that should be used for the remote request. 
* `port`:(int) The port that should be used for the remote request. 
* `protocol`:(string) Protocol for the remote request.* `scp` - Secure Copy Protocol (SCP) to access the file server.* `sftp` - SSH File Transfer Protocol (SFTP) to access file server. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `username`:(string) The username for the remote request. 
 