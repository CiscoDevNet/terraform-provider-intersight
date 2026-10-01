---
subcategory: "fabric"
layout: "intersight"
page_title: "Intersight: intersight_fabric_lan_pin_group"
description: |-
        The LanPinGroup object represents a LAN static pinning group that defines which uplink interface roles are eligible pin targets for LAN traffic.
        #### Purpose
        LanPinGroup enables deterministic LAN pinning behavior by grouping eligible uplink interfaces. It supports consistent and repeatable pinning for LAN traffic (e.g., vNIC pinning), improving predictability and reducing reliance on dynamic selection.
        #### Key Concepts
        - **Static pinning model (LAN):** Defines deterministic uplink selection behavior for LAN traffic.
        - **Interface-role targeting:** References eligible uplink roles (uplink ports or uplink port-channels).
        - **Policy-scoped identity:** Managed within the scope of a port policy.
        - **Traffic predictability:** Promotes stable LAN traffic placement and troubleshooting clarity.

---

# Data Source: intersight_fabric_lan_pin_group
The LanPinGroup object represents a LAN static pinning group that defines which uplink interface roles are eligible pin targets for LAN traffic.
#### Purpose
LanPinGroup enables deterministic LAN pinning behavior by grouping eligible uplink interfaces. It supports consistent and repeatable pinning for LAN traffic (e.g., vNIC pinning), improving predictability and reducing reliance on dynamic selection.
#### Key Concepts
- **Static pinning model (LAN):** Defines deterministic uplink selection behavior for LAN traffic.
- **Interface-role targeting:** References eligible uplink roles (uplink ports or uplink port-channels).
- **Policy-scoped identity:** Managed within the scope of a port policy.
- **Traffic predictability:** Promotes stable LAN traffic placement and troubleshooting clarity.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_fabric_lan_pin_group.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the Pingroup for static pinning. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
