---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_hyper_flex_volume"
description: |-
  A HyperFlex Volume entity.
---

# Data Source: intersight_storage_hyper_flex_volume
A HyperFlex Volume entity.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `capacity`:(int) Provisioned Capacity of the Storage container in bytes. 
* `description`:(string) Short description about the volume. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `naa_id`:(string) NAA id of volume. It is a significant number to identify corresponding lun path in hypervisor. 
* `name`:(string) Named entity of the volume. 
* `serial_number`:(string) Serial number of the volume. 
* `size`:(int) User provisioned volume size. It is the size exposed to host. 
* `uuid`:(string) Uuid of the Datastore/Storage Container. 
