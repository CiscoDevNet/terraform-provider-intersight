---
subcategory: "vnic"
layout: "intersight"
page_title: "Intersight: intersight_vnic_eth_adapter_policy"
description: |-
  An Ethernet adapter policy governs the host-side behavior of the adapter, including how the adapter handles traffic. For each VIC Virtual Ethernet Interface various features like VXLAN, NVGRE, ARFS, Interrupt settings, and TCP Offload settings can be configured.
---

# Resource: intersight_vnic_eth_adapter_policy
An Ethernet adapter policy governs the host-side behavior of the adapter, including how the adapter handles traffic. For each VIC Virtual Ethernet Interface various features like VXLAN, NVGRE, ARFS, Interrupt settings, and TCP Offload settings can be configured.
## Argument Reference
The following arguments are supported:
* `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
* `advanced_filter`:(bool) Enables advanced filtering on the interface. 
* `arfs_settings`:(Array with Maximum of one item) - Settings for Accelerated Receive Flow Steering to reduce the network latency and increase CPU cache efficiency. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
  + `enabled`:(bool) Status of Accelerated Receive Flow Steering on the virtual ethernet interface. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `completion_queue_settings`:(Array with Maximum of one item) - Completion Queue resource settings. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
  + `nr_count`:(int) The number of completion queue resources to allocate. In general, the number of completion queue resources to allocate is equal to the number of transmit queue resources plus the number of receive queue resources. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
  + `ring_size`:(int)(Computed) The number of descriptors in each completion queue. 
* `description`:(string) Description of the policy. 
* `interrupt_scaling`:(bool) Enables Interrupt Scaling on the interface. 
* `interrupt_settings`:(Array with Maximum of one item) - Interrupt Settings for the virtual ethernet interface. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
  + `coalescing_time`:(int) The time to wait between interrupts or the idle period that must be encountered before an interrupt is sent. To turn off interrupt coalescing, enter 0 (zero) in this field. 
  + `coalescing_type`:(string) Interrupt Coalescing Type. This can be one of the following:- MIN  - The system waits for the time specified in the Coalescing Time field before sending another interrupt event IDLE - The system does not send an interrupt until there is a period of no activity lasting as least as long as the time specified in the Coalescing Time field.* `MIN` - The system waits for the time specified in the Coalescing Time field before sending another interrupt event.* `IDLE` - The system does not send an interrupt until there is a period of no activity lasting as least as long as the time specified in the Coalescing Time field. 
  + `nr_count`:(int) The number of interrupt resources to allocate. Typical value is be equal to the number of completion queue resources. 
  + `mode`:(string) Preferred driver interrupt mode. This can be one of the following:- MSIx - Message Signaled Interrupts (MSI) with the optional extension. MSI  - MSI only. INTx - PCI INTx interrupts. MSIx is the recommended option.* `MSIx` - Message Signaled Interrupt (MSI) mechanism with the optional extension (MSIx). MSIx is the recommended and default option.* `MSI` - Message Signaled Interrupt (MSI) mechanism that treats messages as interrupts.* `INTx` - Line-based interrupt (INTx) mechanism similar to the one used in Legacy systems. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `nvgre_settings`:(Array with Maximum of one item) - Network Virtualization using Generic Routing Encapsulation Settings. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
  + `enabled`:(bool) Status of the Network Virtualization using Generic Routing Encapsulation on the virtual ethernet interface. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `organization`:(Array with Maximum of one item) - A reference to a organizationOrganization resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `roce_settings`:(Array with Maximum of one item) - Settings for RDMA over Converged Ethernet. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
  + `class_of_service`:(int) The Class of Service for RoCE on this virtual interface.* `5` - RDMA CoS Service Level 5.* `1` - RDMA CoS Service Level 1.* `2` - RDMA CoS Service Level 2.* `4` - RDMA CoS Service Level 4.* `6` - RDMA CoS Service Level 6. 
  + `enabled`:(bool) If enabled sets RDMA over Converged Ethernet (RoCE) on this virtual interface. 
  + `memory_regions`:(int) The number of memory regions per adapter. Recommended value = integer power of 2. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
  + `queue_pairs`:(int) The number of queue pairs per adapter. Recommended value = integer power of 2. 
  + `resource_groups`:(int) The number of resource groups per adapter. Recommended value = be an integer power of 2 greater than or equal to the number of CPU cores on the system for optimum performance. 
  + `nr_version`:(int) Configure RDMA over Converged Ethernet (RoCE) version on the virtual interface. Only RoCEv1 is supported on Cisco VIC 13xx series adapters and only RoCEv2 is supported on Cisco VIC 14xx series adapters.* `1` - RDMA over Converged Ethernet Protocol Version 1.* `2` - RDMA over Converged Ethernet Protocol Version 2. 
* `rss_hash_settings`:(Array with Maximum of one item) - Receive Side Scaling allows the incoming traffic to be spread across multiple CPU cores. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
  + `ipv4_hash`:(bool) When enabled, the IPv4 address is used for traffic distribution. 
  + `ipv6_ext_hash`:(bool) When enabled, the IPv6 extensions are used for traffic distribution. 
  + `ipv6_hash`:(bool) When enabled, the IPv6 address is used for traffic distribution. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
  + `tcp_ipv4_hash`:(bool) When enabled, both the IPv4 address and TCP port number are used for traffic distribution. 
  + `tcp_ipv6_ext_hash`:(bool) When enabled, both the IPv6 extensions and TCP port number are used for traffic distribution. 
  + `tcp_ipv6_hash`:(bool) When enabled, both the IPv6 address and TCP port number are used for traffic distribution. 
  + `udp_ipv4_hash`:(bool) When enabled, both the IPv4 address and UDP port number are used for traffic distribution. 
  + `udp_ipv6_hash`:(bool) When enabled, both the IPv6 address and UDP port number are used for traffic distribution. 
* `rss_settings`:(bool) Receive Side Scaling allows the incoming traffic to be spread across multiple CPU cores. 
* `rx_queue_settings`:(Array with Maximum of one item) - Receive Queue resouce settings. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
  + `nr_count`:(int) The number of queue resources to allocate. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
  + `ring_size`:(int) The number of descriptors in each queue. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `tcp_offload_settings`:(Array with Maximum of one item) - The TCP offload settings decide whether to offload the TCP related network functions from the CPU to the network hardware or not. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
  + `large_receive`:(bool) Enables the reassembly of segmented packets in hardware before sending them to the CPU. 
  + `large_send`:(bool) Enables the CPU to send large packets to the hardware for segmentation. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
  + `rx_checksum`:(bool) When enabled, the CPU sends all packet checksums to the hardware for validation. 
  + `tx_checksum`:(bool) When enabled, the CPU sends all packets to the hardware so that the checksum can be calculated. 
* `tx_queue_settings`:(Array with Maximum of one item) - Transmit Queue resource settings. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
  + `nr_count`:(int) The number of queue resources to allocate. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
  + `ring_size`:(int) The number of descriptors in each queue. 
* `uplink_failback_timeout`:(int) Uplink Failback Timeout in seconds when uplink failover is enabled for a vNIC. After a vNIC has started using its secondary interface, this setting controls how long the primary interface must be available before the system resumes using the primary interface for the vNIC. 
* `vxlan_settings`:(Array with Maximum of one item) - Virtual Extensible LAN Protocol Settings. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
  + `enabled`:(bool) Status of the Virtual Extensible LAN Protocol on the virtual ethernet interface. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 


## Import
`intersight_vnic_eth_adapter_policy` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_vnic_eth_adapter_policy.example 1234567890987654321abcde
```