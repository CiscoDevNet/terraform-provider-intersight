---
subcategory: "niaapi"
layout: "intersight"
page_title: "Intersight: intersight_niaapi_dcnm_field_notice"
description: |-
  The field notice reporting bug and related software or hardware for DCNM.
---

# Data Source: intersight_niaapi_dcnm_field_notice
The field notice reporting bug and related software or hardware for DCNM.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_niaapi_dcnm_field_notice.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `bugid`:(string) Bug Id associated with this notice. 
* `field_notice_desc`:(string) Field notice Description. 
* `field_notice_id`:(string) Fieldnotice Id of this notice. 
* `field_notice_url`:(string) Field notice URL link to the notice webpage. 
* `headline`:(string) The headline of this field notice. 
* `hwpid`:(string) Hardware PID for affected models. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `sw_release`:(string) Software Release number for affected versions. 
* `workaround_url`:(string) URL of workaround of this notice. 
 