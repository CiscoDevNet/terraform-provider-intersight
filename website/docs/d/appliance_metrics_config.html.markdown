---
subcategory: "appliance"
layout: "intersight"
page_title: "Intersight: intersight_appliance_metrics_config"
description: |-
        MetricsConfig provides system configuration parameters for managing metrics collection on appliance.

---

# Data Source: intersight_appliance_metrics_config
MetricsConfig provides system configuration parameters for managing metrics collection on appliance.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_appliance_metrics_config.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `current_endpoint_count`:(int) Number of discovered endpoints from where metrics is being collected currently. 
* `disk_capacity`:(int) Capacity of the metrics disk /opt/database in bytes. 
* `disk_usage`:(int) Disk usage of the metrics disk /opt/database in bytes. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `endpoint_usage_percent`:(int) Usage percentage of the discovered endpoints. 
* `last_disabled_date`:(string) Disabled date of the metrics collection feature. 
* `last_enabled_date`:(string) Enabled date of the metrics collection feature. 
* `max_endpoint_count`:(int) The maximum number of supported endpoints for an appliance deployment type. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `status_message`:(string) The overall metrics collection Status based on resource constraints. 
* `system_enabled`:(bool) Metric collection state defined by the system. 
* `user_enabled`:(bool) Configured metric collection state by the account administrator. 
 
