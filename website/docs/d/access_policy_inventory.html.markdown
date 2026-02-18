---
subcategory: "access"
layout: "intersight"
page_title: "Intersight: intersight_access_policy_inventory"
description: |-
        The Access Policy object is a reusable policy that configures server and chassis management access options. It allows administrators to define how In-Band and Out-of-Band management IP addresses are assigned and used for endpoints.
        #### Purpose
        The primary purpose of an Access Policy is to standardize the network access configuration for server management controllers (CIMCs). By defining settings for IP address pools and VLANs, this policy ensures that servers and chassis are consistently and correctly configured for management access across an organization.
        #### Key Concepts
        - **Reusable Configuration:** As an object extending policy.AbstractPolicy, this provides a template for access settings that can be applied to multiple server and chassis profiles.
        - **In-Band and Out-of-Band:** The policy supports both In-Band (shared with data traffic) and Out-of-Band (dedicated management network) configurations.
        - **IP Address Management:** It integrates with Intersight's IP Pool (ippool.Pool) and VRF (vrf.Vrf) objects to automate the assignment of management IP addresses.
        - **Profile-Based Application:** This policy is attached to Server or Chassis profiles to apply its configuration to the associated physical hardware.
        - **Impact Awareness:** Changes to this policy are flagged with MgmtNetworkDisconnection, indicating that applying updates may temporarily interrupt management connectivity to the endpoint.

---

# Data Source: intersight_access_policy_inventory
The Access Policy object is a reusable policy that configures server and chassis management access options. It allows administrators to define how In-Band and Out-of-Band management IP addresses are assigned and used for endpoints.
#### Purpose
The primary purpose of an Access Policy is to standardize the network access configuration for server management controllers (CIMCs). By defining settings for IP address pools and VLANs, this policy ensures that servers and chassis are consistently and correctly configured for management access across an organization.
#### Key Concepts
- **Reusable Configuration:** As an object extending policy.AbstractPolicy, this provides a template for access settings that can be applied to multiple server and chassis profiles.
- **In-Band and Out-of-Band:** The policy supports both In-Band (shared with data traffic) and Out-of-Band (dedicated management network) configurations.
- **IP Address Management:** It integrates with Intersight's IP Pool (ippool.Pool) and VRF (vrf.Vrf) objects to automate the assignment of management IP addresses.
- **Profile-Based Application:** This policy is attached to Server or Chassis profiles to apply its configuration to the associated physical hardware.
- **Impact Awareness:** Changes to this policy are flagged with MgmtNetworkDisconnection, indicating that applying updates may temporarily interrupt management connectivity to the endpoint.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_access_policy_inventory.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the policy. 
* `device_mo_id`:(string) Device ID of the entity from where inventory is reported. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `inband_vlan`:(int) VLAN to be used for server access over Inband network. When Inband is enabled, only numbers between 4 to 4093 are allowed. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the inventoried policy object. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
