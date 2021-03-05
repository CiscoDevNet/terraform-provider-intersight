---
subcategory: "hcl"
layout: "intersight"
page_title: "Intersight: intersight_hcl_hyperflex_software_compatibility_info"
description: |-
  Lists software compatibility information between different HperFlex component versions like HyperFlex Data Platform, Hypervisor, Drive Firmware, etc.
---

# Data Source: intersight_hcl_hyperflex_software_compatibility_info
Lists software compatibility information between different HperFlex component versions like HyperFlex Data Platform, Hypervisor, Drive Firmware, etc.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hcl_hyperflex_software_compatibility_info.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `hxdp_version`:(string) HXDP component software version. 
* `hypervisor_type`:(string) Type fo Hypervisor the HyperFlex components versions are compatible with. For example ESX, Hyperv or KVM.* `ESXi` - The hypervisor running on the HyperFlex cluster is a Vmware ESXi hypervisor of any version.* `HyperFlexAp` - The hypervisor running on the HyperFlex cluster is Cisco HyperFlex Application Platform.* `Hyper-V` - The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V.* `Unknown` - The hypervisor running on the HyperFlex cluster is not known. 
* `hypervisor_version`:(string) Hypervisor component software version. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `server_fw_version`:(string) UCS Server Firmware component software version. 
 