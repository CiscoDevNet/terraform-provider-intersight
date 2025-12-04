---
subcategory: "appliance"
layout: "intersight"
page_title: "Intersight: intersight_appliance_metrics_ingestion_utilization"
description: |-
        Measures the utilization of the metrics feature in relation to the available capacity of the system.

---

# Data Source: intersight_appliance_metrics_ingestion_utilization
Measures the utilization of the metrics feature in relation to the available capacity of the system.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_appliance_metrics_ingestion_utilization.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `end_time`:(string) The end of the measurement window. 
* `meta_field`:(string) Uniquely identifies the source of the metric. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `time`:(string) The start of the measurement window. 
* `utilization`:(float) The percent utilization of the metric. 
 
