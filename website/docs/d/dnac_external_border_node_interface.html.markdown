---
subcategory: "dnac"
layout: "intersight"
page_title: "Intersight: intersight_dnac_external_border_node_interface"
description: |-
        The ExternalBorderNodeInterface object serves as the gateway between a network and external entities, focusing on interface management and configuration.
        #### Purpose
        An ExternalBorderNodeInterface object provides the framework for managing external network interfaces, supporting secure and efficient data exchange.
        #### Key Concepts
        - **Interface Management:** Offers detailed configuration options for external interfaces, enhancing control over data flow.
        - **Security and Access Control:** Ensures secure interaction through privilege sets and controlled access mechanisms.
        - **Network Integration:** Interfaces seamlessly with external networks, promoting robust and secure connectivity.

---

# Data Source: intersight_dnac_external_border_node_interface
The ExternalBorderNodeInterface object serves as the gateway between a network and external entities, focusing on interface management and configuration.
#### Purpose
An ExternalBorderNodeInterface object provides the framework for managing external network interfaces, supporting secure and efficient data exchange.
#### Key Concepts
- **Interface Management:** Offers detailed configuration options for external interfaces, enhancing control over data flow.
- **Security and Access Control:** Ensures secure interaction through privilege sets and controlled access mechanisms.
- **Network Integration:** Interfaces seamlessly with external networks, promoting robust and secure connectivity.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_dnac_external_border_node_interface.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `admin_status`:(string) Administrator status for external border node interface. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `external_border_node_id`:(string) External border node's id. 
* `interface_id`:(string) The Moid for the interface in the external border node. 
* `interface_name`:(string) The name for the external border node's interface. 
* `interface_type`:(string) Interface type for external border node interface. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `port_type`:(string) Port type for external border node interface. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
