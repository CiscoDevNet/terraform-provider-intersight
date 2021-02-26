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
The following arguments can be used to get data of already created objects in Intersight appliance:
* `filename`:(string) The name of the file to be imported. 
* `hostname`:(string) The hostname of the machine where the file is located. 
* `is_password_set`:(bool) Indicates whether the value of the 'password' property has been set. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `password`:(string) Password for remote requiest. 
* `path`:(string) The port that should be used for the remote request. 
* `port`:(int) The port that should be used for the remote request. 
* `protocol`:(string) Protocol for the remote request.* `scp` - Secure Copy Protocol (SCP) to access the file server.* `sftp` - SSH File Transfer Protocol (SFTP) to access file server. 
* `username`:(string) The username for the remote request. 
