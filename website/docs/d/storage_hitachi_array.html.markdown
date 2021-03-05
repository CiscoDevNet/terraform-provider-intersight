---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_hitachi_array"
description: |-
  The details of the Hitachi storage array.
---

# Data Source: intersight_storage_hitachi_array
The details of the Hitachi storage array.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_hitachi_array.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `ctl1_ip`:(string) IP address of controller 1 of the storage system. 
* `ctl1_micro_version`:(string) GUM (Gateway for Unified Management) version of the controller 1. 
* `ctl2_ip`:(string) IP address of controller 2 of the storage system. 
* `ctl2_micro_version`:(string) GUM (Gateway for Unified Management) version of the controller 2. 
* `device_id`:(string) ID of the Storage device. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `model`:(string) This field identifies the model of the given component. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Administrator defined name for the device. 
* `revision`:(string) This field identifies the revision of the given component. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field identifies the serial of the given component. 
* `svp_ip`:(string) IP address of the SVP. 
* `target_ctl`:(string) Controller operated by the REST API. 
* `uuid`:(string) Unique identity of the device. 
* `vendor`:(string) This field identifies the vendor of the given component. 
* `nr_version`:(string) Current running software version of the device. 
 