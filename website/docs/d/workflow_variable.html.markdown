---
subcategory: "workflow"
layout: "intersight"
page_title: "Intersight: intersight_workflow_variable"
description: |-
        The Variable object facilitates the management of user-defined entities within workflows, allowing for consistent and reusable data references across different workflow instances.
        #### Purpose
        Variable serves as a central component for defining and managing workflow variables, ensuring consistency and efficiency across workflows. It allows users to set and reference variable values in multiple workflows, promoting data consistency and reducing redundancy in workflow design and execution.
        #### Key Concepts
        - **Reusable Entities:** Supports the creation of reusable variables, facilitating efficient and scalable workflow design.
        - **Data Validation:** Ensures that variable values adhere to defined data types, promoting data integrity and reliability.

---

# Data Source: intersight_workflow_variable
The Variable object facilitates the management of user-defined entities within workflows, allowing for consistent and reusable data references across different workflow instances.  
#### Purpose
Variable serves as a central component for defining and managing workflow variables, ensuring consistency and efficiency across workflows. It allows users to set and reference variable values in multiple workflows, promoting data consistency and reducing redundancy in workflow design and execution.
#### Key Concepts  
- **Reusable Entities:** Supports the creation of reusable variables, facilitating efficient and scalable workflow design. 
- **Data Validation:** Ensures that variable values adhere to defined data types, promoting data integrity and reliability.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_workflow_variable.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `create_user`:(string) The user identifier who created the environment variable. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `mod_user`:(string) The user identifier who last updated the environment variable. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) This defines the name of the variable and it is set by the system. The name used inside the datatype definition will be set as the name of the variable. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
