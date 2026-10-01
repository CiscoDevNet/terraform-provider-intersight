---
subcategory: "fabric"
layout: "intersight"
page_title: "Intersight: intersight_fabric_eth_network_group_policy"
description: |-
        EthNetworkGroupPolicies define the VLAN set that is allowed on an Ethernet-capable virtual interface. They act as the reusable “Ethernet network group” policy that encapsulates VLAN behavior (for example, which VLANs are permitted) and can be attached wherever a standardized VLAN allowance is needed for server or switch profile connectivity.
        #### Purpose
        Provide a centrally managed, reusable policy for controlling allowed VLAN configuration on virtual interfaces, enabling consistent segmentation and connectivity behavior across profiles and deployments.
        #### Key Concepts
        - **Allowed-VLAN policy intent**: Models the VLAN allowance for an interface; the core configuration is captured in `vlanSettings`.
        - **Reusable abstract-policy pattern**: Extends `policy.AbstractPolicy`, supporting reuse and consistent lifecycle management across multiple consumers.
        - **Profile ecosystem integration**: Designed to be used with both server profiles and switch profiles (as reflected in privileges).
        - **Inventory generation**: `generateinventoryobject: true` indicates the system can produce related inventory/observed objects derived from this policy’s application.
        - **Licensed CRUD**: READ/CREATE/UPDATE/DELETE are available under the **Essentials** entitlement.
        - **Platform targeting via tags**: Tagged for applicable management platforms (UCS FI/ISM, unified edge server, FI-attached) and supports shared-object cloning (`clone.allowSharedObject: true`).
        - **Stable identity and discoverability**: Identified by `name`, and indexed by `Description` for search/filtering.

---

# Data Source: intersight_fabric_eth_network_group_policy
EthNetworkGroupPolicies define the VLAN set that is allowed on an Ethernet-capable virtual interface. They act as the reusable “Ethernet network group” policy that encapsulates VLAN behavior (for example, which VLANs are permitted) and can be attached wherever a standardized VLAN allowance is needed for server or switch profile connectivity.
#### Purpose
Provide a centrally managed, reusable policy for controlling allowed VLAN configuration on virtual interfaces, enabling consistent segmentation and connectivity behavior across profiles and deployments.
#### Key Concepts
- **Allowed-VLAN policy intent**: Models the VLAN allowance for an interface; the core configuration is captured in `vlanSettings`.
- **Reusable abstract-policy pattern**: Extends `policy.AbstractPolicy`, supporting reuse and consistent lifecycle management across multiple consumers.
- **Profile ecosystem integration**: Designed to be used with both server profiles and switch profiles (as reflected in privileges).
- **Inventory generation**: `generateinventoryobject: true` indicates the system can produce related inventory/observed objects derived from this policy’s application.
- **Licensed CRUD**: READ/CREATE/UPDATE/DELETE are available under the **Essentials** entitlement.
- **Platform targeting via tags**: Tagged for applicable management platforms (UCS FI/ISM, unified edge server, FI-attached) and supports shared-object cloning (`clone.allowSharedObject: true`).
- **Stable identity and discoverability**: Identified by `name`, and indexed by `Description` for search/filtering.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_fabric_eth_network_group_policy.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the policy. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
