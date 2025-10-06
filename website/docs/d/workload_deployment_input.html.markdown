---
subcategory: "workload"
layout: "intersight"
page_title: "Intersight: intersight_workload_deployment_input"
description: |-
        Store all inputs needed for this deployment along with a unique instance identifier and a generation number to track changes over time. This will include the user provided inputs for the workload definition, merged into the overridden inputs for the workload deployment. The final set of user inputs will be populated into the configuration provided by the blueprints and managed objects representing the polices and profiles will be created which will be used for deployment.

---

# Data Source: intersight_workload_deployment_input
Store all inputs needed for this deployment along with a unique instance identifier and a generation number to track changes over time. This will include the user provided inputs for the workload definition, merged into the overridden inputs for the workload deployment. The final set of user inputs will be populated into the configuration provided by the blueprints and managed objects representing the polices and profiles will be created which will be used for deployment.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_workload_deployment_input.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `gen_number`:(int) A sequential number that increments whenever the input changes. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
