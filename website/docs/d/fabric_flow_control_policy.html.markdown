---
subcategory: "fabric"
layout: "intersight"
page_title: "Intersight: intersight_fabric_flow_control_policy"
description: |-
        The FlowControlPolicy object defines link-level flow control behavior, including priority-based flow control (PFC) intent for ports.
        #### Purpose
        FlowControlPolicy provides a policy-managed mechanism for setting flow control behavior on interfaces. It supports consistent adoption of PFC and flow control direction behaviors where needed for lossless traffic classes and predictable performance.
        #### Key Concepts
        - **Policy-based flow control:** Encapsulates flow-control behavior as reusable intent.
        - **PFC enablement semantics:** Supports consistent lossless behavior expectations where required.
        - **Interface-level applicability:** Typically attached to ports/roles that require specific flow control behavior.
        - **Alignment with QoS:** Works with system QoS intent so lossless priorities can be enforced coherently.
        Priority Flow Control setting for each port.

---

# Data Source: intersight_fabric_flow_control_policy
The FlowControlPolicy object defines link-level flow control behavior, including priority-based flow control (PFC) intent for ports.
#### Purpose
FlowControlPolicy provides a policy-managed mechanism for setting flow control behavior on interfaces. It supports consistent adoption of PFC and flow control direction behaviors where needed for lossless traffic classes and predictable performance.
#### Key Concepts
- **Policy-based flow control:** Encapsulates flow-control behavior as reusable intent.
- **PFC enablement semantics:** Supports consistent lossless behavior expectations where required.
- **Interface-level applicability:** Typically attached to ports/roles that require specific flow control behavior.
- **Alignment with QoS:** Works with system QoS intent so lossless priorities can be enforced coherently.
Priority Flow Control setting for each port.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_fabric_flow_control_policy.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the policy. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `priority_flow_control_mode`:(string) Configure the Priority Flow Control (PFC) for each port to enable the no-drop behavior for the CoS defined by the System QoS Policy and an Ethernet QoS policy. If Auto and On is selected for PFC, the Receive and Send link level flow control will be Off.* `auto` - Enables the no-drop CoS values to be advertised by the DCBXP and negotiated with the peer.A successful negotiation enables PFC on the no-drop CoS.Any failures because of a mismatch in the capability of peers causes the PFC not to be enabled.* `on` - Enables PFC on the local port regardless of the capability of the peers.* `off` - Disable PFC on the local port regardless of the capability of the peers. 
* `receive_direction`:(string) Link-level Flow Control configured in the receive direction.* `Disabled` - Admin configured Disabled State.* `Enabled` - Admin configured Enabled State. 
* `send_direction`:(string) Link-level Flow Control configured in the send direction.* `Disabled` - Admin configured Disabled State.* `Enabled` - Admin configured Enabled State. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
