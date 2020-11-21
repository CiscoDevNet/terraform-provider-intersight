
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
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `driver_iso_url`:(string) URL of the driver ISO images. 
* `management_type`:(string) Type of the UCS version indicating whether it is a UCSM release vesion or a IMC release.* `UCSM` - The server is managed by UCS Manager.* `IMC` - The server is standalone managed by CIMC. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `server_pid`:(string) Three part ID representing the server model as returned by UCSM/CIMC XML APIs. 
