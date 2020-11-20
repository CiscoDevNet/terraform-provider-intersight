
---
layout: "intersight"
page_title: "Intersight: intersight_appliance_diag_setting"
sidebar_current: "docs-intersight-data-source-appliance-diag-setting"
description: |-
DiagSetting model is used for changing the password of the operating system's diagnostic
user account. The diagnostic user account can be used to login to the Intersight Appliance
virtual machine.
The diagnostic user account is protected by two separate authentication mechanisms: user's
password and Cisco CT-engine generated key. Only the Intersight Appliance's local account
administrator has the privileges to use this REST API.
---

# Data Source: intersight_appliance._diag_setting
DiagSetting model is used for changing the password of the operating system's diagnostic
user account. The diagnostic user account can be used to login to the Intersight Appliance
virtual machine.
The diagnostic user account is protected by two separate authentication mechanisms: user's
password and Cisco CT-engine generated key. Only the Intersight Appliance's local account
administrator has the privileges to use this REST API.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `is_password_set`:(bool) Indicates whether the value of the 'password' property has been set. 
* `message`:(string) Status message of the password change operation. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `password`:(string) Password of the Intersight Appliance's OS diagnostic user account. 
