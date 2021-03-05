---
subcategory: "iam"
layout: "intersight"
page_title: "Intersight: intersight_iam_resource_roles"
description: |-
  ResourceRoles provides a way to specify the roles associated with a resource like organization in a permission which can be assigned to a user or user group.
---

# Data Source: intersight_iam_resource_roles
ResourceRoles provides a way to specify the roles associated with a resource like organization in a permission which can be assigned to a user or user group.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iam_resource_roles.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string) The unique identifier of this Managed Object instance. 
 