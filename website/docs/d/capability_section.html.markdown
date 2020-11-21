
---
layout: "intersight"
page_title: "Intersight: intersight_capability_section"
sidebar_current: "docs-intersight-data-source-capability-section"
description: |-
Refers to a section in the capability catalog. A capability catalog is divided into sections which are managed independently.
---

# Data Source: intersight_capability._section
Refers to a section in the capability catalog. A capability catalog is divided into sections which are managed independently.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `action`:(string) Administrative action to initialize/load the catalog section from a particular source.* `None` - Nil value to indicate that no action is required.* `LoadLocal` - Load the catalog file packaged with the service.* `LoadIntersightRepository` - Load a catalog file hosted in the Intersight Repository. 
* `catalog_name`:(string) The catalog name reference. 
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) A unique name for the section inside a catalog. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `source`:(string) The configured source for this section of the catalog.* `Local` - The catalog file is packaged with the service.* `IntersightRepository` - The catalog file is hosted in the Intersight Repository. 
* `version`:(string) Version of the section inside a catalog. 
