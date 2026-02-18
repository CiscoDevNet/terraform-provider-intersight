---
subcategory: "resource"
layout: "intersight"
page_title: "Intersight: intersight_resource_selection_criteria"
description: |-
        The SelectionCriteria object enables the selection of resources based on specific criteria, providing flexibility and precision in resource management.
        #### Purpose
        SelectionCriteria serves as the mechanism for defining conditions that select resources within the system, offering a customizable and efficient approach to resource retrieval and management. It supports both static and dynamic resource selection methods.
        #### Key Concepts
        - **Dynamic Selection:** Allows for the use of OData query filters to dynamically select resources, enhancing flexibility in resource management.
        - **Customization:** Supports the definition of custom criteria for resource selection, providing tailored resource management solutions.
        - **Access Control:** Implements privilege sets to manage access to selection criteria data, ensuring secure and controlled resource selection.

---

# Data Source: intersight_resource_selection_criteria
The SelectionCriteria object enables the selection of resources based on specific criteria, providing flexibility and precision in resource management.
#### Purpose
SelectionCriteria serves as the mechanism for defining conditions that select resources within the system, offering a customizable and efficient approach to resource retrieval and management. It supports both static and dynamic resource selection methods.
#### Key Concepts
- **Dynamic Selection:** Allows for the use of OData query filters to dynamically select resources, enhancing flexibility in resource management.
- **Customization:** Supports the definition of custom criteria for resource selection, providing tailored resource management solutions.
- **Access Control:** Implements privilege sets to manage access to selection criteria data, ensuring secure and controlled resource selection.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_resource_selection_criteria.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) The description of the Resource Selection Criteria. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the Resource Selection Criteria. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
