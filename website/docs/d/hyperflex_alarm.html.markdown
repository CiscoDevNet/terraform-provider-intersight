
---
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_alarm"
sidebar_current: "docs-intersight-data-source-hyperflex-alarm"
description: |-

---

# Data Source: intersight_hyperflex._alarm

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `acknowledged`:(bool)
* `acknowledged_by`:(string)
* `acknowledged_time`:(int)
* `acknowledged_time_as_utc`:(string)
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `description`:(string)
* `entity_data`:(string)
* `entity_name`:(string)
* `entity_type`:(string)
* `entity_uu_id`:(string)
* `message`:(string)
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string)
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `status`:(string)
* `triggered_time`:(int)
* `triggered_time_as_utc`:(string)
* `uuid`:(string)
