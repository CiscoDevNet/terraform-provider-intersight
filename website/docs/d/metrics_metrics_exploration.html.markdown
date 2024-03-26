---
subcategory: "metrics"
layout: "intersight"
page_title: "Intersight: intersight_metrics_metrics_exploration"
description: |-
        Contains info regarding metrics query and templating information.

---

# Data Source: intersight_metrics_metrics_exploration
Contains info regarding metrics query and templating information.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_metrics_metrics_exploration.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) User specified description of this MetricsExploration. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `granularity`:(string) The time unit in which the metrics will be aggregated into. 
* `is_widget`:(bool) True when this MetricsExploration is exposed as a Dashlet widget. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) User specified name of this MetricsExploration. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
