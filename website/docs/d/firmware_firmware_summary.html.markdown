
---
layout: "intersight"
page_title: "Intersight: intersight_firmware_firmware_summary"
sidebar_current: "docs-intersight-data-source-firmware-firmware-summary"
description: |-
Update inventory that contains the details for the firmware running on each component in the compute.Physical object.
---

# Data Source: intersight_firmware._firmware_summary
Update inventory that contains the details for the firmware running on each component in the compute.Physical object.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `bundle_version`:(string) Version details at the bundle level for the each of server. 
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
