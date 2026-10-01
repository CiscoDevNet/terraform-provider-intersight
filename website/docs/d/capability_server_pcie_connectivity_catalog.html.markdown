---
subcategory: "capability"
layout: "intersight"
page_title: "Intersight: intersight_capability_server_pcie_connectivity_catalog"
description: |-
        ServerPcieConnectivityCatalogs are capability-catalog entries that define *valid, supported* physical PCIe connectivity layouts for servers. They model the expected end-to-end topology (CPU → PCIe switch/fabric elements → GPU/adapter endpoints) so the system can validate whether a requested or discovered wiring/layout is supported.
        #### Purpose
        Provide a canonical catalog of supported server-to-PCIe-device physical topologies used for topology validation and compatibility checks.
        #### Key Concepts
        - **Topology validation catalog:** Encodes known-good layouts that can be matched against intended configuration or discovered inventory.
        - **Layout-driven compatibility:** A layout describes how CPUs, slots, and PCIe switches relate to connected GPUs/adapters.
        - **Endpoint mapping primitives:** Uses connection-point groupings (GPU/adapter connection points) to describe where endpoints are attached in the topology.
        - **Capability model integration:** Implemented as a `capability.Capability`- derived object, intended to be curated/managed like other capability catalog entries.

---

# Data Source: intersight_capability_server_pcie_connectivity_catalog
ServerPcieConnectivityCatalogs are capability-catalog entries that define *valid, supported* physical PCIe connectivity layouts for servers. They model the expected end-to-end topology (CPU → PCIe switch/fabric elements → GPU/adapter endpoints) so the system can validate whether a requested or discovered wiring/layout is supported.
#### Purpose
Provide a canonical catalog of supported server-to-PCIe-device physical topologies used for topology validation and compatibility checks.
#### Key Concepts
- **Topology validation catalog:** Encodes known-good layouts that can be matched against intended configuration or discovered inventory.
- **Layout-driven compatibility:** A layout describes how CPUs, slots, and PCIe switches relate to connected GPUs/adapters.
- **Endpoint mapping primitives:** Uses connection-point groupings (GPU/adapter connection points) to describe where endpoints are attached in the topology.
- **Capability model integration:** Implemented as a `capability.Capability`- derived object, intended to be curated/managed like other capability catalog entries.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_capability_server_pcie_connectivity_catalog.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) An unique identifer for a capability descriptor. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
