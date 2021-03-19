---
subcategory: "niatelemetry"
layout: "intersight"
page_title: "Intersight: intersight_niatelemetry_mso_contract_details"
description: |-
  Details of contract configured from the Multi-Site Orchestrator.
---

# Data Source: intersight_niatelemetry_mso_contract_details
Details of contract configured from the Multi-Site Orchestrator.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_niatelemetry_mso_contract_details.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `contract_name`:(string) Name of contract in Multi-Site Orchestrator. 
* `deployed_sites`:(string) Site Ids to which this contract is deployed to. 
* `is_local`:(string) Is the contract local to the Multi-Site Orchestrator. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `reference`:(string) Unique reference for the contract in the Multi-Site Orchestrator. 
* `schema_id`:(string) Schema ID in Multi-Site Orchestrator. 
* `schema_name`:(string) Schema name this contract belongs to in Multi-Site Orchestrator. 
* `template_name`:(string) Template name this contract belongs to in Multi-Site Orchestrator. 
 