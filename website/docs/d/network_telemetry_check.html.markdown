---
subcategory: "network"
layout: "intersight"
page_title: "Intersight: intersight_network_telemetry_check"
description: |-
        ### Overview
        The TelemetryCheck object is an essential element within network management systems, designed to monitor and verify telemetry configurations across devices. It ensures that telemetry data is correctly configured and operational,
        providing insights into system health and performance.
        #### Purpose
        TelemetryCheck acts as a verification mechanism for network telemetry settings, assessing the status and integrity of data collection processes. It aims to detect configuration failures and provide actionable insights,
        facilitating prompt troubleshooting and resolution.
        #### Key Concepts
        - **Monitoring and Verification** - Continuously checks telemetry configurations for failures, ensuring data integrity and system reliability.
        - **Read-Only Access** - Designed for consumption, providing consistent and secure integration with monitoring systems.
        - **Device Relationships** - Associates with registered devices, ensuring that telemetry checks are performed across all relevant network components.
        - **Status Reporting** - Provides clear indicators of configuration health, enabling swift identification and rectification of issues.
        - **Telemetry Data Check** - Ensures configured telemetry data is validated at the endpoint.

---

# Data Source: intersight_network_telemetry_check
### Overview
The TelemetryCheck object is an essential element within network management systems, designed to monitor and verify telemetry configurations across devices. It ensures that telemetry data is correctly configured and operational,  
providing insights into system health and performance.
#### Purpose
TelemetryCheck acts as a verification mechanism for network telemetry settings, assessing the status and integrity of data collection processes. It aims to detect configuration failures and provide actionable insights,  
facilitating prompt troubleshooting and resolution.
#### Key Concepts
- **Monitoring and Verification** - Continuously checks telemetry configurations for failures, ensuring data integrity and system reliability.
- **Read-Only Access** - Designed for consumption, providing consistent and secure integration with monitoring systems.
- **Device Relationships** - Associates with registered devices, ensuring that telemetry checks are performed across all relevant network components.
- **Status Reporting** - Provides clear indicators of configuration health, enabling swift identification and rectification of issues.
- **Telemetry Data Check** - Ensures configured telemetry data is validated at the endpoint.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_network_telemetry_check.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `status`:(string) Failure status for telemetry configured. 
* `telemetry_config`:(string) The telemetry configuration details from endpoint. 
 
