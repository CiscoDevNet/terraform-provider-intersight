---
subcategory: "metrics"
layout: "intersight"
page_title: "Intersight: intersight_metrics_resource_configuration"
description: |-
        The configuration of metric collection for a given resource within an account. Each resource that is capable of collecting metrics into the Intersight system may be configured with options to control the metric collection behavior.

---

# Data Source: intersight_metrics_resource_configuration
The configuration of metric collection for a given resource within an account. Each resource that is capable of collecting metrics into the Intersight system may be configured with options to control the metric collection behavior.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_metrics_resource_configuration.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `enabled`:(bool) Metric collection is enabled for this resource, when enabled all available metrics are collected from the resource into Intersight. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
