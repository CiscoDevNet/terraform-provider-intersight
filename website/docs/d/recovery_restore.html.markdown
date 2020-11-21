
---
layout: "intersight"
page_title: "Intersight: intersight_recovery_restore"
sidebar_current: "docs-intersight-data-source-recovery-restore"
description: |-
Triggers a restore operation on the target endpoint.
---

# Data Source: intersight_recovery._restore
Triggers a restore operation on the target endpoint.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
