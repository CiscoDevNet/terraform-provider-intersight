---
subcategory: "niaapi"
layout: "intersight"
page_title: "Intersight: intersight_niaapi_version_regex"
description: |-
  The regular expression pattern to recongnize the version string.
---

# Data Source: intersight_niaapi_version_regex
The regular expression pattern to recongnize the version string.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_niaapi_version_regex.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `nr_version`:(string) Version number for the Version Regex data, also used as identity. 
 