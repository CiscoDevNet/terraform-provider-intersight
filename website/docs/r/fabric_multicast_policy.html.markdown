
---
layout: "intersight"
page_title: "Intersight: intersight_fabric_multicast_policy"
sidebar_current: "docs-intersight-resource-fabric-multicast-policy"
description: |-
  A policy to configure Multicast settings for all the Virtual LAN networks.
---

# Resource: intersight_fabric._multicast_policy
A policy to configure Multicast settings for all the Virtual LAN networks.
## Argument Reference
The following arguments are supported:
* `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `description`:(string) Description of the policy. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `organization`:(Array with Maximum of one item) - A reference to a organizationOrganization resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `querier_ip_address`:(string) Used to define the IGMP Querier IP address. 
* `querier_ip_address_peer`:(string) Used to define the IGMP Querier IP address of the peer switch. 
* `querier_state`:(string) Administrative state of the IGMP Querier for this VLAN.* `Disabled` - IGMP Querier Disabled State.* `Enabled` - IGMP Querier Enabled State. 
* `snooping_state`:(string) Administrative state of the IGMP Snooping for this VLAN.* `Enabled` - IGMP Snooping Enabled State.* `Disabled` - IGMP Snooping Disabled State. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
