
---
layout: "intersight"
page_title: "Intersight: intersight_virtualization_vmware_datacenter"
sidebar_current: "docs-intersight-data-source-virtualization-vmware-datacenter"
description: |-
Datacenter object in VMware inventory. It is the logical container for all other objects like Datastore, Host, VirtualMachine, etc.
---

# Data Source: intersight_virtualization._vmware_datacenter
Datacenter object in VMware inventory. It is the logical container for all other objects like Datastore, Host, VirtualMachine, etc.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `cluster_count`:(int) Count of all clusters associated with this DC. 
* `datastore_count`:(int) Count of all datastores associated with this DC. 
* `host_count`:(int) Count of all hosts associated with this DC. 
* `identity`:(string) Internally generated identity of this datacenter. This entity is not manipulated by users. It aids in uniquely identifying the datacenter object. For VMware, this is a MOR (managed object reference). 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) User provided name for the datacenter. Usually, this name is subject to manipulations by user. It is not the identity of the datacenter. 
* `network_count`:(int) Count of all networks associated with this datacenter (DC). 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `vm_count`:(int) Count of all virtual machines (VMs) associated with this DC. 
