
---
layout: "intersight"
page_title: "Intersight: intersight_hcl_driver_image"
sidebar_current: "docs-intersight-data-source-hcl-driver-image"
description: |-
Collection used to store the driver ISO urls for each server based on how it is managed.
---

# Data Source: intersight_hcl._driver_image
Collection used to store the driver ISO urls for each server based on how it is managed.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `driver_iso_url`:(string) URL of the driver ISO images. 
* `management_type`:(string) Type of the UCS version indicating whether it is a UCSM release vesion or a IMC release.* `UCSM` - The server is managed by UCS Manager.* `IMC` - The server is standalone managed by CIMC. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `server_pid`:(string) Three part ID representing the server model as returned by UCSM/CIMC XML APIs. 
