---
subcategory: "vrf"
layout: "intersight"
page_title: "Intersight: intersight_vrf_vrf"
description: |-
  Virtual Routing and Forwarding (VRF) is a networking technology that implements an isolated Layer 3 domain.
---

# Data Source: intersight_vrf_vrf
Virtual Routing and Forwarding (VRF) is a networking technology that implements an isolated Layer 3 domain.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_vrf_vrf.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string) Description to help identify or describe this VRF. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the Virtual Routing and Forwarding Instance. 
 