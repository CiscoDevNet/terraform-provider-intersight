---
subcategory: "appliance"
layout: "intersight"
page_title: "Intersight: intersight_appliance_diag_setting"
description: |-
  DiagSetting model is used for changing the password of the operating system's diagnostic
user account. The diagnostic user account can be used to login to the Intersight Appliance
virtual machine.
The diagnostic user account is protected by two separate authentication mechanisms: user's
password and Cisco CT-engine generated key. Only the Intersight Appliance's local account
administrator has the privileges to use this REST API.
---

# Resource: intersight_appliance_diag_setting
DiagSetting model is used for changing the password of the operating system's diagnostic
user account. The diagnostic user account can be used to login to the Intersight Appliance
virtual machine.
The diagnostic user account is protected by two separate authentication mechanisms: user's
password and Cisco CT-engine generated key. Only the Intersight Appliance's local account
administrator has the privileges to use this REST API.
## Argument Reference
The following arguments are supported:
* `account`:(Array with Maximum of one item) - A reference to a iamAccount resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `is_password_set`:(bool)(Computed) Indicates whether the value of the 'password' property has been set. 
* `message`:(string) Status message of the password change operation. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `password`:(string) Password of the Intersight Appliance's OS diagnostic user account. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 


## Import
`intersight_appliance_diag_setting` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_appliance_diag_setting.example 1234567890987654321abcde
```