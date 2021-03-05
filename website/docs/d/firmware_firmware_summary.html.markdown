---
subcategory: "firmware"
layout: "intersight"
page_title: "Intersight: intersight_firmware_firmware_summary"
description: |-
  Update inventory that contains the details for the firmware running on each component in the compute.Physical object.
---

# Data Source: intersight_firmware_firmware_summary
Update inventory that contains the details for the firmware running on each component in the compute.Physical object.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_firmware_firmware_summary.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `bundle_version`:(string) Version details at the bundle level for the each of server. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
 