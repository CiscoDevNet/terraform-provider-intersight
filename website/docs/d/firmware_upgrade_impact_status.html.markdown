---
subcategory: "firmware"
layout: "intersight"
page_title: "Intersight: intersight_firmware_upgrade_impact_status"
description: |-
  Captures the impact for an upgrade.
---

# Data Source: intersight_firmware_upgrade_impact_status
Captures the impact for an upgrade.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_firmware_upgrade_impact_status.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `computation_state`:(string) Captures the status of an upgrade impact calculation. Indicates whether the calculation is complete, in progress or the calculation is impossible due to the absence of the target image on the endpoint.* `Inprogress` - Upgrade impact calculation is in progress.* `Completed` - Upgrade impact calculation is completed.* `Unavailable` - Upgrade impact is not available since image is not present in FI. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `summary`:(string) The summary on the component or components getting impacted by the upgrade.* `Basic` - Summary of a single instance involved in the upgrade operation.* `Detail` - Summary of the collection of single instances for a complex component involved in the upgrade operation. For example, in case of a server upgrade, a detailed summary indicates impact of all the single instances which are part of the server, such as storage controller, drives, and BIOS. 
 