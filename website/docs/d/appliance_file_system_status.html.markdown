
---
layout: "intersight"
page_title: "Intersight: intersight_appliance_file_system_status"
sidebar_current: "docs-intersight-data-source-appliance-file-system-status"
description: |-
Status of a file system on an Intersight Appliance node.
---

# Data Source: intersight_appliance._file_system_status
Status of a file system on an Intersight Appliance node.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `capacity`:(int) Capacity of the file system in bytes. 
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `mountpoint`:(string) Mount point of this file system. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `operational_status`:(string) Operational status of the file system.Operational status is based on the result of the statuschecks. If result of any check is Critical, then itsvalue is Impaired. Otherwise, if result of any check isWarning, then its value is AttentionNeeded. If allchecks are OK, then its value is Operational.* `Unknown` - Operational status of the Intersight Appliance entity is Unknown.* `Operational` - Operational status of the Intersight Appliance entity is Operational.* `Impaired` - Operational status of the Intersight Appliance entity is Impaired.* `AttentionNeeded` - Operational status of the Intersight Appliance entity is AttentionNeeded. 
* `usage`:(float) Percentage of the file system capacity currently in use. 
