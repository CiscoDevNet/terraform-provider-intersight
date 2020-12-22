---
subcategory: "vnic"
layout: "intersight"
page_title: "Intersight: intersight_vnic_fc_qos_policy"
description: |-
  A Fibre Channel Quality of Service (QoS) policy assigns a system class to the outgoing traffic for a vHBA. This system class determines the quality of service for the outgoing traffic. For certain adapters additional controls can also be specified like burst and rate on the outgoing traffic.
---

# Data Source: intersight_vnic_fc_qos_policy
A Fibre Channel Quality of Service (QoS) policy assigns a system class to the outgoing traffic for a vHBA. This system class determines the quality of service for the outgoing traffic. For certain adapters additional controls can also be specified like burst and rate on the outgoing traffic.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `cos`:(int) Class of Service to be associated to the traffic on the virtual interface. 
* `description`:(string) Description of the policy. 
* `max_data_field_size`:(int) The maximum size of the Fibre Channel frame payload bytes that the virtual interface supports. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `priority`:(string) The priortity matching the System QoS specified in the fabric profile.* `Best Effort` - QoS Priority for Best-effort traffic.* `FC` - QoS Priority for FC traffic.* `Platinum` - QoS Priority for Platinum traffic.* `Gold` - QoS Priority for Gold traffic.* `Silver` - QoS Priority for Silver traffic.* `Bronze` - QoS Priority for Bronze traffic. 
* `rate_limit`:(int) The value in Mbps to use for limiting the data rate on the virtual interface. Setting this to zero will turn rate limiting off. 
