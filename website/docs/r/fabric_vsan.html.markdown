---
subcategory: "fabric"
layout: "intersight"
page_title: "Intersight: intersight_fabric_vsan"
description: |-
  Configuration object sent by user to create VSAN configurations.
---

# Resource: intersight_fabric_vsan
Configuration object sent by user to create VSAN configurations.
## Argument Reference
The following arguments are supported:
* `default_zoning`:(string) Enables or Disables the default zoning state.* `Enabled` - Admin configured Enabled State.* `Disabled` - Admin configured Disabled State. 
* `fc_network_policy`:(HashMap) - A reference to a fabricFcNetworkPolicy resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `fc_zone_sharing_mode`:(string) Logical grouping mode for fc ports. 
* `fcoe_vlan`:(int) FCOE Vlan associated to the VSAN configuration. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) User given name for the VSAN configuration. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `vsan_id`:(int) Virtual San Identifier in the switch. 


## Import
`intersight_fabric_vsan` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_fabric_vsan.example 1234567890987654321abcde
``` 