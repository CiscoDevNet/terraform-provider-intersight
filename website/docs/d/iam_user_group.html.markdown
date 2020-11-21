
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
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the user group which the dynamic user belongs to. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
