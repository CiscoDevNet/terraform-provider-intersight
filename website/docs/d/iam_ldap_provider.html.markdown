---
subcategory: "iam"
layout: "intersight"
page_title: "Intersight: intersight_iam_ldap_provider"
description: |-
  LDAP Provider or LDAP Server for user authentication.
---

# Data Source: intersight_iam_ldap_provider
LDAP Provider or LDAP Server for user authentication.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iam_ldap_provider.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `port`:(int) LDAP Server Port for connection establishment. 
* `server`:(string) LDAP Server Address, can be IP address or hostname. 
 