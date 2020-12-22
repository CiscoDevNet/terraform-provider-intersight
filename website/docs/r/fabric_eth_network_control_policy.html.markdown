---
subcategory: "fabric"
layout: "intersight"
page_title: "Intersight: intersight_fabric_eth_network_control_policy"
description: |-
  The features that are applied on a vethernet that is connected to the vNIC.
---

# Resource: intersight_fabric_eth_network_control_policy
The features that are applied on a vethernet that is connected to the vNIC.
## Argument Reference
The following arguments are supported:
* `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
* `cdp_enabled`:(bool) Enables the CDP on an interface. 
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `description`:(string) Description of the policy. 
* `forge_mac`:(string) Determines if the MAC forging is allowed or denied on an interface.* `allow` - Allows mac forging on an interface.* `deny` - Denies mac forging on an interface. 
* `lldp_settings`:(Array with Maximum of one item) - Determines the LLDP setting on an interface on the switch. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
  + `receive_enabled`:(bool) Determines if the LLDP frames can be received by an interface on the switch. 
  + `transmit_enabled`:(bool) Determines if the LLDP frames can be transmitted by an interface on the switch. 
* `mac_registration_mode`:(string) Determines the MAC addresses that have to be registered with the switch.* `nativeVlanOnly` - Register only the MAC addresses learnt on the native VLAN.* `allVlans` - Register all the MAC addresses learnt on all the allowed VLANs. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `network_policy`:(Array) An array of relationships to vnicEthNetworkPolicy resources. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `organization`:(Array with Maximum of one item) - A reference to a organizationOrganization resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `uplink_fail_action`:(string) Determines the state of the virtual interface (vethernet / vfc) on the switch when a suitable uplink is not pinned.* `linkDown` - The vethernet will go down in case a suitable uplink is not pinned.* `warning` - The vethernet will remain up even if a suitable uplink is not pinned. 


## Import
`intersight_fabric_eth_network_control_policy` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_fabric_eth_network_control_policy.example 1234567890987654321abcde
```