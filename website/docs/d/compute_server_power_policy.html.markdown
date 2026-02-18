---
subcategory: "compute"
layout: "intersight"
page_title: "Intersight: intersight_compute_server_power_policy"
description: |-
        The ServerPowerPolicy object dictates the power tasks required during server profile deployments and undeployments. It helps manage server power states according to the defined policy.
        #### Purpose
        ServerPowerPolicy serves as a guideline for determining server power actions, such as power on/off, during server profile operations. It ensures that power state changes are managed consistently and according to policy.
        #### Key Concepts
        - **Policy Definition:** Establishes rules for server power management during profile operations, providing clarity and consistency.
        - **Integration with Profiles:** Connects with server profiles to ensure power policies are applied as part of broader management tasks.
        - **Access Control:** Uses privilege sets to manage policy updates, supporting secure and controlled access.
        - **Relationship Management:** Links with registered devices and servers, ensuring power policies are enforced across the network.

---

# Data Source: intersight_compute_server_power_policy
The ServerPowerPolicy object dictates the power tasks required during server profile deployments and undeployments. It helps manage server power states according to the defined policy.
#### Purpose
ServerPowerPolicy serves as a guideline for determining server power actions, such as power on/off, during server profile operations. It ensures that power state changes are managed consistently and according to policy.
#### Key Concepts
- **Policy Definition:** Establishes rules for server power management during profile operations, providing clarity and consistency.
- **Integration with Profiles:** Connects with server profiles to ensure power policies are applied as part of broader management tasks.
- **Access Control:** Uses privilege sets to manage policy updates, supporting secure and controlled access.
- **Relationship Management:** Links with registered devices and servers, ensuring power policies are enforced across the network.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_compute_server_power_policy.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the policy. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `power_state`:(string) User configured power state of server.* `Policy` - Power state is set to the default value in the policy.* `PowerOn` - Power state of the server is set to On.* `PowerOff` - Power state is the server set to Off.* `PowerCycle` - Power state the server is reset.* `HardReset` - Power state the server is hard reset.* `Shutdown` - Operating system on the server is shut down.* `Reboot` - Power state of IMC is rebooted. 
* `server_name`:(string) The name of the server it is associated with. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
