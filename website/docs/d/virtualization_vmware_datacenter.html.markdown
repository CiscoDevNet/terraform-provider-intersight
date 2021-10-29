---
subcategory: "virtualization"
layout: "intersight"
page_title: "Intersight: intersight_virtualization_vmware_datacenter"
description: |-
  Datacenter object in VMware inventory. It is the logical container for all other objects like Datastore, Host, VirtualMachine, etc.
---

# Data Source: intersight_virtualization_vmware_datacenter
Datacenter object in VMware inventory. It is the logical container for all other objects like Datastore, Host, VirtualMachine, etc.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_virtualization_vmware_datacenter.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `cluster_count`:(int) Count of all clusters associated with this DC. 
* `create_time`:(string) The time when this managed object was created. 
* `datastore_count`:(int) Count of all datastores associated with this DC. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `host_count`:(int) Count of all hosts associated with this DC. 
* `identity`:(string) The internally generated identity of this placement. This entity is not manipulated by users. It aids in uniquely identifying the placement object. 
* `inventory_path`:(string) Inventory path of the DC. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the virtual machine placement. It is the name of the VPC (Virtual Private Cloud) in case of AWSvirtual machine, and datacenter name in case of VMware virtual machine. 
* `network_count`:(int) Count of all networks associated with this datacenter (DC). 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `uuid`:(string) The uuid of this placement. The uuid is internally generated and not user specified. 
* `vm_count`:(int) Count of all virtual machines (VMs) associated with this DC. 
* `vm_template_count`:(int) Count of all virtual machines templates associated with this DC. 
 