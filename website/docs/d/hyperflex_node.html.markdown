---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_node"
description: |-
  A host participating in the cluster. The host consists of a hypervisor installed on a node that manages virtual machines.
---

# Data Source: intersight_hyperflex_node
A host participating in the cluster. The host consists of a hypervisor installed on a node that manages virtual machines.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `build_number`:(string) The build number of the hypervisor running on the host. 
* `display_version`:(string) The user-friendly string representation of the hypervisor version of the host. 
* `host_name`:(string) The hostname configured for the hypervisor running on the host. 
* `hypervisor`:(string) The type of hypervisor running on the host. 
* `lockdown`:(bool) The admin state of lockdown mode on the host. If 'true', lockdown mode is enabled. 
* `model_number`:(string) The model of the host server. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `role`:(string) The role of the host in the HyperFlex cluster. Specifies whether this host is used for compute or for both compute and storage.* `UNKNOWN` - The role of the HyperFlex cluster node is not known.* `STORAGE` - The HyperFlex cluster node provides both storage and compute resources for the cluster.* `COMPUTE` - The HyperFlex cluster node provides compute resources for the cluster. 
* `serial_number`:(string) The serial of the host server. 
* `status`:(string) The status of the host. Indicates whether the hypervisor is online.* `UNKNOWN` - The host status cannot be determined.* `ONLINE` - The host is online and operational.* `OFFLINE` - The host is offline and is currently not participating in the HyperFlex cluster.* `INMAINTENANCE` - The host is not participating in the HyperFlex cluster because of a maintenance operation, such as firmware or data platform upgrade.* `DEGRADED` - The host is degraded and may not be performing in its full operational capacity. 
* `nr_version`:(string) The version of the hypervisor running on the host. 
