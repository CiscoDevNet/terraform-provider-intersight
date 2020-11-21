
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
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `idle_time_out`:(int) The idle timeout interval for the web session in seconds. When a session is not refreshed for this duration, the session is marked as idle and removed. The minimum value is 300 seconds and the maximum value is 18000 seconds (5 hours). The system default value is 1800 seconds. 
* `maximum_limit`:(int) The maximum number of sessions allowed in an account. The default value is 128. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `per_user_limit`:(int) The maximum number of sessions allowed per user. Default value is 32. 
* `session_time_out`:(int) The session expiry duration in seconds. The minimum value is 350 seconds and the maximum value is 31536000 seconds (1 year). The system default value is 57600 seconds. 
