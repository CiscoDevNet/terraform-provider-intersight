
---
layout: "intersight"
page_title: "Intersight: intersight_connectorpack_connector_pack_upgrade"
sidebar_current: "docs-intersight-data-source-connectorpack-connector-pack-upgrade"
description: |-
Used to download or install connector packs on the target device.
---

# Data Source: intersight_connectorpack._connector_pack_upgrade
Used to download or install connector packs on the target device.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `connector_pack_op_type`:(string) The type of operation to be performed on UCS Director.* `Install` - Installs the requisite connector packs on UCS Director.* `Push` - Pushes the requisite connector packs to UCS Director. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
