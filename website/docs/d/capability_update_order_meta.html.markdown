---
subcategory: "capability"
layout: "intersight"
page_title: "Intersight: intersight_capability_update_order_meta"
description: |-
        The UpdateOrderMeta object provides internal metadata to map the required order of operations for firmware updates.
        #### Purpose
        It defines the sequence of versions an endpoint must pass through to reach the target firmware version, ensuring a safe and successful upgrade path.
        #### Key Concepts
        - **Sequencing:** Lists update orders (source -> interim -> target) for specific groups of hardware.
        - **Platform Tagging:** Categorizes update orders by platform and component category.

---

# Data Source: intersight_capability_update_order_meta
The UpdateOrderMeta object provides internal metadata to map the required order of operations for firmware updates.
#### Purpose
It defines the sequence of versions an endpoint must pass through to reach the target firmware version, ensuring a safe and successful upgrade path.
#### Key Concepts
- **Sequencing:** Lists update orders (source -> interim -> target) for specific groups of hardware.
- **Platform Tagging:** Categorizes update orders by platform and component category.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_capability_update_order_meta.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `category`:(string) The category of the model series. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Verbose description regarding this group. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) An unique identifer for a capability descriptor. 
* `platform_tag`:(string) The platform tag value of the model series. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
