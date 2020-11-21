
---
layout: "intersight"
page_title: "Intersight: intersight_uuidpool_universe"
sidebar_current: "docs-intersight-data-source-uuidpool-universe"
description: |-
Universe represents a book keeping container to keep track of all UUIDs for a given account.
---

# Data Source: intersight_uuidpool._universe
Universe represents a book keeping container to keep track of all UUIDs for a given account.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
