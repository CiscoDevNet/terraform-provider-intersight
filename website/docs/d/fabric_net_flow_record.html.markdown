---
subcategory: "fabric"
layout: "intersight"
page_title: "Intersight: intersight_fabric_net_flow_record"
description: |-
        The NetFlowRecord object is a fundamental component of the network telemetry system, designed to define the specific structure and attributes of the traffic data to be collected. It acts as a template that determines which characteristics of a network flow are tracked and recorded.
        ### Purpose
        A NetFlowRecord defines the what of network monitoring. It specifies the criteria used to identify unique flows and the statistics to be gathered for those flows. By configuring a record, administrators can customize visibility into IPv4, IPv6, or Layer 2 traffic based on organizational requirements.
        ### Key Concepts
        *   **Flow Identification (Keys):** Uses specific packet fields—such as source/destination addresses, ports, and protocols—to distinguish one network flow from another.
        *   **Telemetry Collection (Non-Keys):** Captures operational data including packet and byte counters, as well as system timestamps for flow start and end times.
        *   **Protocol Versatility:** Supports multiple network layers, allowing for specialized monitoring of different traffic types (L2, IPv4, and IPv6).

---

# Data Source: intersight_fabric_net_flow_record
The NetFlowRecord object is a fundamental component of the network telemetry system, designed to define the specific structure and attributes of the traffic data to be collected. It acts as a template that determines which characteristics of a network flow are tracked and recorded.
### Purpose
A NetFlowRecord defines the "what" of network monitoring. It specifies the criteria used to identify unique flows and the statistics to be gathered for those flows. By configuring a record, administrators can customize visibility into IPv4, IPv6, or Layer 2 traffic based on organizational requirements.
### Key Concepts
*   **Flow Identification (Keys):** Uses specific packet fields—such as source/destination addresses, ports, and protocols—to distinguish one network flow from another.
*   **Telemetry Collection (Non-Keys):** Captures operational data including packet and byte counters, as well as system timestamps for flow start and end times.
*   **Protocol Versatility:** Supports multiple network layers, allowing for specialized monitoring of different traffic types (L2, IPv4, and IPv6).
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_fabric_net_flow_record.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Netflow record name. Must be a maximum of 63 characters, without spacing. 
* `record_type`:(string) Netflow Record Type IPv4, IPv6 and L2.* `Invalid` - Netflow record invlaid type.* `IPv4` - Netflow record type for IPv4 packet.* `IPv6` - Netflow record type for IPv6 packet.* `L2` - Netflow record type for L2 packet. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
