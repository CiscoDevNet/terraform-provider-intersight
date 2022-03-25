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
To access the ith object of the results obtained, use `data.intersight_niatelemetry_mso_tenant_details.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `deployed_sites`:(string) Site IDs to which this tenant is deployed to. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `number_of_schemas_assigned_per_tenant_in_mso`:(int) Number of schemas assigned to each tenant in Multi-Site Orchestrator. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `sites_each_tenant_is_deployed_to_in_mso`:(int) Number of sites each tenant is deployed to. 
* `tenant_id`:(string) ID of tenant in Multi-Site Orchestrator. 
* `tenant_name`:(string) Name of the tenant in Multi-Site Orchestrator. 
 
