
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
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `enabled`:(bool) Enables the user account on the endpoint. 
* `is_password_set`:(bool) Indicates whether the value of the 'password' property has been set. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `password`:(string) Valid login password of the user. 
