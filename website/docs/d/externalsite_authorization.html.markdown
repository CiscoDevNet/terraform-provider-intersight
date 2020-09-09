
---
layout: "intersight"
page_title: "Intersight: intersight_externalsite_authorization"
sidebar_current: "docs-intersight-data-source-externalsite-authorization"
description: |-
An authentication request that will be used to get authorized from external repository like cisco.com in order to download the image. This MO creation is a one time configuration before calling firmware.Upgrade API. This MO has to be modified with updated details whenever the user has updated the credentials in external repository.
---

# Data Source: intersight_externalsite._authorization
An authentication request that will be used to get authorized from external repository like cisco.com in order to download the image. This MO creation is a one time configuration before calling firmware.Upgrade API. This MO has to be modified with updated details whenever the user has updated the credentials in external repository.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `is_password_set`:(bool) Indicates whether the value of the 'password' property has been set. 
* `is_user_id_set`:(bool) Indicates whether the value of the 'userId' property has been set. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `password`:(string) The password of the given username to download the image from external repository like cisco.com. 
* `repository_type`:(string) The repository type to which this authorization will be requested. Cisco is the only available repository today.* `cisco` - Cisco as an external site from where the resources like image will be downloaded. 
* `user_id`:(string) The username that has permission to download the image from external repository like cisco.com. 
