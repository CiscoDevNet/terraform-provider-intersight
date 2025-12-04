---
subcategory: "metrics"
layout: "intersight"
page_title: "Intersight: intersight_metrics_configuration"
description: |-
        Controls the behavior of metric collection for an account, configuration includes options to enable or disable collection as well how new devices claimed to the account are handled.

---

# Data Source: intersight_metrics_configuration
Controls the behavior of metric collection for an account, configuration includes options to enable or disable collection as well how new devices claimed to the account are handled.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_metrics_configuration.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `collect_new_devices`:(string) The behavior of the system when new resources are added, controls whether metric collection are automatically enabled for the new resources.* `AutoEnable` - Automatically enable metric collection for new resources, up to the limit of resource collection.* `Disabled` - Metrics will not be enabled on new resources, to enable collection requires an explicit user enable. 
* `collection_granularity`:(string) The current supported collection granularity by the system, defined as the lowest granularity supported, with the actual granularity per resource determined by the license tier of the resource. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `enabled`:(bool) Enables metric collection for the account, if disabled metrics will be stopped for all resources in the account. 
* `limit`:(int) The total number of resources that can be enabled for metric collection in this account. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `user_max_ingestion_bytes`:(int) User override of the max ingestion rate for metrics. Provided as an advanced option to override the default limits. Option should only be used in coordination with TAC engineers. 
 
