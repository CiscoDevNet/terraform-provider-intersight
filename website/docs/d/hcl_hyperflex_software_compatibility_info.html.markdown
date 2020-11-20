
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
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `hxdp_version`:(string) HXDP component software version. 
* `hypervisor_type`:(string) Type fo Hypervisor the HyperFlex components versions are compatible with. For example ESX, Hyperv or KVM.* `ESXi` - A Vmware ESXi hypervisor of any version.* `HXAP` - The hypervisor running on the HyperFlex cluster is Cisco HyperFlex Application Platform.* `Hyper-V` - The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V.* `Unknown` - The hypervisor running on the HyperFlex cluster is not known. 
* `hypervisor_version`:(string) Hypervisor component software version. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `server_fw_version`:(string) UCS Server Firmware component software version. 
