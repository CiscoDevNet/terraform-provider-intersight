---
subcategory: "config"
layout: "intersight"
page_title: "Intersight: intersight_config_exporter"
description: |-
  All export operations are captured as Exporter instances. Users shall use this Exporter
mo to track the export operation progress.
---

# Data Source: intersight_config_exporter
All export operations are captured as Exporter instances. Users shall use this Exporter
mo to track the export operation progress.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `download_path`:(string) Pre-signed URL to download the exported package, if the export operation has completed successfully. Regenerated during a GET request, if the existing pre-signed URL has expired. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) An identifier for the exporter instance. 
* `status`:(string) Status of the export operation.* `` - The operation has not started.* `InProgress` - The operation is in progress.* `Success` - The operation has succeeded.* `Failed` - The operation has failed.* `RollBackInitiated` - The rollback has been inititiated for import failure.* `RollBackFailed` - The rollback has failed for import failure.* `RollbackCompleted` - The rollback has completed for import failure.* `RollbackAborted` - The rollback has been aborted for import failure.* `OperationTimedOut` - The operation has timed out. 
* `status_message`:(string) Status message associated with failures or progress indication. 
