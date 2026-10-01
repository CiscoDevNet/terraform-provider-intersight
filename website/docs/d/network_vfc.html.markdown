---
subcategory: "network"
layout: "intersight"
page_title: "Intersight: intersight_network_vfc"
description: |-
        Vfcs represent virtual Fibre Channel interfaces (vHBA peers) configured on a Fabric Interconnect. They track binding/pinning to FC ports or FC port-channels and provide operational state/reason and QoS shaping parameters.
        #### Purpose
        Expose vHBA peer inventory on the Fabric Interconnect so operators can validate SAN bindings, troubleshoot link state, and correlate to server-side host FC interfaces.
        #### Key Concepts
        - **Virtual SAN interface inventory:** Represents a vFC endpoint on the switch side.
        - **Binding and pinning:** Identifies the configured interface DN and pinned uplink DN used for vFC traffic.
        - **Operational diagnostics:** Provides operState and operReason to explain down/no-license/error-disabled conditions.
        - **Adapter correlation:** Can relate back to the corresponding adapter host FC interface.
        - **QoS visibility:** Includes rate limit and burst parameters applied to the vFC.

---

# Data Source: intersight_network_vfc
Vfcs represent virtual Fibre Channel interfaces (vHBA peers) configured on a Fabric Interconnect. They track binding/pinning to FC ports or FC port-channels and provide operational state/reason and QoS shaping parameters.
#### Purpose
Expose vHBA peer inventory on the Fabric Interconnect so operators can validate SAN bindings, troubleshoot link state, and correlate to server-side host FC interfaces.
#### Key Concepts
- **Virtual SAN interface inventory:** Represents a vFC endpoint on the switch side.
- **Binding and pinning:** Identifies the configured interface DN and pinned uplink DN used for vFC traffic.
- **Operational diagnostics:** Provides operState and operReason to explain down/no-license/error-disabled conditions.
- **Adapter correlation:** Can relate back to the corresponding adapter host FC interface.
- **QoS visibility:** Includes rate limit and burst parameters applied to the vFC.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_network_vfc.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `bound_interface_dn`:(string) Port or PortChannel interface configured for Vfcs virtual etherent interface. 
* `burst`:(int) Burst defined on QoS policy. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description for the vHBA. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `oper_reason`:(string) Operational Reason of the Vfcs Vethernet on the Fabric Interconnect. When the operational state is down, Operreason indicates the reason why the vHBA is not operational. Some common reasons are admindown, error disabled. 
* `oper_state`:(string) The operational state of the Vfcs Vethernet peer of a vNIC in Intersight Managed Mode. This state is updated by events from Fabric Interconnect or by periodic updates from Fabric Interconnect. When a Fabric Interconnect is not connected to Intersight or if the Fabric Interconnect is powered down, this property will not be updated.* `unknown` - The operational state of the Vethernet is not known.* `adminDown` - The operational state of the Vethernet is admin down.* `up` - The operational state of the Vethernet is Up.* `down` - The operational state of the Vethernet is Down.* `noLicense` - The operational state of the Vethernet is no license.* `linkUp` - The operational state of the Vethernet is link up.* `hardwareFailure` - The operational state of the Vethernet is hardware failure.* `softwareFailure` - The operational state of the Vethernet is software failure.* `errorDisabled` - The operational state of the Vethernet is error disabled.* `linkDown` - The operational state of the Vethernet is link down.* `sfpNotPresent` - The operational state of the Vethernet is SFP not present.* `udldAggrDown` - The operational state of the Vethernet is UDLD aggregate down. 
* `pinned_interface_dn`:(string) Uplink port or portchannel pinned to a Vfc. 
* `ratelimit`:(int) Rate limit defined on QoS policy. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `vfc_id`:(int) Vfc Identifier on a Fabric Interconnect. 
 
