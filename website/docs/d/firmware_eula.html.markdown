---
subcategory: "firmware"
layout: "intersight"
page_title: "Intersight: intersight_firmware_eula"
description: |-
        The Eulas object tracks the End User License Agreement (EULA) and K9 acceptance status for an account.
        #### Purpose
        It ensures that users have legally accepted the necessary terms and conditions before accessing Cisco software repositories or downloading images.
        #### Key Concepts
        - **Compliance:** Manages the acceptance status for EULA and K9 terms.
        - **Access Control:** Controls access to software download capabilities based on acceptance status.

---

# Data Source: intersight_firmware_eula
The Eulas object tracks the End User License Agreement (EULA) and K9 acceptance status for an account.
 #### Purpose
 It ensures that users have legally accepted the necessary terms and conditions before accessing Cisco software repositories or downloading images.
 #### Key Concepts
 - **Compliance:** Manages the acceptance status for EULA and K9 terms.
 - **Access Control:** Controls access to software download capabilities based on acceptance status.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_firmware_eula.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `accepted`:(bool) Overall acceptance status for the account, both EULA and K9. 
* `account_moid`:(string) The Account ID for this managed object. 
* `content`:(string) Acceptance form content provided by cisco.com. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `eula_accepted`:(bool) EULA acceptance status for the account. 
* `eula_content`:(string) EULA acceptance form content provided by cisco.com. 
* `k9_accepted`:(bool) K9 acceptance status for the account. 
* `k9_content`:(string) K9 acceptance form content provided by cisco.com. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
