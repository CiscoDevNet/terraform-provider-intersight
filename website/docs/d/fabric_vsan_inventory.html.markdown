---
subcategory: "fabric"
layout: "intersight"
page_title: "Intersight: intersight_fabric_vsan_inventory"
description: |-
        Inventory object for VSAN.

---

# Data Source: intersight_fabric_vsan_inventory
Inventory object for VSAN.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_fabric_vsan_inventory.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `admin_state`:(string) Admin state of the VSAN. It can be active or inactive.* `` - Default value to represent the administrative state of isolated vsan.* `Up` - VSAN administrative state is up.* `Down` - VSAN administrative state is down. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `interop_mode`:(string) Interoperability mode of the VSAN. It enables products of multiple vendors to communicate with each other.* `` - Default value to represent the interoperability mode of isolated vsan when it is not configured.* `Default` - Default mode for a VSAN that is communicating between a SAN composed entirely of MDS 9000 switches.* `1` - Allows integration with Brocade and McData switches that have been configured for their own interoperability modes. Brocade and McData switches must be running in interop mode to work with this VSAN mode.* `2` - Allows seamless integration with specific Brocade switches running in their own native mode of operation.* `3` - Allows seamless integration with Brocade switches that contains more than 16 ports.* `4` - Allows seamless integration between MDS VSANs and McData switches running in McData Fabric 1.0 interop mode. 
* `load_balancing`:(string) These attributes indicate the use of the source-destination ID or the originator exchange OX ID for load balancing path selection.* `` - Default value to represent the load balancing scheme of isolated vsan.* `src-id/dst-id` - Directs the switch to use the source and destination ID for its path selection process.* `src-id/dst-id/oxid` - Directs the switch to use the source ID, the destination ID, and the OX ID for its path selection process. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) User-specified name of the VSAN. 
* `oper_state`:(string) Operational state of the VSAN.* `` - Default value to represent the operational state of isolated vsan.* `Up` - VSAN operational state is up.* `Down` - VSAN operational state is down. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `smart_zoning`:(string) Smart zoning status on the VSAN. It can be enabled or disabled.* `` - Default value to represent the smart zoning status of isolated vsan.* `Enabled` - VSAN smart zoning state is enabled.* `Disabled` - VSAN smart zoning state is disabled. 
* `vsan_id`:(int) Identifier for the VSAN. It is an integer from 1 to 4094. 
 
