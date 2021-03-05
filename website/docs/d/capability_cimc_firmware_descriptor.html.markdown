---
subcategory: "capability"
layout: "intersight"
page_title: "Intersight: intersight_capability_cimc_firmware_descriptor"
description: |-
  Descriptor that identifies the server's redfish integration capability using cimc firmware info.
---

# Data Source: intersight_capability_cimc_firmware_descriptor
Descriptor that identifies the server's redfish integration capability using cimc firmware info.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_capability_cimc_firmware_descriptor.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string) Detailed information about the endpoint. 
* `model`:(string) The model of the endpoint, for which this capability information is applicable. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `revision`:(string) Revision information for the server. 
* `vendor`:(string) The vendor of the endpoint, for which this capability information is applicable. 
* `nr_version`:(string) The firmware or software version of the endpoint, for which this capability information is applicable. 
 