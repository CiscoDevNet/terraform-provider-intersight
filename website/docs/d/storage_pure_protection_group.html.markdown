
---
layout: "intersight"
page_title: "Intersight: intersight_storage_pure_protection_group"
sidebar_current: "docs-intersight-data-source-storage-pure-protection-group"
description: |-
Protection group entity in Pure storage array. A volume can be protected by associating with protection group either directly or indirectly (either host or host group). Snapshots are created on protected volume in local array or target array or both as per scheduler configuration.
---

# Data Source: intersight_storage._pure_protection_group
Protection group entity in Pure storage array. A volume can be protected by associating with protection group either directly or indirectly (either host or host group). Snapshots are created on protected volume in local array or target array or both as per scheduler configuration.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the protection Group. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `prefix`:(string) Prefix used for all generated snapshots from the protection group. 
* `replication_enabled`:(bool) Flag to determine if replication is enabled. Snapshots are replicated to the target array if this flag is set. 
* `size`:(int) Overall size of all snapshots in the protection group, represented in bytes. 
* `snapshot_enabled`:(bool) Flag to determine if snapshot creation is enabled. Snapshots are created on local array if this flag is set. 
* `source`:(string) Name of PureStorage array name on which the protection group is created. 
