---
subcategory: "storage"
layout: "intersight"
page_title: "Intersight: intersight_storage_pure_smb_policy_rule"
description: |-
        The PureSmbPolicyRule object represents SMB policy rules within a PureStorage FlashArray, ensuring organized and efficient storage operations.
        #### Purpose
        PureSmbPolicyRule facilitates the management of SMB policies, providing structured control and allocation of storage resources within the FlashArray.
        #### Key Concepts
        - **Policy Management:** Offers structured management of SMB policy rules, optimizing storage operations and resource allocation.
        - **Integration:** Supports integration with directory objects, enhancing storage management capabilities.
        - **Secure Access:** Utilizes privilege sets to secure and manage SMB policy rules within the storage network.

---

# Data Source: intersight_storage_pure_smb_policy_rule
The PureSmbPolicyRule object represents SMB policy rules within a PureStorage FlashArray, ensuring organized and efficient storage operations. 
#### Purpose 
PureSmbPolicyRule facilitates the management of SMB policies, providing structured control and allocation of storage resources within the FlashArray.
#### Key Concepts
- **Policy Management:** Offers structured management of SMB policy rules, optimizing storage operations and resource allocation. 
- **Integration:** Supports integration with directory objects, enhancing storage management capabilities. 
- **Secure Access:** Utilizes privilege sets to secure and manage SMB policy rules within the storage network.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_storage_pure_smb_policy_rule.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `anonymous_access_allowed`:(bool) Anonymous access allowed of the policy rule. 
* `client_name`:(string) Client of the policy rule. 
* `create_time`:(string) The time when this managed object was created. 
* `destroyed`:(bool) A boolean value, if set to true, lists only destroyed objects that are in the eradication pending state. If set to false, lists only objects that are not destroyed. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the SMB policy rule. 
* `policy_name`:(string) Name of the SMB policy rule. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `smb_encryption_required`:(bool) SMB encryption required of the policy rule. 
 
