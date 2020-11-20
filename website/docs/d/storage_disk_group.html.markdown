
---
layout: "intersight"
page_title: "Intersight: intersight_storage_disk_group"
sidebar_current: "docs-intersight-data-source-storage-disk-group"
description: |-
Group of one or more Spans to configure virtual drive.
---

# Data Source: intersight_storage._disk_group
Group of one or more Spans to configure virtual drive.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name to identity this disk group in the controller. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `raid_type`:(string) Raid level of the virtual drives in this diskgroup. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
