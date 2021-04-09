---
subcategory: "license"
layout: "intersight"
page_title: "Intersight: intersight_license_customer_op"
description: |-
  Customer operation object to refresh the registration or re-authenticate, pre-created.
---

# Data Source: intersight_license_customer_op
Customer operation object to refresh the registration or re-authenticate, pre-created.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_license_customer_op.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `active_admin`:(bool) The license administrative state.Set this property to 'true' to activate the license entitlements. 
* `all_devices_to_default_tier`:(bool) Move all licensed devices to default license tier. 
* `create_time`:(string) The time when this managed object was created. 
* `deregister_device`:(bool) Trigger de-registration/disable. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `enable_trial`:(bool) Enable trial for Intersight licensing. 
* `evaluation_period`:(int) The default Trial or Grace period customer is entitled to. 
* `extra_evaluation`:(int) The number of days the trial Trial or Grace period is extended. The trial or grace period can be extended once. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `renew_authorization`:(bool) Trigger renew authorization. 
* `renew_id_certificate`:(bool) Trigger renew registration. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `show_agent_tech_support`:(bool) Trigger show tech support feature. 
 