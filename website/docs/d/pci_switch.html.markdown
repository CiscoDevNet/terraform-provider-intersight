
---
layout: "intersight"
page_title: "Intersight: intersight_pci_switch"
sidebar_current: "docs-intersight-data-source-pci-switch"
description: |-
PCI Switch present in a server connected to two GPUs and one PCIe adapter.
---

# Data Source: intersight_pci._switch
PCI Switch present in a server connected to two GPUs and one PCIe adapter.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `device_id`:(string) The device id of the switch. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `health`:(string) The composite health of the switch. 
* `model`:(string) This field identifies the model of the given component. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `num_of_adaptors`:(string) The number of GPUs and PCI adapters connected the switch. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `pci_address`:(string) The PCI address of the switch. 
* `pci_slot`:(string) The PCI slot name of the switch. 
* `product_name`:(string) The model information for the switch. 
* `product_revision`:(string) The product revision of the switch. 
* `revision`:(string) This field identifies the revision of the given component. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field identifies the serial of the given component. 
* `sub_device_id`:(string) The sub device id of the switch. 
* `sub_vendor_id`:(string) The sub vendor id of the switch. 
* `temperature`:(string) The current temperature of the switch. 
* `type`:(string) The type information of the switch. 
* `vendor`:(string) This field identifies the vendor of the given component. 
* `vendor_id`:(string) The vendor id of the switch. 
