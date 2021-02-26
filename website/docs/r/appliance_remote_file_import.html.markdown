---
subcategory: "appliance"
layout: "intersight"
page_title: "Intersight: intersight_appliance_remote_file_import"
description: |-
  Trigger a remote file import request by configuring this mo.
---

# Resource: intersight_appliance_remote_file_import
Trigger a remote file import request by configuring this mo.
## Argument Reference
The following arguments are supported:
* `account`:(HashMap) -(Computed) A reference to a iamAccount resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `filename`:(string) The name of the file to be imported. 
* `hostname`:(string) The hostname of the machine where the file is located. 
* `is_password_set`:(bool)(Computed) Indicates whether the value of the 'password' property has been set. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `password`:(string) Password for remote requiest. 
* `path`:(string) The port that should be used for the remote request. 
* `port`:(int) The port that should be used for the remote request. 
* `protocol`:(string) Protocol for the remote request.* `scp` - Secure Copy Protocol (SCP) to access the file server.* `sftp` - SSH File Transfer Protocol (SFTP) to access file server. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `username`:(string) The username for the remote request. 


## Import
`intersight_appliance_remote_file_import` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_appliance_remote_file_import.example 1234567890987654321abcde
``` 