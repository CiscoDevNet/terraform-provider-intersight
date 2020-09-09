
---
layout: "intersight"
page_title: "Intersight: intersight_hcl_hyperflex_software_compatibility_info"
sidebar_current: "docs-intersight-resource-hcl-hyperflex-software-compatibility-info"
description: |-
  Lists software compatibility information between different HperFlex component versions like HyperFlex Data Platform, Hypervisor, Drive Firmware, etc.
---

# Resource: intersight_hcl._hyperflex_software_compatibility_info
Lists software compatibility information between different HperFlex component versions like HyperFlex Data Platform, Hypervisor, Drive Firmware, etc.
## Argument Reference
The following arguments are supported:
* `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
* `app_catalog`:(Array with Maximum of one item) - A reference to a hyperflexAppCatalog resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string)(Computed) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `class_id`:(string)(Computed) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `constraints`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string)(Computed) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
  + `constraint_name`:(string) Name or key of the applicable compatibility constraint. 
  + `constraint_value`:(string) Value of the applicable compatibility constraint. Could either be a string value or a regex. 
  + `object_type`:(string) The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types. 
* `hxdp_version`:(string) HXDP component software version. 
* `hypervisor_type`:(string) Type fo Hypervisor the HyperFlex components versions are compatible with. For example ESX, Hyperv or KVM.* `ESXi` - ESXi hypervisor as specified by the user.* `HYPERV` - Hyperv hypervisor as specified by the user.* `KVM` - KVM hypervisor as specified by the user. 
* `hypervisor_version`:(string) Hypervisor component software version. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string)(Computed) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `server_fw_version`:(string) UCS Server Firmware component software version. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
