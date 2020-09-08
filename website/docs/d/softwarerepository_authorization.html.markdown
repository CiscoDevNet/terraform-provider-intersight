
---
layout: "intersight"
page_title: "Intersight: intersight_softwarerepository_authorization"
sidebar_current: "docs-intersight-data-source-softwarerepository-authorization"
description: |-
User's consent for Intersight to contact an external software repository such as cisco.com, on the behalf of the user.
---

# Data Source: intersight_softwarerepository._authorization
User's consent for Intersight to contact an external software repository such as cisco.com, on the behalf of the user.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `is_password_set`:(bool) Indicates whether the value of the 'password' property has been set. 
* `is_user_id_set`:(bool) Indicates whether the value of the 'userId' property has been set. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `password`:(string) The password that will be used by Intersight to create OAuth2 tokens for interacting with the external repository, on the user account's behalf. 
* `repository_type`:(string) The external repository for which this authorization has been provided. The only supported repository today is cisco.com.* `Cisco` - External repository hosted on cisco.com.* `IntersightCloud` - Repository hosted by the Intersight Cloud.* `LocalMachine` - The file is available on the local client machine. Used as an upload source type.* `NetworkShare` - External repository in the customer datacenter. This will typically be a file server. 
* `user_id`:(string) The username that will be used by Intersight to create OAuth2 tokens for interacting with the external repository, on the user account's behalf. 
