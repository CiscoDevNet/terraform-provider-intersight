---
subcategory: "chassis"
layout: "intersight"
page_title: "Intersight: intersight_chassis_config_import"
description: |-
        The ConfigImport object enables the import of configuration from existing hardware (such as servers or chassis) to automatically generate corresponding profiles and policies within the management system.
        #### Purpose
        ConfigImport automates the onboarding process by translating the current hardware state into managed objects, ensuring consistency and rapid integration into centralized workflows. If the import is successful, hardware is assigned to the generated profile and associated policies.
        #### Key Concepts
        - **Automated Profile and Policy Generation:** Converts existing configurations into managed profiles and policies.
        - **Seamless Hardware Onboarding:** Assigns hardware to newly created profiles for immediate management.
        - **Operational Safeguards:** Ensures that profiles and policies are not generated if the import fails.
        - **Controlled Access:** Access to import actions is governed by privilege sets for security.

---

# Data Source: intersight_chassis_config_import
The ConfigImport object enables the import of configuration from existing hardware (such as servers or chassis) to automatically generate corresponding profiles and policies within the management system.
#### Purpose
ConfigImport automates the onboarding process by translating the current hardware state into managed objects, ensuring consistency and rapid integration into centralized workflows. If the import is successful, hardware is assigned to the generated profile and associated policies.
#### Key Concepts
- **Automated Profile and Policy Generation:** Converts existing configurations into managed profiles and policies.
- **Seamless Hardware Onboarding:** Assigns hardware to newly created profiles for immediate management.
- **Operational Safeguards:** Ensures that profiles and policies are not generated if the import fails.
- **Controlled Access:** Access to import actions is governed by privilege sets for security.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_chassis_config_import.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the imported profile. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `policy_prefix`:(string) Policy prefix for the policies of the imported chassis profile. 
* `profile_name`:(string) Profile name for the imported chassis profile. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
