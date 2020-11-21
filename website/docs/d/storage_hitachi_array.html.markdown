
---
layout: "intersight"
page_title: "Intersight: intersight_storage_hitachi_array"
sidebar_current: "docs-intersight-data-source-storage-hitachi-array"
description: |-
The details of the Hitachi storage array.
---

# Data Source: intersight_storage._hitachi_array
The details of the Hitachi storage array.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `ctl1_ip`:(string) IP address of controller 1 of the storage system. 
* `ctl1_micro_version`:(string) GUM (Gateway for Unified Management) version of the controller 1. 
* `ctl2_ip`:(string) IP address of controller 2 of the storage system. 
* `ctl2_micro_version`:(string) GUM (Gateway for Unified Management) version of the controller 2. 
* `device_id`:(string) Storage device ID. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `model`:(string) This field identifies the model of the given component. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Administrator defined name for the device. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `revision`:(string) This field identifies the revision of the given component. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field identifies the serial of the given component. 
* `svp_ip`:(string) IP address of the SVP. 
* `target_ctl`:(string) Controller operated by the REST API. 
* `uuid`:(string) Unique identity of the device. 
* `vendor`:(string) This field identifies the vendor of the given component. 
* `version`:(string) Current running software version of the device. 
