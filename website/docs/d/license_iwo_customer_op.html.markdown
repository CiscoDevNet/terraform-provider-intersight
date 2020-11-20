
---
layout: "intersight"
page_title: "Intersight: intersight_license_iwo_customer_op"
sidebar_current: "docs-intersight-data-source-license-iwo-customer-op"
description: |-
Customer operation object to refresh the registration or re-authenticate, pre-created.
---

# Data Source: intersight_license._iwo_customer_op
Customer operation object to refresh the registration or re-authenticate, pre-created.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `active_admin`:(bool) The workload optimizer license administrative state.Set this property to 'true' to activate the workload optimizer license entitlements. 
* `active_license_type`:(string) Active workload optimizer license tier set by user.* `Base` - Base as a License type. It is default license type.* `Essential` - Essential as a License type.* `Standard` - Standard as a License type.* `Advantage` - Advantage as a License type.* `Premier` - Premier as a License type.* `IWO-Essential` - IWO-Essential as a License type.* `IWO-Advantage` - IWO-Advantage as a License type.* `IWO-Premier` - IWO-Premier as a License type. 
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `enable_trial`:(bool) Enable trial for Intersight licensing. 
* `evaluation_period`:(int) The default Trial or Grace period customer is entitled to. 
* `extra_evaluation`:(int) The number of days the trial Trial or Grace period is extended. The trial or grace period can be extended once. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
