---
subcategory: "iam"
layout: "intersight"
page_title: "Intersight: intersight_iam_resource_permission"
description: |-
  ResourcePermission represents the permissions defined on a resource like organization.
---

# Data Source: intersight_iam_resource_permission
ResourcePermission represents the permissions defined on a resource like organization.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iam_resource_permission.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `target_app`:(string) Name of the service owning the resource. 
 