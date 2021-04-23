---
subcategory: "virtualization"
layout: "intersight"
page_title: "Intersight: intersight_virtualization_vmware_folder"
description: |-
  A folder in a VMware vCenter. Folder can be created directly under the vCenter, under a datacenter, or inside another folder.
---

# Data Source: intersight_virtualization_vmware_folder
A folder in a VMware vCenter. Folder can be created directly under the vCenter, under a datacenter, or inside another folder.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_virtualization_vmware_folder.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `identity`:(string) The internally generated identity of this folder. This entity is not manipulated by users. It aids in uniquely identifying the folder object. For VMware, this is a MOR (managed object reference). 
* `internal`:(bool) If a folder is internal, it will be set to true. 
* `inventory_path`:(string) Inventory path to the folder. Example - /DC/myFolder. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the folder as stored in vCenter. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `typeof_folder`:(string) Determines the type of folder. e.g. vCenter folder, VM and Templete Folder, StorageFolder, NetworkFolder, Host and Cluster Folder.* `Unknown` - The type of the folder is unknown. It may not represent that the folder does not exist but indicates that something might be wrong.* `VMTemplateFolder` - The folder contains VMs and VM templates.* `StorageFolder` - The folder contains storage devices.* `HostClusterFolder` - The folder contains hosts and clusters.* `NetworkFolder` - The folder contains network items.* `VcenterFolder` - The folder created under a vCenter or vCenter folder. 
 