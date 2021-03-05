---
subcategory: "fabric"
layout: "intersight"
page_title: "Intersight: intersight_fabric_pc_operation"
description: |-
  PcOperation objects allows the user to alter the state of the port channel.
---

# Resource: intersight_fabric_pc_operation
PcOperation objects allows the user to alter the state of the port channel.
## Argument Reference
The following arguments are supported:
* `admin_state`:(string) Admin configured state to disable the port channel.* `Enabled` - Admin configured Enabled State.* `Disabled` - Admin configured Disabled State. 
* `config_state`:(string) The configured state of these settings in the target chassis. The value is any one of Applied, Applying, Failed. Applied - This state denotes that the admin state changes are applied successfully in the target FI domain. Applying - This state denotes that the admin state changes are being applied in the target FI domain. Failed - This state denotes that the admin state changes could not be applied in the target FI domain.* `None` - Nil value when no action has been triggered by the user.* `Applied` - User configured settings are in applied state.* `Applying` - User settings are being applied on the target server.* `Failed` - User configured settings could not be applied. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `network_element`:(HashMap) - A reference to a networkElement resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `pc_id`:(int) Port Channel Identifier for the collection of ports. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 


## Import
`intersight_fabric_pc_operation` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_fabric_pc_operation.example 1234567890987654321abcde
``` 