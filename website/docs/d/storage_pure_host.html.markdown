
---
layout: "intersight"
page_title: "Intersight: intersight_storage_pure_host"
sidebar_current: "docs-intersight-data-source-storage-pure-host"
description: |-
A host entity in PureStorage FlashArray. It is an abstraction used by PureStorage to organize the storage network addresses (Fibre Channel worldwide names or iSCSI qualified names) of client computers and to control communications between clients and volumes.
---

# Data Source: intersight_storage._pure_host
A host entity in PureStorage FlashArray. It is an abstraction used by PureStorage to organize the storage network addresses (Fibre Channel worldwide names or iSCSI qualified names) of client computers and to control communications between clients and volumes.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `description`:(string) Short description about the host. 
* `host_group_name`:(string) Name of host group where the host belongs to. Empty if host is not part of any HostGroup. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the host in storage array. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `os_type`:(string) Operating system running on the host. 
