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
The following arguments can be used to get data of already created objects in Intersight appliance:
* `active`:(bool) Indicates if the software upgrade is active or not. 
* `auto_created`:(bool) Indicates that the request was automatically created by the system. 
* `description`:(string) Description of the software upgrade. 
* `elapsed_time`:(int) Elapsed time in seconds during the software upgrade. 
* `end_time`:(string) End date of the software upgrade. 
* `error_code`:(int) Error code for Intersight Appliance's software upgrade. In case of failure - this code will help decide if software upgrade can be retried. 
* `fingerprint`:(string) Software upgrade manifest's fingerprint. 
* `is_rolling_back`:(bool) Track if software upgrade is upgrading or rolling back. 
* `is_user_triggered`:(bool) Indicates if the upgrade is triggered by user or due to schedule. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `rollback_needed`:(bool) Track if rollback is needed. 
* `rollback_status`:(string) Status of the Intersight Appliance's software rollback status. 
* `start_time`:(string) Start date of the software upgrade. UI can modify startTime to re-schedule an upgrade. 
* `status`:(string) Status of the Intersight Appliance's software upgrade. 
* `total_phases`:(int) TotalPhase represents the total number of the upgradePhases for one upgrade. 
* `nr_version`:(string) Software upgrade manifest's version. 
