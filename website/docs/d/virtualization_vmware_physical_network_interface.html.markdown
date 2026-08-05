---
subcategory: "virtualization"
layout: "intersight"
page_title: "Intersight: intersight_virtualization_vmware_physical_network_interface"
description: |-
        The VmwarePhysicalNetworkInterface object delineates the physical network interface entity within VMware environments, detailing attributes for resource allocation and management.
        #### Purpose
        VmwarePhysicalNetworkInterface provides the framework for managing physical network interfaces, supporting detailed configurations and operational settings within VMware environments.
        #### Key Concepts
        - **Resource Allocation:** Defines settings for switch names, driver types, and link speeds, optimizing interface resource management.
        - **Integration:** Interfaces with hosts, facilitating seamless connectivity and resource management for physical network interfaces.
        - **Security and Access:** Utilizes privilege sets for secure read and update operations, maintaining integrity and authorized access.
        - **Operational Settings:** Includes attributes such as MAC addresses and PCI info, supporting tailored interface operations.

---

# Data Source: intersight_virtualization_vmware_physical_network_interface
The VmwarePhysicalNetworkInterface object delineates the physical network interface entity within VMware environments, detailing attributes for resource allocation and management.
#### Purpose
VmwarePhysicalNetworkInterface provides the framework for managing physical network interfaces, supporting detailed configurations and operational settings within VMware environments.
#### Key Concepts
- **Resource Allocation:** Defines settings for switch names, driver types, and link speeds, optimizing interface resource management.
- **Integration:** Interfaces with hosts, facilitating seamless connectivity and resource management for physical network interfaces.
- **Security and Access:** Utilizes privilege sets for secure read and update operations, maintaining integrity and authorized access.
- **Operational Settings:** Includes attributes such as MAC addresses and PCI info, supporting tailored interface operations.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_virtualization_vmware_physical_network_interface.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `driver`:(string) Driver type associated with physical network interface. 
* `identity`:(string) The internally generated identity of physical network interface. This entity cannot manipulated by users. It aids in uniquely identifying the physical network interface object. For VMware, this is MOR (managed object reference). 
* `link_speed`:(int) Link speed of the physical network interface. 
* `mac_address`:(string) Standard MAC address assigned to physical network interface. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) User-provided name to identify the physical network interface. 
* `pci`:(string) PCI info for physical network interface. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `switch_name`:(string) Switch associated with the physical network interface. 
 
