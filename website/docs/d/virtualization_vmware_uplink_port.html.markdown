---
subcategory: "virtualization"
layout: "intersight"
page_title: "Intersight: intersight_virtualization_vmware_uplink_port"
description: |-
        The VmwareUplinkPort object represents the VMware uplink port entity, detailing attributes for network connectivity and management.
        #### Purpose
        VmwareUplinkPort serves as the critical component for managing uplink port configurations, optimizing network resource allocation and control within VMware environments.
        #### Key Concepts
        - **Network Connectivity:** Supports settings for port names, keys, and identities, enhancing network management and control.
        - **Integration:** Interfaces with distributed networks, hosts, and physical interfaces, facilitating seamless network management and resource distribution.
        - **Security and Access** Utilizes privilege sets for secure read and update operations, maintaining integrity and authorized access.
        - **Operational Features:** Includes attributes for network relationships, supporting resilient and adaptive network operations.

---

# Data Source: intersight_virtualization_vmware_uplink_port
The VmwareUplinkPort object represents the VMware uplink port entity, detailing attributes for network connectivity and management.
#### Purpose
VmwareUplinkPort serves as the critical component for managing uplink port configurations, optimizing network resource allocation and control within VMware environments.
#### Key Concepts
- **Network Connectivity:** Supports settings for port names, keys, and identities, enhancing network management and control.
- **Integration:** Interfaces with distributed networks, hosts, and physical interfaces, facilitating seamless network management and resource distribution.   
- **Security and Access** Utilizes privilege sets for secure read and update operations, maintaining integrity and authorized access.
- **Operational Features:** Includes attributes for network relationships, supporting resilient and adaptive network operations.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_virtualization_vmware_uplink_port.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `identity`:(string) The VMware managed object reference as a string. 
* `key`:(string) The internally assigned key of this uplink port object. This entity is not manipulated by users. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) User-provided name to identify the uplink port object. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
