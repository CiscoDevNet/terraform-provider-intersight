---
subcategory: "dnac"
layout: "intersight"
page_title: "Intersight: intersight_dnac_transit"
description: |-
        The Transit object is integral to network connectivity, representing transit paths and configurations that facilitate communication between different network segments.
        #### Purpose
        A Transit object defines the connectivity pathways within the network, enabling efficient routing and communication between distinct network areas.
        #### Key Concepts
        - **Connectivity Management:** Facilitates the establishment and management of transit paths, optimizing network traffic flow.
        - **Identity Control:** Utilizes unique identifiers for precise tracking and management of transit entities.
        - **Integration with Network Elements:** Seamlessly interfaces with other network components, enhancing overall network performance.

---

# Data Source: intersight_dnac_transit
The Transit object is integral to network connectivity, representing transit paths and configurations that facilitate communication between different network segments.
#### Purpose
A Transit object defines the connectivity pathways within the network, enabling efficient routing and communication between distinct network areas.
#### Key Concepts
- **Connectivity Management:** Facilitates the establishment and management of transit paths, optimizing network traffic flow.
- **Identity Control:** Utilizes unique identifiers for precise tracking and management of transit entities.
- **Integration with Network Elements:** Seamlessly interfaces with other network components, enhancing overall network performance.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_dnac_transit.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `transit_id`:(string) Identification for the Transit. 
* `transit_name`:(string) Name identifier for the Transit. 
* `transit_type`:(string) Transit type for the transit. 
 
