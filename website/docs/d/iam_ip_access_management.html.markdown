
---
layout: "intersight"
page_title: "Intersight: intersight_iam_ip_access_management"
sidebar_current: "docs-intersight-data-source-iam-ip-access-management"
description: |-
The access management based on IP address.
---

# Data Source: intersight_iam._ip_access_management
The access management based on IP address.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `enable`:(bool) Flag stores the state of IP address based access management. Access management is enabled when it's true. 
* `last_recovery_time`:(string) The access to account gets locked out if wrong IP addresses are configured. Account Administrators have privilege to unblock the account. It stores the time when the account was last recovered from lock out. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
