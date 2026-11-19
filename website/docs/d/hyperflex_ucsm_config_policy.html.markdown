---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_ucsm_config_policy"
description: |-
        UcsmConfigPolicies define UCS Manager-related configuration settings used when configuring servers in a Fabric Interconnect environment. The policy captures key UCSM settings such as the service profile organization, KVM IP address assignments, and MAC prefix configuration, enabling consistent UCSM-integrated provisioning as part of HyperFlex cluster profile operations.
        #### Purpose
        Centralize and standardize UCS Manager configuration inputs required for FI-based server configuration, so HyperFlex cluster provisioning and ongoing operations can apply consistent UCSM settings through cluster profiles.
        #### Key Concepts
        - **UCSM configuration intent**: Specifies UCS Manager settings (service profile org, KVM IPs, MAC prefix) needed to drive consistent server configuration.
        - **FI/HyperFlex attachability**: Tagged for FI-based management (`managementplatform.fi: true`) and attachable to a HyperFlex cluster profile, indicating it is meant to be consumed in HyperFlex profile-driven workflows.
        - **Reusable policy pattern**: Extends `policy.AbstractPolicy`, supporting reuse and standardized policy lifecycle.
        - **Governed lifecycle**: Supports create/update/delete under the Manage HyperFlex Cluster Profiles privilege set, ensuring that only authorized administrators change UCSM configuration baselines.
        - **Organization-scoped**: `owner: organization` indicates it is tenant-owned and managed within an org boundary.

---

# Data Source: intersight_hyperflex_ucsm_config_policy
UcsmConfigPolicies define UCS Manager-related configuration settings used when configuring servers in a Fabric Interconnect environment. The policy captures key UCSM settings such as the service profile organization, KVM IP address assignments, and MAC prefix configuration, enabling consistent UCSM-integrated provisioning as part of HyperFlex cluster profile operations.
#### Purpose
Centralize and standardize UCS Manager configuration inputs required for FI-based server configuration, so HyperFlex cluster provisioning and ongoing operations can apply consistent UCSM settings through cluster profiles.
#### Key Concepts
- **UCSM configuration intent**: Specifies UCS Manager settings (service profile org, KVM IPs, MAC prefix) needed to drive consistent server configuration.
- **FI/HyperFlex attachability**: Tagged for FI-based management (`managementplatform.fi: true`) and attachable to a HyperFlex cluster profile, indicating it is meant to be consumed in HyperFlex profile-driven workflows.
- **Reusable policy pattern**: Extends `policy.AbstractPolicy`, supporting reuse and standardized policy lifecycle.
- **Governed lifecycle**: Supports create/update/delete under the Manage HyperFlex Cluster Profiles privilege set, ensuring that only authorized administrators change UCSM configuration baselines.
- **Organization-scoped**: `owner: organization` indicates it is tenant-owned and managed within an org boundary.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_ucsm_config_policy.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the policy. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `server_firmware_version`:(string) The server firmware bundle version used for server components such as CIMC, adapters, BIOS, etc. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
