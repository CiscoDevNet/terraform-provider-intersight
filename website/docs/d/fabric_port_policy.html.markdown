
---
layout: "intersight"
page_title: "Intersight: intersight_fabric_port_policy"
sidebar_current: "docs-intersight-data-source-fabric-port-policy"
description: |-
A policy for all the physical ports of the Fabric Interconnect.
---

# Data Source: intersight_fabric._port_policy
A policy for all the physical ports of the Fabric Interconnect.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `description`:(string) Description of the policy. 
* `device_model`:(string) This field specifies the device model that this Port Policy is being configured for.* `UCS-FI-6454` - The standard 4th generation UCS Fabric Interconnect with 54 ports.* `UCS-FI-64108` - The expanded 4th generation UCS Fabric Interconnect with 108 ports.* `unknown` - Unknown device type, usage is TBD. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
