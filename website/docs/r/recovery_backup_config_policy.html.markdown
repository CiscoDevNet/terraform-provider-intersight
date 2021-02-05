---
subcategory: "recovery"
layout: "intersight"
page_title: "Intersight: intersight_recovery_backup_config_policy"
description: |-
  Backup config policy which contains all the required inputs to do backup on a local or remote server.
---

# Resource: intersight_recovery_backup_config_policy
Backup config policy which contains all the required inputs to do backup on a local or remote server.
## Argument Reference
The following arguments are supported:
* `backup_profiles`:(Array) An array of relationships to recoveryBackupProfile resources. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `description`:(string) Description of the policy. 
* `file_name_prefix`:(string) The file name for the backup image. This name is added as a prefix in the name for the backup image. A unique file name for the backup image is created along with a timestamp. For example: prefix-1572431305418. 
* `is_password_set`:(bool)(Computed) Indicates whether the value of the 'password' property has been set. 
* `location_type`:(string) Specifies whether the backup will be stored locally or remotely.* `Network Share` - The backup is stored remotely on a separate server.* `Local Storage` - The backup is stored locally on the endpoint. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `organization`:(HashMap) - A reference to a organizationOrganization resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `password`:(string) Password of Backup server. 
* `path`:(string) The file system path where the backup images must be stored. Include the IP address/hostname of the network share location and the complete file system path. For example: 172.29.109.234/var/backups/. 
* `protocol`:(string) Protocol for transferring the backup image to the network share location.* `SCP` - Secure Copy Protocol (SCP) to access the file server.* `SFTP` - SSH File Transfer Protocol (SFTP) to access file server.* `FTP` - File Transfer Protocol (FTP) to access file server. 
* `retention_count`:(int) Number of backup copies maintained on the local or remote server. When the created backup files exceed this number, the initial backup files are overwritten in a sequential manner. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `user_name`:(string) Username for the backup server. 


## Import
`intersight_recovery_backup_config_policy` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_recovery_backup_config_policy.example 1234567890987654321abcde
``` 