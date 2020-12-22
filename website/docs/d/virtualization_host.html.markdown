---
subcategory: "virtualization"
layout: "intersight"
page_title: "Intersight: intersight_virtualization_host"
description: |-
  Depicts operations to control the life cycle of a Hypervisor Host.
---

# Data Source: intersight_virtualization_host
Depicts operations to control the life cycle of a Hypervisor Host.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `action`:(string) Action to be performed on a host (Create, PowerState, Migrate, Clone etc).* `None` - A place holder for the default value.* `PowerOffStorageController` - Power off HyperFlex storage controller node running on selected hypervisor host.* `PowerOnStorageController` - Power on HyperFlex storage controller node running on selected hypervisor host. 
* `hypervisor_type`:(string) Identifies the broad product type of the hypervisor but without any version information. It is here to easily identify the type of the virtual machine. There are other entities (Host, Cluster, etc.) that can be indirectly used to determine the hypervisor but a direct attribute makes it easier to work with.* `ESXi` - The hypervisor running on the HyperFlex cluster is a Vmware ESXi hypervisor of any version.* `HyperFlexAp` - The hypervisor running on the HyperFlex cluster is Cisco HyperFlex Application Platform.* `Hyper-V` - The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V.* `Unknown` - The hypervisor running on the HyperFlex cluster is not known. 
* `identity`:(string) Unique identifier assigned to the hypervisor host. 
* `model`:(string) Commercial model information about this hardware. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the hypervisor host. It must be unique within the targer endpoint. 
* `serial`:(string) Serial number of this host (internally generated). 
* `vendor`:(string) Commercial vendor details of this hardware. 
