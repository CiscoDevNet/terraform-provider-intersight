---
subcategory: "kvm"
layout: "intersight"
page_title: "Intersight: intersight_kvm_tunneled_kvm_policy"
description: |-
        The TunneledKvmPolicy object manages policies controlling tunneled vKVM access within the account. This provides configuration options for enabling or disabling tunneled access and configuration within the account.
        #### Purpose
        TunneledKvmPolicy object enables organizations to manage and control tunneled vKVM access through policy configuration. It ensures secure and efficient access by allowing account-specific configurations.
        #### Key Concepts
        - **Policy Management:** Manages policies for tunneled vKVM access, ensuring secure and efficient access operations.
        - **Configuration Options:** Provides options to enable or disable tunneled vKVM access and tunneled vKVM access configuration within the account, supporting tailored access control.

---

# Data Source: intersight_kvm_tunneled_kvm_policy
The TunneledKvmPolicy object manages policies controlling tunneled vKVM access within the account. This provides configuration options for enabling or disabling tunneled access and configuration within the account.
#### Purpose
TunneledKvmPolicy object enables organizations to manage and control tunneled vKVM access through policy configuration. It ensures secure and efficient access by allowing account-specific configurations.
#### Key Concepts
- **Policy Management:** Manages policies for tunneled vKVM access, ensuring secure and efficient access operations.
- **Configuration Options:** Provides options to enable or disable tunneled vKVM access and tunneled vKVM access configuration within the account, supporting tailored access control.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_kvm_tunneled_kvm_policy.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `tunneled_kvm_configuration`:(bool) Enable or Disable configuration of tunneled KVM for a specific account. 
* `tunneled_kvm_launch`:(bool) Enable or Disable launching tunneled KVM for a specific account. 
 
