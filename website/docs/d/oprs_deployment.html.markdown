---
subcategory: "oprs"
layout: "intersight"
page_title: "Intersight: intersight_oprs_deployment"
description: |-
  Monitors the status of operator deployed in the assist.
---

# Data Source: intersight_oprs_deployment
Monitors the status of operator deployed in the assist.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_oprs_deployment.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `available_replicas`:(int) Available number of replicas. 
* `create_time`:(string) The time when this managed object was created. 
* `desired_replicas`:(int) The expected number of replicas. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `event`:(string) The type of event which was triggered. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Agent name for which the event is triggered. 
* `namespace`:(string) Name space in which the agents are running. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `status`:(string) Status which shows if the resource is healthy or not.* `` - An Unknown status indicates that the resource status is not known.* `Healthy` - A healthy status indicates that the resource is healthy and running as per spec.* `Unhealthy` - An unhealthy status indicates that the resource is down. 
* `time_stamp`:(string) The time at which the event was generated. Date is accurate to Intersights clock. This time will be used to identify order of events. 
* `unavailable_replicas`:(int) Number of replicas Unavailable. 
 