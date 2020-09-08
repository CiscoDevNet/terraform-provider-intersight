
---
layout: "intersight"
page_title: "Intersight: intersight_config_exported_item"
sidebar_current: "docs-intersight-data-source-config-exported-item"
description: |-
A single managed object that is being exported.
---

# Data Source: intersight_config._exported_item
A single managed object that is being exported.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `file_name`:(string) Name of the file corresponding to item MO. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) MO item identity (the moref corresponding to item) expressed as a string. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `service_version`:(string) Version of the service that owns the item MO. 
* `status`:(string) Status of the item's export operation.* `` - The operation has not started.* `InProgress` - The operation is in progress.* `Success` - The operation has succeeded.* `Failed` - The operation has failed.* `RollBackInitiated` - The rollback has been inititiated for import failure.* `RollBackFailed` - The rollback has failed for import failure.* `RollbackCompleted` - The rollback has completed for import failure.* `RollbackAborted` - The rollback has been aborted for import failure.* `OperationTimedOut` - The operation has timed out. 
* `status_message`:(string) Progress or error message for the MO's export operation. 
