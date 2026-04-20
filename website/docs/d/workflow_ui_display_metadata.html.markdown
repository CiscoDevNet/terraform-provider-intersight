---
subcategory: "workflow"
layout: "intersight"
page_title: "Intersight: intersight_workflow_ui_display_metadata"
description: |-
        The UiDisplayMetadata object captures the visual layout metadata for user interfaces in workflow systems. It focuses on defining form and view layouts, contributing to user-friendly and efficient workflow interactions.
        #### Purpose
        UiDisplayMetadata is responsible for managing the visual presentation of workflow interfaces, including both input forms and data views. It facilitates UI layout configuration to enhance user interaction, streamline navigation, and improve data visualization within workflows, ensuring a consistent and intuitive user experience.
        #### Key concepts
        - **Form and View Integration:** Provides metadata for input forms and views, promoting cohesive and intuitive user experiences.
        - **Customizability:** Offers flexibility in UI design, allowing workflows to adapt to varying user and organizational needs.

---

# Data Source: intersight_workflow_ui_display_metadata
The UiDisplayMetadata object captures the visual layout metadata for user interfaces in workflow systems. It focuses on defining form and view layouts, contributing to user-friendly and efficient workflow interactions. 
#### Purpose  
UiDisplayMetadata is responsible for managing the visual presentation of workflow interfaces, including both input forms and data views. It facilitates UI layout configuration to enhance user interaction, streamline navigation, and improve data visualization within workflows, ensuring a consistent and intuitive user experience.
#### Key concepts 
- **Form and View Integration:** Provides metadata for input forms and views, promoting cohesive and intuitive user experiences. 
- **Customizability:** Offers flexibility in UI design, allowing workflows to adapt to varying user and organizational needs.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_workflow_ui_display_metadata.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
