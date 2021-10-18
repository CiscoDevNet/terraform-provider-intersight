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
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_virtualization_host.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `action`:(string) Action to be performed on a host (Create, PowerState, Migrate, Clone etc).* `None` - A place holder for the default value.* `PowerOffStorageController` - Power off HyperFlex storage controller node running on selected hypervisor host.* `PowerOnStorageController` - Power on HyperFlex storage controller node running on selected hypervisor host. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `hypervisor_type`:(string) Identifies the broad product type of the hypervisor but without any version information. It is here to easily identify the type of the virtual machine. There are other entities (Host, Cluster, etc.) that can be indirectly used to determine the hypervisor but a direct attribute makes it easier to work with.* `ESXi` - The hypervisor running on the HyperFlex cluster is a Vmware ESXi hypervisor of any version.* `HyperFlexAp` - The hypervisor of the virtualization platform is Cisco HyperFlex Application Platform.* `IWE` - The hypervisor of the virtualization platform is Cisco Intersight Workload Engine.* `Hyper-V` - The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V.* `Unknown` - The hypervisor running on the HyperFlex cluster is not known. 
* `identity`:(string) Unique identifier assigned to the hypervisor host. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) Commercial model information about this hardware. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the hypervisor host. It must be unique within the target endpoint. 
* `serial`:(string) Serial number of this host (internally generated). 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `vendor`:(string) Commercial vendor details of this hardware. 
 