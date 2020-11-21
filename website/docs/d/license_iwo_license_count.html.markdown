
---
layout: "intersight"
page_title: "Intersight: intersight_license_iwo_license_count"
sidebar_current: "docs-intersight-data-source-license-iwo-license-count"
description: |-
Customer operation object to request reservation code.
---

# Data Source: intersight_license._iwo_license_count
Customer operation object to request reservation code.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `vm_license_count`:(int) The total number of devices claimed in the Intersight account. 
