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
The following arguments can be used to get data of already created objects in Intersight appliance:
* `active_admin`:(bool) The license administrative state.Set this property to 'true' to activate the license entitlements. 
* `all_devices_to_default_tier`:(bool) Move all licensed devices to default license tier. 
* `deregister_device`:(bool) Trigger de-registration/disable. 
* `enable_trial`:(bool) Enable trial for Intersight licensing. 
* `evaluation_period`:(int) The default Trial or Grace period customer is entitled to. 
* `extra_evaluation`:(int) The number of days the trial Trial or Grace period is extended. The trial or grace period can be extended once. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `renew_authorization`:(bool) Trigger renew authorization. 
* `renew_id_certificate`:(bool) Trigger renew registration. 
* `show_agent_tech_support`:(bool) Trigger show tech support feature. 
