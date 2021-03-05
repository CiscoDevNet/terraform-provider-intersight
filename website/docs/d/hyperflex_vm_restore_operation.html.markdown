---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_vm_restore_operation"
description: |-
  Invoke Virtual Machine restore operation.
---

# Data Source: intersight_hyperflex_vm_restore_operation
Invoke Virtual Machine restore operation.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_vm_restore_operation.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `new_name`:(string) New name for the Virtual Machine after recovery. 
* `power_on`:(bool) Power on the Virtual Machine after recovery. 
 