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
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_asset_subscription_device_contract_information.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `device_id`:(string) Unique identifier of the Cisco device. 
* `device_pid`:(string) Product identifier for the specified Cisco device. It is used to distinguish between HyperFlex and UCS devices. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `subscription_ref_id`:(string) Identifies the consumption-based subscription. 
 
