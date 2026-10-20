---
subcategory: "niatelemetry"
layout: "intersight"
page_title: "Intersight: intersight_niatelemetry_mso_site_details"
description: |-
        The MsoSiteDetails object centralizes site management within Multi-Site Orchestrator environments, capturing site IDs, names, versions, and operational metrics.
        #### Purpose
        MsoSiteDetails provides visibility into site deployment, configuration, and operational state across orchestrated networks.
        #### Key Concepts
        - **Site Inventory**: Tracks site IDs, names, controller versions, and operational metrics.
        - **Operational Analytics**: Supports planning, monitoring, and troubleshooting site deployments.
        - **CloudSec Integration**: Documents cloud security enablement across orchestrator sites.

---

# Data Source: intersight_niatelemetry_mso_site_details
The MsoSiteDetails object centralizes site management within Multi-Site Orchestrator environments, capturing site IDs, names, versions, and operational metrics.
#### Purpose
MsoSiteDetails provides visibility into site deployment, configuration, and operational state across orchestrated networks.
#### Key Concepts
- **Site Inventory**: Tracks site IDs, names, controller versions, and operational metrics.
- **Operational Analytics**: Supports planning, monitoring, and troubleshooting site deployments.
- **CloudSec Integration**: Documents cloud security enablement across orchestrator sites.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_niatelemetry_mso_site_details.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `is_cloud_sec_enabled`:(string) Status of cloudSec on Multi-Site Orchestrator site. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `number_of_leafs_per_site_in_mso`:(int) Number of leafs per site in Multi-Site Orchestrator. 
* `number_of_pods_per_site_in_mso`:(int) Number of pods per site in Multi-Site Orchestrator. 
* `number_of_spines_per_site_in_mso`:(int) Number of spines per site in Multi-Site Orchestrator. 
* `record_type`:(string) Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `site_id`:(string) ID of site in Multi-Site Orchestrator. 
* `site_name`:(string) Name of the site in Multi-Site Orchestrator. 
* `site_version`:(string) Version of the controller in the site. 
 
