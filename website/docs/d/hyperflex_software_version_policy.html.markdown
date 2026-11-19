---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_software_version_policy"
description: |-
        SoftwareVersionPolicies capture the set of software versions a user intends to run across the different components that make up a HyperFlex cluster (for example, HyperFlex Data Platform and the hypervisor). The policy acts as a single place to define version selection so cluster deployments and lifecycle operations can apply a consistent, compatible software stack.
        #### Purpose
        Enable administrators to standardize and apply component software versions to HyperFlex clusters via cluster profiles, supporting controlled upgrades, repeatable deployments, and compatibility alignment across the stack.
        #### Key Concepts
        - **Version intent for multiple components**: Encapsulates desired versions for several HyperFlex-compatible components rather than a single software artifact.
        - **Profile-driven application**: Designed to be attached/applied as part of HyperFlex cluster profile workflows.
        - **Lifecycle management**: Supports controlled create/update/delete so version baselines can evolve over time (e.g., during upgrades).
        - **Reusable policy object**: Extends `policy.AbstractPolicy`, enabling reuse across clusters and consistency across environments.
        - **Role-based access**: Read access is broadly available; create/update/delete is restricted to users who manage HyperFlex cluster profiles.

---

# Data Source: intersight_hyperflex_software_version_policy
SoftwareVersionPolicies capture the set of software versions a user intends to run across the different components that make up a HyperFlex cluster (for example, HyperFlex Data Platform and the hypervisor). The policy acts as a single place to define version selection so cluster deployments and lifecycle operations can apply a consistent, compatible software stack.
#### Purpose
Enable administrators to standardize and apply component software versions to HyperFlex clusters via cluster profiles, supporting controlled upgrades, repeatable deployments, and compatibility alignment across the stack.
#### Key Concepts
- **Version intent for multiple components**: Encapsulates desired versions for several HyperFlex-compatible components rather than a single software artifact.
- **Profile-driven application**: Designed to be attached/applied as part of HyperFlex cluster profile workflows.
- **Lifecycle management**: Supports controlled create/update/delete so version baselines can evolve over time (e.g., during upgrades).
- **Reusable policy object**: Extends `policy.AbstractPolicy`, enabling reuse across clusters and consistency across environments.
- **Role-based access**: Read access is broadly available; create/update/delete is restricted to users who manage HyperFlex cluster profiles.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_software_version_policy.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the policy. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `hxdp_version`:(string) Desired HyperFlex Data Platform software version to apply on the HyperFlex cluster. 
* `hypervisor_version`:(string) Desired  hypervisor version to apply for all the nodes on the HyperFlex cluster. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `server_firmware_version`:(string) Desired server firmware version to apply on the HyperFlex Cluster. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
