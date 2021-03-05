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

# Data Source: intersight_appliance_diag_setting
DiagSetting model is used for changing the password of the operating system's diagnostic
user account. The diagnostic user account can be used to login to the Intersight Appliance
virtual machine.
The diagnostic user account is protected by two separate authentication mechanisms: user's
password and Cisco CT-engine generated key. Only the Intersight Appliance's local account
administrator has the privileges to use this REST API.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_appliance_diag_setting.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `is_password_set`:(bool) Indicates whether the value of the 'password' property has been set. 
* `message`:(string) Status message of the password change operation. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `password`:(string) Password of the Intersight Appliance's OS diagnostic user account. 
 