---
subcategory: "tam"
layout: "intersight"
page_title: "Intersight: intersight_tam_advisory_info"
description: |-
        ### Overview
        The AdvisoryInfo object captures the state of Security Advisories and Field Notices, including whether an advisory is visible in the Intersight account, as depicted by its acknowledgement status.
        #### Purpose
        AdvisoryInfo serves as a personalized repository of advisory states for each account, allowing users to manage and track advisories relevant to their environment.
        #### Key Concepts
        - **State Management** - Tracks whether advisories are active or acknowledged, reflecting user preferences for updates.
        - **Account Association** - Directly links advisory states to specific accounts for tailored management.
        - **Shared Resource** - Supports collaborative management of advisories within shared organizational contexts.
        - **Controlled Permissions** - Restricts creation, updating, and deletion of advisory states to Account Administrators.

---

# Data Source: intersight_tam_advisory_info
### Overview
The AdvisoryInfo object captures the state of Security Advisories and Field Notices, including whether an advisory is visible in the Intersight account, as depicted by its acknowledgement status.
#### Purpose
AdvisoryInfo serves as a personalized repository of advisory states for each account, allowing users to manage and track advisories relevant to their environment.
#### Key Concepts
- **State Management** - Tracks whether advisories are active or acknowledged, reflecting user preferences for updates.
- **Account Association** - Directly links advisory states to specific accounts for tailored management.
- **Shared Resource** - Supports collaborative management of advisories within shared organizational contexts.
- **Controlled Permissions** - Restricts creation, updating, and deletion of advisory states to Account Administrators.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_tam_advisory_info.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `state`:(string) Current state of the advisory for the owner. Indicates if the user is interested in getting updates for the advisory.* `active` - Advisory is currently active and the user wants to receive updates for this advisory.* `acknowledged` - Advisory is seen and acknowledged by the user and she no longer wants to recieve updates. 
 
