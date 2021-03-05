---
subcategory: "iam"
layout: "intersight"
page_title: "Intersight: intersight_iam_system"
description: |-
  System is the top level object in the Intersight. All other objects which can be accessed globally are added to system object, like privilege sets and privileges can be shared by multiple roles and privilege sets.
---

# Data Source: intersight_iam_system
System is the top level object in the Intersight. All other objects which can be accessed globally are added to system object, like privilege sets and privileges can be shared by multiple roles and privilege sets.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iam_system.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string) The unique identifier of this Managed Object instance. 
 