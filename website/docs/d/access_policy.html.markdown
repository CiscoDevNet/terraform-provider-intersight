
---
layout: "intersight"
page_title: "Intersight: intersight_access_policy"
sidebar_current: "docs-intersight-data-source-access-policy"
description: |-
Policy to configure server management options via CIMC.
---

# Data Source: intersight_access._policy
Policy to configure server management options via CIMC.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `description`:(string) Description of the policy. 
* `inband_vlan`:(int) VLAN to be used for server access over Inband network. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
