
---
layout: "intersight"
page_title: "Intersight: intersight_appliance_upgrade_policy"
sidebar_current: "docs-intersight-data-source-appliance-upgrade-policy"
description: |-
UpgradePolicy stores the Intersight Appliance's software upgrade policy. UpgradePolicy
is a sinlgeton managed object. A default upgrade policy is created during the Intersight
Appliance setup, and it is configured with an automatic upgrade policy.
Automatic upgrade policy lets the system start software upgrade after a pre-defined
number of seconds set in the software manifest.
Scheduled upgrade policy lets the user start software upgrade at a specified schedule.
However, scheduled time cannot be beyond the time limit enforced by the upgrade grace
period set in the software manifest.
---

# Data Source: intersight_appliance._upgrade_policy
UpgradePolicy stores the Intersight Appliance's software upgrade policy. UpgradePolicy
is a sinlgeton managed object. A default upgrade policy is created during the Intersight
Appliance setup, and it is configured with an automatic upgrade policy.
Automatic upgrade policy lets the system start software upgrade after a pre-defined
number of seconds set in the software manifest.
Scheduled upgrade policy lets the user start software upgrade at a specified schedule.
However, scheduled time cannot be beyond the time limit enforced by the upgrade grace
period set in the software manifest.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `auto_upgrade`:(bool) Indicates if the upgrade service is set to automatically start the software upgrade or not. If autoUpgrade is true, then the value of the schedule field is ignored. 
* `blackout_dates_enabled`:(bool) If enabled, allows the user to define a blackout period during which the appliance will not be upgraded. 
* `blackout_end_date`:(string) End date of the black out period. 
* `blackout_start_date`:(string) Start date of the black out period. The appliance will not be upgraded during this period. 
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
