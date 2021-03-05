---
subcategory: "iam"
layout: "intersight"
page_title: "Intersight: intersight_iam_session_limits"
description: |-
  The session related configuration limits.
---

# Data Source: intersight_iam_session_limits
The session related configuration limits.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iam_session_limits.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `idle_time_out`:(int) The idle timeout interval for the web session in seconds. When a session is not refreshed for this duration, the session is marked as idle and removed. The minimum value is 300 seconds and the maximum value is 18000 seconds (5 hours). The system default value is 1800 seconds. 
* `maximum_limit`:(int) The maximum number of sessions allowed in an account. The default value is 128. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `per_user_limit`:(int) The maximum number of sessions allowed per user. Default value is 32. 
* `session_time_out`:(int) The session expiry duration in seconds. The minimum value is 350 seconds and the maximum value is 31536000 seconds (1 year). The system default value is 57600 seconds. 
 