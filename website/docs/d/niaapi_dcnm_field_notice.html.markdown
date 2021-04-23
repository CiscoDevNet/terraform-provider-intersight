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
To access the ith object of the results obtained, use `data.intersight_niaapi_dcnm_field_notice.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `bugid`:(string) Bug Id associated with this notice. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `field_notice_desc`:(string) Field notice Description. 
* `field_notice_id`:(string) Fieldnotice Id of this notice. 
* `field_notice_url`:(string) Field notice URL link to the notice webpage. 
* `headline`:(string) The headline of this field notice. 
* `hwpid`:(string) Hardware PID for affected models. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `sw_release`:(string) Software Release number for affected versions. 
* `workaround_url`:(string) URL of workaround of this notice. 
 