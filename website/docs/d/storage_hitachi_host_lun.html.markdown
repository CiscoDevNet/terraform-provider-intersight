
---
layout: "intersight"
page_title: "Intersight: intersight_storage_hitachi_host_lun"
sidebar_current: "docs-intersight-data-source-storage-hitachi-host-lun"
description: |-
A host LUN entity in Hitachi storage array. It exists only if the volume has a connection to host group. A host lun provides public connection to all hosts associated within host group. Hitachi assign same HLU for all the host.
---

# Data Source: intersight_storage._hitachi_host_lun
A host LUN entity in Hitachi storage array. It exists only if the volume has a connection to host group. A host lun provides public connection to all hosts associated within host group. Hitachi assign same HLU for all the host.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `hlu`:(int) Logical unit number (LUN) by which hosts address specified volume. Hlu is a decimal representation of the LUN from the endpoint. 
* `host_name`:(string) Name of the host associated with LUN. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `port_id`:(string) Port ID of the lun. 
* `volume_name`:(string) Name of the storage volume associated with LUN. 
