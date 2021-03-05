---
subcategory: "top"
layout: "intersight"
page_title: "Intersight: intersight_top_system"
description: |-
  Root container for all UCSM / CIMC MOs.
---

# Data Source: intersight_top_system
Root container for all UCSM / CIMC MOs.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_top_system.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `ipv4_address`:(string) The IPv4 address of system. 
* `ipv6_address`:(string) The IPv6 address of system. 
* `mode`:(string) The current mode of the system. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The admin configured name of the system. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `time_zone`:(string) The operational timezone of the system, empty indicates no timezone has been set specifically. 
 