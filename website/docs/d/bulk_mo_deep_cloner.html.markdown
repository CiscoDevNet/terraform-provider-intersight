---
subcategory: "bulk"
layout: "intersight"
page_title: "Intersight: intersight_bulk_mo_deep_cloner"
description: |-
        The MODeepCloner interface facilitates making n number of deep copies of any resource instance which supports the CREATE operation.
        The MO to be cloned should be specified as an MoRef object in the Source.
        The Targets array should contain n JSON documents each compliant to RFC 7386.
        For each target MO to be created, the user can specify the following -
        - new values for the identity properties, if applicable
        - new values for specific properties or references of the source MO which need to be overridden in the cloned object.

---

# Data Source: intersight_bulk_mo_deep_cloner
The MODeepCloner interface facilitates making n number of deep copies of any resource instance which supports the CREATE operation.
The MO to be cloned should be specified as an MoRef object in the "Source".
The "Targets" array should contain n JSON documents each compliant to RFC 7386. 
For each target MO to be created, the user can specify the following -
- new values for the identity properties, if applicable
- new values for specific properties or references of the source MO which need to be overridden in the cloned object.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_bulk_mo_deep_cloner.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `reference_name_suffix`:(string) Name suffix to be applied to all the MOs being cloned when ReferencePolicy chosen is CreateNew. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-) or an underscore (_). 
* `reference_policy`:(string) User selected reference clone behavior. Applies to all the MOs being cloned.* `ReuseAll` - Any policies in the destination organization whose name matches the policy referenced in the cloned policy will be attached. If no policyin the destination organization matches by name, a policy will be cloned with the same name.Pool references will always be matched by name. If not found, the pool will be cloned in the destination organization, but no identifierblocks will be created.* `CreateNew` - New policies will be created for the source and all the attached policies. If a policy of the same name and type already exists in thedestination organization or any organization from which it shares policies, a clone will be created with the provided suffix added to the name.Pool references will always be matched by name. If not found, the pool will be cloned in the destination organization, but no identifierblocks will be created. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `workflow_name_suffix`:(string) A user-friendly short name to identify the workflow. Name can only contain letters (a-z, A-Z),numbers (0-9), hyphen (-), period (.), colon (:), space ( ), forward slash (/), comma or an underscore (_). 
 
