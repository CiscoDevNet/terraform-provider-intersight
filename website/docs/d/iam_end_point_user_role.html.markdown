
---
layout: "intersight"
page_title: "Intersight: intersight_iam_end_point_user_role"
sidebar_current: "docs-intersight-data-source-iam-end-point-user-role"
description: |-
Mapping of endpoint user to endpoint roles.
---

# Data Source: intersight_iam._end_point_user_role
Mapping of endpoint user to endpoint roles.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `change_password`:(bool) Denotes whether password has changed. 
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `enabled`:(bool) Enables the user account on the endpoint. 
* `is_password_set`:(bool) Indicates whether the value of the 'password' property has been set. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `password`:(string) Valid login password of the user. 
