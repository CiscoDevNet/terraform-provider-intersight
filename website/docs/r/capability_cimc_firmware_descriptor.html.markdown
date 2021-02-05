---
subcategory: "capability"
layout: "intersight"
page_title: "Intersight: intersight_capability_cimc_firmware_descriptor"
description: |-
  Descriptor that identifies the server's redfish integration capability using cimc firmware info.
---

# Resource: intersight_capability_cimc_firmware_descriptor
Descriptor that identifies the server's redfish integration capability using cimc firmware info.
## Argument Reference
The following arguments are supported:
* `capabilities`:(Array) An array of relationships to capabilityCapability resources. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `description`:(string) Detailed information about the endpoint. 
* `model`:(string) The model of the endpoint, for which this capability information is applicable. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `revision`:(string) Revision information for the server. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `vendor`:(string) The vendor of the endpoint, for which this capability information is applicable. 
* `nr_version`:(string) The firmware or software version of the endpoint, for which this capability information is applicable. 


## Import
`intersight_capability_cimc_firmware_descriptor` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_capability_cimc_firmware_descriptor.example 1234567890987654321abcde
``` 