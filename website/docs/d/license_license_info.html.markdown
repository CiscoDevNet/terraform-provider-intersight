---
subcategory: "license"
layout: "intersight"
page_title: "Intersight: intersight_license_license_info"
description: |-
        License state information for a specific license entitlement. Essentials license entitlement is supported currently.
        licenseState attribute is used for license enforcement. When license state is one of TrialPeriod, Compliance, or OutOfCompliance,
        the feature set defined for the license entitlement is granted to the customer.

---

# Data Source: intersight_license_license_info
License state information for a specific license entitlement. Essentials license entitlement is supported currently.
licenseState attribute is used for license enforcement. When license state is one of TrialPeriod, Compliance, or OutOfCompliance,
the feature set defined for the license entitlement is granted to the customer.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_license_license_info.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `active_admin`:(bool) The license administrative state.Set this property to 'true' to activate the license entitlements. 
* `balance`:(int) The total balance we have for licenses. 
* `create_time`:(string) The time when this managed object was created. 
* `days_left`:(int) The number of days left for licenseState to stay in TrialPeriod or OutOfCompliance state. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `end_time`:(string) The date and time when the trial period expires.The value of the 'endTime' property is set when the account enters the TrialPeriod or OutOfCompliance state. 
* `enforce_mode`:(string) The entitlement mode reported by Cisco Smart Software Manager. 
* `error_desc`:(string) The detailed error message when there is any error related to this licensing entitlement. 
* `evaluation_period`:(int) The default Trial or Grace period customer is entitled to. 
* `expire_time`:(string) The date and time when the next expiration time of license subscription. 
* `extra_evaluation`:(int) The number of days the trial Trial or Grace period is extended.The trial or grace period can be extended once. 
* `license_count`:(int) The total number of license consumed in the Intersight account. 
* `license_count_purchased`:(int) The total number of license purchased from cisco. 
* `license_state`:(string) The license state defined by Intersight.The value may be one of NotLicensed, TrialPeriod, OutOfCompliance, Compliance, GraceExpired, or TrialExpired.* `NotLicensed` - The license token is neither activated nor registered.* `GraceExpired` - The license grace period has expired.* `TrialPeriod` - The 90 days of trial period.* `OutOfCompliance` - The license is out of compliance.* `Compliance` - The license is in compliance.* `TrialExpired` - The trial period of 90 days has expired. 
* `license_type`:(string) The name of the Intersight license entitlement.For example, this property may be set to 'Essential'.* `Base` - Base as a License type. It is default license type.* `Essential` - Essential as a License type.* `Standard` - Standard as a License type.* `Advantage` - Advantage as a License type.* `Premier` - Premier as a License type.* `IWO-Essential` - IWO-Essential as a License type.* `IWO-Advantage` - IWO-Advantage as a License type.* `IWO-Premier` - IWO-Premier as a License type.* `IKS-Advantage` - IKS-Advantage as a License type.* `INC-Premier-1GFixed` - Premier 1G Fixed license tier for Intersight Nexus Cloud.* `INC-Premier-10GFixed` - Premier 10G Fixed license tier for Intersight Nexus Cloud.* `INC-Premier-100GFixed` - Premier 100G Fixed license tier for Intersight Nexus Cloud.* `INC-Premier-Mod4Slot` - Premier Modular 4 slot license tier for Intersight Nexus Cloud.* `INC-Premier-Mod8Slot` - Premier Modular 8 slot license tier for Intersight Nexus Cloud.* `INC-Premier-D2OpsFixed` - Premier D2Ops fixed license tier for Intersight Nexus Cloud.* `INC-Premier-D2OpsMod` - Premier D2Ops modular license tier for Intersight Nexus Cloud.* `IntersightTrial` - Virtual dummy license type to indicate trial. Used for UI display of trial mode Intersight tiers.* `IWOTrial` - Virtual dummy license type to indicate trial. Used for UI display of trial mode IKS tiers.* `IKSTrial` - Virtual dummy license type to indicate trial. Used for UI display of trial mode IWO tiers.* `INCTrial` - Virtual dummy license type to indicate trial. Used for UI display of trial mode Nexus tiers. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `net_substitution`:(int) The total number of substituted licenses added or removed. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `start_time`:(string) The date and time when the licenseState entered the TrialPeriod or OutOfCompliance state. 
* `subscription_id`:(string) The id of license subscription. 
* `trial_admin`:(bool) The administrative state of the trial license.When the LicenseState is set to 'NotLicensed', 'trialAdmin' can be set to true to start the trial period,i.e. licenseState is set to be TrialPeriod. 
 
