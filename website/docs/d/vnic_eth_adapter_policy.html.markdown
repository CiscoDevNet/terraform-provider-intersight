
---
layout: "intersight"
page_title: "Intersight: intersight_vnic_eth_adapter_policy"
sidebar_current: "docs-intersight-data-source-vnic-eth-adapter-policy"
description: |-
An Ethernet adapter policy governs the host-side behavior of the adapter, including how the adapter handles traffic. For each VIC Virtual Ethernet Interface various features like VXLAN, NVGRE, ARFS, Interrupt settings, and TCP Offload settings can be configured.
---

# Data Source: intersight_vnic._eth_adapter_policy
An Ethernet adapter policy governs the host-side behavior of the adapter, including how the adapter handles traffic. For each VIC Virtual Ethernet Interface various features like VXLAN, NVGRE, ARFS, Interrupt settings, and TCP Offload settings can be configured.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `advanced_filter`:(bool) Enables advanced filtering on the interface. 
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `description`:(string) Description of the policy. 
* `interrupt_scaling`:(bool) Enables Interrupt Scaling on the interface. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `rss_settings`:(bool) Receive Side Scaling allows the incoming traffic to be spread across multiple CPU cores. 
* `uplink_failback_timeout`:(int) Uplink Failback Timeout in seconds when uplink failover is enabled for a vNIC. After a vNIC has started using its secondary interface, this setting controls how long the primary interface must be available before the system resumes using the primary interface for the vNIC. 
