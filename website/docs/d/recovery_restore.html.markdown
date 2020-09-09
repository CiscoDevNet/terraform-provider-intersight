
---
layout: "intersight"
page_title: "Intersight: intersight_recovery_restore"
sidebar_current: "docs-intersight-data-source-recovery-restore"
description: |-
Triggers a restore operation on the target endpoint.
---

# Data Source: intersight_recovery._restore
Triggers a restore operation on the target endpoint.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
