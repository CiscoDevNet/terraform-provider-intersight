
---
layout: "intersight"
page_title: "Intersight: intersight_management_interface"
sidebar_current: "docs-intersight-data-source-management-interface"
description: |-
Interface that provides access to the management controller.
---

# Data Source: intersight_management._interface
Interface that provides access to the management controller.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `gateway`:(string) Default gateway for the interface. 
* `host_name`:(string) Hostname configured for the interface. 
* `ip_address`:(string) IP address of the interface. 
* `ipv4_address`:(string) IPv4 address of the interface. 
* `ipv4_gateway`:(string) IPv4 default gateway for the interface. 
* `ipv4_mask`:(string) IPv4 Netmask for the interface. 
* `ipv6_address`:(string) IPv6 address of the interface. 
* `ipv6_gateway`:(string) IPv6 default gateway for the interface. 
* `ipv6_prefix`:(int) IPv6 prefix for the interface. 
* `mac_address`:(string) MAC address configured for the interface. 
* `mask`:(string) Netmask for the interface. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `switch_id`:(string) Switch Id connected to the interface. 
* `uem_conn_status`:(string) The event channel connection status for the interface. 
* `virtual_host_name`:(string) Virtual hostname configured for the interface in case of clustered environment. 
