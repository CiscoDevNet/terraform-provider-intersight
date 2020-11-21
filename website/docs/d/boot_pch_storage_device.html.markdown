
---
layout: "intersight"
page_title: "Intersight: intersight_boot_pch_storage_device"
sidebar_current: "docs-intersight-data-source-boot-pch-storage-device"
description: |-
Pch Storage Boot Device configured on the server.
---

# Data Source: intersight_boot._pch_storage_device
Pch Storage Boot Device configured on the server.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `model`:(string) This field identifies the model of the given component. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the boot device configured in the boot policy. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `order`:(int) The order of the boot device configured in the boot policy. 
* `revision`:(string) This field identifies the revision of the given component. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field identifies the serial of the given component. 
* `state`:(string) The state of the boot device configured in the boot policy. 
* `type`:(string) The type of the boot device configured in the boot policy. 
* `vendor`:(string) This field identifies the vendor of the given component. 
