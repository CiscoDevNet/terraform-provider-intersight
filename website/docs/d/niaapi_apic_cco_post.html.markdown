---
subcategory: "niaapi"
layout: "intersight"
page_title: "Intersight: intersight_niaapi_apic_cco_post"
description: |-
  The post reporting a new release is available for APIC.
---

# Data Source: intersight_niaapi_apic_cco_post
The post reporting a new release is available for APIC.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_niaapi_apic_cco_post.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `post_date`:(string) The date when this new release notice is posted. 
* `post_type`:(string) The document type of this post. 
* `postid`:(string) Identificator of this inbox post. 
* `revision`:(string) Revision number of this notice. 
 