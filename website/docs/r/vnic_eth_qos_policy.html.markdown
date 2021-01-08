---
subcategory: "vnic"
layout: "intersight"
page_title: "Intersight: intersight_vnic_eth_qos_policy"
description: |-
  An Ethernet Quality of Service (QoS) policy assigns a system class to the outgoing traffic for a vNIC. This system class determines the quality of service for the outgoing traffic. For certain adapters additional controls can be specified like burst and rate on the outgoing traffic.
---

# Resource: intersight_vnic_eth_qos_policy
An Ethernet Quality of Service (QoS) policy assigns a system class to the outgoing traffic for a vNIC. This system class determines the quality of service for the outgoing traffic. For certain adapters additional controls can be specified like burst and rate on the outgoing traffic.
## Argument Reference
The following arguments are supported:
* `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `cos`:(int) Class of Service to be associated to the traffic on the virtual interface. 
* `description`:(string) Description of the policy. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `mtu`:(int) The Maximum Transmission Unit (MTU) or packet size that the virtual interface accepts. 
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
* `priority`:(string) The priortity matching the System QoS specified in the fabric profile.* `Best Effort` - QoS Priority for Best-effort traffic.* `FC` - QoS Priority for FC traffic.* `Platinum` - QoS Priority for Platinum traffic.* `Gold` - QoS Priority for Gold traffic.* `Silver` - QoS Priority for Silver traffic.* `Bronze` - QoS Priority for Bronze traffic. 
* `rate_limit`:(int) The value in Mbps (0-10G/40G/100G depending on Adapter Model) to use for limiting the data rate on the virtual interface. Setting this to zero will turn rate limiting off. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `trust_host_cos`:(bool) Enables usage of the Class of Service provided by the operating system. 


## Import
`intersight_vnic_eth_qos_policy` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_vnic_eth_qos_policy.example 1234567890987654321abcde
```