---
subcategory: "license"
layout: "intersight"
page_title: "Intersight: intersight_license_license_info"
description: |-
  License state information for a specific license entitlement. Essentials license entitlement is supported currently.
licenseState attribute is used for license enforcement. When license state is one of TrialPeriod, Compliance, or OutOfCompliance,
the feature set defined for the license entitlement is granted to the customer.
---

# Resource: intersight_license_license_info
License state information for a specific license entitlement. Essentials license entitlement is supported currently.
licenseState attribute is used for license enforcement. When license state is one of TrialPeriod, Compliance, or OutOfCompliance,
the feature set defined for the license entitlement is granted to the customer.
## Argument Reference
The following arguments are supported:
* `account_license_data`:(HashMap) - A reference to a licenseAccountLicenseData resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `active_admin`:(bool)(Computed) The license administrative state.Set this property to 'true' to activate the license entitlements. 
* `days_left`:(int)(Computed) The number of days left for licenseState to stay in TrialPeriod or OutOfCompliance state. 
* `end_time`:(string)(Computed) The date and time when the trial period expires.The value of the 'endTime' property is set when the account enters the TrialPeriod or OutOfCompliance state. 
* `enforce_mode`:(string)(Computed) The entitlement mode reported by Cisco Smart Software Manager. 
* `error_desc`:(string)(Computed) The detailed error message when there is any error related to this licensing entitlement. 
* `evaluation_period`:(int) The default Trial or Grace period customer is entitled to. 
* `extra_evaluation`:(int) The number of days the trial Trial or Grace period is extended.The trial or grace period can be extended once. 
* `license_count`:(int)(Computed) The total number of devices claimed in the Intersight account. 
* `license_state`:(string)(Computed) The license state defined by Intersight.The value may be one of NotLicensed, TrialPeriod, OutOfCompliance, Compliance, GraceExpired, or TrialExpired.* `NotLicensed` - The license token is neither activated nor registered.* `GraceExpired` - The license grace period has expired.* `TrialPeriod` - The 90 days of trial period.* `OutOfCompliance` - The license is out of compliance.* `Compliance` - The license is in compliance.* `TrialExpired` - The trial period of 90 days has expired. 
* `license_type`:(string)(Computed) The name of the Intersight license entitlement.For example, this property may be set to 'Essential'.* `Base` - Base as a License type. It is default license type.* `Essential` - Essential as a License type.* `Standard` - Standard as a License type.* `Advantage` - Advantage as a License type.* `Premier` - Premier as a License type.* `IWO-Essential` - IWO-Essential as a License type.* `IWO-Advantage` - IWO-Advantage as a License type.* `IWO-Premier` - IWO-Premier as a License type. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `start_time`:(string)(Computed) The date and time when the licenseState entered the TrialPeriod or OutOfCompliance state. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `trial_admin`:(bool)(Computed) The administrative state of the trial license.When the LicenseState is set to 'NotLicensed', 'trialAdmin' can be set to true to start the trial period,i.e. licenseState is set to be TrialPeriod. 


## Import
`intersight_license_license_info` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_license_license_info.example 1234567890987654321abcde
``` 