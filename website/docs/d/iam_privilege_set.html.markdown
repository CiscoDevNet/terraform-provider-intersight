---
subcategory: "iam"
layout: "intersight"
page_title: "Intersight: intersight_iam_privilege_set"
description: |-
        A collection of privileges.

---

# Data Source: intersight_iam_privilege_set
A collection of privileges.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iam_privilege_set.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `allow_future_updates`:(bool) Flag used by UI to keep track of the user selection option for future updates of privilege sets. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the privilege set. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `is_privilege_names_updated`:(bool) Flag to indicate if the privilege names are updated. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the privilege set. 
* `privilege_set_type`:(string) Type of the privilege set.* `Internal` - Privilege set is internal to the system.* `SystemPackaged` - Privilege set is packaged by the system and user can use it as a reference for custom privilege set creation.* `SystemDefined` - Privilege set is defined by the system.* `TreeNode` - Privilege set is a tree node in the custom privilege set creation hierarchy.* `TreeRoot` - Privilege set is a tree root in the custom privilege set creation hierarchy.* `TreeLeaf` - Privilege set is a tree leaf in the custom privilege set creation hierarchy.* `UserCreated` - Privilege set is created by the user. 
* `scope`:(string) The scope of the privilege set.* `Generic` - Privilege set can be added to account wide permission or organization permissions.* `Account` - Privilege set can be added to account wide permission only. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `uuid`:(string) UUID of the privilege set. 
 
