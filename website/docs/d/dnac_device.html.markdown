---
subcategory: "dnac"
layout: "intersight"
page_title: "Intersight: intersight_dnac_device"
description: |-
        The Device object is a comprehensive representation of network devices, facilitating the management and operational integration of network hardware within the system.
        #### Purpose
        A Device object provides detailed insights into the configuration, status, and management of network devices, supporting efficient network operations and maintenance.
        #### Key Concepts
        - **Device Management:** Supports detailed device configuration and status tracking, enhancing network reliability and efficiency.
        - **Identity and Access Control:** Utilizes unique identifiers and privilege sets to ensure secure and controlled interaction with device data.
        - **Operational Integration:** Seamlessly interfaces with other network components, promoting cohesive and efficient network operations.

---

# Data Source: intersight_dnac_device
The Device object is a comprehensive representation of network devices, facilitating the management and operational integration of network hardware within the system.
#### Purpose
A Device object provides detailed insights into the configuration, status, and management of network devices, supporting efficient network operations and maintenance.
#### Key Concepts
- **Device Management:** Supports detailed device configuration and status tracking, enhancing network reliability and efficiency.
- **Identity and Access Control:** Utilizes unique identifiers and privilege sets to ensure secure and controlled interaction with device data.
- **Operational Integration:** Seamlessly interfaces with other network components, promoting cohesive and efficient network operations.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_dnac_device.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `device_id`:(string) Unique identity of the network device. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `host_name`:(string) Host name of the network device. 
* `management_ip_address`:(string) IP address of the network device. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
