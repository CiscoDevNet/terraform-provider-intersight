
---
layout: "intersight"
page_title: "Intersight: intersight_appliance_upgrade"
sidebar_current: "docs-intersight-data-source-appliance-upgrade"
description: |-
Upgrade tracks the Intersight Appliance's software upgrades. Intersight Appliance's
upgrade service automatically creates an Upgrade managed object when there is a
pending software upgrade. User may modify an active Upgrade managed object to reset
the software upgrade start time. However, user may not be able to postpone an upgrade
beyond the limit enforced by the upgrade grace period set in the software manifest.
---

# Data Source: intersight_appliance._upgrade
Upgrade tracks the Intersight Appliance's software upgrades. Intersight Appliance's
upgrade service automatically creates an Upgrade managed object when there is a
pending software upgrade. User may modify an active Upgrade managed object to reset
the software upgrade start time. However, user may not be able to postpone an upgrade
beyond the limit enforced by the upgrade grace period set in the software manifest.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `active`:(bool) Indicates if the software upgrade is active or not. 
* `auto_created`:(bool) Indicates that the request was automatically created by the system. 
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `description`:(string) Description of the software upgrade. 
* `elapsed_time`:(int) Elapsed time in seconds during the software upgrade. 
* `end_time`:(string) End date of the software upgrade. 
* `error_code`:(int) Error code for Intersight Appliance's software upgrade. In case of failure - this code will help decide if software upgrade can be retried. 
* `fingerprint`:(string) Software upgrade manifest's fingerprint. 
* `is_rolling_back`:(bool) Track if software upgrade is upgrading or rolling back. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `rollback_needed`:(bool) Track if rollback is needed. 
* `rollback_status`:(string) Status of the Intersight Appliance's software rollback status. 
* `start_time`:(string) Start date of the software upgrade. UI can modify startTime to re-schedule an upgrade. 
* `status`:(string) Status of the Intersight Appliance's software upgrade. 
* `total_phases`:(int) TotalPhase represents the total number of the upgradePhases for one upgrade. 
* `version`:(string) Software upgrade manifest's version. 
