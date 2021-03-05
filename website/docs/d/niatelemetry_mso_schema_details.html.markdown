---
subcategory: "niatelemetry"
layout: "intersight"
page_title: "Intersight: intersight_niatelemetry_mso_schema_details"
description: |-
  Details of schema in Multi-Site Orchestrator.
---

# Data Source: intersight_niatelemetry_mso_schema_details
Details of schema in Multi-Site Orchestrator.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_niatelemetry_mso_schema_details.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `number_of_policy_objects_per_schema`:(int) Number of policy objects per schema. 
* `number_of_templates_per_schema`:(int) Number of tenants assigned per schema in Multi-Site Orchestrator. 
* `schema_id`:(string) Schema ID in Multi-Site Orchestrator. 
* `schema_name`:(string) Schema name in Multi-Site Orchestrator. 
 