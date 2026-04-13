---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_vm_import_operation"
description: |-
        The VmImportOperation object is designed to facilitate the import of Virtual Machines into the HyperFlex environment. It serves as a bridge for integrating external Virtual Machines into the HyperFlex ecosystem, ensuring seamless transitions and resource management.
        #### Purpose
        The object enables the systematic import of Virtual Machines, allowing for the expansion and integration of existing infrastructure within HyperFlex. It addresses both the operational aspects of importing and the necessary configurations required for successful deployment.
        #### Key Concepts
        - **Integration:** Supports the incorporation of Virtual Machines from external environments, broadening the operational scope of HyperFlex clusters.
        - **Configuration Management:** Provides detailed configurations to ensure that imported Virtual Machines align with existing cluster policies and resources.
        - **Operational Efficiency:** Streamlines the import process, reducing downtime and complexity associated with integrating new Virtual Machines.
        - **Security and Permissions:** Incorporates robust access controls to safeguard the import process, ensuring that only authorized personnel can initiate and manage imports.

---

# Data Source: intersight_hyperflex_vm_import_operation
The VmImportOperation object is designed to facilitate the import of Virtual Machines into the HyperFlex environment. It serves as a bridge for integrating external Virtual Machines into the HyperFlex ecosystem, ensuring seamless transitions and resource management.
#### Purpose
The object enables the systematic import of Virtual Machines, allowing for the expansion and integration of existing infrastructure within HyperFlex. It addresses both the operational aspects of importing and the necessary configurations required for successful deployment.
#### Key Concepts
- **Integration:** Supports the incorporation of Virtual Machines from external environments, broadening the operational scope of HyperFlex clusters.
- **Configuration Management:** Provides detailed configurations to ensure that imported Virtual Machines align with existing cluster policies and resources.
- **Operational Efficiency:** Streamlines the import process, reducing downtime and complexity associated with integrating new Virtual Machines.
- **Security and Permissions:** Incorporates robust access controls to safeguard the import process, ensuring that only authorized personnel can initiate and manage imports.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_vm_import_operation.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
