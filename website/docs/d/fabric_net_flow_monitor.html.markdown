---
subcategory: "fabric"
layout: "intersight"
page_title: "Intersight: intersight_fabric_net_flow_monitor"
description: |-
        The NetFlowMonitor object is the orchestrating element that applies monitoring logic to the network. It functions as the operational bridge, linking data collection definitions with export configurations to enable active traffic analysis on specific interfaces.
        ### Purpose
        A NetFlowMonitor serves as the deployment mechanism for telemetry. By combining a NetFlowRecord with one or more NetFlowExporters, it creates a complete monitoring profile that can be applied to virtual or physical interfaces to track real-time traffic patterns.
        ### Key Concepts
        *   **Functional Association:** Binds a single data template (Record) to specific delivery destinations (Exporters), creating a unified monitoring workflow.
        *   **Interface Application:** Acts as the primary object referenced when enabling NetFlow on network interfaces or vNICs.
        *   **Scalable Export:** Supports association with multiple exporters, allowing the same flow data to be sent to different collectors for redundancy or specialized analysis.
        *   **Usage Tracking:** Maintains visibility into how many network components are actively utilizing the monitor, aiding in fabric-wide resource management.

---

# Data Source: intersight_fabric_net_flow_monitor
The NetFlowMonitor object is the orchestrating element that applies monitoring logic to the network. It functions as the operational bridge, linking data collection definitions with export configurations to enable active traffic analysis on specific interfaces.
 ### Purpose
 A NetFlowMonitor serves as the deployment mechanism for telemetry. By combining a NetFlowRecord with one or more NetFlowExporters, it creates a complete monitoring profile that can be applied to virtual or physical interfaces to track real-time traffic patterns.
 ### Key Concepts
 *   **Functional Association:** Binds a single data template (Record) to specific delivery destinations (Exporters), creating a unified monitoring workflow.
 *   **Interface Application:** Acts as the primary object referenced when enabling NetFlow on network interfaces or vNICs.
 *   **Scalable Export:** Supports association with multiple exporters, allowing the same flow data to be sent to different collectors for redundancy or specialized analysis.
 *   **Usage Tracking:** Maintains visibility into how many network components are actively utilizing the monitor, aiding in fabric-wide resource management.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_fabric_net_flow_monitor.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Netflow Monitor name, must be a maximum of 63 characters, without spacing. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `vnic_usage_count`:(int) The count of the NetFlow monitor usage on vNICs. 
 
