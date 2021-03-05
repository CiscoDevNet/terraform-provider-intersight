---
subcategory: "iam"
layout: "intersight"
page_title: "Intersight: intersight_iam_session"
description: |-
  The web session of a user. After a user logs into Intersight, a session object is created. Session object is deleted upon logout, idle timeout, expiry timeout, or manual deletion.
---

# Data Source: intersight_iam_session
The web session of a user. After a user logs into Intersight, a session object is created. Session object is deleted upon logout, idle timeout, expiry timeout, or manual deletion.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iam_session.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `client_ip_address`:(string) The user agent IP address from which the session is launched. 
* `expiration`:(string) Expiration time for the session. 
* `idle_time_expiration`:(string) Idle time expiration for the session. 
* `last_login_client`:(string) The client address from which last login is initiated. 
* `last_login_time`:(string) The last login time for user. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
 