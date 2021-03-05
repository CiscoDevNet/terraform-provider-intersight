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
To access the ith object of the results obtained, use `data.intersight_niatelemetry_nexus_dashboards.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `cluster_name`:(string) Nexus Dashboard can onboard multiple APIC clusters/sites. 
* `is_cluster_healthy`:(string) Health of Nexus Dashboard cluster. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `nd_cluster_size`:(int) Number of nodes in Nexus Dashboard cluster. 
* `nd_type`:(string) Node type in Nexus Dashboard cluster. 
* `nd_version`:(string) Version running on Nexus Dashboard. 
* `number_of_apps`:(int) Number of applications installed in the Nexus Dashboard. 
* `number_of_schemas_in_mso`:(int) Number of total schemas in Multi-Site Orchestrator. 
* `number_of_sites_in_mso`:(int) Number of sites in Multi-Site Orchestrator. 
* `number_of_sites_serviced`:(int) Number of sites serviced by ND. 
* `number_of_tenants_in_mso`:(int) Number of total tenants in Multi-Site Orchestrator. 
* `type_of_site_in_mso`:(string) Type of site added to Multi-Site Orchestrator. 
 