
---
layout: "intersight"
page_title: "Intersight: intersight_iam_user_group"
sidebar_current: "docs-intersight-data-source-iam-user-group"
description: |-
User Group provides a way to assign permissions to a group of users based on the IdP attributes received after authentication.
---

# Data Source: intersight_iam._user_group
User Group provides a way to assign permissions to a group of users based on the IdP attributes received after authentication.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the user group which the dynamic user belongs to. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
