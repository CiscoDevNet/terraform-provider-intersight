---
subcategory: "niatelemetry"
layout: "intersight"
page_title: "Intersight: intersight_niatelemetry_nexus_dashboards"
description: |-
        Object is available for Nexus Dashboard devices.

---

# Data Source: intersight_niatelemetry_nexus_dashboards
Object is available for Nexus Dashboard devices.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_niatelemetry_nexus_dashboards.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `cluster_name`:(string) Nexus Dashboard can onboard multiple APIC clusters/sites. 
* `cluster_uuid`:(string) UUID of the Nexus Dashboard cluster. 
* `create_time`:(string) The time when this managed object was created. 
* `dn`:(string) Dn of the objects present for Nexus Dashboard devices. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `is_cluster_healthy`:(string) Health of Nexus Dashboard cluster. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `nd_cluster_size`:(int) Number of nodes in Nexus Dashboard cluster. 
* `nd_type`:(string) Node type in Nexus Dashboard cluster. 
* `nd_version`:(string) Version running on Nexus Dashboard. 
* `number_of_apps`:(int) Number of applications installed in the Nexus Dashboard. 
* `number_of_insight_groups`:(int) Number of total insight groups in ND. 
* `number_of_nir_dashboards`:(int) Number of total NIR dashboards in ND. 
* `number_of_schemas_in_mso`:(int) Number of total schemas in Multi-Site Orchestrator. 
* `number_of_sites_in_mso`:(int) Number of sites in Multi-Site Orchestrator. 
* `number_of_sites_serviced`:(int) Number of sites serviced by ND. 
* `number_of_tenants_in_mso`:(int) Number of total tenants in Multi-Site Orchestrator. 
* `number_of_vxlan_fabric_sites_in_mso`:(int) Number of sites with vxLan type fabric in Multi-Site Orchestrator. 
* `record_type`:(string) Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `type_of_site_in_mso`:(string) Type of site added to Multi-Site Orchestrator. 
 
