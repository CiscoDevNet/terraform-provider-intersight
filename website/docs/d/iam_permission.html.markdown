
---
layout: "intersight"
page_title: "Intersight: intersight_iam_permission"
sidebar_current: "docs-intersight-data-source-iam-permission"
description: |-
Permission provides a way to assign roles to a user or user group to perform operations on object hierarchy.
---

# Data Source: intersight_iam._permission
Permission provides a way to assign roles to a user or user group to perform operations on object hierarchy.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `description`:(string) The informative description about each permission. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the permission which has to be granted to user. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
