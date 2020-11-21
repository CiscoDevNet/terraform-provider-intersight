
---
layout: "intersight"
page_title: "Intersight: intersight_tam_advisory_instance"
sidebar_current: "docs-intersight-data-source-tam-advisory-instance"
description: |-
Instance of an Intersight advisory applicable for an Intersight managed object. An advisory instance is created when a given advisory is found applicable for an Intersight managed object. An advisory instance is retained for some time even after being cleared for historical purposes. A 'cleared' advisory instance is deleted after the retention time is elaspsed.
---

# Data Source: intersight_tam._advisory_instance
Instance of an Intersight advisory applicable for an Intersight managed object. An advisory instance is created when a given advisory is found applicable for an Intersight managed object. An advisory instance is retained for some time even after being cleared for historical purposes. A 'cleared' advisory instance is deleted after the retention time is elaspsed.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `affected_object_moid`:(string) Moid of the Intersight MO affected by the alert. Deprecated now and will be removed in subsequent releases. 
* `affected_object_type`:(string) Object type of the Intersight MO affected by the alert. Deprecated now and will be removed in subsequent releases. 
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `last_state_change_time`:(string) Timestamp when a state change was observed on this advisory instnace. 
* `last_verified_time`:(string) Timestamp when this advisory was last evaluated. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `state`:(string) Current state of the advisory instance (Active/Cleared/Unknown etc.).* `unknown` - Intersight is unable to determine if the Advisory instance is applicable for the affected managed object.* `active` - Advisory instance is currently active and applicable for the affected managed object.* `cleared` - Advisory instance is no longer applicable for the affected managed object. 
