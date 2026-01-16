---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_pure_host_group"
description: |-
        The PureHostGroup object represents a collection of hosts within a PureStorage FlashArray, facilitating common connectivity to storage volumes.
        #### Purpose
        PureHostGroup is used to manage and establish shared connections between hosts and volumes, ensuring efficient access and data sharing across multiple hosts.
        #### Key Concepts
        - **Group Connectivity:** PureHostGroup allows hosts within the group to access volumes through shared connections, streamlining data access and communication.
        - **Access Management:** It supports various privilege sets to control and secure interactions with storage volumes.
        - **Integration:** PureHostGroup integrates with other objects such as PureVolume and PureHost, providing a cohesive storage management solution.

---

# Data Source: intersight_storage_pure_host_group
The PureHostGroup object represents a collection of hosts within a PureStorage FlashArray, facilitating common connectivity to storage volumes.  
#### Purpose  
PureHostGroup is used to manage and establish shared connections between hosts and volumes, ensuring efficient access and data sharing across multiple hosts.  
#### Key Concepts 
- **Group Connectivity:** PureHostGroup allows hosts within the group to access volumes through shared connections, streamlining data access and communication. 
- **Access Management:** It supports various privilege sets to control and secure interactions with storage volumes. 
- **Integration:** PureHostGroup integrates with other objects such as PureVolume and PureHost, providing a cohesive storage management solution.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_pure_host_group.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Short description about the host group. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the host group in storage array. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
