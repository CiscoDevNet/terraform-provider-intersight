
---
layout: "intersight"
page_title: "Intersight: intersight_iam_session"
sidebar_current: "docs-intersight-data-source-iam-session"
description: |-
The web session of a user. After a user logs into Intersight, a session object is created. Session object is deleted upon logout, idle timeout, expiry timeout, or manual deletion.
---

# Data Source: intersight_iam._session
The web session of a user. After a user logs into Intersight, a session object is created. Session object is deleted upon logout, idle timeout, expiry timeout, or manual deletion.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `client_ip_address`:(string) The user agent IP address from which the session is launched. 
* `expiration`:(string) Expiration time for the session. 
* `idle_time_expiration`:(string) Idle time expiration for the session. 
* `last_login_client`:(string) The client address from which last login is initiated. 
* `last_login_time`:(string) The last login time for user. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
