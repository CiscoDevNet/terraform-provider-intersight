
---
layout: "intersight"
page_title: "Intersight: intersight_storage_hyper_flex_volume"
sidebar_current: "docs-intersight-data-source-storage-hyper-flex-volume"
description: |-
A HyperFlex Volume entity.
---

# Data Source: intersight_storage._hyper_flex_volume
A HyperFlex Volume entity.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `capacity`:(int) Provisioned Capacity of the Storage container in bytes. 
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `description`:(string) Short description about the volume. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `naa_id`:(string) NAA id of volume. It is a significant number to identify corresponding lun path in hypervisor. 
* `name`:(string) Named entity of the volume. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `serial_number`:(string) Serial number of the volume. 
* `size`:(int) User provisioned volume size. It is the size exposed to host. 
* `uuid`:(string) Uuid of the Datastore/Storage Container. 
