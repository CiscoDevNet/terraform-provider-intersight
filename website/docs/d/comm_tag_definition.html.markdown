---
subcategory: "comm"
layout: "intersight"
page_title: "Intersight: intersight_comm_tag_definition"
description: |-
        An arbitrary string. Tag can be assigned to any managed object. The key will be copied to the managed object.
        The tag key is a string representation of the tag. The tag key can be hierarchical or key-value pair.

---

# Data Source: intersight_comm_tag_definition
An arbitrary string. Tag can be assigned to any managed object. The key will be copied to the managed object.
The tag key is a string representation of the tag. The tag key can be hierarchical or key-value pair.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_comm_tag_definition.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `enable_propagation`:(bool) If this flag is enabled, the tag will be propagated to related managed objects.Propagation is supported in a limited manner for path tags and it is not supported for key value. Rules for propagation areconfigured by Intersight and cannot be configured by user. 
* `key`:(string) The string representation of the tag key. If the tag is of path type, then \ /\  will be interpreted as path delimiters.The tag key must be unique within the account. The tag key is case sensitive and must not be empty. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `restrict_values`:(bool) If this flag is enabled, then values of the KeyValue tag is restricted to values present in the allowedValues list. RestrictValues is not applicable to path tags. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `sys_tag`:(bool) Specifies whether the tag is user-defined or owned by the system. 
* `type`:(string) An enum type that defines the type of tag. Only path tags are supported for now, and the type is set to path by default.* `KeyValue` - KeyValue type of tag. Key is required for these tags. Value is optional.* `PathTag` - Key contain path information. Value is not present for these tags. The path is created by using the '/' character as a delimiter.For example, if the tag is \ A/B/C\ , then \ A\  is the parent tag, \ B\  is the child tag of \ A\  and \ C\  is the child tag of \ B\ . 
 
