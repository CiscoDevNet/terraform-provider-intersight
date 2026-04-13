---
subcategory: "os"
layout: "intersight"
page_title: "Intersight: intersight_os_distribution"
description: |-
        The Distribution object encompasses the details of operating system distributions that are validated and supported by Intersight for automated installation. It acts as a repository of supported OS versions and their associated metadata.
        #### Purpose
        Distribution objects serve as a reference for all supported operating system distributions within Intersight, ensuring compatibility and reliability of automated installations. They provide essential metadata required for managing and deploying these OS distributions efficiently.
        #### Key Concepts
        - **OS Validation:** Contains validated OS distribution details, ensuring reliability and compatibility with Intersight's automated installation services.
        - **Intersight-Supported:** Lists OS distributions that are officially supported for automated installation, backed by thorough validation processes.
        - **Integration with Configuration Files:** Works in conjunction with OS configuration files and post-install scripts to facilitate seamless installation processes.
        - **Access Control:** Privileges such as System Administrator and Account Administrator ensure secure access and management of distribution details.

---

# Data Source: intersight_os_distribution
The Distribution object encompasses the details of operating system distributions that are validated and supported by Intersight for automated installation. It acts as a repository of supported OS versions and their associated metadata.
#### Purpose
Distribution objects serve as a reference for all supported operating system distributions within Intersight, ensuring compatibility and reliability of automated installations. They provide essential metadata required for managing and deploying these OS distributions efficiently.
#### Key Concepts
- **OS Validation:** Contains validated OS distribution details, ensuring reliability and compatibility with Intersight's automated installation services.
- **Intersight-Supported:** Lists OS distributions that are officially supported for automated installation, backed by thorough validation processes.
- **Integration with Configuration Files:** Works in conjunction with OS configuration files and post-install scripts to facilitate seamless installation processes.
- **Access Control:** Privileges such as System Administrator and Account Administrator ensure secure access and management of distribution details.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_os_distribution.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `is_supported`:(bool) An internal property that is used to denote if the OS Distribution is supportedby Intersight for Automated Installation. 
* `label`:(string) The label of the OS distribution such as ESXi, CentOS to be displayed. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the OS distribution such as ESXi, CentOS. 
* `scu_supported`:(bool) An internal property that is used to denote if the OS Distribution is supportedby the Server Configuration Utility. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
