
---
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_cisco_hypervisor_manager"
sidebar_current: "docs-intersight-data-source-hyperflex-cisco-hypervisor-manager"
description: |-
A hypervisor manager to manage Cisco HyperFlex compute clusters and is available per user account.
---

# Data Source: intersight_hyperflex._cisco_hypervisor_manager
A hypervisor manager to manage Cisco HyperFlex compute clusters and is available per user account.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `identity`:(string) Identity of the hypervisor (not manipulated by user). It could be a UUID too. For example, c917093f-5443-4748-bc09-eec72ded7608. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The user provided name for the hypervisor manager. For example, vCenterIreland. Usually, this name is subject to manipulation by the user. It is not the identity of the hypervisor. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `version`:(string) Release version of the Hypervisor Manger (VMware vCenter Server 6.0.0 build-4541947). 
