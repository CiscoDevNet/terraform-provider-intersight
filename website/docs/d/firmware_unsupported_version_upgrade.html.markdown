---
subcategory: "firmware"
layout: "intersight"
page_title: "Intersight: intersight_firmware_unsupported_version_upgrade"
description: |-
  This represents an operation managed object used for upgrading equipment that cannot be discovered due to unsupported firmware. Currently, it only supports blade upgrades.
---

# Data Source: intersight_firmware_unsupported_version_upgrade
This represents an operation managed object used for upgrading equipment that cannot be discovered due to unsupported firmware. Currently, it only supports blade upgrades.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_firmware_unsupported_version_upgrade.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `download_error`:(string) Any error encountered. Set to empty when download is in progress or completed. 
* `download_progress`:(int) The download progress of the file represented as a percentage between 0% and 100%. If progress reporting is not possible a value of -1 is sent. 
* `download_retries`:(int) The number of retries the plugin attempted before succeeding or failing the download. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `upgrade_status`:(string) Workflow status of firmware upgrade.* `None` - Upgrade status is none when upgrade is in progress.* `Completed` - Upgrade completed successfully.* `Failed` - Upgrade status is failed when upgrade has failed. 
 