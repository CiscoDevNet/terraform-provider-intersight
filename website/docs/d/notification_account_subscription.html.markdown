---
subcategory: "notification"
layout: "intersight"
page_title: "Intersight: intersight_notification_account_subscription"
description: |-
        ### Overview
        The AccountSubscription object is a crucial component of the Intersight notification framework,   designed to facilitate account-level subscriptions managed by Account Administrators.
        It provides a structured way to handle notifications based on specific events within an account, ensuring administrators can effectively configure and manage subscription settings.
        #### Purpose
        The AccountSubscription object enables administrators to set up notifications for various account events. It serves as the blueprint for creating, updating, and managing subscriptions, ensuring that notifications are sent according to predefined conditions and actions.
        #### Key Concepts
        - **Account-Level Management** - Designed for account administrators, allowing easy configuration of subscription settings at the account level.
        - **Integration with Notification Methods** - Supports multiple notification methods such as email and webhook to meet diverse requirements.
        - **Conditional Notifications** - Allows definition of conditions that trigger notifications, ensuring only relevant alerts are sent.
        - **Access Control** - Enforces privilege sets so only authorized users can create, update, or delete subscriptions, maintaining security and integrity.
        - **Flexibility and Scalability** - Scales notification features to support complex use cases and evolving business needs.

---

# Data Source: intersight_notification_account_subscription
### Overview
The AccountSubscription object is a crucial component of the Intersight notification framework,   designed to facilitate account-level subscriptions managed by Account Administrators.  
It provides a structured way to handle notifications based on specific events within an account, ensuring administrators can effectively configure and manage subscription settings.
#### Purpose
The AccountSubscription object enables administrators to set up notifications for various account events. It serves as the blueprint for creating, updating, and managing subscriptions, ensuring that notifications are sent according to predefined conditions and actions.
#### Key Concepts
- **Account-Level Management** - Designed for account administrators, allowing easy configuration of subscription settings at the account level.
- **Integration with Notification Methods** - Supports multiple notification methods such as email and webhook to meet diverse requirements.
- **Conditional Notifications** - Allows definition of conditions that trigger notifications, ensuring only relevant alerts are sent.
- **Access Control** - Enforces privilege sets so only authorized users can create, update, or delete subscriptions, maintaining security and integrity.
- **Flexibility and Scalability** - Scales notification features to support complex use cases and evolving business needs.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_notification_account_subscription.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `condition_operator`:(string) Operation that binds all the different conditions together.* `All` - All is an AND condition applied against the individual conditions.* `Any` - Any is an OR condition applied against the individual conditions. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) The description for the subscription. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `enabled`:(bool) Subscription can be switched on/off without necessity to change the subscriptionsettings: notification methods, conditions, etc.Ex.: Subscription MO can be configured, but switched off. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the subscription. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `type`:(string) The chosen subscription type imposes it is own validation rules.When 'email' type is chosen, actions array can contain only one entry and it is entry should be of canbe only notification.SendEmail; conditions can contain only notification.AlarmMoCondition and conditiontypes should be unique.When the 'webhook' type is chosen, the actions array can contain only one entry and it is entry should be of canbe only notification.TriggerWebhook; conditions can contain up to a limited amount of entries and all of themshould be of type notification.MoCondition.* `email` - Email type requires usage of notification.SendEmail complex types for actionsand notification.AlarmMoCondition complex types for conditions.* `webhook` - Webhook type requires usage of notification.TriggerWebhook complex types for actionsand notification.MoCondition complex types for conditions. 
* `verify`:(string) Used to verify the actions of the Subscription MO. For a 'webhook' type Ping event is sent to verifythat the webhook server is accessible. For an 'email' type there will be a verification email sent.* `none` - No actions will be verified. Default value.* `all` - All actions will be re-verified. The previous state of the verification will be preserved. 
 
