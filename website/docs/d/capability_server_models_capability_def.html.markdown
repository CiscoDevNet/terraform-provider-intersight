---
subcategory: "capability"
layout: "intersight"
page_title: "Intersight: intersight_capability_server_models_capability_def"
description: |-
        The ServerModelsCapabilityDef object categorizes server models into logical groupings (for example, by server family or generation).
        #### Purpose
        ServerModelsCapabilityDef provides a catalog-backed classification mechanism for server models, enabling feature targeting, policy applicability checks, or UI grouping based on server type.
        #### Key Concepts
        - **Model categorization:** Groups many model strings under a meaningful server type label.
        - **Feature targeting:** Enables applying behavior or validation to a category rather than individual model strings.
        - **Scalable cataloging:** Reduces duplication by centralizing model-to-type mapping.
        - **Cross-service consistency:** Improves consistent interpretation of “server family” across workflows.

---

# Data Source: intersight_capability_server_models_capability_def
The ServerModelsCapabilityDef object categorizes server models into logical groupings (for example, by server family or generation).
#### Purpose
ServerModelsCapabilityDef provides a catalog-backed classification mechanism for server models, enabling feature targeting, policy applicability checks, or UI grouping based on server type.
#### Key Concepts
- **Model categorization:** Groups many model strings under a meaningful server type label.
- **Feature targeting:** Enables applying behavior or validation to a category rather than individual model strings.
- **Scalable cataloging:** Reduces duplication by centralizing model-to-type mapping.
- **Cross-service consistency:** Improves consistent interpretation of “server family” across workflows.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_capability_server_models_capability_def.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) An unique identifer for a capability descriptor. 
* `server_type`:(string) Type of the server. Example, BladeM6, RackM5. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
