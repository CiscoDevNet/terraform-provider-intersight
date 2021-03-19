---
subcategory: "niatelemetry"
layout: "intersight"
page_title: "Intersight: intersight_niatelemetry_lc"
description: |-
  Object is available at Line Card scope.
---

# Data Source: intersight_niatelemetry_lc
Object is available at Line Card scope.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_niatelemetry_lc.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string) Description of the line cards present. 
* `dn`:(string) Dn value for the line cards present. 
* `hardware_version`:(string) Hardware version of the line cards present. 
* `model`:(string) Model of the line cards present. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `node_id`:(int) Node Id of the line card present. 
* `operational_state`:(string) Opretaional state of the line cards present. 
* `power_state`:(string) Power state of the line cards present. 
* `record_type`:(string) Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. 
* `record_version`:(string) Version of record being pushed. This determines what was the API version for data available from the device. 
* `redundancy_state`:(string) Redundancy state of the line cards present. 
* `serial_number`:(string) Serial number of the line card present. 
* `site_name`:(string) The Site name represents an APIC cluster. Service Engine can onboard multiple APIC clusters / sites. 
 