---
subcategory: "niatelemetry"
layout: "intersight"
page_title: "Intersight: intersight_niatelemetry_mso_epg_details"
description: |-
        The MsoEpgDetails object provides detailed information about Endpoint Groups (EPGs) configured via Multi-Site Orchestrator.
        #### Purpose
        MsoEpgDetails enables comprehensive tracking and management of EPG deployments, templates, and site associations.
        #### Key Concepts
        - **EPG Analytics**: Documents EPG names, schema associations, templates, and deployment sites.
        - **Reference Management**: Supports unique referencing of EPGs across orchestrator environments.
        - **Operational Visibility**: Enhances EPG management and troubleshooting.

---

# Data Source: intersight_niatelemetry_mso_epg_details
The MsoEpgDetails object provides detailed information about Endpoint Groups (EPGs) configured via Multi-Site Orchestrator.
#### Purpose
MsoEpgDetails enables comprehensive tracking and management of EPG deployments, templates, and site associations.
#### Key Concepts
- **EPG Analytics**: Documents EPG names, schema associations, templates, and deployment sites.
- **Reference Management**: Supports unique referencing of EPGs across orchestrator environments.
- **Operational Visibility**: Enhances EPG management and troubleshooting.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_niatelemetry_mso_epg_details.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `deployed_sites`:(string) Site Ids to which this EPG is deployed to. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `epg_name`:(string) Name of EPG in Multi-Site Orchestrator. 
* `is_local`:(string) Is the EPG local to the Multi-Site Orchestrator. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `reference`:(string) Unique reference for the EPG in the Multi-Site Orchestrator. 
* `schema_id`:(string) Schema ID in Multi-Site Orchestrator. 
* `schema_name`:(string) Schema name in Multi-Site Orchestrator. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `template_name`:(string) Template name in Multi-Site Orchestrator. 
 
