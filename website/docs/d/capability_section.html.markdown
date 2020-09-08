
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
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) A unique name for the section inside a catalog. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `source`:(string) The configured source for this section of the catalog.* `Local` - The catalog file is packaged with the service.* `IntersightRepository` - The catalog file is hosted in the Intersight Repository. 
* `version`:(string) Version of the section inside a catalog. 
