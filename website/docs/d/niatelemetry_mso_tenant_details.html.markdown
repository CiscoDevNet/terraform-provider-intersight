---
subcategory: "niatelemetry"
layout: "intersight"
page_title: "Intersight: intersight_niatelemetry_mso_tenant_details"
description: |-
  Details of tenant in Multi-Site Orchestrator.
---

# Data Source: intersight_niatelemetry_mso_tenant_details
Details of tenant in Multi-Site Orchestrator.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_niatelemetry_mso_tenant_details.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `number_of_schemas_assigned_per_tenant_in_mso`:(int) Number of schemas assigned to each tenant in Multi-Site Orchestrator. 
* `sites_each_tenant_is_deployed_to_in_mso`:(int) Number of sites each tenant is deployed to. 
* `tenant_id`:(string) ID of tenant in Multi-Site Orchestrator. 
* `tenant_name`:(string) Name of the tenant in Multi-Site Orchestrator. 
 