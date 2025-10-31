---
subcategory: "terminal"
layout: "intersight"
page_title: "Intersight: intersight_terminal_audit_log"
description: |-
        ### Overview
        The AuditLog object captures and records the audit trail of remote terminal user sessions. This provides a comprehensive history of user interactions with terminal sessions, ensuring accountability and traceability of actions performed.
        #### Purpose
        The primary purpose of an AuditLog is to maintain a detailed record of all remote terminal user sessions. This record is crucial for security, compliance, and troubleshooting, allowing administrators to review who accessed what, when, and for how long, thereby ensuring operational transparency and accountability.
        #### Key Concepts
        - **Session Tracking** - Records key events and timestamps related to the opening and closing of terminal sessions.
        - **Accountability** - The audit log object records only the session metadata, specifically the start and end times for a given user and device. It does not capture or retain the detailed contents of the session itself.
        - **Security and Compliance** - Supports auditing requirements by offering a clear and immutable log of session activities.

---

# Data Source: intersight_terminal_audit_log
### Overview  
The AuditLog object captures and records the audit trail of remote terminal user sessions. This provides a comprehensive history of user interactions with terminal sessions, ensuring accountability and traceability of actions performed.
#### Purpose
The primary purpose of an AuditLog is to maintain a detailed record of all remote terminal user sessions. This record is crucial for security, compliance, and troubleshooting, allowing administrators to review who accessed what, when, and for how long, thereby ensuring operational transparency and accountability.
#### Key Concepts
- **Session Tracking** - Records key events and timestamps related to the opening and closing of terminal sessions. 
- **Accountability** - The audit log object records only the session metadata, specifically the start and end times for a given user and device. It does not capture or retain the detailed contents of the session itself. 
- **Security and Compliance** - Supports auditing requirements by offering a clear and immutable log of session activities.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_terminal_audit_log.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `end_time`:(string) The time the terminal was closed. If terminal has not closed, value is zero time. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `start_time`:(string) The time the terminal session was opened. 
 
