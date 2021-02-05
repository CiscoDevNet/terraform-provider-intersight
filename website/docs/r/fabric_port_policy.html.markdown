---
subcategory: "fabric"
layout: "intersight"
page_title: "Intersight: intersight_fabric_port_policy"
description: |-
  A policy for all the physical ports of the Fabric Interconnect.
---

# Resource: intersight_fabric_port_policy
A policy for all the physical ports of the Fabric Interconnect.
## Argument Reference
The following arguments are supported:
* `description`:(string) Description of the policy. 
* `device_model`:(string) This field specifies the device model that this Port Policy is being configured for.* `UCS-FI-6454` - The standard 4th generation UCS Fabric Interconnect with 54 ports.* `UCS-FI-64108` - The expanded 4th generation UCS Fabric Interconnect with 108 ports.* `unknown` - Unknown device type, usage is TBD. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `organization`:(HashMap) - A reference to a organizationOrganization resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `profiles`:(Array) An array of relationships to fabricSwitchProfile resources. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 


## Import
`intersight_fabric_port_policy` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_fabric_port_policy.example 1234567890987654321abcde
``` 