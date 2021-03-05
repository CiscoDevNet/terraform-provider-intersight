---
subcategory: "inventory"
layout: "intersight"
page_title: "Intersight: intersight_inventory_generic_inventory_holder"
description: |-
  A container class for generic inventory.
---

# Data Source: intersight_inventory_generic_inventory_holder
A container class for generic inventory.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_inventory_generic_inventory_holder.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `endpoint`:(string) The endpoint represented by this holder. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
 