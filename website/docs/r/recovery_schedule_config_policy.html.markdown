---
subcategory: "recovery"
layout: "intersight"
page_title: "Intersight: intersight_recovery_schedule_config_policy"
description: |-
  Base Schedule config which contains all the required inputs to do schedule on a local or remote server.
---

# Resource: intersight_recovery_schedule_config_policy
Base Schedule config which contains all the required inputs to do schedule on a local or remote server.
## Argument Reference
The following arguments are supported:
* `backup_profiles`:(Array) An array of relationships to recoveryBackupProfile resources. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `description`:(string) Description of the policy. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `organization`:(HashMap) - A reference to a organizationOrganization resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `schedule`:(HashMap) - Schedule to create a backup on the target device. Minimum is 4 hours and Max is 1440 hours (30 Days). 
This complex property has following sub-properties:
  + `execution_time`:(string) The time at which the backup is to be run on a given day. This is used when the frequency unit is daily. 
  + `frequency_unit`:(string) The frequency at which the backup schedule must run.* `Daily` - Allows the user to run the backup daily at a given time.* `Periodic` - Allows the user to run the backup after a certain number of hours. 
  + `hours`:(int) The frequency, in hours, at which the backup schedule runs.* `8` - * `4` - * `12` - * `16` - * `20` - 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 


## Import
`intersight_recovery_schedule_config_policy` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_recovery_schedule_config_policy.example 1234567890987654321abcde
``` 