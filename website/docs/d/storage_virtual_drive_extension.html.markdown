
---
layout: "intersight"
page_title: "Intersight: intersight_storage_virtual_drive_extension"
sidebar_current: "docs-intersight-data-source-storage-virtual-drive-extension"
description: |-
Information of virtual drives as reported by a storage controller. In certain cases like S-series servers, virtual drive information will be reported by the controller separately and this represents such information.
---

# Data Source: intersight_storage._virtual_drive_extension
Information of virtual drives as reported by a storage controller. In certain cases like S-series servers, virtual drive information will be reported by the controller separately and this represents such information.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `bootable`:(string) The ability to boot from the virtual drive. 
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `container_id`:(int) The container id of the virtual drive. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `drive_state`:(string) The state of the virtual drive. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the Virtual drive. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `oper_device_id`:(string) The operational device id of the virtual drive. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `uuid`:(string) The UUID assigned to the virtual drive. 
* `vendor_uuid`:(string) The UUID value of the vendor assigned to the virtual drive. 
* `virtual_drive_dn`:(string) The distinguished name of the virtual drive for which the extended data is provided. 
* `virtual_drive_id`:(string) The Id of the virtual drive. 
