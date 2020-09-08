
---
layout: "intersight"
page_title: "Intersight: intersight_capability_switch_capability"
sidebar_current: "docs-intersight-data-source-capability-switch-capability"
description: |-
Type to represent additional switch specific capabilities.
---

# Data Source: intersight_capability._switch_capability
Type to represent additional switch specific capabilities.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `default_fcoe_vlan`:(int) Default Fcoe VLAN associated with this switch. 
* `dynamic_vifs_supported`:(bool) Dynamic VIFs support on this switch. 
* `fan_modules_supported`:(bool) Fan Modules support on this switch. 
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
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `pid`:(string) Product Identifier for a Switch/Fabric-Interconnect. 
* `sereno_netflow_supported`:(bool) Sereno Adaptor with Netflow support on this switch. 
* `sku`:(string) SKU information for Switch/Fabric-Interconnect. 
* `unified_rule`:(string) The Slider rule for Unified ports on this switch. 
* `vid`:(string) VID information for Switch/Fabric-Interconnect. 
* `vp_compression_supported`:(bool) VP Compression support on this switch. 
