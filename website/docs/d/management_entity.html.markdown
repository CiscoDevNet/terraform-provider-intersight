---
subcategory: "management"
layout: "intersight"
page_title: "Intersight: intersight_management_entity"
description: |-
  Logical representation that captures the role of each Fabric Interconnect in UCS Manager.
---

# Data Source: intersight_management_entity
Logical representation that captures the role of each Fabric Interconnect in UCS Manager.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_management_entity.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `cluster_link_state`:(string) Cluster link state between the Fabric Interconnects. 
* `cluster_readiness`:(string) Cluster readiness of the Fabric Interconnect. 
* `cluster_state`:(string) Cluster state of the Fabric Interconnect. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `entity_id`:(string) Identity of the Fabric Interconnect - A/B. 
* `leadership`:(string) Role (Primary / Subordinate) of the Fabric Interconnect. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
 