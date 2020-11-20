
---
layout: "intersight"
page_title: "Intersight: intersight_capability_catalog"
sidebar_current: "docs-intersight-data-source-capability-catalog"
description: |-
Container for capability information of managed systems.
This catalog will be managed by devops using a specific role in the Catalog Admin account.
---

# Data Source: intersight_capability._catalog
Container for capability information of managed systems.
This catalog will be managed by devops using a specific role in the Catalog Admin account.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) A unique name for the catalog. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
