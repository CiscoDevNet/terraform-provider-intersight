
---
layout: "intersight"
page_title: "Intersight: intersight_appliance_group_status"
sidebar_current: "docs-intersight-data-source-appliance-group-status"
description: |-
Status of a group of applications.
---

# Data Source: intersight_appliance._group_status
Status of a group of applications.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `description`:(string) Description of the group. 
* `group_name`:(string) The name of group, which includes Identity Management, Device Connector Service, Core Service, Analytics, Internal and Appliance. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `overall_status`:(string) The overall API status from this group's applications. 
