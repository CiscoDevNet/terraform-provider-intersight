
---
layout: "intersight"
page_title: "Intersight: intersight_uuidpool_pool"
sidebar_current: "docs-intersight-data-source-uuidpool-pool"
description: |-
Pool represents a collection of UUID items that can be allocated to server profiles.
---

# Data Source: intersight_uuidpool._pool
Pool represents a collection of UUID items that can be allocated to server profiles.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `assigned`:(int) Number of IDs that are currently assigned. 
* `assignment_order`:(string) Assignment order decides the order in which the next identifier is allocated.* `sequential` - Identifiers are assigned in a sequential order.* `default` - Assignment order is decided by the system. 
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `description`:(string) Description of the policy. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `prefix`:(string) The UUID prefix must be in hexadecimal format xxxxxxxx-xxxx-xxxx. 
* `size`:(int) Total number of identifiers in this pool. 
