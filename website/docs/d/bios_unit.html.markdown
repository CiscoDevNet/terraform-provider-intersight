
---
layout: "intersight"
page_title: "Intersight: intersight_bios_unit"
sidebar_current: "docs-intersight-data-source-bios-unit"
description: |-
The BIOS Unit present on a server.
---

# Data Source: intersight_bios._unit
The BIOS Unit present on a server.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `init_seq`:(string) The initSeq of the equipment. 
* `init_ts`:(string) The initTs of the equipment. 
* `model`:(string) This field identifies the model of the given component. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `revision`:(string) This field identifies the revision of the given component. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field identifies the serial of the given component. 
* `vendor`:(string) This field identifies the vendor of the given component. 
