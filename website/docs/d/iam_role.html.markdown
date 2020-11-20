
---
layout: "intersight"
page_title: "Intersight: intersight_iam_role"
sidebar_current: "docs-intersight-data-source-iam-role"
description: |-
A role is a collection of privilege sets that are assigned to a user using a permission object.
---

# Data Source: intersight_iam._role
A role is a collection of privilege sets that are assigned to a user using a permission object.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `description`:(string) Informative description about each role. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the role which has to be granted to user. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
