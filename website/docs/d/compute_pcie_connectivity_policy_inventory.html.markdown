---
subcategory: "compute"
layout: "intersight"
page_title: "Intersight: intersight_compute_pcie_connectivity_policy_inventory"
description: |-
        PcieConnectivityPolicies define the intended PCIe connectivity for a server profile by describing one or more PCIe zones. Each zone specifies a root PCIe endpoint (CPU) and a set of PCIe endpoint targets (GPUs and/or adapters) selected via property filters such as model and count.
        #### Purpose
        Express user intent for mapping PCIe devices (GPUs/adapters) to server CPUs so the platform can apply and enforce a supported PCIe connectivity configuration.
        #### Key Concepts
        - **Zone-based intent model:** A policy is composed of one or more zones that group endpoints under a root CPU selection.
        - **Root vs target endpoints:** Distinguishes initiators (CPU/root endpoint) from targets (GPU/adapter endpoints).
        - **Property-filtered selection:** Endpoints can be selected by model and count, enabling reusable intent across compatible hardware.
        - **Profile attachability:** Designed to attach to server profiles and participate in deployment workflows where policy intent becomes applied configuration.

---

# Data Source: intersight_compute_pcie_connectivity_policy_inventory
PcieConnectivityPolicies define the intended PCIe connectivity for a server profile by describing one or more PCIe zones. Each zone specifies a root PCIe endpoint (CPU) and a set of PCIe endpoint targets (GPUs and/or adapters) selected via property filters such as model and count.
#### Purpose
Express user intent for mapping PCIe devices (GPUs/adapters) to server CPUs so the platform can apply and enforce a supported PCIe connectivity configuration.
#### Key Concepts
- **Zone-based intent model:** A policy is composed of one or more zones that group endpoints under a root CPU selection.
- **Root vs target endpoints:** Distinguishes initiators (CPU/root endpoint) from targets (GPU/adapter endpoints).
- **Property-filtered selection:** Endpoints can be selected by model and count, enabling reusable intent across compatible hardware.
- **Profile attachability:** Designed to attach to server profiles and participate in deployment workflows where policy intent becomes applied configuration.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_compute_pcie_connectivity_policy_inventory.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the policy. 
* `device_mo_id`:(string) Device ID of the entity from where inventory is reported. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the inventoried policy object. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
