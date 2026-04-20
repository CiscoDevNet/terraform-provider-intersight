---
subcategory: "capability"
layout: "intersight"
page_title: "Intersight: intersight_capability_policy_constants"
description: |-
        The PolicyConstants object is a specialized component within the capability catalog, designed to manage and aggregate validation rules for server policy properties. It serves as a centralized reference point for ensuring that server policies adhere to predefined standards and constraints, facilitating the consistent enforcement of management rules across the system.
        #### Purpose
        The primary purpose of PolicyConstants is to provide a blueprint for policy validation. By aggregating specific rules and restrictions—such as reserved keywords for device naming—it ensures that server policies are configured correctly before they are applied. This centralized approach reduces configuration errors and maintains the integrity of server profiles and orchestration processes.
        #### Key Concepts
        - **Standardized Validation:** Acts as a repository for rules that govern how policy properties are defined, ensuring uniformity across different server configurations and policy types.
        - **Policy Integrity:** By enforcing restrictions on naming and parameters, it helps maintain the operational stability of critical system configurations, such as boot order policies.
        - **Capability Integration:** As an extension of the Capability object, it is part of a broader catalog of pre-defined server policy types, allowing it to be seamlessly integrated into the system's management framework.
        - **Role-Based Access:** Access to these constants is governed by a comprehensive set of privilege sets, ensuring that technical, administrative, and marketing roles have the appropriate level of visibility into policy constraints.

---

# Data Source: intersight_capability_policy_constants
The PolicyConstants object is a specialized component within the capability catalog, designed to manage and aggregate validation rules for server policy properties. It serves as a centralized reference point for ensuring that server policies adhere to predefined standards and constraints, facilitating the consistent enforcement of management rules across the system.
#### Purpose
The primary purpose of PolicyConstants is to provide a blueprint for policy validation. By aggregating specific rules and restrictions—such as reserved keywords for device naming—it ensures that server policies are configured correctly before they are applied. This centralized approach reduces configuration errors and maintains the integrity of server profiles and orchestration processes.
#### Key Concepts
- **Standardized Validation:** Acts as a repository for rules that govern how policy properties are defined, ensuring uniformity across different server configurations and policy types.
- **Policy Integrity:** By enforcing restrictions on naming and parameters, it helps maintain the operational stability of critical system configurations, such as boot order policies.
- **Capability Integration:** As an extension of the Capability object, it is part of a broader catalog of pre-defined server policy types, allowing it to be seamlessly integrated into the system's management framework.
- **Role-Based Access:** Access to these constants is governed by a comprehensive set of privilege sets, ensuring that technical, administrative, and marketing roles have the appropriate level of visibility into policy constraints.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_capability_policy_constants.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `boot_device_reserved_keywords`:(string) List of reserved keywords for boot device policies. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) An unique identifer for a capability descriptor. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
