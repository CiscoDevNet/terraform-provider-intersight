
---
layout: "intersight"
page_title: "Intersight: intersight_capability_sioc_module_descriptor"
sidebar_current: "docs-intersight-data-source-capability-sioc-module-descriptor"
description: |-
Descriptor that uniquely identifies an SIOC module.
---

# Data Source: intersight_capability._sioc_module_descriptor
Descriptor that uniquely identifies an SIOC module.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `description`:(string) Detailed information about the endpoint. 
* `model`:(string) The model of the endpoint, for which this capability information is applicable. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `revision`:(string) Revision for the SIOC module. 
* `vendor`:(string) The vendor of the endpoint, for which this capability information is applicable. 
* `version`:(string) The firmware or software version of the endpoint, for which this capability information is applicable. 
