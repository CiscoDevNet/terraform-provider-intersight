---
subcategory: "dnac"
layout: "intersight"
page_title: "Intersight: intersight_dnac_virtual_network_fabric_site"
description: |-
        The VirtualNetworkFabricSite object encapsulates the virtual networking aspects of a fabric site, enabling advanced network configurations and interactions. It plays a pivotal role in the deployment and management of virtual networks within the fabric.
        #### Purpose
        A VirtualNetworkFabricSite object facilitates the establishment and management of virtual networks associated with a fabric site, offering flexibility and scalability in network architecture.
        #### Key Concepts
        - **Virtual Network Management:** Supports the creation, monitoring, and modification of virtual networks, enhancing network adaptability.
        - **Identity and Access Control:** Ensures secure and efficient management through unique identifiers and privilege sets.
        - **Network Integration:** Seamlessly integrates with physical and logical network elements, fostering cohesive network operations.

---

# Data Source: intersight_dnac_virtual_network_fabric_site
The VirtualNetworkFabricSite object encapsulates the virtual networking aspects of a fabric site, enabling advanced network configurations and interactions. It plays a pivotal role in the deployment and management of virtual networks within the fabric.
#### Purpose
A VirtualNetworkFabricSite object facilitates the establishment and management of virtual networks associated with a fabric site, offering flexibility and scalability in network architecture.
#### Key Concepts
- **Virtual Network Management:** Supports the creation, monitoring, and modification of virtual networks, enhancing network adaptability.
- **Identity and Access Control:** Ensures secure and efficient management through unique identifiers and privilege sets.
- **Network Integration:** Seamlessly integrates with physical and logical network elements, fostering cohesive network operations.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_dnac_virtual_network_fabric_site.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `fabric_site_name_hierarchy`:(string) Fabric site name hierarchy. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `site_id`:(string) Site id for the virtual network fabric site. 
* `virtual_network_id`:(string) Virtual network id fro the virtual network fabric site. 
* `virtual_network_name`:(string) Virtual network name for the virtual network fabric site. 
 
