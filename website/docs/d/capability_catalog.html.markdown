---
subcategory: "capability"
layout: "intersight"
page_title: "Intersight: intersight_capability_catalog"
description: |-
  Container for capability information of managed systems.
This catalog will be managed by devops using a specific role in the Catalog Admin account.
---

# Data Source: intersight_capability_catalog
Container for capability information of managed systems.
This catalog will be managed by devops using a specific role in the Catalog Admin account.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_capability_catalog.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) A unique name for the catalog. 
 