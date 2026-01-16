---
subcategory: "server"
layout: "intersight"
page_title: "Intersight: intersight_server_diagnostics"
description: |-
        The Diagnostics object is designed to facilitate health checks on server hardware components. It plays a crucial role in ensuring the operational integrity and performance of server infrastructure by performing diagnostics on various hardware elements.
        #### Purpose
        The Diagnostics object is intended to conduct quick and comprehensive evaluations of server hardware, identifying potential issues and ensuring that components function as expected. This provides a structured approach to initiating, managing, and reporting on diagnostic tests for server components.
        #### Key Concepts
        - **Comprehensive Testing**: Supports a wide range of diagnostic tests, covering multiple hardware components such as CPU, memory, storage, and more.
        - **Action Management**: Facilitates the start, cancellation, and completion of diagnostic processes, allowing for flexible management of diagnostic operations.
        - **Access Control**: Employs privilege sets to manage who can initiate and manage diagnostics, ensuring that only authorized users perform these operations.

---

# Data Source: intersight_server_diagnostics
The Diagnostics object is designed to facilitate health checks on server hardware components. It plays a crucial role in ensuring the operational integrity and performance of server infrastructure by performing diagnostics on various hardware elements.
#### Purpose
The Diagnostics object is intended to conduct quick and comprehensive evaluations of server hardware, identifying potential issues and ensuring that components function as expected. This provides a structured approach to initiating, managing, and reporting on diagnostic tests for server components.
#### Key Concepts
- **Comprehensive Testing**: Supports a wide range of diagnostic tests, covering multiple hardware components such as CPU, memory, storage, and more.
- **Action Management**: Facilitates the start, cancellation, and completion of diagnostic processes, allowing for flexible management of diagnostic operations.
- **Access Control**: Employs privilege sets to manage who can initiate and manage diagnostics, ensuring that only authorized users perform these operations.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_server_diagnostics.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `action`:(string) The action to be performed on the server whether to start or to cancel the diagnostics.* `Start` - Mark the start of the diagnostics on the server.* `Cancel` - Mark the cancellation of the diagnostics on the server.* `Complete` - Mark the completion of the diagnostics on the server. 
* `create_time`:(string) The time when this managed object was created. 
* `diagnostics_type`:(string) Type of diagnostics to be performed on the server hardware components.* `Quick` - Perform fast and limited diagnostics on server hardware components.* `Comprehensive` - Perform slow and extensive diagnostics on server hardware components. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the diagnostics being run. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
