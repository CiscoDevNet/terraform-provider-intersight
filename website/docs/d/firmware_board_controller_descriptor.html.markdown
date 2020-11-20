
---
layout: "intersight"
page_title: "Intersight: intersight_firmware_board_controller_descriptor"
sidebar_current: "docs-intersight-data-source-firmware-board-controller-descriptor"
description: |-
Descriptor to uniquely identify a board controller.
---

# Data Source: intersight_firmware._board_controller_descriptor
Descriptor to uniquely identify a board controller.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `brand_string`:(string) The brand string of the endpoint for which this capability information is applicable. 
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `description`:(string) Detailed information about the endpoint. 
* `label`:(string) The label type for the component. 
* `model`:(string) The model of the endpoint, for which this capability information is applicable. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `revision`:(string) The revision for the component. 
* `vendor`:(string) The vendor of the endpoint, for which this capability information is applicable. 
* `version`:(string) The firmware or software version of the endpoint, for which this capability information is applicable. 
