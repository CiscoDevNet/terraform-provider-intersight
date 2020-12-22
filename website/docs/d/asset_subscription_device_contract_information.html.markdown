---
subcategory: "asset"
layout: "intersight"
page_title: "Intersight: intersight_asset_subscription_device_contract_information"
description: |-
  Contains information about Cisco devices associated with consumption-based subscriptions. In addition to device installation status and customer address, information about returns and replacements is also recorded here. We listen to messages sent by Cisco Install Base and create/update an instance of this object.
---

# Data Source: intersight_asset_subscription_device_contract_information
Contains information about Cisco devices associated with consumption-based subscriptions. In addition to device installation status and customer address, information about returns and replacements is also recorded here. We listen to messages sent by Cisco Install Base and create/update an instance of this object.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `device_id`:(string) Unique identifier of the Cisco device. 
* `device_pid`:(string) Product identifier for the specified Cisco device. It is used to distinguish between HyperFlex and UCS devices. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `subscription_ref_id`:(string) Identifies the consumption-based subscription. 
