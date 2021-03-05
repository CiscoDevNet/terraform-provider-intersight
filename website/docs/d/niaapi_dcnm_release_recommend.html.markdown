---
subcategory: "niaapi"
layout: "intersight"
page_title: "Intersight: intersight_niaapi_dcnm_release_recommend"
description: |-
  The recommend version information for each release on DCNM.
---

# Data Source: intersight_niaapi_dcnm_release_recommend
The recommend version information for each release on DCNM.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_niaapi_dcnm_release_recommend.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `cll`:(string) Current long-lived release. 
* `crr`:(string) Customer recommended releases. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `pid`:(string) Hardware model identificator. 
* `ull`:(string) Upcoming long-lived release. 
 