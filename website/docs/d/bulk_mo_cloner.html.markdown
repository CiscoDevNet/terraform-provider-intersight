---
subcategory: "bulk"
layout: "intersight"
page_title: "Intersight: intersight_bulk_mo_cloner"
description: |-
        ### Overview
        The MoCloner object is an interface designed to facilitate the cloning of managed objects (MOs) in bulk operations. It enables the creation of multiple copies of a specified resource instance.
        #### Purpose
        MoCloner is used primarily for template derivation operations, allowing users to replicate existing templates into new derived instances while applying specific overrides to identity properties or other configurations.
        #### Key Concepts
        - **Bulk Cloning:** - Supports the generation of multiple clones in a single operation, optimizing resource duplication processes.
        - **Template Derivation:** - Facilitates the derivation of server profiles, vNIC templates, and vHBA templates, ensuring consistency across cloned instances.
        - **Configuration Override:** - Allows for customization of cloned instances by specifying new values for identity and other properties.
        - **Async Processing:** - Supports asynchronous operation, enabling efficient handling of large-scale cloning tasks.

---

# Data Source: intersight_bulk_mo_cloner
### Overview
The MoCloner object is an interface designed to facilitate the cloning of managed objects (MOs) in bulk operations. It enables the creation of multiple copies of a specified resource instance.   
#### Purpose  
MoCloner is used primarily for template derivation operations, allowing users to replicate existing templates into new derived instances while applying specific overrides to identity properties or other configurations.   
#### Key Concepts 
- **Bulk Cloning:** - Supports the generation of multiple clones in a single operation, optimizing resource duplication processes. 
- **Template Derivation:** - Facilitates the derivation of server profiles, vNIC templates, and vHBA templates, ensuring consistency across cloned instances. 
- **Configuration Override:** - Allows for customization of cloned instances by specifying new values for identity and other properties. 
- **Async Processing:** - Supports asynchronous operation, enabling efficient handling of large-scale cloning tasks.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_bulk_mo_cloner.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `workflow_name_suffix`:(string) A user-friendly short name to identify the workflow. Name can only contain letters (a-z, A-Z),numbers (0-9), hyphen (-), period (.), colon (:), space ( ), forward slash (/), comma or an underscore (_). 
 
