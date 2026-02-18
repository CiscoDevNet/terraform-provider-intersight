---
subcategory: "deviceconnector"
layout: "intersight"
page_title: "Intersight: intersight_deviceconnector_policy"
description: |-
        The Device Connector Policy is a reusable policy that controls whether configuration changes are permitted directly on a server's local management interface (Cisco IMC) when it is being managed by Intersight.
        #### Purpose
        The purpose of this policy is to enforce a single source of truth for server configuration. When a server is managed by Intersight, this policy can be used to lock the configuration on the endpoint, preventing out-of-band changes that could cause configuration drift and conflicts with the settings defined in the Intersight server profile.
        #### Key Concepts
        - **Configuration Lockout:** Its primary feature is the lockoutEnabled flag, which, when set to true, prevents administrators from making configuration changes through the local Cisco IMC interface.
        - **Standalone Server Management:** This policy is specifically designed for servers operating in Standalone Mode (not connected to Fabric Interconnect, i.e. not in IMM mode).
        - **Profile-Based Enforcement:** It is applied to servers by including it in a Server Profile.
        - **Impact Awareness:** Applying this policy can cause a brief management network disconnection, which is flagged by the MgmtNetworkDisconnection disruption type.

---

# Data Source: intersight_deviceconnector_policy
The Device Connector Policy is a reusable policy that controls whether configuration changes are permitted directly on a server's local management interface (Cisco IMC) when it is being managed by Intersight.
#### Purpose
The purpose of this policy is to enforce a single source of truth for server configuration. When a server is managed by Intersight, this policy can be used to "lock" the configuration on the endpoint, preventing out-of-band changes that could cause configuration drift and conflicts with the settings defined in the Intersight server profile.
#### Key Concepts
- **Configuration Lockout:** Its primary feature is the lockoutEnabled flag, which, when set to true, prevents administrators from making configuration changes through the local Cisco IMC interface.
- **Standalone Server Management:** This policy is specifically designed for servers operating in Standalone Mode (not connected to Fabric Interconnect, i.e. not in IMM mode).
- **Profile-Based Enforcement:** It is applied to servers by including it in a Server Profile.
- **Impact Awareness:** Applying this policy can cause a brief management network disconnection, which is flagged by the MgmtNetworkDisconnection disruption type.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_deviceconnector_policy.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the policy. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `lockout_enabled`:(bool) Enables configuration lockout on the endpoint. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
