
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
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `description`:(string) Description of the policy. 
* `inband_vlan`:(int) VLAN to be used for server access over Inband network. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
