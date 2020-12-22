---
subcategory: "iam"
layout: "intersight"
page_title: "Intersight: intersight_iam_ip_access_management"
description: |-
  The access management based on IP address.
---

# Data Source: intersight_iam_ip_access_management
The access management based on IP address.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `enable`:(bool) Flag stores the state of IP address based access management. Access management is enabled when it's true. 
* `last_recovery_time`:(string) The access to account gets locked out if wrong IP addresses are configured. Account Administrators have privilege to unblock the account. It stores the time when the account was last recovered from lock out. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
