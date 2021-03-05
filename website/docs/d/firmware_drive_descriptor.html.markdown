---
subcategory: "firmware"
layout: "intersight"
page_title: "Intersight: intersight_firmware_drive_descriptor"
description: |-
  Descriptor to uniquely identify a Drive.
---

# Data Source: intersight_firmware_drive_descriptor
Descriptor to uniquely identify a Drive.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_firmware_drive_descriptor.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `brand_string`:(string) The brand string of the endpoint for which this capability information is applicable. 
* `description`:(string) Detailed information about the endpoint. 
* `label`:(string) The label type for the component. 
* `model`:(string) The model of the endpoint, for which this capability information is applicable. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `revision`:(string) The revision for the component. 
* `vendor`:(string) The vendor of the endpoint, for which this capability information is applicable. 
* `nr_version`:(string) The firmware or software version of the endpoint, for which this capability information is applicable. 
 