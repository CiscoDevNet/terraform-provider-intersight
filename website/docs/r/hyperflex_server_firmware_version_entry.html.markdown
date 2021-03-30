---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_server_firmware_version_entry"
description: |-
  An entry specifying supported server firmware version in regex format.
---

# Resource: intersight_hyperflex_server_firmware_version_entry
An entry specifying supported server firmware version in regex format.
## Argument Reference
The following arguments are supported:
* `constraint`:(HashMap) - The conditions that must be satisfied before applying the AppSetting. 
This complex property has following sub-properties:
  + `hxdp_version`:(string) The supported HyperFlex Data Platform version in regex format. 
  + `hypervisor_type`:(string) The hypervisor type for the HyperFlex cluster.* `ESXi` - The hypervisor running on the HyperFlex cluster is a Vmware ESXi hypervisor of any version.* `HyperFlexAp` - The hypervisor running on the HyperFlex cluster is Cisco HyperFlex Application Platform.* `Hyper-V` - The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V.* `Unknown` - The hypervisor running on the HyperFlex cluster is not known. 
  + `mgmt_platform`:(string) The supported management platform for the HyperFlex Cluster.* `FI` - The host servers used in the cluster deployment are managed by a UCS Fabric Interconnect.* `EDGE` - The host servers used in the cluster deployment are standalone severs. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `server_model`:(string) The supported server models in regex format. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `server_firmware_version`:(HashMap) - A reference to a hyperflexServerFirmwareVersion resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `server_platform`:(string) The server platform type that is applicable for the server firmware bundle version.* `M5` - M5 generation of UCS server.* `M3` - M3 generation of UCS server.* `M4` - M4 generation of UCS server. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `nr_version`:(string) The server firmware bundle version. 


## Import
`intersight_hyperflex_server_firmware_version_entry` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_hyperflex_server_firmware_version_entry.example 1234567890987654321abcde
``` 