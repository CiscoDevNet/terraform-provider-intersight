---
subcategory: "adapter"
layout: "intersight"
page_title: "Intersight: intersight_adapter_config_policy"
description: |-
        The ConfigPolicy object is a core element for managing adapter configurations on standalone rack servers. It defines and maintains Ethernet and Fibre Channel settings for Cisco VIC adapters. This policy is not supported on Intersight Managed Mode (IMM) servers.
        #### Purpose
        A ConfigPolicy serves as a policy framework for configuring vital adapter settings, allowing administrators to manage and deploy configurations efficiently. It supports operations such as read, create, update, and delete, facilitating seamless integration with server management workflows.
        #### Key Concepts
        - **Adapter Configuration:** Focuses on configuring Ethernet and Fibre-Channel settings, addressing the needs of various server environments.
        - **Access Control:** Ensures that only authorized users can perform configuration tasks, enhancing security and reliability.

---

# Data Source: intersight_adapter_config_policy
The ConfigPolicy object is a core element for managing adapter configurations on standalone rack servers. It defines and maintains Ethernet and Fibre Channel settings for Cisco VIC adapters. This policy is not supported on Intersight Managed Mode (IMM) servers.
#### Purpose
A ConfigPolicy serves as a policy framework for configuring vital adapter settings, allowing administrators to manage and deploy configurations efficiently. It supports operations such as read, create, update, and delete, facilitating seamless integration with server management workflows.
#### Key Concepts
- **Adapter Configuration:** Focuses on configuring Ethernet and Fibre-Channel settings, addressing the needs of various server environments.
- **Access Control:** Ensures that only authorized users can perform configuration tasks, enhancing security and reliability.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_adapter_config_policy.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the policy. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
