---
subcategory: "view"
layout: "intersight"
page_title: "Intersight: intersight_view_health_status"
description: |-
        The HealthStatus object is designed to provide a high-level, aggregated status of Intersight components for a given account user. It informs users about potential issues that may require their attention, offering a comprehensive view of the system's health status.
        #### Purpose
        The HealthStatus object serves as a central indicator of the health of various components within the Intersight ecosystem. It consolidates status information for multiple categories like 'Licensing', 'Advisories', and 'Alarms', making it easier for users to understand the overall condition of their system and take necessary actions.
        #### Key Concepts
        - **Aggregated Status:** Combines the status of multiple Intersight categories, providing a concise snapshot of system health.
        - **Dashboard Integration:** Specifically designed to be easily consumed by external dashboards for at-a-glance status visualization.
        - **Read-Only Access:** Ensures that the status information is reliable and consistent, allowing users to focus on monitoring rather than modifying.

---

# Data Source: intersight_view_health_status
The HealthStatus object is designed to provide a high-level, aggregated status of Intersight components for a given account user. It informs users about potential issues that may require their attention, offering a comprehensive view of the system's health status.
#### Purpose
The HealthStatus object serves as a central indicator of the health of various components within the Intersight ecosystem. It consolidates status information for multiple categories like 'Licensing', 'Advisories', and 'Alarms', making it easier for users to understand the overall condition of their system and take necessary actions.
#### Key Concepts
- **Aggregated Status:** Combines the status of multiple Intersight categories, providing a concise snapshot of system health.
- **Dashboard Integration:** Specifically designed to be easily consumed by external dashboards for at-a-glance status visualization.
- **Read-Only Access:** Ensures that the status information is reliable and consistent, allowing users to focus on monitoring rather than modifying.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_view_health_status.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
