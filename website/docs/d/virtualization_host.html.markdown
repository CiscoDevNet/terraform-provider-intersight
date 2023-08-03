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
* `action`:(string) Action to be performed on a host (Create, PowerState, Migrate, Clone etc).* `None` - A place holder for the default value.* `EnterMaintenanceMode` - Put a host into maintenance mode.* `ExitMaintenanceMode` - Put a host into active mode.* `PowerOffStorageController` - Power off HyperFlex storage controller node running on selected hypervisor host.* `PowerOnStorageController` - Power on HyperFlex storage controller node running on selected hypervisor host. 
* `create_time`:(string) The time when this managed object was created. 
* `discovered`:(bool) Flag to indicate whether the configuration is created from inventory object. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `evacuate`:(bool) If true, move powered-off and suspended virtual machines to other hosts in the cluster. 
* `hypervisor_type`:(string) Identifies the broad product type of the hypervisor but without any version information. It is here to easily identify the type of the virtual machine. There are other entities (Host, Cluster, etc.) that can be indirectly used to determine the hypervisor but a direct attribute makes it easier to work with.* `ESXi` - The hypervisor running on the HyperFlex cluster is a Vmware ESXi hypervisor of any version.* `Hyper-V` - The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V.* `Unknown` - The hypervisor running on the HyperFlex cluster is not known. 
* `identity`:(string) Unique identifier assigned to the hypervisor host. 
* `maintenance_state`:(string) Expected state of host. An action on the host (e.g., Enter Maintenance) may cause the host to be put into maintenance mode.* `None` - A place holder for the default value.* `Enter` - Power action is performed on the virtual machine.* `Exit` - The virtual machine will be migrated from existing node to a different node in cluster. The behavior depends on the underlying hypervisor. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) Commercial model information about this hardware. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the hypervisor host. It must be unique within the target endpoint. 
* `serial`:(string) Serial number of this host (internally generated). 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `vendor`:(string) Commercial vendor details of this hardware. 
 
