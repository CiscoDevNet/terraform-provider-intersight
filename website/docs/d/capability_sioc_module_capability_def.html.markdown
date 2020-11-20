
---
layout: "intersight"
page_title: "Intersight: intersight_capability_sioc_module_capability_def"
sidebar_current: "docs-intersight-data-source-capability-sioc-module-capability-def"
description: |-
Chassis SIOC module capabilities.
---

# Data Source: intersight_capability._sioc_module_capability_def
Chassis SIOC module capabilities.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `dc_supported`:(bool) Device connector support on SIOC. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) An unique identifer for a capability descriptor. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
