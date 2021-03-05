---
subcategory: "kubernetes"
layout: "intersight"
page_title: "Intersight: intersight_kubernetes_virtual_machine_node_profile"
description: |-
  A profile specifying configuration settings for a Virtual Machine node. It is auto-generated based on the NodeGroupProfile and VirtualMachineNodePolicy configuration. Users can do operations like Drain, Cordon, Rebuild on a node.
---

# Data Source: intersight_kubernetes_virtual_machine_node_profile
A profile specifying configuration settings for a Virtual Machine node. It is auto-generated based on the NodeGroupProfile and VirtualMachineNodePolicy configuration. Users can do operations like Drain, Cordon, Rebuild on a node.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_kubernetes_virtual_machine_node_profile.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `cloud_provider`:(string) Cloud provider for this node profile.* `noProvider` - Enables the use of no cloud provider.* `external` - Out of tree cloud provider, e.g. CPI for vsphere. 
* `description`:(string) Description of the profile. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete profile. 
* `type`:(string) Defines the type of the profile. Accepted value is instance.* `instance` - The profile defines the configuration for a specific instance of a target. 
 