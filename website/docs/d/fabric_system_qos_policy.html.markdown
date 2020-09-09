
---
layout: "intersight"
page_title: "Intersight: intersight_fabric_system_qos_policy"
sidebar_current: "docs-intersight-data-source-fabric-system-qos-policy"
description: |-
Configuration object sent by user to setup Quality of Service (QoS) for this switch.
---

# Data Source: intersight_fabric._system_qos_policy
Configuration object sent by user to setup Quality of Service (QoS) for this switch.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `description`:(string) Description of the policy. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
