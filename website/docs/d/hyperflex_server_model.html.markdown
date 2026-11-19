---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_server_model"
description: |-
        ServerModels define the set of supported HyperFlex server models. Instead of listing only exact model strings, the supported models are expressed as one or more regex-based entries, allowing the platform to match and validate a broader set of model variants consistently.
        #### Purpose
        Provide an authoritative, system-managed list of supported server model patterns used by HyperFlex workflows (such as validation and compatibility checks) to determine whether a given server model is supported.
        #### Key Concepts
        - **Supported model catalog**: Represents the supported server model set for HyperFlex.
        - **Regex-based matching**: `serverModelEntries` stores supported model definitions in regex format to match families/variants of server models.

---

# Data Source: intersight_hyperflex_server_model
ServerModels define the set of supported HyperFlex server models. Instead of listing only exact model strings, the supported models are expressed as one or more regex-based entries, allowing the platform to match and validate a broader set of model variants consistently.
#### Purpose
Provide an authoritative, system-managed list of supported server model patterns used by HyperFlex workflows (such as validation and compatibility checks) to determine whether a given server model is supported.
#### Key Concepts
- **Supported model catalog**: Represents the supported server model set for HyperFlex.
- **Regex-based matching**: `serverModelEntries` stores supported model definitions in regex format to match families/variants of server models.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_server_model.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
