---
subcategory: "capability"
layout: "intersight"
page_title: "Intersight: intersight_capability_adapter_unit_descriptor"
description: |-
  Descriptor that uniquely identifies an adaptor.
---

# Resource: intersight_capability_adapter_unit_descriptor
Descriptor that uniquely identifies an adaptor.
## Argument Reference
The following arguments are supported:
* `capabilities`:(Array) An array of relationships to capabilityCapability resources. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `connectivity_order`:(string) Order in which the ports are connected; sequential or cyclic. 
* `description`:(string) Detailed information about the endpoint. 
* `ethernet_port_speed`:(int) The port speed for ethernet ports in Mbps. 
* `fibre_channel_port_speed`:(int) The port speed for fibre channel ports in Mbps. 
* `fibre_channel_scsi_ioq_limit`:(int) The number of SCSI I/O Queue resources to allocate. 
* `model`:(string) The model of the endpoint, for which this capability information is applicable. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `num_dce_ports`:(int) Number of Dce Ports for the adaptor. 
* `prom_card_type`:(string) Prom card type for the adaptor. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `vendor`:(string) The vendor of the endpoint, for which this capability information is applicable. 
* `nr_version`:(string) The firmware or software version of the endpoint, for which this capability information is applicable. 


## Import
`intersight_capability_adapter_unit_descriptor` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_capability_adapter_unit_descriptor.example 1234567890987654321abcde
``` 