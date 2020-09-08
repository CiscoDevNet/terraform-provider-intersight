
---
layout: "intersight"
page_title: "Intersight: intersight_iam_session_limits"
sidebar_current: "docs-intersight-data-source-iam-session-limits"
description: |-
The session related configuration limits.
---

# Data Source: intersight_iam._session_limits
The session related configuration limits.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `idle_time_out`:(int) The idle timeout interval for the web session in seconds. When a session is not refreshed for this duration, the session is marked as idle and removed. 
* `maximum_limit`:(int) The maximum number of sessions allowed in an account. The default value is 128. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `per_user_limit`:(int) The maximum number of sessions allowed per user. Default value is 32. 
* `session_time_out`:(int) The session expiry duration in seconds. 
