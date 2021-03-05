---
subcategory: "uuidpool"
layout: "intersight"
page_title: "Intersight: intersight_uuidpool_uuid_lease"
description: |-
  UuidLease represents a single UUID that is part of the universe, allocated either from a pool or through static assignment.
---

# Data Source: intersight_uuidpool_uuid_lease
UuidLease represents a single UUID that is part of the universe, allocated either from a pool or through static assignment.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_uuidpool_uuid_lease.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `uuid`:(string) UUID Prefix+Suffix numbers. 
 