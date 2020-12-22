---
subcategory: "externalsite"
layout: "intersight"
page_title: "Intersight: intersight_externalsite_authorization"
description: |-
  An authentication request that will be used to get authorized from external repository like cisco.com in order to download the image. This MO creation is a one time configuration before calling firmware.Upgrade API. This MO has to be modified with updated details whenever the user has updated the credentials in external repository.
---

# Data Source: intersight_externalsite_authorization
An authentication request that will be used to get authorized from external repository like cisco.com in order to download the image. This MO creation is a one time configuration before calling firmware.Upgrade API. This MO has to be modified with updated details whenever the user has updated the credentials in external repository.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `is_password_set`:(bool) Indicates whether the value of the 'password' property has been set. 
* `is_user_id_set`:(bool) Indicates whether the value of the 'userId' property has been set. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `password`:(string) The password of the given username to download the image from external repository like cisco.com. 
* `repository_type`:(string) The repository type to which this authorization will be requested. Cisco is the only available repository today.* `cisco` - Cisco as an external site from where the resources like image will be downloaded. 
* `user_id`:(string) The username that has permission to download the image from external repository like cisco.com. 
