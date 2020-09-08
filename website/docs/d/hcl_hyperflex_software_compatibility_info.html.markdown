
---
layout: "intersight"
page_title: "Intersight: intersight_hcl_hyperflex_software_compatibility_info"
sidebar_current: "docs-intersight-data-source-hcl-hyperflex-software-compatibility-info"
description: |-
Lists software compatibility information between different HperFlex component versions like HyperFlex Data Platform, Hypervisor, Drive Firmware, etc.
---

# Data Source: intersight_hcl._hyperflex_software_compatibility_info
Lists software compatibility information between different HperFlex component versions like HyperFlex Data Platform, Hypervisor, Drive Firmware, etc.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `hxdp_version`:(string) HXDP component software version. 
* `hypervisor_type`:(string) Type fo Hypervisor the HyperFlex components versions are compatible with. For example ESX, Hyperv or KVM.* `ESXi` - ESXi hypervisor as specified by the user.* `HYPERV` - Hyperv hypervisor as specified by the user.* `KVM` - KVM hypervisor as specified by the user. 
* `hypervisor_version`:(string) Hypervisor component software version. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `server_fw_version`:(string) UCS Server Firmware component software version. 
