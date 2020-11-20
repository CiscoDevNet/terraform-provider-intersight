
---
layout: "intersight"
page_title: "Intersight: intersight_macpool_universe"
sidebar_current: "docs-intersight-data-source-macpool-universe"
description: |-
Universe represents a book keeping container to keep track of all IDs for a given account and pool type.
---

# Data Source: intersight_macpool._universe
Universe represents a book keeping container to keep track of all IDs for a given account and pool type.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
