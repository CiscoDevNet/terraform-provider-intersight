
---
layout: "intersight"
page_title: "Intersight: intersight_iam_o_auth_token"
sidebar_current: "docs-intersight-data-source-iam-o-auth-token"
description: |-
The meta data for generating OAuth2 token of a user.
It is created when user logged in via OAuth2 using Authorization Code grant
and deleted upon logout, expiration timeout or manual deletion.
---

# Data Source: intersight_iam.o_auth_token
The meta data for generating OAuth2 token of a user.
It is created when user logged in via OAuth2 using Authorization Code grant
and deleted upon logout, expiration timeout or manual deletion.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `access_expiration_time`:(string) Expiration time for the JWT token to which it can be used for api calls. 
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `client_id`:(string) The identifier of the registered application to which the token belongs. 
* `client_ip_address`:(string) The user agent IP address from which the auth token is launched. 
* `client_name`:(string) The name of the registered application to which the token belongs. 
* `expiration_time`:(string) Expiration time for the JWT token to which it can be refreshed. 
* `last_login_client`:(string) The client address from which last login is initiated. 
* `last_login_time`:(string) The last login time for user. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `token_id`:(string) Token identifier. Not the Access Token itself. 
