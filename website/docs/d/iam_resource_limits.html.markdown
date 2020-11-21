
---
layout: "intersight"
page_title: "Intersight: intersight_iam_resource_limits"
sidebar_current: "docs-intersight-data-source-iam-resource-limits"
description: |-
The resource limits used to limit resources such as User accounts.
---

# Data Source: intersight_iam._resource_limits
The resource limits used to limit resources such as User accounts.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `per_account_user_limit`:(int) The maximum number of users allowed in an account. The default value is 200. 
