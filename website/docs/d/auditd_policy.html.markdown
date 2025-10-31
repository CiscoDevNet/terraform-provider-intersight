---
subcategory: "auditd"
layout: "intersight"
page_title: "Intersight: intersight_auditd_policy"
description: |-
        A policy to configure the kernel level AuditD settings in the Fabric Interconnect/Servers.

---

# Data Source: intersight_auditd_policy
A policy to configure the kernel level AuditD settings in the Fabric Interconnect/Servers.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_auditd_policy.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `admin_state`:(string) Admin state for the AuditD feature.* `Disabled` - Admin configured Disabled State.* `Enabled` - Admin configured Enabled State. 
* `auditd_log_level`:(string) The log level for the AuditD feature. The default value is \ notifications\ .* `notifications` - Generated logs are of Notification level, providing information about normal but significant system events requiring awareness.* `emergencies` - Generated logs will be of Emergency log level, indicating a critical and unstable system state.* `alerts` - Generated logs are of Alert level, indicating critical issues needing immediate attention to prevent system disruption.* `critical` - Generated logs are of Critical level, signaling severe issues that may cause system failure if not addressed immediately.* `errors` - Generated logs are of Error level, indicating significant problems that affect functionality but do not cause system failure.* `warnings` - Generated logs are of Warning level, highlighting potential issues that require attention but do not yet impact functionality.* `information` - Generated logs are of Information level, detailing routine operational messages without indicating any issues or errors.* `debugging` - Generated logs are of Debugging level, providing detailed information to help diagnose and troubleshoot system issues. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the policy. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
