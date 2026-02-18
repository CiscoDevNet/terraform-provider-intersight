---
subcategory: "fabric"
layout: "intersight"
page_title: "Intersight: intersight_fabric_span_session"
description: |-
        The SpanSession object defines SPAN (Switched Port Analyzer) sessions used for traffic monitoring and analysis. It supports network visibility by directing selected traffic to monitoring interfaces. The operational state and reason fields reflect the status of the SPAN session itself, not underlying network issues.
        #### Purpose
        SpanSession serves to establish and manage SPAN sessions, allowing administrators to capture and analyze network traffic for troubleshooting and performance monitoring. It plays a crucial role in network diagnostics and operational efficiency.
        #### Key Concepts
        - **Traffic Monitoring:** Allows the configuration of sources and destinations for traffic capture, providing insights into network performance and issues.
        - **Dynamic Configuration:** Supports the dynamic setup and adjustment of SPAN sessions, enabling flexible and responsive network management.
        - **Integrated Network Management:** Works in conjunction with network elements to ensure comprehensive monitoring capabilities, enhancing overall network oversight.
        - **Operational State Tracking:** Provides detailed information on the SPAN session's operational state and reasons, helping administrators quickly identify and address SPAN session-related issues.

---

# Data Source: intersight_fabric_span_session
The SpanSession object defines SPAN (Switched Port Analyzer) sessions used for traffic monitoring and analysis. It supports network visibility by directing selected traffic to monitoring interfaces. The operational state and reason fields reflect the status of the SPAN session itself, not underlying network issues.
#### Purpose
SpanSession serves to establish and manage SPAN sessions, allowing administrators to capture and analyze network traffic for troubleshooting and performance monitoring. It plays a crucial role in network diagnostics and operational efficiency.
#### Key Concepts
- **Traffic Monitoring:** Allows the configuration of sources and destinations for traffic capture, providing insights into network performance and issues.
- **Dynamic Configuration:** Supports the dynamic setup and adjustment of SPAN sessions, enabling flexible and responsive network management.
- **Integrated Network Management:** Works in conjunction with network elements to ensure comprehensive monitoring capabilities, enhancing overall network oversight.
- **Operational State Tracking:** Provides detailed information on the SPAN session's operational state and reasons, helping administrators quickly identify and address SPAN session-related issues.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_fabric_span_session.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `admin_state`:(string) Admin state to enable or disable the SPAN session.* `Disabled` - Admin configured Disabled State.* `Enabled` - Admin configured Enabled State. 
* `config_state`:(string) The configured state of SPAN configuration. If the configuration fails to deploy to the Fabric Interconnect, it can be redeployed by toggling the admin state.* `None` - Nil value when no action has been triggered by the user.* `Applied` - User configured settings are in applied state.* `Applying` - User settings are being applied on the target server.* `Failed` - User configured settings could not be applied.* `Scheduled` - User configured settings are scheduled to be applied. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the SPAN session. The name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.) or an underscore (_). 
* `oper_state`:(string) Operational state of the SPAN session.* `Unknown` - SPAN session operational state is Unknown.* `Up` - SPAN session operational state is Up.* `Down` - SPAN session operational state is Down.* `Error` - SPAN session operational state is Error. 
* `oper_state_reason`:(string) Operational state reason of the SPAN session.* `Unknown` - Unknown operational state.* `Active` - Active and operational SPAN session.* `NoHardwareResource` - No hardware resources available.* `NoOperationalSrcDst` - No operational SPAN source or destination.* `GenericError` - Generic operational error.* `NoSourcesConfigured` - No source interfaces configured.* `NoDestinationConfigured` - No destination port configured.* `NoSourceDestinationConfigured` - No source or destination interface configured.* `SessionAdminShut` - Session is administratively disabled.* `WrongDestinationMode` - Wrong Destination mode configured.* `WrongSourceMode` - Wrong Source mode configured.* `TunnelMisconfDown` - Tunnel Misconfigured or Down.* `NoFlowIdSpecified` - No Flow ID specified for ERSPAN. 
* `session_id`:(int) Session ID identifies the SPAN session on the Fabric Interconnect. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `source_count`:(int) Count of total number of sources in the SPAN session. 
* `span_control_packets`:(string) Admin state to enable or disable spanning control packets.* `Disabled` - Admin configured Disabled State.* `Enabled` - Admin configured Enabled State. 
* `switch_id`:(string) Switch ID (A or B) corresponding to network element.* `A` - Switch Identifier of Fabric Interconnect A.* `B` - Switch Identifier of Fabric Interconnect B. 
 
