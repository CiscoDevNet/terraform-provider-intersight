
---
layout: "intersight"
page_title: "Intersight: intersight_iam_system"
sidebar_current: "docs-intersight-data-source-iam-system"
description: |-
System is the top level object in the Intersight. All other objects which can be accessed globally are added to system object, like privilege sets and privileges can be shared by multiple roles and privilege sets.
---

# Data Source: intersight_iam._system
System is the top level object in the Intersight. All other objects which can be accessed globally are added to system object, like privilege sets and privileges can be shared by multiple roles and privilege sets.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
