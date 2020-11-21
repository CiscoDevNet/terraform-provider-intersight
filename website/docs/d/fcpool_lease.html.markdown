
---
layout: "intersight"
page_title: "Intersight: intersight_fcpool_lease"
sidebar_current: "docs-intersight-data-source-fcpool-lease"
description: |-
Lease represents a single WWN ID that is part of the universe, allocated either from a pool or through static assignment.
---

# Data Source: intersight_fcpool._lease
Lease represents a single WWN ID that is part of the universe, allocated either from a pool or through static assignment.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `pool_purpose`:(string) Purpose of this WWN pool. 
* `wwn_id`:(string) WWN ID allocated for pool based allocation. 
