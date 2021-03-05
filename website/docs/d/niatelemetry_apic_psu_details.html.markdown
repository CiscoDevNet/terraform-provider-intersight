---
subcategory: "niatelemetry"
layout: "intersight"
page_title: "Intersight: intersight_niatelemetry_apic_psu_details"
description: |-
  Object to capture PSU details in APIC.
---

# Data Source: intersight_niatelemetry_apic_psu_details
Object to capture PSU details in APIC.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_niatelemetry_apic_psu_details.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `dn`:(string) Dn for the PSU in the inventory. 
* `model_number`:(string) Model number of the PSU in APIC. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `node_id`:(string) Node id for the PSU in the inventory. 
* `record_type`:(string) Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. 
* `record_version`:(string) Version of record being pushed. This determines what was the API version for data available from the device. 
* `serial_number`:(string) Serial number of the PSU in APIC. 
* `site_name`:(string) Name of the APIC site from which this data is being collected. 
* `vendor_name`:(string) Vendor name of the PSU in APIC. 
 