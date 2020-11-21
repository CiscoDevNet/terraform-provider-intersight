
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
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `description`:(string) Description of the policy. 
* `enabled`:(bool) State of the Virtual Media service on the endpoint. 
* `encryption`:(bool) If enabled, allows encryption of all Virtual Media communications. 
* `low_power_usb`:(bool) If enabled, the virtual drives appear on the boot selection menu after mapping the image and rebooting the host. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
