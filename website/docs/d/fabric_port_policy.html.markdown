---
subcategory: "fabric"
layout: "intersight"
page_title: "Intersight: intersight_fabric_port_policy"
description: |-
  A policy for all the physical ports of the Fabric Interconnect.
---

# Data Source: intersight_fabric_port_policy
A policy for all the physical ports of the Fabric Interconnect.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string) Description of the policy. 
* `device_model`:(string) This field specifies the device model that this Port Policy is being configured for.* `UCS-FI-6454` - The standard 4th generation UCS Fabric Interconnect with 54 ports.* `UCS-FI-64108` - The expanded 4th generation UCS Fabric Interconnect with 108 ports.* `unknown` - Unknown device type, usage is TBD. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
