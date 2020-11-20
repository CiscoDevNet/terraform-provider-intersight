
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
* `address`:(string) The Trusted IP range's address. IP address, CIDR range, and IP address range formats are supported. For example '12.13.14.15', '12.13.14.0/24', and '12.13.14.15-12.13.14.200'. Reserved IP ranges '127.0.0.1', '10.0.0.0/8', '172.16.0.0/12', and '192.168.0.0/16' are not allowed. 
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `description`:(string) Description of Trusted IP address range. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
