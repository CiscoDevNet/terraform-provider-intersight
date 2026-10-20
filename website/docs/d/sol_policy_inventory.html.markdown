---
subcategory: "sol"
layout: "intersight"
page_title: "Intersight: intersight_sol_policy_inventory"
description: |-
        The SOL Policy is a reusable policy for configuring the Serial Over LAN (SOL) service on a server's management controller.
        #### Purpose
        The purpose of this policy is to enable and standardize out-of-band access to a server's serial console. This provides a low-level, text-based interface for server management and troubleshooting, which is particularly useful when the primary network or graphical console is unavailable.
        #### Key Concepts
        - **Service Control:** The policy allows administrators to enable or disable the SOL service on the endpoint.
        - **Communication Parameters:** This provides options to configure the baudRate and comPort to ensure compatibility with the server's operating system and terminal clients.
        - **Direct SSH Access:** A key feature is the ability to configure a dedicated sshPort for SOL, which allows administrators to bypass the main CIMC shell and connect directly to the serial console via SSH.
        - **Profile-Based Application:** The policy is attached to a Server Profile to apply a consistent SOL configuration across multiple servers.

---

# Data Source: intersight_sol_policy_inventory
The SOL Policy is a reusable policy for configuring the Serial Over LAN (SOL) service on a server's management controller.
#### Purpose
The purpose of this policy is to enable and standardize out-of-band access to a server's serial console. This provides a low-level, text-based interface for server management and troubleshooting, which is particularly useful when the primary network or graphical console is unavailable.
#### Key Concepts
- **Service Control:** The policy allows administrators to enable or disable the SOL service on the endpoint.
- **Communication Parameters:** This provides options to configure the baudRate and comPort to ensure compatibility with the server's operating system and terminal clients.
- **Direct SSH Access:** A key feature is the ability to configure a dedicated sshPort for SOL, which allows administrators to bypass the main CIMC shell and connect directly to the serial console via SSH.
- **Profile-Based Application:** The policy is attached to a Server Profile to apply a consistent SOL configuration across multiple servers.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_sol_policy_inventory.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `baud_rate`:(int) Baud Rate used for Serial Over LAN communication.* `9600` - Use baud rate 9600 for communication.* `19200` - Use baud rate 19200 for communication.* `38400` - Use baud rate 38400 for communication.* `57600` - Use baud rate 57600 for communication.* `115200` - Use baud rate 115200 for communication. 
* `com_port`:(string) Serial port through which the system routes Serial Over LAN communication. This field is available only on some Cisco UCS C-Series servers. If it is unavailable, the server uses COM port 0 by default.* `com0` - Use serial port com0 for communication.* `com1` - Use serial port com1 for communication. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the policy. 
* `device_mo_id`:(string) Device ID of the entity from where inventory is reported. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `enabled`:(bool) State of Serial Over LAN service on the endpoint. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the inventoried policy object. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `ssh_port`:(int) SSH port used to access Serial Over LAN directly. Enables bypassing Cisco IMC shell to provide direct access to Serial Over LAN. 
 
