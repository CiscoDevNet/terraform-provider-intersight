---
subcategory: "iaas"
layout: "intersight"
page_title: "Intersight: intersight_iaas_service_request"
description: |-
  Gets last six months Service Requests from UCSD.
---

# Data Source: intersight_iaas_service_request
Gets last six months Service Requests from UCSD.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iaas_service_request.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `duration`:(string) Service request duration. 
* `initiating_user`:(string) Service Request initiating user. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `request_end_time`:(string) Service request end time. 
* `request_id`:(string) Service request id of an SR. 
* `request_start_time`:(string) Service request start time. 
* `request_type`:(string) Service request type of an SR. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `status`:(string) UCSD service request status. 
* `workflow_name`:(string) Executed workflow name for an SR. 
 