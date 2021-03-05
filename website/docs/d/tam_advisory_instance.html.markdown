---
subcategory: "tam"
layout: "intersight"
page_title: "Intersight: intersight_tam_advisory_instance"
description: |-
  Instance of an Intersight advisory applicable for an Intersight managed object. An advisory instance is created when a given advisory is found applicable for an Intersight managed object. An advisory instance is retained for some time even after being cleared for historical purposes. A 'cleared' advisory instance is deleted after the retention time is elaspsed.
---

# Data Source: intersight_tam_advisory_instance
Instance of an Intersight advisory applicable for an Intersight managed object. An advisory instance is created when a given advisory is found applicable for an Intersight managed object. An advisory instance is retained for some time even after being cleared for historical purposes. A 'cleared' advisory instance is deleted after the retention time is elaspsed.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_tam_advisory_instance.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `affected_object_moid`:(string) Moid of the Intersight MO affected by the alert. Deprecated now and will be removed in subsequent releases. 
* `affected_object_type`:(string) Object type of the Intersight MO affected by the alert. Deprecated now and will be removed in subsequent releases. 
* `last_state_change_time`:(string) Timestamp when a state change was observed on this advisory instnace. 
* `last_verified_time`:(string) Timestamp when this advisory was last evaluated. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `state`:(string) Current state of the advisory instance (Active/Cleared/Unknown etc.).* `unknown` - Intersight is unable to determine if the Advisory instance is applicable for the affected managed object.* `active` - Advisory instance is currently active and applicable for the affected managed object.* `cleared` - Advisory instance is no longer applicable for the affected managed object. 
 