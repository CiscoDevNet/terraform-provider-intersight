
---
layout: "intersight"
page_title: "Intersight: intersight_fcpool_pool"
sidebar_current: "docs-intersight-data-source-fcpool-pool"
description: |-
Pool represents a collection of WWN addresses that can be allocated to VHBAs of a server profile.
---

# Data Source: intersight_fcpool._pool
Pool represents a collection of WWN addresses that can be allocated to VHBAs of a server profile.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `assigned`:(int) Number of IDs that are currently assigned. 
* `assignment_order`:(string) Assignment order decides the order in which the next identifier is allocated.* `sequential` - Identifiers are assigned in a sequential order.* `default` - Assignment order is decided by the system. 
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `description`:(string) Description of the policy. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `pool_purpose`:(string) Purpose of this WWN pool. 
* `size`:(int) Total number of identifiers in this pool. 
