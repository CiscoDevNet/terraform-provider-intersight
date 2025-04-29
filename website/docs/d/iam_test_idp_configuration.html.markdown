---
subcategory: "iam"
layout: "intersight"
page_title: "Intersight: intersight_iam_test_idp_configuration"
description: |-
        API used to test an Intersight IdP configuration.

---

# Data Source: intersight_iam_test_idp_configuration
API used to test an Intersight IdP configuration.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iam_test_idp_configuration.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `error_details`:(string) Error returned by the IdP when the configuration test fails. 
* `idp_entity_id`:(string) Entity ID of the IdP whose configuration needs to be tested. 
* `is_password_set`:(bool) Indicates whether the value of the 'password' property has been set. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `no_of_groups`:(int) Total number of groups of the user received from the IdP after successful authentication. 
* `password`:(string) The password of the test user for testing the IdP configuration settings. It can be any string that adheres to the following constraints. It can have character except spaces, tabs, line breaks. It cannot be more than 256 characters. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `status`:(string) IdP configuration test result.* `Unverified` - The IdP configuration is yet to be verified.* `Verified` - Successfully authenticated the user via the configured IdP and verified the configuration parameters.* `Pending` - Pending verification of the IdP configuration.* `Failed` - Could not authenticate the user via the configured IdP and verify its configuration. 
* `trigger_test`:(bool) Trigger property used to initiate an IdP configuration test. 
* `username`:(string) Email or userId of the test user for testing the IdP configuration settings. 
 
