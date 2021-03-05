---
subcategory: "iam"
layout: "intersight"
page_title: "Intersight: intersight_iam_security_holder"
description: |-
  Holder for organization aggregated permissions and global account permissions.
User configures permissions for entire account or subset of organizations and specifies associated roles with each organization.
Intersight aggregates all the permissions and stores per organization aggregate permissions in iam.ResourcePermission object.
---

# Data Source: intersight_iam_security_holder
Holder for organization aggregated permissions and global account permissions.
User configures permissions for entire account or subset of organizations and specifies associated roles with each organization.
Intersight aggregates all the permissions and stores per organization aggregate permissions in iam.ResourcePermission object.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iam_security_holder.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string) The unique identifier of this Managed Object instance. 
 