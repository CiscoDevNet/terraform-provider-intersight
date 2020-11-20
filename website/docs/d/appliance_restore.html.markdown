
---
layout: "intersight"
page_title: "Intersight: intersight_appliance_restore"
sidebar_current: "docs-intersight-data-source-appliance-restore"
description: |-
Restore tracks requests to restore the Intersight Appliance. There will be only
one Restore managed object with a 'Started' state at any time. All other Restore
managed objects will be in terminal states.
---

# Data Source: intersight_appliance._restore
Restore tracks requests to restore the Intersight Appliance. There will be only
one Restore managed object with a 'Started' state at any time. All other Restore
managed objects will be in terminal states.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `elapsed_time`:(int) Elapsed time in seconds since the restore process has started. 
* `end_time`:(string) End date and time of the restore process. 
* `filename`:(string) Backup filename to backup or restore. 
* `is_password_set`:(bool) Indicates whether the value of the 'password' property has been set. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `password`:(string) Password for authenticating with the file server. 
* `protocol`:(string) Communication protocol used by the file server (e.g. scp or sftp).* `scp` - Secure Copy Protocol (SCP) to access the file server.* `sftp` - SSH File Transfer Protocol (SFTP) to access file server. 
* `remote_host`:(string) Hostname of the remote file server. 
* `remote_path`:(string) File server directory to copy the file. 
* `remote_port`:(int) Remote TCP port on the file server (e.g. 22 for scp). 
* `start_time`:(string) Start date and time of the restore process. 
* `status`:(string) Status of the restore managed object.* `Started` - Backup or restore process has started.* `Created` - Backup or restore is in created state.* `Failed` - Backup or restore process has failed.* `Completed` - Backup or restore process has completed.* `Copied` - Backup file has been copied. 
* `username`:(string) Username to authenticate the fileserver. 
