
---
layout: "intersight"
page_title: "Intersight: intersight_vmedia_policy"
sidebar_current: "docs-intersight-data-source-vmedia-policy"
description: |-
Policy to configure virtual media settings on endpoint.
---

# Data Source: intersight_vmedia._policy
Policy to configure virtual media settings on endpoint.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `description`:(string) Description of the policy. 
* `enabled`:(bool) State of the Virtual Media service on the endpoint. 
* `encryption`:(bool) If enabled, allows encryption of all Virtual Media communications. 
* `low_power_usb`:(bool) If enabled, the virtual drives appear on the boot selection menu after mapping the image and rebooting the host. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
