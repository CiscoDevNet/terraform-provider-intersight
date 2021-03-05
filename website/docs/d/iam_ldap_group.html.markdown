---
subcategory: "iam"
layout: "intersight"
page_title: "Intersight: intersight_iam_ldap_group"
description: |-
  Mapping of LDAP Group to EndPointRoles.
---

# Data Source: intersight_iam_ldap_group
Mapping of LDAP Group to EndPointRoles.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iam_ldap_group.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `domain`:(string) LDAP server domain the Group resides in. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) LDAP Group name in the LDAP server database. 
 