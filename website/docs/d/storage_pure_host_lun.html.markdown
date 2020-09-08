
---
layout: "intersight"
page_title: "Intersight: intersight_storage_pure_host_lun"
sidebar_current: "docs-intersight-data-source-storage-pure-host-lun"
description: |-
A host LUN entity in Pure storage array. It exists only if the volume has a connection to host or host group. For host group mapping, it provides public connection to all hosts associated within host group. A volume can have private connection to host as well. It cannot have public and private connection. Pure assign same HLU for all the host in case if it is connected through host group.
---

# Data Source: intersight_storage._pure_host_lun
A host LUN entity in Pure storage array. It exists only if the volume has a connection to host or host group. For host group mapping, it provides public connection to all hosts associated within host group. A volume can have private connection to host as well. It cannot have public and private connection. Pure assign same HLU for all the host in case if it is connected through host group.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `hlu`:(int) Logical unit number (LUN) by which hosts address specified volume. Hlu is a decimal representation of the LUN from the endpoint. 
* `host_group_name`:(string) Name of the host group associated with LUN. 
* `host_name`:(string) Name of the host associated with LUN. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `shared`:(bool) Kind of volume connection to host. True if it is connected through host group. False in case of direct host connection. 
* `volume_name`:(string) Name of the storage volume associated with LUN. 
