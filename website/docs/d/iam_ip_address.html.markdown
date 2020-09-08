
---
layout: "intersight"
page_title: "Intersight: intersight_iam_ip_address"
sidebar_current: "docs-intersight-data-source-iam-ip-address"
description: |-
Add an IP address to enable IP address based access management.
---

# Data Source: intersight_iam._ip_address
Add an IP address to enable IP address based access management.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `address`:(string) The Trusted IP range's address. IP address, CIDR range, and IP address range formats are supported. For example ’12.13.14.15’, ’12.13.14.0/24’, and ’12.13.14.15-12.13.14.200’. Reserved IP ranges ‘127.0.0.1, ‘10.0.0.0/8’, ‘172.16.0.0/12’, and ‘192.168.0.0/16’ are not allowed. 
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `description`:(string) Description of Trusted IP address range. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
