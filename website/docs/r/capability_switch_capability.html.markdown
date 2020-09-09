
---
layout: "intersight"
page_title: "Intersight: intersight_capability_switch_capability"
sidebar_current: "docs-intersight-resource-capability-switch-capability"
description: |-
  Type to represent additional switch specific capabilities.
---

# Resource: intersight_capability._switch_capability
Type to represent additional switch specific capabilities.
## Argument Reference
The following arguments are supported:
* `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
* `class_id`:(string)(Computed) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `default_fcoe_vlan`:(int) Default Fcoe VLAN associated with this switch. 
* `dynamic_vifs_supported`:(bool) Dynamic VIFs support on this switch. 
* `fan_modules_supported`:(bool) Fan Modules support on this switch. 
* `fc_end_host_mode_reserved_vsans`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string)(Computed) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
  + `end_port_id`:(int) Ending Port ID in this range of ports. 
  + `end_slot_id`:(int) Ending Slot ID in this range of ports. 
  + `object_type`:(string) The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types. 
  + `start_port_id`:(int) Starting Port ID in this range of ports. 
  + `start_slot_id`:(int) Starting Slot ID in this range of ports. 
* `fc_uplink_ports_auto_negotiation_supported`:(bool) Fc Uplink ports auto negotiation speed support on this switch. 
* `locator_beacon_supported`:(bool) Locator Beacon LED support on this switch. 
* `max_active_span_sessions`:(int) Maximum allowed Traffic Monitoring (SPAN) sessions on this switch. 
* `max_ethernet_port_channel_members`:(int) Maximum allowed Ethernet Uplink Port-channel members for each Uplink Port-channel on this switch. 
* `max_ethernet_port_channels`:(int) Maximum allowed Ethernet Uplink Port-channels on this switch. 
* `max_ethernet_uplink_ports`:(int) Maximum allowed Ethernet Uplink Ports on this switch. 
* `max_fc_fcoe_port_channels`:(int) Total maximum Fc and Fcoe Port-channels allowed on this switch. 
* `max_fc_port_channel_members`:(int) Maximum allowed FC Uplink Port-channel members for each FCoE Port-channel on this switch. 
* `max_fcoe_port_channel_members`:(int) Maximum allowed FCoE Uplink Port-channel members for each FCoE Port-channel on this switch. 
* `max_ports`:(int) Maximum allowed physical ports on this switch. 
* `max_slots`:(int) Maximum allowed physical slots on this switch. 
* `max_vsans_supported`:(int) Maximum number of Vsans supported on this switch. 
* `min_active_fans`:(int) Minimum number of fans needed to be active/running on this switch. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) An unique identifer for a capability descriptor. 
* `object_type`:(string)(Computed) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `pid`:(string) Product Identifier for a Switch/Fabric-Interconnect. 
* `ports_supporting100g_speed`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string)(Computed) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
  + `end_port_id`:(int) Ending Port ID in this range of ports. 
  + `end_slot_id`:(int) Ending Slot ID in this range of ports. 
  + `object_type`:(string) The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types. 
  + `start_port_id`:(int) Starting Port ID in this range of ports. 
  + `start_slot_id`:(int) Starting Slot ID in this range of ports. 
* `ports_supporting10g_speed`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string)(Computed) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
  + `end_port_id`:(int) Ending Port ID in this range of ports. 
  + `end_slot_id`:(int) Ending Slot ID in this range of ports. 
  + `object_type`:(string) The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types. 
  + `start_port_id`:(int) Starting Port ID in this range of ports. 
  + `start_slot_id`:(int) Starting Slot ID in this range of ports. 
* `ports_supporting1g_speed`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string)(Computed) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
  + `end_port_id`:(int) Ending Port ID in this range of ports. 
  + `end_slot_id`:(int) Ending Slot ID in this range of ports. 
  + `object_type`:(string) The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types. 
  + `start_port_id`:(int) Starting Port ID in this range of ports. 
  + `start_slot_id`:(int) Starting Slot ID in this range of ports. 
* `ports_supporting25g_speed`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string)(Computed) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
  + `end_port_id`:(int) Ending Port ID in this range of ports. 
  + `end_slot_id`:(int) Ending Slot ID in this range of ports. 
  + `object_type`:(string) The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types. 
  + `start_port_id`:(int) Starting Port ID in this range of ports. 
  + `start_slot_id`:(int) Starting Slot ID in this range of ports. 
* `ports_supporting40g_speed`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string)(Computed) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
  + `end_port_id`:(int) Ending Port ID in this range of ports. 
  + `end_slot_id`:(int) Ending Slot ID in this range of ports. 
  + `object_type`:(string) The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types. 
  + `start_port_id`:(int) Starting Port ID in this range of ports. 
  + `start_slot_id`:(int) Starting Slot ID in this range of ports. 
* `ports_supporting_breakout`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string)(Computed) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
  + `end_port_id`:(int) Ending Port ID in this range of ports. 
  + `end_slot_id`:(int) Ending Slot ID in this range of ports. 
  + `object_type`:(string) The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types. 
  + `start_port_id`:(int) Starting Port ID in this range of ports. 
  + `start_slot_id`:(int) Starting Slot ID in this range of ports. 
* `ports_supporting_fcoe`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string)(Computed) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
  + `end_port_id`:(int) Ending Port ID in this range of ports. 
  + `end_slot_id`:(int) Ending Slot ID in this range of ports. 
  + `object_type`:(string) The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types. 
  + `start_port_id`:(int) Starting Port ID in this range of ports. 
  + `start_slot_id`:(int) Starting Slot ID in this range of ports. 
* `ports_supporting_server_role`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string)(Computed) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
  + `end_port_id`:(int) Ending Port ID in this range of ports. 
  + `end_slot_id`:(int) Ending Slot ID in this range of ports. 
  + `object_type`:(string) The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types. 
  + `start_port_id`:(int) Starting Port ID in this range of ports. 
  + `start_slot_id`:(int) Starting Slot ID in this range of ports. 
* `reserved_vsans`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string)(Computed) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
  + `end_port_id`:(int) Ending Port ID in this range of ports. 
  + `end_slot_id`:(int) Ending Slot ID in this range of ports. 
  + `object_type`:(string) The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types. 
  + `start_port_id`:(int) Starting Port ID in this range of ports. 
  + `start_slot_id`:(int) Starting Slot ID in this range of ports. 
* `section`:(Array with Maximum of one item) - A reference to a capabilitySection resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string)(Computed) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `sereno_netflow_supported`:(bool) Sereno Adaptor with Netflow support on this switch. 
* `sku`:(string) SKU information for Switch/Fabric-Interconnect. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `unified_ports`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string)(Computed) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
  + `end_port_id`:(int) Ending Port ID in this range of ports. 
  + `end_slot_id`:(int) Ending Slot ID in this range of ports. 
  + `object_type`:(string) The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types. 
  + `start_port_id`:(int) Starting Port ID in this range of ports. 
  + `start_slot_id`:(int) Starting Slot ID in this range of ports. 
* `unified_rule`:(string) The Slider rule for Unified ports on this switch. 
* `vid`:(string) VID information for Switch/Fabric-Interconnect. 
* `vp_compression_supported`:(bool) VP Compression support on this switch. 
