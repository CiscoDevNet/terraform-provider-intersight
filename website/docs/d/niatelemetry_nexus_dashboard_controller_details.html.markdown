---
subcategory: "niatelemetry"
layout: "intersight"
page_title: "Intersight: intersight_niatelemetry_nexus_dashboard_controller_details"
description: |-
        The NexusDashboardControllerDetails object delivers focused information about controllers serviced by Nexus Dashboard, including health, version, and operational status.
        #### Purpose
        NexusDashboardControllerDetails assists administrators in tracking controller health, firmware, and site attributes, providing actionable insights for dashboard-serviced environments.
        #### Key Concepts
        - **Controller Monitoring**: Tracks health and version of controllers within the dashboard ecosystem.
        - **Site-Level Analytics**: Associates controller data with serviced sites for operational clarity.
        - **Inventory Integration**: Enhances the broader Nexus Dashboard inventory with controller-specific insights.

---

# Data Source: intersight_niatelemetry_nexus_dashboard_controller_details
The NexusDashboardControllerDetails object delivers focused information about controllers serviced by Nexus Dashboard, including health, version, and operational status.
#### Purpose
NexusDashboardControllerDetails assists administrators in tracking controller health, firmware, and site attributes, providing actionable insights for dashboard-serviced environments.
#### Key Concepts
- **Controller Monitoring**: Tracks health and version of controllers within the dashboard ecosystem.
- **Site-Level Analytics**: Associates controller data with serviced sites for operational clarity.
- **Inventory Integration**: Enhances the broader Nexus Dashboard inventory with controller-specific insights.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_niatelemetry_nexus_dashboard_controller_details.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `site_health`:(int) Health of the site serviced by ND. 
* `site_name`:(string) Name of fabric domain of the controller. 
* `version_of_controller`:(string) Version of the controller serviced by ND. 
 
