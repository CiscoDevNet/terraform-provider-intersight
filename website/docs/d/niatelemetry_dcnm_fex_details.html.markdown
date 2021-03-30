---
subcategory: "niatelemetry"
layout: "intersight"
page_title: "Intersight: intersight_niatelemetry_dcnm_fex_details"
description: |-
  Inventory Object available for DCNM FEX.
---

# Data Source: intersight_niatelemetry_dcnm_fex_details
Inventory Object available for DCNM FEX.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_niatelemetry_dcnm_fex_details.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the fex in the fabric inventory. 
* `product_id`:(string) Product ID of the fex in the fabric inventory. 
* `record_type`:(string) Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. 
* `record_version`:(string) Version of record being pushed. This determines what was the API version for data available from the device. 
* `serial_number`:(string) Serial number of the fex in the fabric inventory. 
* `vendor_id`:(string) Vendor Id of the fex in the fabric inventory. 
 