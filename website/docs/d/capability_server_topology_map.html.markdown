---
subcategory: "capability"
layout: "intersight"
page_title: "Intersight: intersight_capability_server_topology_map"
description: |-
        ServerTopologyMaps map specific server models (and related infrastructure components like XFM and PCIe nodes) to a supported PCIe topology configuration and an associated handler identifier. This forms a compatibility matrix that links “what hardware is present” to “which topology definition applies.”
        #### Purpose
        Associate hardware model inventory (server/XFM/PCIe node) with the appropriate supported PCIe topology configuration used for validation and orchestration.
        #### Key Concepts
        - **Model-to-topology mapping:** Connects a server model definition to the topology rules it supports.
        - **Device inventory descriptors:** Uses inventory-like descriptors (model + min/max version bounds) to define compatibility.
        - **Cross-component compatibility:** Captures that topology support depends on the combination of server, XFM, and PCIe node models/versions.
        - **Handler-based processing:** Includes a handler identifier to route topology interpretation/validation to the correct logic.

---

# Data Source: intersight_capability_server_topology_map
ServerTopologyMaps map specific server models (and related infrastructure components like XFM and PCIe nodes) to a supported PCIe topology configuration and an associated handler identifier. This forms a compatibility matrix that links “what hardware is present” to “which topology definition applies.”
#### Purpose
Associate hardware model inventory (server/XFM/PCIe node) with the appropriate supported PCIe topology configuration used for validation and orchestration.
#### Key Concepts
- **Model-to-topology mapping:** Connects a server model definition to the topology rules it supports.
- **Device inventory descriptors:** Uses inventory-like descriptors (model + min/max version bounds) to define compatibility.
- **Cross-component compatibility:** Captures that topology support depends on the combination of server, XFM, and PCIe node models/versions.
- **Handler-based processing:** Includes a handler identifier to route topology interpretation/validation to the correct logic.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_capability_server_topology_map.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `handler`:(string) Handler identifier for managing this topology configuration. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) An unique identifer for a capability descriptor. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `supported_topology_name`:(string) Server model information for which this topology configuration is defined. 
 
