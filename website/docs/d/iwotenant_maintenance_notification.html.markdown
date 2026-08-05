---
subcategory: "iwotenant"
layout: "intersight"
page_title: "Intersight: intersight_iwotenant_maintenance_notification"
description: |-
        MaintenanceNotifications represent maintenance-related messages intended to be shown to users as a UI banner when they log in. The notification is scoped to an IWO tenant (by `iwoId`) and can be tied to an account/tenant relationship for targeted communication.
        #### Purpose
        Communicate planned IWO maintenance windows to customers through a time-bounded banner message, improving transparency and reducing surprise during service-impacting operations.
        #### Key Concepts
        - **UI banner communication:** Stores message content (or an i18n key) for consistent UI display.
        - **Time window control:** Defines when to show the message (`showFromTime` → `showUntilTime`) and when maintenance actually starts (`maintenanceStartTime`).
        - **Tenant identification:** Uses `iwoId` to bind the notification to the intended tenant namespace/account context.
        - **Lifecycle management:** Admin-controlled create/delete, enabling planned messaging only when needed.

---

# Data Source: intersight_iwotenant_maintenance_notification
MaintenanceNotifications represent maintenance-related messages intended to be shown to users as a UI banner when they log in. The notification is scoped to an IWO tenant (by `iwoId`) and can be tied to an account/tenant relationship for targeted communication.
#### Purpose
Communicate planned IWO maintenance windows to customers through a time-bounded banner message, improving transparency and reducing surprise during service-impacting operations.
#### Key Concepts
- **UI banner communication:** Stores message content (or an i18n key) for consistent UI display.
- **Time window control:** Defines when to show the message (`showFromTime` → `showUntilTime`) and when maintenance actually starts (`maintenanceStartTime`).
- **Tenant identification:** Uses `iwoId` to bind the notification to the intended tenant namespace/account context.
- **Lifecycle management:** Admin-controlled create/delete, enabling planned messaging only when needed.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iwotenant_maintenance_notification.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `i18n_key`:(string) Any i18n (internationalization) key defined the message content. If the key already exists then the  message content will be picked based on the key, otherwise provided message value will be used. 
* `iwo_id`:(string) The iwoId uniquely identifies a IWO tenant. The iwoId is used as part of namespace, (logical) database names, policies in vault and many others. As of now, accountMoid has to be provided as the iwoId. 
* `maintenance_start_time`:(string) The date/time from which the actual maintenance operations will be performed for a Customer's account. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `ntfn_message`:(string) The notification message content is to display in the UI banner after the Customer's login to inform about planned maintenance operations on IWO. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `show_from_time`:(string) The date/time from which the maintenance banner message will be shown to the Customer after login in to  Intersight UI. 
* `show_until_time`:(string) The date/time until which the maintenance banner message will be shown to the Customer after login into  Intersight UI. This will also be the time actual maintenance operation is planned for the finish of a  Customer's account. 
 
