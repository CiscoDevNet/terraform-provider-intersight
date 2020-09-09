
---
layout: "intersight"
page_title: "Intersight: intersight_uuidpool_pool_member"
sidebar_current: "docs-intersight-data-source-uuidpool-pool-member"
description: |-
PoolMember represents a single UUID that is part of a pool.
---

# Data Source: intersight_uuidpool._pool_member
PoolMember represents a single UUID that is part of a pool.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `assigned`:(bool) Boolean to represent whether the ID is assigned or not. 
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `uuid`:(string) UUID Prefix+Suffix of this PoolMember. 
