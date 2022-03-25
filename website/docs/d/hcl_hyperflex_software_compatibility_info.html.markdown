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
To access the ith object of the results obtained, use `data.intersight_hcl_hyperflex_software_compatibility_info.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `hxdp_version`:(string) HXDP component software version. 
* `hypervisor_type`:(string) Type fo Hypervisor the HyperFlex components versions are compatible with. For example ESX, Hyperv or KVM.* `ESXi` - The hypervisor running on the HyperFlex cluster is a Vmware ESXi hypervisor of any version.* `HyperFlexAp` - The hypervisor of the virtualization platform is Cisco HyperFlex Application Platform.* `IWE` - The hypervisor of the virtualization platform is Cisco Intersight Workload Engine.* `Hyper-V` - The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V.* `Unknown` - The hypervisor running on the HyperFlex cluster is not known. 
* `hypervisor_version`:(string) Hypervisor component software version. 
* `is_mgmt_build`:(string) Type of the HXDP bundle mgmt or full. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `server_fw_version`:(string) UCS Server Firmware component software version. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
