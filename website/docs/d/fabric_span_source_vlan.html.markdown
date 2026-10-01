---
subcategory: "fabric"
layout: "intersight"
page_title: "Intersight: intersight_fabric_span_source_vlan"
description: |-
        The SpanSourceVlan object represents a VLAN configured as a SPAN traffic source.
        #### Purpose
        SpanSourceVlan enables mirroring of traffic that matches a VLAN context into a SPAN session, supporting VLAN-scoped troubleshooting and visibility use cases.
        #### Key Concepts
        - **VLAN-scoped mirroring:** Mirrors traffic associated with a VLAN identifier.
        - **Session-scoped source membership:** Exists as a source attached to a particular SPAN session.
        - **Fabric-wide visibility aid:** Useful for diagnosing issues tied to VLAN segmentation.
        - **Directional capture semantics:** Aligns with the SPAN source direction model.

---

# Data Source: intersight_fabric_span_source_vlan
The SpanSourceVlan object represents a VLAN configured as a SPAN traffic source.
#### Purpose
SpanSourceVlan enables mirroring of traffic that matches a VLAN context into a SPAN session, supporting VLAN-scoped troubleshooting and visibility use cases.
#### Key Concepts
- **VLAN-scoped mirroring:** Mirrors traffic associated with a VLAN identifier.
- **Session-scoped source membership:** Exists as a source attached to a particular SPAN session.
- **Fabric-wide visibility aid:** Useful for diagnosing issues tied to VLAN segmentation.
- **Directional capture semantics:** Aligns with the SPAN source direction model.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_fabric_span_source_vlan.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `direction`:(string) Direction of the source SPAN.* `Receive` - SPAN incoming traffic on the SPAN source interface.* `Transmit` - SPAN outgoing traffic on the SPAN source interface.* `Both` - SPAN incoming and outgoing traffic on the SPAN source interface. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `vlan_id`:(int) SPAN source VLAN Identifier. 
 
