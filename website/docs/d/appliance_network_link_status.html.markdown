---
subcategory: "appliance"
layout: "intersight"
page_title: "Intersight: intersight_appliance_network_link_status"
description: |-
        The NetworkLinkStatus object examines the network connection between nodes in an Intersight Appliance cluster. It is vital for ensuring seamless communication and data transfer.
        #### Purpose
        NetworkLinkStatus provides insights into the connectivity and network performance between cluster nodes, facilitating efficient data flow and system integration.
        #### Key Concepts
        - **Link Health:** Assesses the network link status, reporting on connectivity and ping times between nodes.
        - **Communication Efficiency:** Ensures robust and efficient data exchange, crucial for maintaining overall system functionality.
        - **System Integration:** Integrates network link status within broader system monitoring to ensure comprehensive oversight.

---

# Data Source: intersight_appliance_network_link_status
The NetworkLinkStatus object examines the network connection between nodes in an Intersight Appliance cluster. It is vital for ensuring seamless communication and data transfer.
#### Purpose
NetworkLinkStatus provides insights into the connectivity and network performance between cluster nodes, facilitating efficient data flow and system integration.
#### Key Concepts
- **Link Health:** Assesses the network link status, reporting on connectivity and ping times between nodes.
- **Communication Efficiency:** Ensures robust and efficient data exchange, crucial for maintaining overall system functionality.
- **System Integration:** Integrates network link status within broader system monitoring to ensure comprehensive oversight.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_appliance_network_link_status.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `destination_hostname`:(string) Hostname of the destination endpoint. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `ping_time`:(float) Time to reach the destination endpoint in milliseconds from the source endpoint. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `source_hostname`:(string) Hostname of the source endpoint. 
 
