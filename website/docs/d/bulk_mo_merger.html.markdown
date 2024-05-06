---
subcategory: "bulk"
layout: "intersight"
page_title: "Intersight: intersight_bulk_mo_merger"
description: |-
        The MO Merger interface facilitates merging of all or selected properties of any MO instance to one or more MO instances.
        The Sources array should contain the list of source MO instances as MoRef objects.
        The Targets array should contain the list of target MO instances as MoRef objects.
        The TargetConfig property is applicable only for a merge operation. If a configuration action needs to be applied on all target MOs, it can be specified using this property.
        Currently this API is used to synchronize template update to all its derived instances for the Server Profile Templates, vNIC Templates and vHBA Templates.

---

# Data Source: intersight_bulk_mo_merger
The MO Merger interface facilitates merging of all or selected properties of any MO instance to one or more MO instances.
The "Sources" array should contain the list of source MO instances as MoRef objects.
The "Targets" array should contain the list of target MO instances as MoRef objects.
The "TargetConfig" property is applicable only for a merge operation. If a configuration action needs to be applied on all target MOs, it can be specified using this property.
Currently this API is used to synchronize template update to all its derived instances for the Server Profile Templates, vNIC Templates and vHBA Templates.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_bulk_mo_merger.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `merge_action`:(string) The type of merge action to be applied on the target MOs. * `Merge` - The null properties/relationships of the source MO will be ignored for the target MO. The non-null properties/relationships of the source will override the target MO properties/relationships.* `Replace` - Merge action as described in RFC 7386. The null properties/relationships of the source MO will be deleted on the target MO.The non-null properties/relationships of the source will override the target MO properties/relationships.When source object type is different from target, only the properties common to both source and target  will be affected.Other properties on the target will be ignored. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `workflow_name_suffix`:(string) A user-friendly short name to identify the workflow. Name can only contain letters (a-z, A-Z),numbers (0-9), hyphen (-), period (.), colon (:), space ( ), forward slash (/), comma or an underscore (_). 
 
