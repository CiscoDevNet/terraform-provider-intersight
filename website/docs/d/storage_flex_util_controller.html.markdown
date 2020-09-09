
---
layout: "intersight"
page_title: "Intersight: intersight_storage_flex_util_controller"
sidebar_current: "docs-intersight-data-source-storage-flex-util-controller"
description: |-
Storage Flex Util Adapter.
---

# Data Source: intersight_storage._flex_util_controller
Storage Flex Util Adapter.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `controller_name`:(string) Name of the Flex Util Controller. 
* `controller_status`:(string) The current status of the controller. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `ff_controller_id`:(string) Identifier for the Storage Flex Util Controller. 
* `internal_state`:(string) The internal state of the controller. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
