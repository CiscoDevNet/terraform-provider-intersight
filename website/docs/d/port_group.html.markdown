---
subcategory: "port"
layout: "intersight"
page_title: "Intersight: intersight_port_group"
description: |-
  Holder for multiple ports. A switch card will have one or more port groups.
---

# Data Source: intersight_port_group
Holder for multiple ports. A switch card will have one or more port groups.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_port_group.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `transport`:(string) Type of port group. Values are Eth or Fc. 
 