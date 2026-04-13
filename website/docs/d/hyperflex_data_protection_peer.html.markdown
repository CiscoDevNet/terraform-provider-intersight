---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_data_protection_peer"
description: |-
        The DataProtectionPeer object describes the cluster pair from the HyperFlex backend, emphasizing data protection relationships.
        #### Purpose
        DataProtectionPeer focuses on the relationships between cluster pairs, providing essential information for managing data protection across HyperFlex environments.
        #### Key Concepts
        - **Peer Relationships:** Defines the connections between clusters, ensuring coherent data protection strategies.
        - **Status Tracking:** Offers insights into the status and health of peer clusters, supporting informed decision-making.
        - **Controlled Access:** Manages access through privilege sets, maintaining security and data integrity in peer operations.

---

# Data Source: intersight_hyperflex_data_protection_peer
The DataProtectionPeer object describes the cluster pair from the HyperFlex backend, emphasizing data protection relationships.
#### Purpose
DataProtectionPeer focuses on the relationships between cluster pairs, providing essential information for managing data protection across HyperFlex environments.
#### Key Concepts
- **Peer Relationships:** Defines the connections between clusters, ensuring coherent data protection strategies.
- **Status Tracking:** Offers insights into the status and health of peer clusters, supporting informed decision-making.
- **Controlled Access:** Manages access through privilege sets, maintaining security and data integrity in peer operations.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_data_protection_peer.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
