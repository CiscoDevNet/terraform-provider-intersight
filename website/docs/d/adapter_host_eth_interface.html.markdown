
---
layout: "intersight"
page_title: "Intersight: intersight_adapter_host_eth_interface"
sidebar_current: "docs-intersight-data-source-adapter-host-eth-interface"
description: |-
Physical / Virtual port of an adapter as seen by the host.
---

# Data Source: intersight_adapter._host_eth_interface
Physical / Virtual port of an adapter as seen by the host.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `admin_state`:(string) Admin state of the Host Ethernet Interface. 
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `ep_dn`:(string) The Endpoint Config Dn of the Host Ethernet Interface. 
* `host_eth_interface_id`:(int) Unique Identifier for an Host Ethernet Interface within the adapter object. 
* `interface_type`:(string) Type of External Ethernet Interface. 
* `mac_address`:(string) Mac address of the Host Ethernet Interface. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of Host Ethernet Interface. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `oper_state`:(string) Operational state of an Interface. 
* `operability`:(string) Operability status of Host Ethernet Channel Interface. 
* `original_mac_address`:(string) The factory default Mac address of the Host Ethernet Interface. 
* `pci_addr`:(string) The PCI address of the Host Ethernet Interface. 
* `peer_dn`:(string) The distinguished name of the peer endpoint connected to the Host Ethernet interface. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `virtualization_preference`:(string) Virtualization Preference of the Host Ethernet Interface indicating if virtualization is enabled or not. 
* `vnic_dn`:(string) The Virtual Ethernet Interface DN connected to the Host Ethernet Interface. 
