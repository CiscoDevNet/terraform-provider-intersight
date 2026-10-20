---
subcategory: "fabric"
layout: "intersight"
page_title: "Intersight: intersight_fabric_remote_span_exporter"
description: |-
        Configures an origin interface for ERSPAN tunnel frames on the Fabric Interconnect.

---

# Data Source: intersight_fabric_remote_span_exporter
Configures an origin interface for ERSPAN tunnel frames on the Fabric Interconnect.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_fabric_remote_span_exporter.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `config_state`:(string) The configured state of Origin Interface configuration. If the configuration fails to deploy to the Fabric Interconnect, it can be redeployed by toggling the admin state.* `None` - Nil value when no action has been triggered by the user.* `Applied` - User configured settings are in applied state.* `Applying` - User settings are being applied on the target server.* `Failed` - User configured settings could not be applied.* `Scheduled` - User configured settings are scheduled to be applied. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `gateway`:(string) Source gateway that acts as the origin interface for ERSPAN tunnel frames on the Fabric Interconnect. 
* `ip_address_source`:(string) Indicates whether the effective ERSPAN exporter IP address is sourced from this configuration or shared from a NetFlow exporter on the same VLAN. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `oper_state`:(string) Operational state of the Origin Interface.* `Unknown` - SPAN session operational state is Unknown.* `Up` - SPAN session operational state is Up.* `Down` - SPAN session operational state is Down.* `Error` - SPAN session operational state is Error. 
* `oper_state_reason`:(string) Operational state reason of the SPAN session.* `Unknown` - Unknown operational state.* `Active` - Active and operational SPAN session.* `NoHardwareResource` - No hardware resources available.* `NoOperationalSrcDst` - No operational SPAN source or destination.* `GenericError` - Generic operational error.* `NoSourcesConfigured` - No source interfaces configured.* `NoDestinationConfigured` - No destination port configured.* `NoSourceDestinationConfigured` - No source or destination interface configured.* `SessionAdminShut` - Session is administratively disabled.* `WrongDestinationMode` - Wrong Destination mode configured.* `WrongSourceMode` - Wrong Source mode configured.* `TunnelMisconfDown` - Tunnel Misconfigured or Down.* `NoFlowIdSpecified` - No Flow ID specified for ERSPAN.* `NoErspanIdSpecified` - No ERSPAN ID is configured.* `MultiDestinationDenied` - Multiple destination is not allowed.* `VrfInvalid` - No valid VRF is configured.* `IpInvalid` - No valid destination IP address configured.* `GlobalIpInvalid` - No valid global origin IP address configured.* `NoRoute` - No route to destination IP address.* `ArpNotResolved` - ARP entry for destination is not resolved.* `NoEgressInterface` - Egress interface is not resolved.* `NoSviMember` - SVI member was not retrieved. 
* `operational_ip_address`:(string) Effective IP address programmed on the Fabric Interconnect for the ERSPAN exporter origin interface. When the exporter VLAN is shared with NetFlow, this reflects the NetFlow-owned SVI address; otherwise it matches sourceIp. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `source_ip`:(string) Source Interface IP Address that acts as the origin interface for ERSPAN tunnel frames on the Fabric Interconnect. 
* `source_netmask`:(string) Source Interface Netmask that acts as the origin interface for ERSPAN tunnel frames on the Fabric Interconnect. 
* `source_vlan_id`:(int) Source Interface VLAN ID that acts as the origin interface for ERSPAN tunnel frames on the Fabric Interconnect. 
* `switch_id`:(string) Switch ID (A or B) corresponding to network element. 
 
