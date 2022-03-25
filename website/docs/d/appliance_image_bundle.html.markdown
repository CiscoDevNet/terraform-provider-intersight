---
subcategory: "appliance"
layout: "intersight"
page_title: "Intersight: intersight_appliance_image_bundle"
description: |-
        ImageBundle keeps track of all the software bundles installed in the Intersight Appliances.
        Each ImageBundle managed object is derived from a software upgrade manifest. ImageBundle has
        additional properties computed during the manifest processing. Additional properties are the
        dynamic attributes of the software packages declared in the software manifest. For example,
        SHA256 values of the software packages are computed during the software manifest processing.
        An ImageBundle managed object named 'current' is always present in the Intersight Appliance.
        The software upgrade service creates another ImageBundle managed object named 'pending'
        when there is a pending software upgrade. The upgrade service renames the 'pending' bundle
        to the 'current' bundle after the software upgrade is successful.

---

# Data Source: intersight_appliance_image_bundle
ImageBundle keeps track of all the software bundles installed in the Intersight Appliances.
Each ImageBundle managed object is derived from a software upgrade manifest. ImageBundle has
additional properties computed during the manifest processing. Additional properties are the
dynamic attributes of the software packages declared in the software manifest. For example,
SHA256 values of the software packages are computed during the software manifest processing.
An ImageBundle managed object named 'current' is always present in the Intersight Appliance.
The software upgrade service creates another ImageBundle managed object named 'pending'
when there is a pending software upgrade. The upgrade service renames the 'pending' bundle
to the 'current' bundle after the software upgrade is successful.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_appliance_image_bundle.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `auto_upgrade`:(bool) Indicates that the software upgrade was automatically initiated by the Intersight Appliance. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Short description of the software upgrade bundle. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `fingerprint`:(string) Fingerprint of the software manifest from which this bundle is created. Fingerprint is calculated using the SHA256 algorithm. 
* `has_error`:(bool) Indicates that the ImageBundle has errors. The upgrade service sets this field when it encounters errors during the manifest processing. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the software upgrade bundle. 
* `notes`:(string) Detailed description of the software upgrade bundle. 
* `priority`:(string) Software upgrade manifest's upgrade priority. The upgrade service supports two priorities, Normal and Critical. Normal priority is used for regular software upgrades, and the upgrade service uses the Upgrade Policy to compute upgrade start time. Critical priority is used for the critical software security patches, and the upgrade service ignores the Upgrade Policy when it computes the upgrade start time.* `Normal` - Normal upgrade priority is used for all the software upgrades except for the critical security updates. The upgrade service of Intersight Appliance uses the Software Upgrade Policy settings to start the upgrade process.* `Critical` - Critical upgrade priority is used for critical updates such as security patches. The upgrade service of the Intersight Appliance starts the upgrade as specified by the upgrade properties in the software manifest file. The upgrade service will not use the settings specified in the Software Upgrade Policy. 
* `release_time`:(string) Software upgrade manifest's release date and time. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `status_message`:(string) Status message set during the manifest processing. 
* `upgrade_end_time`:(string) End date of the software upgrade process. 
* `upgrade_grace_period`:(int) Grace period in seconds before the automatic upgrade is initiated. The upgrade service uses the grace period to compute the upgrade start time when it receives an upgrade notfication from the Intersight. If there is an Upgrade Policy configured for the Intersight Appliance, then the upgrade service uses the policy to compute the upgrade start time. However, the upgrade start time cannot not exceed the limit enforced by the grace period. 
* `upgrade_impact_duration`:(int) Duration (in minutes) for which services will be disrupted. 
* `upgrade_impact_enum`:(string) UpgradeImpactEnum is used to indicate the kind of impact the upgrade has on currently running services on the appliance.* `None` - The upgrade has no effect on the system.* `Disruptive` - The services will not be functional during the upgrade.* `Disruptive-reboot` - The upgrade needs a reboot. 
* `upgrade_start_time`:(string) Start date of the software upgrade process. 
* `nr_version`:(string) Software upgrade manifest's version. 
 
