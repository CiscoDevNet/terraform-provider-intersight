
---
layout: "intersight"
page_title: "Intersight: intersight_resource_membership_holder"
sidebar_current: "docs-intersight-data-source-resource-membership-holder"
description: |-
A holder of REST resources and their membership.
---

# Data Source: intersight_resource._membership_holder
A holder of REST resources and their membership.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of this resource membership holder. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
