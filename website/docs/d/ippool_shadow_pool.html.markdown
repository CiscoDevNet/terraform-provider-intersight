
---
layout: "intersight"
page_title: "Intersight: intersight_ippool_shadow_pool"
sidebar_current: "docs-intersight-data-source-ippool-shadow-pool"
description: |-
Shadow Pool is a tracking object created on behalf of an IP pool, for each VRF.
---

# Data Source: intersight_ippool._shadow_pool
Shadow Pool is a tracking object created on behalf of an IP pool, for each VRF.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `assigned`:(int) Number of IDs that are currently assigned. 
* `assignment_order`:(string) Assignment order decides the order in which the next identifier is allocated.* `sequential` - Identifiers are assigned in a sequential order.* `default` - Assignment order is decided by the system. 
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `description`:(string) Description of the policy. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `size`:(int) Total number of identifiers in this pool. 
* `v4_assigned`:(int) Number of IPv4 addresses currently in use. 
* `v4_size`:(int) Number of IPv4 addresses in this pool. 
* `v6_assigned`:(int) Number of IPv6 addresses currently in use. 
* `v6_size`:(int) Number of IPv6 addresses in this pool. 
