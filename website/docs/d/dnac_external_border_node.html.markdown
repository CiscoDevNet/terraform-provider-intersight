---
subcategory: "dnac"
layout: "intersight"
page_title: "Intersight: intersight_dnac_external_border_node"
description: |-
        The ExternalBorderNode object represents the external boundaries of a network, focusing on the management and configuration of nodes that interact with external network entities.
        #### Purpose
        An ExternalBorderNode object is essential for defining and managing the interfaces between internal and external networks, ensuring secure and efficient communication.
        #### Key Concepts
        - **Boundary Definition:** Specifies the roles and configurations of external border nodes, clarifying their position within the network architecture.
        - **Security and Roles Management:** Enables detailed role assignment and security configuration, safeguarding network boundaries.
        - **Interoperability:** Designed to facilitate interaction with external networks, promoting seamless and secure connectivity.

---

# Data Source: intersight_dnac_external_border_node
The ExternalBorderNode object represents the external boundaries of a network, focusing on the management and configuration of nodes that interact with external network entities.
#### Purpose
An ExternalBorderNode object is essential for defining and managing the interfaces between internal and external networks, ensuring secure and efficient communication.
#### Key Concepts
- **Boundary Definition:** Specifies the roles and configurations of external border nodes, clarifying their position within the network architecture.
- **Security and Roles Management:** Enables detailed role assignment and security configuration, safeguarding network boundaries.
- **Interoperability:** Designed to facilitate interaction with external networks, promoting seamless and secure connectivity.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_dnac_external_border_node.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `external_border_node_id`:(string) External border node's id. 
* `external_border_node_name`:(string) External border node's name. 
* `fabric_site_id`:(string) Fabric Site id in UUID format. 
* `import_external_routes`:(bool) Flag to determine if Border Node is External or Internal. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
