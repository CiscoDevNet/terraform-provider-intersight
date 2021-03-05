---
subcategory: "resource"
layout: "intersight"
page_title: "Intersight: intersight_resource_membership"
description: |-
  ResourceMembership represents a resource's associated groups, organizations and the permissions associated to this resource via organizations.
---

# Data Source: intersight_resource_membership
ResourceMembership represents a resource's associated groups, organizations and the permissions associated to this resource via organizations.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_resource_membership.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `target_app`:(string) Name of the Service owning the resource. 
 