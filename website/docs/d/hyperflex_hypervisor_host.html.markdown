---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_hypervisor_host"
description: |-
        ### Overview
        The HypervisorHost object represents a physical or virtual host appliance within the HyperFlex Cluster. This provides a detailed interface for managing host resources, including status, role, and configuration, supporting efficient cluster operation and management.
        #### Purpose
        A HypervisorHost serves as a critical node within the HyperFlex infrastructure, enabling the deployment and management of virtualized environments through resource allocation and operational control.
        #### Key Concepts
        - **Status Monitoring:** - Tracks host status, including operational state and maintenance mode, to ensure reliable cluster performance.
        - **Role Definition:** - Defines the specific role of the host within the HyperFlex environment, such as storage or compute.
        - **Cluster Association:** - Ensures seamless integration with the HyperFlex Cluster, supporting coordinated resource management and operational efficiency.

---

# Data Source: intersight_hyperflex_hypervisor_host
### Overview
The HypervisorHost object represents a physical or virtual host appliance within the HyperFlex Cluster. This provides a detailed interface for managing host resources, including status, role, and configuration, supporting efficient cluster operation and management.
#### Purpose
A HypervisorHost serves as a critical node within the HyperFlex infrastructure, enabling the deployment and management of virtualized environments through resource allocation and operational control.
#### Key Concepts
- **Status Monitoring:** - Tracks host status, including operational state and maintenance mode, to ensure reliable cluster performance.
- **Role Definition:** - Defines the specific role of the host within the HyperFlex environment, such as storage or compute.
- **Cluster Association:** - Ensures seamless integration with the HyperFlex Cluster, supporting coordinated resource management and operational efficiency.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_hypervisor_host.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `configured_memory`:(int) Memory configured for controller virtual machine. 
* `controller_vm_display_version`:(string) The display version of HyperFlex software running on the controller VM. 
* `controller_vm_hxdp_display_version`:(string) Platform storage software product display version running on controller VM. 
* `controller_vm_hxdp_version`:(string) Platform storage software product version running on controller VM. 
* `controller_vm_uuid`:(string) The UUID of the controller VM which belongs to this host. 
* `controller_vm_version`:(string) The version of HyperFlex software running on the controller VM. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `host_status`:(string) The status of the HyperFlex host.* `UNKNOWN` - Current status of the HyperFlex host is unknown.* `ONLINE` - The HyperFlex host is online.* `OFFLINE` - The HyperFlex host is offline.* `INMAINTENANCE` - The HyperFlex host is in maintenance mode.* `DEGRADED` - Current status of the HyperFlex virtual machine is degraded. 
* `hypervisor`:(string) The hypervisor type of the host. 
* `hypervisor_type`:(string) Identifies the broad type of the underlying hypervisor.* `ESXi` - The hypervisor running on the HyperFlex cluster is a Vmware ESXi hypervisor of any version.* `Hyper-V` - The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V.* `Unknown` - The hypervisor running on the HyperFlex cluster is not known. 
* `identity`:(string) The internally generated identity of this host. This entity is not manipulated by users. It aids in uniquely identifying the datacenter object. For VMware, this is an MOR (managed object reference). 
* `lockdown`:(bool) Flag indicating whether the HyperFlex host is in lockdown mode. 
* `maintenance_mode`:(bool) Is this host in maintenance mode. Set to true or false. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) Commercial model information about this hardware. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of this host supplied by user. It is not the identity of the host. The name is subject to user manipulations. 
* `os_version`:(string) The operation system version of the controller VM. 
* `role`:(string) The role of the HyperFlex host.* `UNKNOWN` - The role of the HyperFlex host is unknown.* `STORAGE` - The HyperFlex host's role is storage.* `COMPUTE` - The HyperFlex host's role is compute. 
* `serial`:(string) Serial number of this host (internally generated). 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `status`:(string) Host health status, as reported by the hypervisor platform.* `Unknown` - Entity status is unknown.* `Degraded` - State is degraded, and might impact normal operation of the entity.* `Critical` - Entity is in a critical state, impacting operations.* `Ok` - Entity status is in a stable state, operating normally. 
* `template_version`:(string) The controller virtual machine template version. 
* `up_time`:(string) The uptime of the host, stored as Duration (from w3c). 
* `uuid`:(string) Universally unique identity of this host (example b3d4483b-5560-9342-8309-b486c9236610). Internally generated. 
* `vendor`:(string) Commercial vendor details of this hardware. 
* `virtual_cpus`:(int) Configured number of virtual CPUs for Controller virtual machine. 
 
