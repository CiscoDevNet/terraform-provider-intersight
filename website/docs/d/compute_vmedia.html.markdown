---
subcategory: "compute"
layout: "intersight"
page_title: "Intersight: intersight_compute_vmedia"
description: |-
  Inventory of Virtual Media configuration and images uploaded.
---

# Data Source: intersight_compute_vmedia
Inventory of Virtual Media configuration and images uploaded.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_compute_vmedia.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `enabled`:(bool) State of the Virtual Media service on the server. 
* `encryption`:(bool) If enabled, allows encryption of all Virtual Media communications. 
* `low_power_usb`:(bool) If enabled, the virtual drives appear on the boot selection menu after mapping the image and rebooting the host. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
 