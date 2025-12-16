---
subcategory: "appliance"
layout: "intersight"
page_title: "Intersight: intersight_appliance_group_op_status"
description: |-
        The GroupOpStatus object describes the status of a group of applications within the Intersight Appliance system. It consolidates application-specific status to provide a unified view of group health.
        #### Purpose
        GroupOpStatus aggregates application status information to offer insights into the collective health and functionality of system features, aiding in efficient management and troubleshooting.
        #### Key Concepts
        - **Group Health:** Represents the overall status of applications grouped by features like Identity Management and Core Service.
        - **Application Aggregation:** Collates individual application statuses, providing a consolidated view for easier monitoring and intervention.
        - **Operational Insights:** Offers descriptions and status indicators vital for maintaining system operations and performance.

---

# Data Source: intersight_appliance_group_op_status
The GroupOpStatus object describes the status of a group of applications within the Intersight Appliance system. It consolidates application-specific status to provide a unified view of group health.
#### Purpose
GroupOpStatus aggregates application status information to offer insights into the collective health and functionality of system features, aiding in efficient management and troubleshooting.
#### Key Concepts
- **Group Health:** Represents the overall status of applications grouped by features like Identity Management and Core Service.
- **Application Aggregation:** Collates individual application statuses, providing a consolidated view for easier monitoring and intervention.
- **Operational Insights:** Offers descriptions and status indicators vital for maintaining system operations and performance.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_appliance_group_op_status.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the group. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `group_name`:(string) The name of group, which includes Identity Management, Device Connector Service, Core Service, Analytics, Infra and Appliance. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `overall_status`:(string) The overall API status from this group's applications. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
