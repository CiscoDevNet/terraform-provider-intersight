---
subcategory: "asset"
layout: "intersight"
page_title: "Intersight: intersight_asset_subscription_device_contract_information"
description: |-
        The SubscriptionDeviceContractInformation object is a crucial component of asset management within the Cisco ecosystem, designed to encapsulate information regarding devices associated with consumption-based subscriptions. It plays a vital role in tracking device installation status, customer details, and managing returns and replacements.
        #### Purpose
        SubscriptionDeviceContractInformation serves as a comprehensive repository for device-related information under consumption-based subscriptions. It ensures a seamless integration with Cisco Install Base, providing up-to-date data on device status and associated customer information.
        #### Key Concepts
        - **Device Association:** Centralizes information about Cisco devices under subscriptions, facilitating efficient management and tracking.
        - **Integration with Install Base:** Listens to messages from Cisco Install Base for automated updates, ensuring accurate and current data representation.
        - **Ownership and Access:** Managed by the system with defined access privileges, supporting secure and controlled access by authorized users.
        - **Information Persistence:** Ensures data integrity and reliability by maintaining detailed records of device installation status and transaction history.

---

# Data Source: intersight_asset_subscription_device_contract_information
The SubscriptionDeviceContractInformation object is a crucial component of asset management within the Cisco ecosystem, designed to encapsulate information regarding devices associated with consumption-based subscriptions. It plays a vital role in tracking device installation status, customer details, and managing returns and replacements.
#### Purpose
SubscriptionDeviceContractInformation serves as a comprehensive repository for device-related information under consumption-based subscriptions. It ensures a seamless integration with Cisco Install Base, providing up-to-date data on device status and associated customer information.
#### Key Concepts
- **Device Association:** Centralizes information about Cisco devices under subscriptions, facilitating efficient management and tracking.
- **Integration with Install Base:** Listens to messages from Cisco Install Base for automated updates, ensuring accurate and current data representation.
- **Ownership and Access:** Managed by the system with defined access privileges, supporting secure and controlled access by authorized users.
- **Information Persistence:** Ensures data integrity and reliability by maintaining detailed records of device installation status and transaction history.
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
 
