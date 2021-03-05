---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_sas_port"
description: |-
  Sas Port details of the SAS endpoint.
---

# Data Source: intersight_storage_sas_port
Sas Port details of the SAS endpoint.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_sas_port.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `address`:(string) The SAS Address assigned to storage port. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `disk_id`:(int) The unique disk identifier. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `end_point_id`:(int) The end-point Id assigned to storage port. 
* `link_description`:(string) The description for the link. 
* `link_speed`:(string) The link speed negotiated for communication. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
 