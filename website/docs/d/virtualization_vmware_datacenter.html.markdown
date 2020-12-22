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
The following arguments can be used to get data of already created objects in Intersight appliance:
* `cluster_count`:(int) Count of all clusters associated with this DC. 
* `datastore_count`:(int) Count of all datastores associated with this DC. 
* `host_count`:(int) Count of all hosts associated with this DC. 
* `identity`:(string) Internally generated identity of this datacenter. This entity is not manipulated by users. It aids in uniquely identifying the datacenter object. For VMware, this is a MOR (managed object reference). 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) User provided name for the datacenter. Usually, this name is subject to manipulations by user. It is not the identity of the datacenter. 
* `network_count`:(int) Count of all networks associated with this datacenter (DC). 
* `vm_count`:(int) Count of all virtual machines (VMs) associated with this DC. 
