
---
layout: "intersight"
page_title: "Intersight: intersight_iam_privilege_set"
sidebar_current: "docs-intersight-data-source-iam-privilege-set"
description: |-
A collection of privileges.
---

# Data Source: intersight_iam._privilege_set
A collection of privileges.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `description`:(string) Description of the privilege set. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the privilege set. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
