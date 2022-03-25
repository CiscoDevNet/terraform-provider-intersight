---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_net_app_lun"
description: |-
        NetApp LUN (logical unit number) is an identifier for a device called a logical unit addressed by a SAN protocol.

---

# Data Source: intersight_storage_net_app_lun
NetApp LUN (logical unit number) is an identifier for a device called a logical unit addressed by a SAN protocol.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_net_app_lun.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Short description about the volume. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `key`:(string) Unique identifier of Lun across data center. 
* `mapped`:(bool) Reports if the LUN is mapped to one or more initiator groups. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `naa_id`:(string) NAA id of volume. It is a significant number to identify corresponding lun path in hypervisor. 
* `name`:(string) Named entity of the volume. 
* `os_type`:(string) The operating system (OS) type for this LUN.* `Linux` - Family of open source Unix-like operating systems based on the Linux kernel.* `AIX` - Advanced Interactive Executive (AIX).* `HP-UX` - HP-UX is implementation of the Unix operating system, based on Unix System V.* `Hyper-V` - Windows Server 2008 or Windows Server 2012 Hyper-V.* `OpenVMS` - OpenVMS is multi-user, multiprocessing virtual memory-based operating system.* `Solaris` - Solaris is a Unix operating system.* `NetWare` - NetWare is a computer network operating system.* `VMware` - An enterprise-class, type-1 hypervisor developed by VMware for deploying and serving virtual computers.* `Windows` - Single-partition Windows disk using the Master Boot Record (MBR) partitioning style.* `Xen` - Xen is a type-1 hypervisor, providing services that allow multiple computer operating systems to execute on the same computer hardware concurrently. 
* `path`:(string) Path where the LUN is mounted. 
* `serial`:(string) Serial number for the provisioned LUN. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `size`:(int) User provisioned volume size. It is the size exposed to host. 
* `state`:(string) The administrative state of a LUN.* `offline` - The LUN is administratively offline, or a more detailed offline reason is not available.* `online` - The state of the LUN is online. 
* `uuid`:(string) Universally unique identifier of the LUN. 
 
