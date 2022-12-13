---
subcategory: "license"
layout: "intersight"
page_title: "Intersight: intersight_license_license_registration_status"
description: |-
        Current step of the registration status for licensing.

---

# Data Source: intersight_license_license_registration_status
Current step of the registration status for licensing.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_license_license_registration_status.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_creation_state`:(bool) Stores information on whether account has gone through the registration wizard. True if has not. 
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `is_new_account`:(bool) Stores information on whether account is new. This data is used for UI theme upgrade, where existing users will be shown a slightly different screen. True if new. 
* `license_registration_state`:(string) Stores information on the current flow of license registration.* `RegistrationNotStarted` - The license registration state to chose between trial and registration.* `RegistrationStarted` - The license registration state during set up flow.* `RegistrationComplete` - The license registration state after completion. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `trial_registration_complete`:(bool) Stores information on whether trial flow has been completed. True if trial registration finish. 
 
