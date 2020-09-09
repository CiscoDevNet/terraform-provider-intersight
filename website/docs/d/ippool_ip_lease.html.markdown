
---
layout: "intersight"
page_title: "Intersight: intersight_ippool_ip_lease"
sidebar_current: "docs-intersight-data-source-ippool-ip-lease"
description: |-
IpLease represents an IP address that is allocated from a pool to a specific entity like server profile.
---

# Data Source: intersight_ippool._ip_lease
IpLease represents an IP address that is allocated from a pool to a specific entity like server profile.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `ip_v4_address`:(string) IPv4 Address given as a lease to an external entity like server profiles. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
