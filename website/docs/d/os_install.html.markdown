---
subcategory: "os"
layout: "intersight"
page_title: "Intersight: intersight_os_install"
description: |-
        The Install object is the execution point for initiating operating system installations on target servers. It models the necessary configuration and relationships required to start OS installations through Intersight's workflow system.
        #### Purpose
        The Install object captures all required properties and relationships for initiating OS installations. It acts as a launch point for installation workflows, ensuring that all necessary configurations are validated and tracked through Intersight's workflow engine.
        #### Key Concepts
        - **Workflow Integration:** Initiates OS installation workflows, allowing tracking and management through the workflow engine.
        - **Configuration Management:** Encapsulates all necessary installation configurations, ensuring comprehensive setup before deployment.
        - **Access Control:** Enforces privilege sets such as Account Administrator and Manage Servers, safeguarding the initiation and management of installations.
        - **Relationship Management:** Links to critical components such as OS images and configuration files, ensuring seamless installation processes.

---

# Data Source: intersight_os_install
The Install object is the execution point for initiating operating system installations on target servers. It models the necessary configuration and relationships required to start OS installations through Intersight's workflow system.
#### Purpose
The Install object captures all required properties and relationships for initiating OS installations. It acts as a launch point for installation workflows, ensuring that all necessary configurations are validated and tracked through Intersight's workflow engine.
#### Key Concepts
- **Workflow Integration:** Initiates OS installation workflows, allowing tracking and management through the workflow engine.
- **Configuration Management:** Encapsulates all necessary installation configurations, ensuring comprehensive setup before deployment.
- **Access Control:** Enforces privilege sets such as Account Administrator and Manage Servers, safeguarding the initiation and management of installations.
- **Relationship Management:** Links to critical components such as OS images and configuration files, ensuring seamless installation processes.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_os_install.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) User provided description about the OS install configuration. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `error_msg`:(string) The failure message of the API. 
* `install_method`:(string) The install method to be used for OS installation - vMedia, iPXE. Only vMedia is supported as of now.* `vMedia` - OS image is mounted as vMedia in target server for OS installation. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the OS install configuration. 
* `oper_state`:(string) Denotes API operating status as pending, in_progress, completed_ok, completed_error based on the execution status.* `Pending` - The initial value of the OperStatus.* `InProgress` - The OperStatus value will be InProgress during execution.* `CompletedOk` - The API is successful with operation then OperStatus will be marked as CompletedOk.* `CompletedError` - The API is failed with operation then OperStatus will be marked as CompletedError.* `CompletedWarning` - The API is completed with some warning then OperStatus will be CompletedWarning. 
* `override_secure_boot`:(bool) ESXi Secure Boot installation is currently not supported. As a workaround, Secure Boot will be disabled before installation and restored after installation is complete. Enable to Override Secure Boot Configuration. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
