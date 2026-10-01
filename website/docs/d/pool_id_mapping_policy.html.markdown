---
subcategory: "pool"
layout: "intersight"
page_title: "Intersight: intersight_pool_id_mapping_policy"
description: |-
        The IdMappingPolicies object defines a grouping of deployment targets (resource groups and organizations) that can be mapped to specific ID blocks.
        #### Purpose
        It provides a way to restrict which ID blocks can be used by specific deployment targets, ensuring logical separation and controlled resource distribution.
        #### Key Concepts
        - **Target Mapping:** Maps ID blocks to specific organizations or resource groups.
        - **Usage Tracking:** Monitors how many ID pools are attached to the policy.

---

# Data Source: intersight_pool_id_mapping_policy
The IdMappingPolicies object defines a grouping of deployment targets (resource groups and organizations) that can be mapped to specific ID blocks.
#### Purpose
It provides a way to restrict which ID blocks can be used by specific deployment targets, ensuring logical separation and controlled resource distribution.
#### Key Concepts
- **Target Mapping:** Maps ID blocks to specific organizations or resource groups.
- **Usage Tracking:** Monitors how many ID pools are attached to the policy.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_pool_id_mapping_policy.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the policy. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `usage_count`:(int) The number of ID pools to which this ID mapping policy is attached. 
 
