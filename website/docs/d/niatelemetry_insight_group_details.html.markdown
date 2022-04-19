---
subcategory: "niatelemetry"
layout: "intersight"
page_title: "Intersight: intersight_niatelemetry_insight_group_details"
description: |-
        Insight group details in ND.

---

# Data Source: intersight_niatelemetry_insight_group_details
Insight group details in ND.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_niatelemetry_insight_group_details.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `alert_rules_count`:(int) Alert rules count of the Insight group. 
* `analysis_settings_status`:(string) Analysis setting status of the Insight group. 
* `bug_scan_settings_status`:(string) Bug scan settings status of the Insight group. 
* `create_time`:(string) The time when this managed object was created. 
* `delta_analysis_job_count`:(int) Delta analysis job count of the Insight group. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `email_settings_count`:(int) Email settings count of the Insight group. 
* `flow_settings_count`:(int) Flow setting count of the Insight group. 
* `flow_settings_status`:(string) Flow setting status of the Insight group. 
* `group_name`:(string) Name of the Insight group. 
* `kafka_settings_count`:(int) Kafka settings count of the Insight group. 
* `micro_burst_settings_status`:(string) Microburst setting status of the Insight group. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `prechange_analysis_count`:(int) Prechange analysis count of the Insight group. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `tac_collection_config_count`:(int) TAC collection config count of the Insight group. 
 
