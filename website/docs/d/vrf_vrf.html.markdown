
---
layout: "intersight"
page_title: "Intersight: intersight_vrf_vrf"
sidebar_current: "docs-intersight-data-source-vrf-vrf"
description: |-
Virtual Routing and Forwarding (VRF) is a networking technology that implements an isolated Layer 3 domain.
---

# Data Source: intersight_vrf._vrf
Virtual Routing and Forwarding (VRF) is a networking technology that implements an isolated Layer 3 domain.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `description`:(string) Description to help identify or describe this VRF. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the Virtual Routing and Forwarding Instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
