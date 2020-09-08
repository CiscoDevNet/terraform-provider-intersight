
---
layout: "intersight"
page_title: "Intersight: intersight_firmware_upgrade_impact_status"
sidebar_current: "docs-intersight-data-source-firmware-upgrade-impact-status"
description: |-
Captures the impact for an upgrade.
---

# Data Source: intersight_firmware._upgrade_impact_status
Captures the impact for an upgrade.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `computation_state`:(string) Captures the status of an upgrade impact calculation. Indicates whether the calculation is complete, in progress or the calculation is impossible due to the absence of the target image on the endpoint.* `Inprogress` - Upgrade impact calculation is in progress.* `Completed` - Upgrade impact calculation is completed.* `Unavailable` - Upgrade impact is not available since image is not present in FI. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `summary`:(string) The summary on the component or components getting impacted by the upgrade.* `Basic` - Summary of a single instance involved in the upgrade operation.* `Detail` - Summary of the collection of single instances for a complex component involved in the upgrade operation. For example, in case of a server upgrade, a detailed summary indicates impact of all the single instances which are part of the server, such as storage controller, drives, and BIOS. 
