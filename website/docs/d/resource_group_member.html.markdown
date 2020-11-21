
---
layout: "intersight"
page_title: "Intersight: intersight_resource_group_member"
sidebar_current: "docs-intersight-data-source-resource-group-member"
description: |-
A resolved member of a resource group.
---

# Data Source: intersight_resource._group_member
A resolved member of a resource group.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
