---
subcategory: "iam"
layout: "intersight"
page_title: "Intersight: intersight_iam_privilege"
description: |-
  Privilege represents an action which can be performed in Intersight such as creating server profile, deleting a user etc.
---

# Data Source: intersight_iam_privilege
Privilege represents an action which can be performed in Intersight such as creating server profile, deleting a user etc.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iam_privilege.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `hostname_prefix`:(string) The hostname prefix of the resource corresponding to this privilege. For example \\'sentry\\' in https://sentry.intersight.com . 
* `method`:(string) The API method on the rest resource corresponding to privilege. For example READ, CREATE, UPDATE etc. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the privilege reported by microservice. 
* `rest_path`:(string) The REST API path of the resource corresponding to this privilege. For example /v1/iam/Accounts, /v1/iam/Sessions. 
* `url_prefix`:(string) The URL path prefix of the resource corresponding to this privilege. For example /devops/kibana, /devops/grafana etc. 
 