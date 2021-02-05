---
subcategory: "hcl"
layout: "intersight"
page_title: "Intersight: intersight_hcl_hyperflex_software_compatibility_info"
description: |-
  Lists software compatibility information between different HperFlex component versions like HyperFlex Data Platform, Hypervisor, Drive Firmware, etc.
---

# Resource: intersight_hcl_hyperflex_software_compatibility_info
Lists software compatibility information between different HperFlex component versions like HyperFlex Data Platform, Hypervisor, Drive Firmware, etc.
## Argument Reference
The following arguments are supported:
* `app_catalog`:(HashMap) - A reference to a hyperflexAppCatalog resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `constraints`:(Array)
This complex property has following sub-properties:
  + `constraint_name`:(string) Name or key of the applicable compatibility constraint. 
  + `constraint_value`:(string) Value of the applicable compatibility constraint. Could either be a string value or a regex. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `hxdp_version`:(string) HXDP component software version. 
* `hypervisor_type`:(string) Type fo Hypervisor the HyperFlex components versions are compatible with. For example ESX, Hyperv or KVM.* `ESXi` - The hypervisor running on the HyperFlex cluster is a Vmware ESXi hypervisor of any version.* `HyperFlexAp` - The hypervisor running on the HyperFlex cluster is Cisco HyperFlex Application Platform.* `Hyper-V` - The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V.* `Unknown` - The hypervisor running on the HyperFlex cluster is not known. 
* `hypervisor_version`:(string) Hypervisor component software version. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `server_fw_version`:(string) UCS Server Firmware component software version. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 


## Import
`intersight_hcl_hyperflex_software_compatibility_info` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_hcl_hyperflex_software_compatibility_info.example 1234567890987654321abcde
``` 