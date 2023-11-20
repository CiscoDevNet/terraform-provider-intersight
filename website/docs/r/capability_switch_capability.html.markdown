---
subcategory: "capability"
layout: "intersight"
page_title: "Intersight: intersight_capability_switch_capability"
description: |-
        Type to represent additional switch specific capabilities.

---

# Resource: intersight_capability_switch_capability
Type to represent additional switch specific capabilities.
## Usage Example
### Resource Creation

```hcl
resource "intersight_capability_switch_capability" "capability_switch_capability1" {
  name                     = "capability_switch_capability1"
  pid                      = "UCS-FI-6454"
  sku                      = "CON-NCF4P-FI6454CH"
  vid                      = "V00"
  dynamic_vifs_supported   = true
  fan_modules_supported    = true
  locator_beacon_supported = true
}
```
## Argument Reference
The following arguments are supported:
* `account_moid`:(string)(ReadOnly) The Account ID for this managed object. 
* `ancestors`:(Array)(ReadOnly) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `create_time`:(string)(ReadOnly) The time when this managed object was created. 
* `default_fcoe_vlan`:(int)(ReadOnly) Default Fcoe VLAN associated with this switch. 
* `domain_group_moid`:(string)(ReadOnly) The DomainGroup ID for this managed object. 
* `dynamic_vifs_supported`:(bool)(ReadOnly) Dynamic VIFs support on this switch. 
* `fan_modules_supported`:(bool)(ReadOnly) Fan Modules support on this switch. 
* `fc_end_host_mode_reserved_vsans`:(Array)
This complex property has following sub-properties:
  + `end_port_id`:(int) Ending Port ID in this range of ports. 
  + `end_slot_id`:(int) Ending Slot ID in this range of ports. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `start_port_id`:(int) Starting Port ID in this range of ports. 
  + `start_slot_id`:(int) Starting Slot ID in this range of ports. 
* `fc_uplink_ports_auto_negotiation_supported`:(bool)(ReadOnly) Fc Uplink ports auto negotiation speed support on this switch. 
* `locator_beacon_supported`:(bool)(ReadOnly) Locator Beacon LED support on this switch. 
* `max_ports`:(int)(ReadOnly) Maximum allowed physical ports on this switch. 
* `max_slots`:(int)(ReadOnly) Maximum allowed physical slots on this switch. 
* `min_version_with_breakout_support`:(string)(ReadOnly) Minimum firmware version supported for breakout ports on this switch. 
* `min_version_with_locator_led_support`:(string)(ReadOnly) Minimum firmware version supported for locator leds on this switch. 
* `min_version_with_neg_auto25g`:(string)(ReadOnly) Minimum firmware version supported for 'negotiate auto 25000' port admin speed on this switch. 
* `mod_time`:(string)(ReadOnly) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) An unique identifer for a capability descriptor. 
* `network_limits`:(HashMap) -(ReadOnly) List of network limitations for this switch. 
This complex property has following sub-properties:
  + `max_compressed_port_vlan_count`:(int) Maximum Compressed configurable VLANs on Switch/Fabric-Interconnect. 
  + `max_uncompressed_port_vlan_count`:(int) Maximum configurable VLANs on Switch/Fabric-Interconnect. 
  + `maximum_active_traffic_monitoring_sessions`:(int) Maximum configured and enabled Traffic Monitoring sessions on this Switch/Fabric-Interconnect. 
  + `maximum_ethernet_port_channels`:(int) Maximum configurable Ethernet port-channels on Switch/Fabric-Interconnect. 
  + `maximum_ethernet_uplink_ports`:(int) Maximum configurable Ethernet Uplink ports on Switch/Fabric-Interconnect. 
  + `maximum_fc_port_channel_members`:(int) Maximum configurable Fibre Channel port-channel member ports on Switch/Fabric-Interconnect. 
  + `maximum_fc_port_channels`:(int) Maximum configurable Fibre Channel port-channels on Switch/Fabric-Interconnect. 
  + `maximum_igmp_groups`:(int) Maximum configurable IGMP Groups on Switch/Fabric-Interconnect. 
  + `maximum_port_channel_members`:(int) Maximum configurable ports per each port-channel on Switch/Fabric-Interconnect. 
  + `maximum_primary_vlan`:(int) Maximum configurable Primary Private VLANs on Switch/Fabric-Interconnect. 
  + `maximum_secondary_vlan`:(int) Maximum configurable Secondary Private VLANs on Switch/Fabric-Interconnect. 
  + `maximum_secondary_vlan_per_primary`:(int) Maximum configurable Secondary VLANs per each Primary VLAN on Switch/Fabric-Interconnect. 
  + `maximum_vifs`:(int) Maximum allowes VIFs on Switch/Fabric-Interconnect. 
  + `maximum_vlans`:(int) Maximum configurable VLANs on Switch/Fabric-Interconnect. 
  + `minimum_active_fans`:(int) Minimum required fans in 'active' state for this Switch/Fabric-Interconnect. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `owners`:
                (Array of schema.TypeString) -(ReadOnly)
* `parent`:(HashMap) -(ReadOnly) A reference to a moBaseMo resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `permission_resources`:(Array)(ReadOnly) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `pid`:(string) Product Identifier for a Switch/Fabric-Interconnect.* `UCS-FI-6454` - The standard 4th generation UCS Fabric Interconnect with 54 ports.* `UCS-FI-64108` - The expanded 4th generation UCS Fabric Interconnect with 108 ports.* `UCS-FI-6536` - The standard 5th generation UCS Fabric Interconnect with 36 ports.* `unknown` - Unknown device type, usage is TBD. 
* `ports_supporting100g_speed`:(Array)
This complex property has following sub-properties:
  + `end_port_id`:(int) Ending Port ID in this range of ports. 
  + `end_slot_id`:(int) Ending Slot ID in this range of ports. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `start_port_id`:(int) Starting Port ID in this range of ports. 
  + `start_slot_id`:(int) Starting Slot ID in this range of ports. 
* `ports_supporting10g_speed`:(Array)
This complex property has following sub-properties:
  + `end_port_id`:(int) Ending Port ID in this range of ports. 
  + `end_slot_id`:(int) Ending Slot ID in this range of ports. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `start_port_id`:(int) Starting Port ID in this range of ports. 
  + `start_slot_id`:(int) Starting Slot ID in this range of ports. 
* `ports_supporting1g_speed`:(Array)
This complex property has following sub-properties:
  + `end_port_id`:(int) Ending Port ID in this range of ports. 
  + `end_slot_id`:(int) Ending Slot ID in this range of ports. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `start_port_id`:(int) Starting Port ID in this range of ports. 
  + `start_slot_id`:(int) Starting Slot ID in this range of ports. 
* `ports_supporting25g_speed`:(Array)
This complex property has following sub-properties:
  + `end_port_id`:(int) Ending Port ID in this range of ports. 
  + `end_slot_id`:(int) Ending Slot ID in this range of ports. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `start_port_id`:(int) Starting Port ID in this range of ports. 
  + `start_slot_id`:(int) Starting Slot ID in this range of ports. 
* `ports_supporting40g_speed`:(Array)
This complex property has following sub-properties:
  + `end_port_id`:(int) Ending Port ID in this range of ports. 
  + `end_slot_id`:(int) Ending Slot ID in this range of ports. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `start_port_id`:(int) Starting Port ID in this range of ports. 
  + `start_slot_id`:(int) Starting Slot ID in this range of ports. 
* `ports_supporting_breakout`:(Array)
This complex property has following sub-properties:
  + `end_port_id`:(int) Ending Port ID in this range of ports. 
  + `end_slot_id`:(int) Ending Slot ID in this range of ports. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `start_port_id`:(int) Starting Port ID in this range of ports. 
  + `start_slot_id`:(int) Starting Slot ID in this range of ports. 
* `ports_supporting_fcoe`:(Array)
This complex property has following sub-properties:
  + `end_port_id`:(int) Ending Port ID in this range of ports. 
  + `end_slot_id`:(int) Ending Slot ID in this range of ports. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `start_port_id`:(int) Starting Port ID in this range of ports. 
  + `start_slot_id`:(int) Starting Slot ID in this range of ports. 
* `ports_supporting_server_role`:(Array)
This complex property has following sub-properties:
  + `end_port_id`:(int) Ending Port ID in this range of ports. 
  + `end_slot_id`:(int) Ending Slot ID in this range of ports. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `start_port_id`:(int) Starting Port ID in this range of ports. 
  + `start_slot_id`:(int) Starting Slot ID in this range of ports. 
* `reserved_vsans`:(Array)
This complex property has following sub-properties:
  + `end_port_id`:(int) Ending Port ID in this range of ports. 
  + `end_slot_id`:(int) Ending Slot ID in this range of ports. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `start_port_id`:(int) Starting Port ID in this range of ports. 
  + `start_slot_id`:(int) Starting Slot ID in this range of ports. 
* `sereno_netflow_supported`:(bool)(ReadOnly) Sereno Adaptor with Netflow support on this switch. 
* `server_role_supported_on_breakout`:
                (Array of schema.TypeString) -
* `shared_scope`:(string)(ReadOnly) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `sku`:(string) SKU information for Switch/Fabric-Interconnect. 
* `storage_limits`:(HashMap) -(ReadOnly) List of storage limitations for this switch. 
This complex property has following sub-properties:
  + `maximum_user_zone_count`:(int) Maximum user zones per Switch/Fabric-Interconnect. 
  + `maximum_virtual_fc_interfaces`:(int) Maximum configurable Virtual Fibre Channel interfaces on Switch/Fabric-Interconnect. 
  + `maximum_virtual_fc_interfaces_per_blade_server`:(int) Maximum configurable Virtual Fibre Channel interfaces per blade. 
  + `maximum_vsans`:(int) Maximum configurable VSANs on Switch/Fabric-Interconnect. 
  + `maximum_zone_count`:(int) Zone limit per Switch/Fabric-Interconnect. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `switching_mode_capabilities`:(Array)
This complex property has following sub-properties:
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `switching_mode`:(string) Switching mode type (endhost, switch) of the switch.* `end-host` - In end-host mode, the fabric interconnects appear to the upstream devices as end hosts with multiple links.In this mode, the switch does not run Spanning Tree Protocol and avoids loops by following a set of rules for traffic forwarding.In case of ethernet switching mode - Ethernet end-host mode is also known as Ethernet host virtualizer.* `switch` - In switch mode, the switch runs Spanning Tree Protocol to avoid loops, and broadcast and multicast packets are handled in the traditional way.This is the traditional switch mode. 
  + `vp_compression_supported`:(bool) VP Compression support on this switch. 
* `system_limits`:(HashMap) -(ReadOnly) List of system limitations for this switch. 
This complex property has following sub-properties:
  + `maximum_chassis_count`:(int) Maximum UCS chassis that can be connected to this Switch/Fabric-Interconnect. 
  + `maximum_fex_per_domain`:(int) Maximum UCS Fabric-extenders (FEX) per Switch/Fabric-Interconnect. 
  + `maximum_servers_per_domain`:(int) Maximum UCS servers per Switch/Fabric-Interconnect. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `unified_ports`:(Array)
This complex property has following sub-properties:
  + `end_port_id`:(int) Ending Port ID in this range of ports. 
  + `end_slot_id`:(int) Ending Slot ID in this range of ports. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `start_port_id`:(int) Starting Port ID in this range of ports. 
  + `start_slot_id`:(int) Starting Slot ID in this range of ports. 
* `unified_rule`:(string)(ReadOnly) The Slider rule for Unified ports on this switch. 
* `uplink_admin_port_speed_neg_auto25_gbps_supported`:(bool)(ReadOnly) 'Negotiate Auto 25000' admin speed support on this switch for port or port-channelwith Ethernet Uplink/Appliance/FCoE Uplink roles. 
* `version_context`:(HashMap) -(ReadOnly) The versioning info for this managed object. 
This complex property has following sub-properties:
  + `interested_mos`:(Array)
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `marked_for_deletion`:(bool)(ReadOnly) The flag to indicate if snapshot is marked for deletion or not. If flag is set then snapshot will be removed after the successful deployment of the policy. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `ref_mo`:(HashMap) -(ReadOnly) A reference to the original Managed Object. 
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `timestamp`:(string)(ReadOnly) The time this versioned Managed Object was created. 
  + `nr_version`:(string)(ReadOnly) The version of the Managed Object, e.g. an incrementing number or a hash id. 
  + `version_type`:(string)(ReadOnly) Specifies type of version. Currently the only supported value is \ Configured\ that is used to keep track of snapshots of policies and profiles that are intendedto be configured to target endpoints.* `Modified` - Version created every time an object is modified.* `Configured` - Version created every time an object is configured to the service profile.* `Deployed` - Version created for objects related to a service profile when it is deployed. 
* `vid`:(string) VID information for Switch/Fabric-Interconnect. 


## Import
`intersight_capability_switch_capability` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_capability_switch_capability.example 1234567890987654321abcde
``` 
