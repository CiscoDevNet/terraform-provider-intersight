---
subcategory: "pci"
layout: "intersight"
page_title: "Intersight: intersight_pci_coprocessor_card"
description: |-
  PCIe Compression and Cryptographic CPU Offload Card.
---

# Data Source: intersight_pci_coprocessor_card
PCIe Compression and Cryptographic CPU Offload Card.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_pci_coprocessor_card.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `card_id`:(int) The id of the coprocessor card. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `model`:(string) This field identifies the model of the given component. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `pci_slot`:(string) The PCI slot name for the coprocessor card. 
* `revision`:(string) This field identifies the revision of the given component. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field identifies the serial of the given component. 
* `vendor`:(string) This field identifies the vendor of the given component. 
 