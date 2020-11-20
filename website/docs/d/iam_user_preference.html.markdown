
---
layout: "intersight"
page_title: "Intersight: intersight_iam_user_preference"
sidebar_current: "docs-intersight-data-source-iam-user-preference"
description: |-
Holder for UI preferences such as theme, columns.
---

# Data Source: intersight_iam._user_preference
Holder for UI preferences such as theme, columns.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `user_unique_identifier`:(string) Unique id of the user used by the identity provider to store the user. 
