
---
layout: "intersight"
page_title: "Intersight: intersight_license_customer_op"
sidebar_current: "docs-intersight-data-source-license-customer-op"
description: |-
Customer operation object to refresh the registration or re-authenticate, pre-created.
---

# Data Source: intersight_license._customer_op
Customer operation object to refresh the registration or re-authenticate, pre-created.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `active_admin`:(bool) The license administrative state.Set this property to 'true' to activate the license entitlements. 
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `deregister_device`:(bool) Trigger de-registration/disable. 
* `enable_trial`:(bool) Enable trial for Intersight licensing. 
* `evaluation_period`:(int) The default Trial or Grace period customer is entitled to. 
* `extra_evaluation`:(int) The number of days the trial Trial or Grace period is extended. The trial or grace period can be extended once. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `renew_authorization`:(bool) Trigger renew authorization. 
* `renew_id_certificate`:(bool) Trigger renew registration. 
* `show_agent_tech_support`:(bool) Trigger show tech support feature. 
