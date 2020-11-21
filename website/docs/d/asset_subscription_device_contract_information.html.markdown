
---
layout: "intersight"
page_title: "Intersight: intersight_asset_subscription_device_contract_information"
sidebar_current: "docs-intersight-data-source-asset-subscription-device-contract-information"
description: |-
Contains information about Cisco devices associated with consumption-based subscriptions. In addition to device installation status and customer address, information about returns and replacements is also recorded here. We listen to messages sent by Cisco Install Base and create/update an instance of this object.
---

# Data Source: intersight_asset._subscription_device_contract_information
Contains information about Cisco devices associated with consumption-based subscriptions. In addition to device installation status and customer address, information about returns and replacements is also recorded here. We listen to messages sent by Cisco Install Base and create/update an instance of this object.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `device_id`:(string) Unique identifier of the Cisco device. 
* `device_pid`:(string) Product identifier for the specified Cisco device. It is used to distinguish between HyperFlex and UCS devices. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `subscription_ref_id`:(string) Identifies the consumption-based subscription. 
