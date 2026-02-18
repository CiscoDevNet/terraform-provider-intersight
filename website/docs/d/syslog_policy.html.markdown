---
subcategory: "syslog"
layout: "intersight"
page_title: "Intersight: intersight_syslog_policy"
description: |-
        The Syslog Policy is a reusable policy for configuring how system log messages are handled on an endpoint.
        #### Purpose
        The purpose of this policy is to centralize and standardize the logging configuration for managed devices. It allows administrators to forward system logs to one or more remote syslog servers for aggregation, analysis, and long-term storage, which is essential for monitoring, troubleshooting, and security auditing.
        #### Key Concepts
        - **Remote and Local Logging:** The policy supports configuring both remote logging clients (to forward logs to a central server) and local logging clients (to control how logs are handled on the device itself).
        - **Multiple Destinations:** Administrators can define a list of remote syslog servers, each with its own specific configuration.
        - **Flexible Transport:** For remote logging, the policy supports both UDP and TCP as transport protocols.
        - **Severity Filtering:** It allows for the configuration of a minSeverity level for both local and remote logs, enabling administrators to filter out less important messages and reduce noise.
        - **Profile-Based Application:** The policy is attached to a Server Profile to apply a consistent logging configuration across multiple servers.

---

# Data Source: intersight_syslog_policy
The Syslog Policy is a reusable policy for configuring how system log messages are handled on an endpoint.
#### Purpose
The purpose of this policy is to centralize and standardize the logging configuration for managed devices. It allows administrators to forward system logs to one or more remote syslog servers for aggregation, analysis, and long-term storage, which is essential for monitoring, troubleshooting, and security auditing.
#### Key Concepts
- **Remote and Local Logging:** The policy supports configuring both remote logging clients (to forward logs to a central server) and local logging clients (to control how logs are handled on the device itself).
- **Multiple Destinations:** Administrators can define a list of remote syslog servers, each with its own specific configuration.
- **Flexible Transport:** For remote logging, the policy supports both UDP and TCP as transport protocols.
- **Severity Filtering:** It allows for the configuration of a minSeverity level for both local and remote logs, enabling administrators to filter out less important messages and reduce noise.
- **Profile-Based Application:** The policy is attached to a Server Profile to apply a consistent logging configuration across multiple servers.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_syslog_policy.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the policy. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
