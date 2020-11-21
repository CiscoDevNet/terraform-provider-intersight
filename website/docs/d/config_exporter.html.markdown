
---
layout: "intersight"
page_title: "Intersight: intersight_config_exporter"
sidebar_current: "docs-intersight-data-source-config-exporter"
description: |-
All export operations are captured as Exporter instances. Users shall use this Exporter
mo to track the export operation progress.
---

# Data Source: intersight_config._exporter
All export operations are captured as Exporter instances. Users shall use this Exporter
mo to track the export operation progress.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `download_path`:(string) Pre-signed URL to download the exported package, if the export operation has completed successfully. Regenerated during a GET request, if the existing pre-signed URL has expired. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) An identifier for the exporter instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `status`:(string) Status of the export operation.* `` - The operation has not started.* `InProgress` - The operation is in progress.* `Success` - The operation has succeeded.* `Failed` - The operation has failed.* `RollBackInitiated` - The rollback has been inititiated for import failure.* `RollBackFailed` - The rollback has failed for import failure.* `RollbackCompleted` - The rollback has completed for import failure.* `RollbackAborted` - The rollback has been aborted for import failure.* `OperationTimedOut` - The operation has timed out. 
* `status_message`:(string) Status message associated with failures or progress indication. 
