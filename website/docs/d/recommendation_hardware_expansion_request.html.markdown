---
subcategory: "recommendation"
layout: "intersight"
page_title: "Intersight: intersight_recommendation_hardware_expansion_request"
description: |-
        Entity representing the user request for HyperFlex cluster expansion.

---

# Data Source: intersight_recommendation_hardware_expansion_request
Entity representing the user request for HyperFlex cluster expansion.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_recommendation_hardware_expansion_request.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `action`:(string) Action to be triggered for computing recommendation. Default value is None.* `None` - The Enum value None represents that no action is triggered on the forecast Instance managed object.* `Evaluate` - The Enum value Evaluate represents that a re-evaluation of the forecast needs to be triggered. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `message`:(string) User visible error message for the Hardware Expansion request. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the expansion specification. 
* `request_time`:(string) The time when the expansion was requested. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `status`:(string) Status of the Hardware Expansion request.* `None` - The Enum value None represents the default status value before any API call is made.* `Success` - The Enum value Success represents that the API call returned with success.* `Fail` - The Enum value Fail represents that the API call returned with a failure. 
 
