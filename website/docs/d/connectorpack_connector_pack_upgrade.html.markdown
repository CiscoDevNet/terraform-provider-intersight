
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
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `connector_pack_op_type`:(string) The type of operation to be performed on UCS Director.* `Install` - Installs the requisite connector packs on UCS Director.* `Push` - Pushes the requisite connector packs to UCS Director. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
