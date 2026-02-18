---
subcategory: "ssh"
layout: "intersight"
page_title: "Intersight: intersight_ssh_policy_inventory"
description: |-
        The SSH Policy is a reusable policy for configuring the Secure Shell (SSH) service on an endpoint's management controller.
        #### Purpose
        The purpose of this policy is to enable and standardize secure command-line access to managed devices. It allows administrators to control the state of the SSH service, the network port it listens on, and the session timeout period, ensuring consistent and secure remote access.
        #### Key Concepts
        - **Service Control:** The policy's primary function is to enable or disable the SSH service on the endpoint.
        - **Port and Timeout Configuration:** It allows administrators to customize the SSH port and the session timeout duration in seconds.
        - **Reusable and Profile-Based:** As a policy object, it can be attached to a Server Profile to apply a consistent SSH configuration across multiple servers.
        - **Fundamental Security:** Configuring SSH access is a fundamental step in securing the management plane of a server.

---

# Data Source: intersight_ssh_policy_inventory
The SSH Policy is a reusable policy for configuring the Secure Shell (SSH) service on an endpoint's management controller.
#### Purpose
The purpose of this policy is to enable and standardize secure command-line access to managed devices. It allows administrators to control the state of the SSH service, the network port it listens on, and the session timeout period, ensuring consistent and secure remote access.
#### Key Concepts
- **Service Control:** The policy's primary function is to enable or disable the SSH service on the endpoint.
- **Port and Timeout Configuration:** It allows administrators to customize the SSH port and the session timeout duration in seconds.
- **Reusable and Profile-Based:** As a policy object, it can be attached to a Server Profile to apply a consistent SSH configuration across multiple servers.
- **Fundamental Security:** Configuring SSH access is a fundamental step in securing the management plane of a server.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_ssh_policy_inventory.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the policy. 
* `device_mo_id`:(string) Device ID of the entity from where inventory is reported. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `enabled`:(bool) State of SSH service on the endpoint. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the inventoried policy object. 
* `port`:(int) Port used for secure shell access. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `timeout`:(int) Number of seconds to wait before the system considers a SSH request to have timed out. 
 
