
---
layout: "intersight"
page_title: "Intersight: intersight_storage_pure_controller"
sidebar_current: "docs-intersight-data-source-storage-pure-controller"
description: |-
A storage controller entity in Pure FlashArray.
---

# Data Source: intersight_storage._pure_controller
A storage controller entity in Pure FlashArray.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `model`:(string) This field identifies the model of the given component. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Storage array controller name. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `operational_mode`:(string) Controller running mode, Primary or Secondary.* `Unknown` - Component operational state is unknown.* `Primary` - Component operates in primary mode and accepts workloads.* `Secondary` - Component is running as a secondary or standby mode.* `Maintenance` - Component is in maintenance mode for upgrade. During maintenance mode, component does not perform any workload. 
* `revision`:(string) This field identifies the revision of the given component. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field identifies the serial of the given component. 
* `status`:(string) Status of the storage controller.* `Unknown` - Component status is not available.* `Ok` - Component is healthy and no issues found.* `Degraded` - Functioning, but not at full capability due to a non-fatal failure.* `Critical` - Not functioning or requiring immediate attention.* `Offline` - Component is installed, but powered off.* `Identifying` - Component is in initialization process.* `NotAvailable` - Component is not installed or not available.* `Updating` - Software update is in progress.* `Unrecognized` - Component is not recognized. It may be because the component is not installed properly or it is not supported. 
* `vendor`:(string) This field identifies the vendor of the given component. 
* `version`:(string) Software version running on a storage controller. 
