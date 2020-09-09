
---
layout: "intersight"
page_title: "Intersight: intersight_firmware_dimm_descriptor"
sidebar_current: "docs-intersight-data-source-firmware-dimm-descriptor"
description: |-
Descriptor to uniquely identify a DIMM.
---

# Data Source: intersight_firmware._dimm_descriptor
Descriptor to uniquely identify a DIMM.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `brand_string`:(string) The brand string of the endpoint for which this capability information is applicable. 
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `description`:(string) Detailed information about the endpoint. 
* `label`:(string) The label type for the component. 
* `model`:(string) The model of the endpoint, for which this capability information is applicable. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `revision`:(string) The revision for the component. 
* `vendor`:(string) The vendor of the endpoint, for which this capability information is applicable. 
* `version`:(string) The firmware or software version of the endpoint, for which this capability information is applicable. 
