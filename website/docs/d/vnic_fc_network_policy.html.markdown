
---
layout: "intersight"
page_title: "Intersight: intersight_vnic_fc_network_policy"
sidebar_current: "docs-intersight-data-source-vnic-fc-network-policy"
description: |-
A Fibre Channel Network policy governs the VSAN configuration for the virtual interfaces.
---

# Data Source: intersight_vnic._fc_network_policy
A Fibre Channel Network policy governs the VSAN configuration for the virtual interfaces.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `description`:(string) Description of the policy. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
