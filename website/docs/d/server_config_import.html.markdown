---
subcategory: "server"
layout: "intersight"
page_title: "Intersight: intersight_server_config_import"
description: |-
        The ConfigImport object automates the translation of existing hardware configurations into managed system profiles and policies, streamlining the integration of hardware into centralized control.
        #### Purpose
        ConfigImport is designed to facilitate the rapid adoption of hardware by importing its configuration and generating the necessary management objects. Successful imports result in the hardware being assigned to the generated profile, establishing immediate operational control.
        #### Key Concepts
        - **Efficient Onboarding:** Translates hardware configurations into managed objects.
        - **Integrated Assignment:** Assigns imported hardware to new profiles and policies.
        - **Integrity and Reliability:** Ensures objects are only created on successful import.
        - **Privilege Governance:** Secure import actions managed through access controls.

---

# Data Source: intersight_server_config_import
The ConfigImport object automates the translation of existing hardware configurations into managed system profiles and policies, streamlining the integration of hardware into centralized control.
#### Purpose
ConfigImport is designed to facilitate the rapid adoption of hardware by importing its configuration and generating the necessary management objects. Successful imports result in the hardware being assigned to the generated profile, establishing immediate operational control.
#### Key Concepts
- **Efficient Onboarding:** Translates hardware configurations into managed objects.
- **Integrated Assignment:** Assigns imported hardware to new profiles and policies.
- **Integrity and Reliability:** Ensures objects are only created on successful import.
- **Privilege Governance:** Secure import actions managed through access controls.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_server_config_import.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the imported profile. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `policy_prefix`:(string) Policy prefix for the policies of the imported server profile. 
* `profile_name`:(string) Profile name for the imported server profile. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
