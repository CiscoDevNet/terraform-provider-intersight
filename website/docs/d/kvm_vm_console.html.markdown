---
subcategory: "kvm"
layout: "intersight"
page_title: "Intersight: intersight_kvm_vm_console"
description: |-
  API to launch the virtual machine console.
---

# Data Source: intersight_kvm_vm_console
API to launch the virtual machine console.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_kvm_vm_console.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `client_url`:(string) The multiplexer URL for the client to connect on. 
* `kvm_mux_url`:(string) The URL of the KVM MUX to connect to. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
 