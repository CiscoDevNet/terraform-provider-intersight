
---
layout: "intersight"
page_title: "Intersight: intersight_pci_link"
sidebar_current: "docs-intersight-data-source-pci-link"
description: |-
The PCI Switch Link connected to PCIe Switch.
---

# Data Source: intersight_pci._link
The PCI Switch Link connected to PCIe Switch.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `adapter`:(string) The name of the PCI device. 
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `link_speed`:(string) The upstream link speed of the PCI device. 
* `link_status`:(string) The upstream link status of the PCI device. 
* `link_width`:(string) The upstream link width of the PCI device. 
* `model`:(string) This field identifies the model of the given component. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `pci_slot`:(string) The slot name of the PCI device. 
* `revision`:(string) This field identifies the revision of the given component. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field identifies the serial of the given component. 
* `slot_status`:(string) The health information of the PCI device. 
* `vendor`:(string) This field identifies the vendor of the given component. 
