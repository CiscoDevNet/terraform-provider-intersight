---
subcategory: "bulk"
layout: "intersight"
page_title: "Intersight: intersight_bulk_mo_merger"
description: |-
        ### Overview
        The MoMerger object is an interface designed to merge properties of managed object (MO) instances. It allows for the synchronization and updating of multiple target MOs based on specified source configurations.
        #### Purpose
        MoMerger is instrumental in template synchronization, enabling bulk updates of derived instances with changes made to their parent templates, ensuring consistency across configurations.
        #### Key Concepts
        - **Property Merging: Supports the merging of selected properties from source to target instances, allowing for efficient updates and configuration synchronization.
        - **Template Synchronization: Facilitates the propagation of template changes to derived profiles, ensuring uniformity and adherence to updated configurations.
        - **Configuration Application: Offers the ability to apply specific configuration changes across all target instances during the merge process.
        - **Async Processing: Operates asynchronously, accommodating large-scale merge operations without impacting system performance.

---

# Data Source: intersight_bulk_mo_merger
### Overview
The MoMerger object is an interface designed to merge properties of managed object (MO) instances. It allows for the synchronization and updating of multiple target MOs based on specified source configurations.   
#### Purpose  
MoMerger is instrumental in template synchronization, enabling bulk updates of derived instances with changes made to their parent templates, ensuring consistency across configurations.   
#### Key Concepts 
- **Property Merging: Supports the merging of selected properties from source to target instances, allowing for efficient updates and configuration synchronization. 
- **Template Synchronization: Facilitates the propagation of template changes to derived profiles, ensuring uniformity and adherence to updated configurations. 
- **Configuration Application: Offers the ability to apply specific configuration changes across all target instances during the merge process. 
- **Async Processing: Operates asynchronously, accommodating large-scale merge operations without impacting system performance.
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
 
