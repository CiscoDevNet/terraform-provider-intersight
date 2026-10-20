---
subcategory: "fabric"
layout: "intersight"
page_title: "Intersight: intersight_fabric_vlan_inventory"
description: |-
        The VlanInventory object represents VLAN configuration and status inventoried from a network device, exposing what is actually present and reported by the endpoint.
        #### Purpose
        VlanInventory provides a read-only inventory view of VLANs as they exist on the device, enabling operational visibility, auditing, and reconciliation against intended policy configuration.
        #### Key Concepts
        - **Device-sourced truth:** Represents the device’s reported VLAN state rather than desired configuration.
        - **Operational troubleshooting:** Supports diagnosing mismatches and deployment issues by showing actual device status.
        - **Inventory correlation:** Links VLAN inventory to the associated network element and device registration context.
        - **Consumption-oriented:** Built primarily for read/monitor workflows and reporting.

---

# Data Source: intersight_fabric_vlan_inventory
The VlanInventory object represents VLAN configuration and status inventoried from a network device, exposing what is actually present and reported by the endpoint.
#### Purpose
VlanInventory provides a read-only inventory view of VLANs as they exist on the device, enabling operational visibility, auditing, and reconciliation against intended policy configuration.
#### Key Concepts
- **Device-sourced truth:** Represents the device’s reported VLAN state rather than desired configuration.
- **Operational troubleshooting:** Supports diagnosing mismatches and deployment issues by showing actual device status.
- **Inventory correlation:** Links VLAN inventory to the associated network element and device registration context.
- **Consumption-oriented:** Built primarily for read/monitor workflows and reporting.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_fabric_vlan_inventory.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `admin_state`:(string) Admin State of the VLAN attribute in the switch. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the VLAN Inventory for the switch. 
* `port`:(string) Ports associated to the VLAN. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `status`:(string) Status of the VLAN attribute in the switch. 
* `vlan_id`:(string) The Identifier for VLAN in the switch. 
 
