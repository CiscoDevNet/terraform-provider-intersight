
---
layout: "intersight"
page_title: "Intersight: intersight_config_imported_item"
sidebar_current: "docs-intersight-data-source-config-imported-item"
description: |-
A single managed object that is being imported.
---

# Data Source: intersight_config._imported_item
A single managed object that is being imported.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `is_shared`:(bool) Specifies whether this item MO was in shared scope or user scope when exported. 
* `is_updated`:(bool) Specifies whether this item MO was updated or created while importing in desired service. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) MO item identity (the moref corresponding to item) expressed as a string. 
* `new_moid`:(string) Moid of the MO created/updated during import for the item. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `service_version`:(string) Version of the service that owned the item MO when the item was exported. 
* `status`:(string) Status of the item's import operation.* `` - The operation has not started.* `InProgress` - The operation is in progress.* `Success` - The operation has succeeded.* `Failed` - The operation has failed.* `RollBackInitiated` - The rollback has been inititiated for import failure.* `RollBackFailed` - The rollback has failed for import failure.* `RollbackCompleted` - The rollback has completed for import failure.* `RollbackAborted` - The rollback has been aborted for import failure.* `OperationTimedOut` - The operation has timed out. 
* `status_message`:(string) Progress or error message for the MO's import operation. 
