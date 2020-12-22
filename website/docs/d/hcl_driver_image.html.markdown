---
subcategory: "hcl"
layout: "intersight"
page_title: "Intersight: intersight_hcl_driver_image"
description: |-
  Collection used to store the driver ISO urls for each server based on how it is managed.
---

# Data Source: intersight_hcl_driver_image
Collection used to store the driver ISO urls for each server based on how it is managed.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `driver_iso_url`:(string) URL of the driver ISO images. 
* `management_type`:(string) Type of the UCS version indicating whether it is a UCSM release vesion or a IMC release.* `UCSM` - The server is managed by UCS Manager.* `IMC` - The server is standalone managed by CIMC. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `server_pid`:(string) Three part ID representing the server model as returned by UCSM/CIMC XML APIs. 
