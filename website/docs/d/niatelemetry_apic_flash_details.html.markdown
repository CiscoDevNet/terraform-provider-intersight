---
subcategory: "niatelemetry"
layout: "intersight"
page_title: "Intersight: intersight_niatelemetry_apic_flash_details"
description: |-
  Object to capture the flash details in APIC.
---

# Data Source: intersight_niatelemetry_apic_flash_details
Object to capture the flash details in APIC.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_niatelemetry_apic_flash_details.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `model_number`:(string) Model number of the flash in APIC. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `record_type`:(string) Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. 
* `record_version`:(string) Version of record being pushed. This determines what was the API version for data available from the device. 
* `revision`:(string) Revision of the flash Mo in APIC. 
* `serial_number`:(string) Serial number of the flash in APIC. 
* `site_name`:(string) Name of the APIC site from which this data is being collected. 
 