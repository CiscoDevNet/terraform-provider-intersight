---
subcategory: "dnac"
layout: "intersight"
page_title: "Intersight: intersight_dnac_fabric_site"
description: |-
        The FabricSite object is a critical component within the network management framework, designed to represent and manage the physical or logical sites within the network fabric. This provides a structured approach to organizing and accessing site-specific details, aiding in efficient network operations and resource allocation.
        #### Purpose
        A FabricSite facilitates the management and operational understanding of network sites. This provides a holistic view of site configurations, enabling seamless integration and interaction with other network entities and services.
        #### Key Concepts
        - **Hierarchy Representation:** FabricSite objects maintain a hierarchical structure, allowing for organized and intuitive navigation of site relationships.
        - **Access Control:** Privilege sets ensure secure and controlled access, allowing only authorized users to interact with FabricSite data.
        - **Integration:** FabricSite objects are integral to workflows and network operations, supporting robust interactions with devices, interfaces, and services.

---

# Data Source: intersight_dnac_fabric_site
The FabricSite object is a critical component within the network management framework, designed to represent and manage the physical or logical sites within the network fabric. This provides a structured approach to organizing and accessing site-specific details, aiding in efficient network operations and resource allocation.
#### Purpose
A FabricSite facilitates the management and operational understanding of network sites. This provides a holistic view of site configurations, enabling seamless integration and interaction with other network entities and services.
#### Key Concepts
- **Hierarchy Representation:** FabricSite objects maintain a hierarchical structure, allowing for organized and intuitive navigation of site relationships.
- **Access Control:** Privilege sets ensure secure and controlled access, allowing only authorized users to interact with FabricSite data.
- **Integration:** FabricSite objects are integral to workflows and network operations, supporting robust interactions with devices, interfaces, and services.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_dnac_fabric_site.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `authentication_profile_name`:(string) Authentication profile name. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `fabric_site_id`:(string) UUID for the Fabric Site. 
* `fabric_site_name_hierarchy`:(string) Fabric site name hierarchy. 
* `is_pub_sub_enabled`:(string) Pub sub for the fabric site. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `site_id`:(string) Site id for the fabric site. 
 
