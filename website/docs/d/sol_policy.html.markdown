
---
layout: "intersight"
page_title: "Intersight: intersight_sol_policy"
sidebar_current: "docs-intersight-data-source-sol-policy"
description: |-
Policy for configuring Serial Over LAN settings on endpoint.
---

# Data Source: intersight_sol._policy
Policy for configuring Serial Over LAN settings on endpoint.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `baud_rate`:(int) Baud Rate used for Serial Over LAN communication.* `9600` - Use baud rate 9600 for communication.* `19200` - Use baud rate 19200 for communication.* `38400` - Use baud rate 38400 for communication.* `57600` - Use baud rate 57600 for communication.* `115200` - Use baud rate 115200 for communication. 
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `com_port`:(string) Serial port through which the system routes Serial Over LAN communication. This field is available only on some Cisco UCS C-Series servers. If it is unavailable, the server uses COM port 0 by default.* `com0` - Use serial port com0 for communication.* `com1` - Use serial port com1 for communication. 
* `description`:(string) Description of the policy. 
* `enabled`:(bool) State of Serial Over LAN service on the endpoint. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `ssh_port`:(int) SSH port used to access Serial Over LAN directly. Enables bypassing Cisco IMC shell to provide direct access to Serial Over LAN. 
