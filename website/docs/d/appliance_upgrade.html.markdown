---
subcategory: "appliance"
layout: "intersight"
page_title: "Intersight: intersight_appliance_upgrade"
description: |-
        Upgrade tracks the Intersight Appliance's software upgrades. Intersight Appliance's
        upgrade service automatically creates an Upgrade managed object when there is a
        pending software upgrade. User may modify an active Upgrade managed object to reset
        the software upgrade start time. However, user may not be able to postpone an upgrade
        beyond the limit enforced by the upgrade grace period set in the software manifest.

---

# Data Source: intersight_appliance_upgrade
Upgrade tracks the Intersight Appliance's software upgrades. Intersight Appliance's
upgrade service automatically creates an Upgrade managed object when there is a
pending software upgrade. User may modify an active Upgrade managed object to reset
the software upgrade start time. However, user may not be able to postpone an upgrade
beyond the limit enforced by the upgrade grace period set in the software manifest.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_appliance_upgrade.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `active`:(bool) Indicates if the software upgrade is active or not. 
* `all_nodes_pingable`:(bool) True if all nodes in cluster are pingable, otherwise false. 
* `auto_created`:(bool) Indicates that the request was automatically created by the system. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the software upgrade. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `elapsed_time`:(int) Elapsed time in seconds during the software upgrade. 
* `end_time`:(string) End date of the software upgrade. 
* `error_code`:(int) Error code for Intersight Appliance's software upgrade. In case of failure - this code will help decide if software upgrade can be retried. 
* `fingerprint`:(string) Software upgrade manifest's fingerprint. 
* `is_rolling_back`:(bool) Track if software upgrade is upgrading or rolling back. 
* `is_user_triggered`:(bool) Indicates if the upgrade is triggered by user or due to schedule. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `previous_install_attempts`:(int) The number of previous failed install attempts of the same upgrade version. 
* `rollback_needed`:(bool) Track if rollback is needed. 
* `rollback_status`:(string) Status of the Intersight Appliance's software rollback status. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `start_time`:(string) Start date of the software upgrade. UI can modify startTime to re-schedule an upgrade. 
* `status`:(string) Status of the Intersight Appliance's software upgrade. 
* `total_nodes`:(int) Total number of nodes this upgrade will run on. 
* `total_phases`:(int) TotalPhase represents the total number of the upgradePhases for one upgrade. 
* `nr_version`:(string) Software upgrade manifest's version. 
 
