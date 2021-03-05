---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_net_app_cluster"
description: |-
  NetApp cluster consists of one or more nodes grouped together as HA pairs to form a scalable cluster.
---

# Data Source: intersight_storage_net_app_cluster
NetApp cluster consists of one or more nodes grouped together as HA pairs to form a scalable cluster.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_net_app_cluster.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `location`:(string) Location of the storage controller. 
* `management_address`:(string) FQDN or IP Address of Storage Cluster. 
* `model`:(string) This field identifies the model of the given component. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Administrator defined name for the device. 
* `revision`:(string) This field identifies the revision of the given component. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field identifies the serial of the given component. 
* `uuid`:(string) Unique identity of the device. 
* `vendor`:(string) This field identifies the vendor of the given component. 
* `nr_version`:(string) Current running software version of the device. 
 