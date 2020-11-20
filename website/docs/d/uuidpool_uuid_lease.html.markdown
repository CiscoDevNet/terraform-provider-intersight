
---
layout: "intersight"
page_title: "Intersight: intersight_uuidpool_uuid_lease"
sidebar_current: "docs-intersight-data-source-uuidpool-uuid-lease"
description: |-
UuidLease represents a single UUID that is part of the universe, allocated either from a pool or through static assignment.
---

# Data Source: intersight_uuidpool._uuid_lease
UuidLease represents a single UUID that is part of the universe, allocated either from a pool or through static assignment.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `uuid`:(string) UUID Prefix+Suffix numbers. 
