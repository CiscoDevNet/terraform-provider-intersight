
---
layout: "intersight"
page_title: "Intersight: intersight_iam_user"
sidebar_current: "docs-intersight-data-source-iam-user"
description: |-
The Intersight account user.
---

# Data Source: intersight_iam._user
The Intersight account user.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `client_ip_address`:(string) IP address from which the user last logged in to Intersight. 
* `email`:(string) Email of the user. Users are added to Intersight using the email configured in the IdP. 
* `first_name`:(string) First name of the user. This field is populated from the IdP attributes received after authentication. 
* `last_login_time`:(string) Last successful login time for user. 
* `last_name`:(string) Last name of the user. This field is populated from the IdP attributes received after authentication. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name as configured in the IdP. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `user_id_or_email`:(string) UserID or email as configured in the IdP. 
* `user_type`:(string) Type of the User. If a user is added manually by specifying the email address, or has logged in using groups, based on the IdP attributes received during authentication. If added manually, the user type will be static, otherwise dynamic. 
* `user_unique_identifier`:(string) Unique id of the user used by the identity provider to store the user. 
