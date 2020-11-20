
---
layout: "intersight"
page_title: "Intersight: intersight_storage_controller"
sidebar_current: "docs-intersight-data-source-storage-controller"
description: |-
Storage Controller present in a server.
---

# Data Source: intersight_storage._controller
Storage Controller present in a server.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `controller_flags`:(string) The flags for the storage controller. 
* `controller_id`:(string) The Id of the storage controller. 
* `controller_status`:(string) The current status of controller. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `foreign_config_present`:(bool) Storage controller has detected disks in foreign config. 
* `hw_revision`:(string) The hardware revision of controller. 
* `interface_type`:(string) Interface types are Sas, Sata, PCH. 
* `max_volumes_supported`:(int) Maximum virtual drives that can be created on this Storage Controller. 
* `model`:(string) This field identifies the model of the given component. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `oob_interface_supported`:(string) The CIMC support for out-of-band configuration of controller. 
* `oper_state`:(string) The current operational state of controller. 
* `operability`:(string) Operability state of the storage controller. 
* `pci_addr`:(string) The current pci address of controller. 
* `pci_slot`:(string) The pci slot name for the controller. 
* `presence`:(string) Physical Presence State for the Storage Controller. 
* `raid_support`:(string) The RAID levels supported by controller. 
* `rebuild_rate`:(string) Logical volume or RAID rebuild rate of Storage Controller. 
* `revision`:(string) This field identifies the revision of the given component. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `self_encrypt_enabled`:(string) Storage controller disk self encryption state. 
* `serial`:(string) This field identifies the serial of the given component. 
* `type`:(string) Controller types are Raid, FlexFlash. 
* `vendor`:(string) This field identifies the vendor of the given component. 
