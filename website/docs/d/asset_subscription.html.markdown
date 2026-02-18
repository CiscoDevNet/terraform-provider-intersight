---
subcategory: "asset"
layout: "intersight"
page_title: "Intersight: intersight_asset_subscription"
description: |-
        The Subscription object is central to managing consumption-based subscriptions within the Cisco technology ecosystem. It encapsulates details about the subscriptions linked to Cisco devices, facilitating automated updates and management through integration with Cisco Install Base.
        #### Purpose
        Subscription provides a comprehensive structure for managing consumption-based subscriptions, ensuring seamless integration with Cisco Install Base for accurate data representation and updates. This information is also queried by downstream OPUS team for billing purposes.
        #### Key Concepts
        - **Integration:** Automates subscription management through real-time updates from Cisco Install Base, ensuring data accuracy and currency.
        - **Data Management:** Supports efficient tracking and management of consumption-based subscriptions.
        - **Access Control:** Ensures secure handling of subscription data through defined privileges and permissions.
        - **Relationship Management:** Maintains robust linkage with associated accounts and deployments, supporting comprehensive subscription management.

---

# Data Source: intersight_asset_subscription
The Subscription object is central to managing consumption-based subscriptions within the Cisco technology ecosystem. It encapsulates details about the subscriptions linked to Cisco devices, facilitating automated updates and management through integration with Cisco Install Base.
#### Purpose
Subscription provides a comprehensive structure for managing consumption-based subscriptions, ensuring seamless integration with Cisco Install Base for accurate data representation and updates. This information is also queried by downstream OPUS team for billing purposes. 
#### Key Concepts
- **Integration:** Automates subscription management through real-time updates from Cisco Install Base, ensuring data accuracy and currency.
- **Data Management:** Supports efficient tracking and management of consumption-based subscriptions.
- **Access Control:** Ensures secure handling of subscription data through defined privileges and permissions.
- **Relationship Management:** Maintains robust linkage with associated accounts and deployments, supporting comprehensive subscription management.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_asset_subscription.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `application_name`:(string) Application name reported by Cisco Install Base. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `subscription_ref_id`:(string) Identifies the consumption-based subscription. 
 
