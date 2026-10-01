---
subcategory: "fabric"
layout: "intersight"
page_title: "Intersight: intersight_fabric_span_source_eth_port"
description: |-
        The SpanSourceEthPort object represents an Ethernet physical port configured as a SPAN traffic source for a session.
        #### Purpose
        SpanSourceEthPort enables a specific Ethernet port’s traffic (ingress/egress/both) to be mirrored into a SPAN session, supporting port-level traffic analysis and diagnostics.
        #### Key Concepts
        - **Port-level mirroring:** Mirrors traffic from a specific Ethernet port.
        - **Direction semantics:** Supports selecting ingress, egress, or both directions for capture.
        - **Session-scoped membership:** Acts as a source member within a SPAN session.
        - **Topology-aware targeting:** Uses slot/port/breakout identifiers consistent with port modeling.

---

# Data Source: intersight_fabric_span_source_eth_port
The SpanSourceEthPort object represents an Ethernet physical port configured as a SPAN traffic source for a session.
#### Purpose
SpanSourceEthPort enables a specific Ethernet port’s traffic (ingress/egress/both) to be mirrored into a SPAN session, supporting port-level traffic analysis and diagnostics.
#### Key Concepts
- **Port-level mirroring:** Mirrors traffic from a specific Ethernet port.
- **Direction semantics:** Supports selecting ingress, egress, or both directions for capture.
- **Session-scoped membership:** Acts as a source member within a SPAN session.
- **Topology-aware targeting:** Uses slot/port/breakout identifiers consistent with port modeling.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_fabric_span_source_eth_port.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `aggregate_port_id`:(int) Breakout port Identifier of the Switch Interface.When a port is not configured as a breakout port, the aggregatePortId is set to 0, and unused.When a port is configured as a breakout port, the 'aggregatePortId' port number as labeled on the equipment,e.g. the id of the port on the switch. 
* `create_time`:(string) The time when this managed object was created. 
* `direction`:(string) Direction of the source SPAN.* `Receive` - SPAN incoming traffic on the SPAN source interface.* `Transmit` - SPAN outgoing traffic on the SPAN source interface.* `Both` - SPAN incoming and outgoing traffic on the SPAN source interface. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `port_id`:(int) Port Identifier of the Switch Interface.When a port is not configured as a breakout port, the portId is the port number as labeled on the equipment,e.g. the id of the port on the switch, FEX or chassis.When a port is configured as a breakout port, the 'portId' represents the port id on the fanout side of the breakout cable. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `slot_id`:(int) Slot Identifier of the Switch Interface. 
* `source_role`:(string) Role of the port configured as SPAN source.* `Uplink` - Uplink Role corresponding to PortRole in PortPolicy.* `FcoeUplink` - FcoeUplink Role corresponding to PortRole in PortPolicy.* `FcUplink` - FcoeUplink Role corresponding to PortRole in PortPolicy.* `Appliance` - FcoeUplink Role corresponding to PortRole in PortPolicy. 
 
