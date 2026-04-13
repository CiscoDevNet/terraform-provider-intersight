---
subcategory: "os"
layout: "intersight"
page_title: "Intersight: intersight_os_bulk_install_info"
description: |-
        The BulkInstallInfo object is a critical element in the automation of operating system installations across multiple servers. It models the content of CSV files that hold configuration parameters for bulk OS installations, ensuring streamlined and efficient deployment.
        #### Purpose
        BulkInstallInfo facilitates the mass deployment of operating systems by encapsulating necessary parameters within a CSV file. This supports validation processes and manages installation configurations, enabling administrators to efficiently deploy OS across multiple servers.
        #### Key Concepts
        - **CSV-Based Deployment:** Utilizes CSV files to store OS installation parameters, providing a scalable solution for deploying operating systems to numerous servers.
        - **Validation Process:** Includes built-in validation mechanisms to ensure the accuracy and completeness of installation configurations.
        - **Access Control:** Enforces privilege sets such as Account Administrator and Manage Servers, ensuring that only authorized personnel can create and manage bulk installations.
        - **Server-Specific Configurations:** Manages configurations specific to each server, allowing customized installations tailored to individual server requirements.

---

# Data Source: intersight_os_bulk_install_info
The BulkInstallInfo object is a critical element in the automation of operating system installations across multiple servers. It models the content of CSV files that hold configuration parameters for bulk OS installations, ensuring streamlined and efficient deployment.
#### Purpose
BulkInstallInfo facilitates the mass deployment of operating systems by encapsulating necessary parameters within a CSV file. This supports validation processes and manages installation configurations, enabling administrators to efficiently deploy OS across multiple servers.
#### Key Concepts
- **CSV-Based Deployment:** Utilizes CSV files to store OS installation parameters, providing a scalable solution for deploying operating systems to numerous servers.
- **Validation Process:** Includes built-in validation mechanisms to ensure the accuracy and completeness of installation configurations.
- **Access Control:** Enforces privilege sets such as Account Administrator and Manage Servers, ensuring that only authorized personnel can create and manage bulk installations.
- **Server-Specific Configurations:** Manages configurations specific to each server, allowing customized installations tailored to individual server requirements.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_os_bulk_install_info.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `file_content`:(string) The content of the entire CSV file is stored as value. The content can hold complete OS install parameters in two sections.The first section holds generic information about the OS Install like OS Image, SCU Image etc. The second section holdsparameters which are specific to each server level data. 
* `is_file_content_set`:(bool) Indicates whether the value of the 'fileContent' property has been set. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the CSV file, which holds the OS install parameters. 
* `oper_state`:(string) Denotes if the operating is pending, in_progress, completed_ok, completed_error.* `Pending` - The initial value of the OperStatus.* `InProgress` - The OperStatus value will be InProgress during execution.* `CompletedOk` - The API is successful with operation then OperStatus will be marked as CompletedOk.* `CompletedError` - The API is failed with operation then OperStatus will be marked as CompletedError.* `CompletedWarning` - The API is completed with some warning then OperStatus will be CompletedWarning. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
