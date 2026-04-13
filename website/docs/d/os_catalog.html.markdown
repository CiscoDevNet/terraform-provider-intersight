---
subcategory: "os"
layout: "intersight"
page_title: "Intersight: intersight_os_catalog"
description: |-
        The Catalog object is an integral component of the operating system management framework, specifically designed to handle collections of operating system-related artifacts such as configuration files and post-installation scripts. It is pivotal in organizing and managing these artifacts within user accounts and system contexts.
        #### Purpose
        A Catalog serves as a repository for operating system configuration files and scripts, facilitating both private storage by users and shared access managed by Cisco. It enables efficient OS installations by providing validated files and scripts, which users can leverage in their Intersight account environments.
        #### Key Concepts
        - **Object Organization:** Catalogs offer structured storage for configuration files and scripts, ensuring easy access and management within user accounts.
        - **Cisco-Managed Shared Catalogs:** Cisco manages shared catalogs, allowing users to read and utilize validated files and scripts for OS installations. Contributions to shared catalogs are published by Cisco at its discretion.
        - **Access Control:** Different privilege sets are enforced to ensure secure and authorized access, with Account Administrators and View Servers privileges being essential for reading catalog contents.
        - **Relationship Management:** Catalog objects maintain relationships with other components, such as configuration files and OS distributions, fostering comprehensive management and integration.

---

# Data Source: intersight_os_catalog
The Catalog object is an integral component of the operating system management framework, specifically designed to handle collections of operating system-related artifacts such as configuration files and post-installation scripts. It is pivotal in organizing and managing these artifacts within user accounts and system contexts.
#### Purpose
A Catalog serves as a repository for operating system configuration files and scripts, facilitating both private storage by users and shared access managed by Cisco. It enables efficient OS installations by providing validated files and scripts, which users can leverage in their Intersight account environments.
#### Key Concepts
 - **Object Organization:** Catalogs offer structured storage for configuration files and scripts, ensuring easy access and management within user accounts.
 - **Cisco-Managed Shared Catalogs:** Cisco manages shared catalogs, allowing users to read and utilize validated files and scripts for OS installations. Contributions to shared catalogs are published by Cisco at its discretion.
 - **Access Control:** Different privilege sets are enforced to ensure secure and authorized access, with Account Administrators and View Servers privileges being essential for reading catalog contents.
 - **Relationship Management:** Catalog objects maintain relationships with other components, such as configuration files and OS distributions, fostering comprehensive management and integration.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_os_catalog.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The catalog name. There will be one for system and one for each user account. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
