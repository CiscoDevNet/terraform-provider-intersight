---
subcategory: "management"
layout: "intersight"
page_title: "Intersight: intersight_management_entity"
description: |-
        Entities (management) represent the logical role and clustering state of each Fabric Interconnect in UCS Manager (e.g., Primary/Subordinate and cluster readiness/state/link state).
        #### Purpose
        Expose FI role and cluster state so administrators can understand HA posture and interconnect leadership.
        #### Key Concepts
        - **FI role modeling:** Captures leadership (Primary/Subordinate).
        - **Cluster state visibility:** Includes readiness, cluster state, and cluster link/umbilical state.
        - **UCSM logical layer:** Represents the UCSM logical perspective of FI clustering.

---

# Data Source: intersight_management_entity
Entities (management) represent the logical role and clustering state of each Fabric Interconnect in UCS Manager (e.g., Primary/Subordinate and cluster readiness/state/link state).
#### Purpose
Expose FI role and cluster state so administrators can understand HA posture and interconnect leadership.
#### Key Concepts
- **FI role modeling:** Captures leadership (Primary/Subordinate).
- **Cluster state visibility:** Includes readiness, cluster state, and cluster link/umbilical state.
- **UCSM logical layer:** Represents the UCSM logical perspective of FI clustering.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_management_entity.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `cluster_link_state`:(string) Cluster link state between the Fabric Interconnects. 
* `cluster_readiness`:(string) Cluster readiness of the Fabric Interconnect. 
* `cluster_state`:(string) Cluster state of the Fabric Interconnect. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `entity_id`:(string) Identity of the Fabric Interconnect - A/B. 
* `leadership`:(string) Role (Primary / Subordinate) of the Fabric Interconnect. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
