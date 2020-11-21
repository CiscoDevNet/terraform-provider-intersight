
---
layout: "intersight"
page_title: "Intersight: intersight_virtualization_vmware_vcenter"
sidebar_current: "docs-intersight-data-source-virtualization-vmware-vcenter"
description: |-
VMware vCenter entity. The vCenter has a name assigned by user in Intersight.
---

# Data Source: intersight_virtualization._vmware_vcenter
VMware vCenter entity. The vCenter has a name assigned by user in Intersight.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `identity`:(string) Identity of the hypervisor (not manipulated by user). It could be a UUID too. For example, c917093f-5443-4748-bc09-eec72ded7608. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The user provided name for the hypervisor manager. For example, vCenterIreland. Usually, this name is subject to manipulation by the user. It is not the identity of the hypervisor. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `version`:(string) Release version of the Hypervisor Manger (VMware vCenter Server 6.0.0 build-4541947). 
