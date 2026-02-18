---
subcategory: "asset"
layout: "intersight"
page_title: "Intersight: intersight_asset_subscription_account"
description: |-
        The SubscriptionAccount object is a foundational element within the subscription management framework, acting as a placeholder for the account collection. Its primary function is to ensure stability and continuity in subscription-related operations by preventing data updates during account change events.
        #### Purpose
        SubscriptionAccount acts as a stabilizing component within the subscription management system, safeguarding data integrity during account modifications. This provides a structured approach to manage and reference account-related subscription data.
        #### Key Concepts
        - **Data Stability:** Protects subscription data from unintended updates during account changes, ensuring operational consistency.
        - **Account Integration:** Serves as a reference point for subscription data, maintaining a clear linkage between accounts and subscriptions.
        - **Access Control:** Leveraging permissions, it ensures only authorized users can interact with account-related subscription data.
        - **Dependency Management:** Plays a key role in managing relationships and dependencies within the subscription ecosystem.

---

# Data Source: intersight_asset_subscription_account
The SubscriptionAccount object is a foundational element within the subscription management framework, acting as a placeholder for the account collection. Its primary function is to ensure stability and continuity in subscription-related operations by preventing data updates during account change events.
#### Purpose
SubscriptionAccount acts as a stabilizing component within the subscription management system, safeguarding data integrity during account modifications. This provides a structured approach to manage and reference account-related subscription data.
#### Key Concepts
- **Data Stability:** Protects subscription data from unintended updates during account changes, ensuring operational consistency.
- **Account Integration:** Serves as a reference point for subscription data, maintaining a clear linkage between accounts and subscriptions.
- **Access Control:** Leveraging permissions, it ensures only authorized users can interact with account-related subscription data.
- **Dependency Management:** Plays a key role in managing relationships and dependencies within the subscription ecosystem.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_asset_subscription_account.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
