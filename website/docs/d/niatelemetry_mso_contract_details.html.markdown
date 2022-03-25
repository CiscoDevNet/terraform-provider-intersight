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
To access the ith object of the results obtained, use `data.intersight_niatelemetry_mso_contract_details.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `contract_name`:(string) Name of contract in Multi-Site Orchestrator. 
* `create_time`:(string) The time when this managed object was created. 
* `deployed_sites`:(string) Site Ids to which this contract is deployed to. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `is_local`:(string) Is the contract local to the Multi-Site Orchestrator. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `reference`:(string) Unique reference for the contract in the Multi-Site Orchestrator. 
* `schema_id`:(string) Schema ID in Multi-Site Orchestrator. 
* `schema_name`:(string) Schema name this contract belongs to in Multi-Site Orchestrator. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `template_name`:(string) Template name this contract belongs to in Multi-Site Orchestrator. 
 
