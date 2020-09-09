
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
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `device_id`:(string) The device id of the switch. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `health`:(string) The composite health of the switch. 
* `model`:(string) This field identifies the model of the given component. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `num_of_adaptors`:(string) The number of GPUs and PCI adapters connected the switch. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
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
