---
subcategory: "vmrc"
layout: "intersight"
page_title: "Intersight: intersight_vmrc_console"
description: |-
  API to launch VMRC console to a VMware virtual machine.
---

# Data Source: intersight_vmrc_console
API to launch VMRC console to a VMware virtual machine.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_vmrc_console.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `client_url`:(string) The multiplexer URL for the client to connect on. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
 