---
subcategory: "externalsite"
layout: "intersight"
page_title: "Intersight: intersight_externalsite_authorization"
description: |-
  An authentication request that will be used to get authorized from external repository like cisco.com in order to download the image. This MO creation is a one time configuration before calling firmware.Upgrade API. This MO has to be modified with updated details whenever the user has updated the credentials in external repository.
---

# Resource: intersight_externalsite_authorization
An authentication request that will be used to get authorized from external repository like cisco.com in order to download the image. This MO creation is a one time configuration before calling firmware.Upgrade API. This MO has to be modified with updated details whenever the user has updated the credentials in external repository.
## Argument Reference
The following arguments are supported:
* `account`:(HashMap) -(Computed) A reference to a iamAccount resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `is_password_set`:(bool)(Computed) Indicates whether the value of the 'password' property has been set. 
* `is_user_id_set`:(bool)(Computed) Indicates whether the value of the 'userId' property has been set. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `password`:(string) The password of the given username to download the image from external repository like cisco.com. 
* `repository_type`:(string) The repository type to which this authorization will be requested. Cisco is the only available repository today.* `cisco` - Cisco as an external site from where the resources like image will be downloaded. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `user_id`:(string) The username that has permission to download the image from external repository like cisco.com. 


## Import
`intersight_externalsite_authorization` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_externalsite_authorization.example 1234567890987654321abcde
``` 