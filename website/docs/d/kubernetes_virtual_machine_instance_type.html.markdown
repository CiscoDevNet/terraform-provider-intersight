---
subcategory: "kubernetes"
layout: "intersight"
page_title: "Intersight: intersight_kubernetes_virtual_machine_instance_type"
description: |-
  A policy specifying node configuration for a Virtual Machine.
---

# Data Source: intersight_kubernetes_virtual_machine_instance_type
A policy specifying node configuration for a Virtual Machine.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `cpu`:(int) Number of CPUs allocated to virtual machine. 
* `description`:(string) Description of the policy. 
* `disk_size`:(int) Ephemeral disk capacity to be provided with units example - 10Gi. 
* `memory`:(int) Virtual machine memory defined in mebibytes (MiB). 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
