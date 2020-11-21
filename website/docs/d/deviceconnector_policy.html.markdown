
---
layout: "intersight"
page_title: "Intersight: intersight_deviceconnector_policy"
sidebar_current: "docs-intersight-data-source-deviceconnector-policy"
description: |-
Policy to control configuration changes allowed from Cisco IMC.
---

# Data Source: intersight_deviceconnector._policy
Policy to control configuration changes allowed from Cisco IMC.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `description`:(string) Description of the policy. 
* `lockout_enabled`:(bool) Enables configuration lockout on the endpoint. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
