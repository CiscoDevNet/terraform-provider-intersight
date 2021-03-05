---
subcategory: "license"
layout: "intersight"
page_title: "Intersight: intersight_license_iwo_customer_op"
description: |-
  Customer operation object to refresh the registration or re-authenticate, pre-created.
---

# Data Source: intersight_license_iwo_customer_op
Customer operation object to refresh the registration or re-authenticate, pre-created.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_license_iwo_customer_op.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `active_admin`:(bool) The workload optimizer license administrative state.Set this property to 'true' to activate the workload optimizer license entitlements. 
* `active_license_type`:(string) Active workload optimizer license tier set by user.* `Base` - Base as a License type. It is default license type.* `Essential` - Essential as a License type.* `Standard` - Standard as a License type.* `Advantage` - Advantage as a License type.* `Premier` - Premier as a License type.* `IWO-Essential` - IWO-Essential as a License type.* `IWO-Advantage` - IWO-Advantage as a License type.* `IWO-Premier` - IWO-Premier as a License type. 
* `enable_trial`:(bool) Enable trial for Intersight licensing. 
* `evaluation_period`:(int) The default Trial or Grace period customer is entitled to. 
* `extra_evaluation`:(int) The number of days the trial Trial or Grace period is extended. The trial or grace period can be extended once. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
 