---
subcategory: "capability"
layout: "intersight"
page_title: "Intersight: intersight_capability_template_catalog"
description: |-
        The TemplateCatalog object lists the settings that can be overridden in vNIC and vHBA templates, allowing flexible template management and customization.
        #### Purpose
        A TemplateCatalog serves as an authoritative source for allowed overrides in template configurations, facilitating efficient management of vNIC and vHBA resources. It supports read operations to ensure consistent access to override options.
        #### Key Concepts
        - **Override Management:** Catalogs properties that can be overridden, enhancing flexibility and customization in template configurations.
        - **Template Integration:** Focuses on vNIC and vHBA templates, ensuring streamlined management and deployment within server environments.
        - **Access Control:** Managed through privilege sets to ensure secure and reliable access to template catalog information.
        - **Versioning and Updates:** Supports meta version updates, ensuring that override capabilities are current and aligned with system requirements.

---

# Data Source: intersight_capability_template_catalog
The TemplateCatalog object lists the settings that can be overridden in vNIC and vHBA templates, allowing flexible template management and customization.
#### Purpose
A TemplateCatalog serves as an authoritative source for allowed overrides in template configurations, facilitating efficient management of vNIC and vHBA resources. It supports read operations to ensure consistent access to override options.
#### Key Concepts
- **Override Management:** Catalogs properties that can be overridden, enhancing flexibility and customization in template configurations.
- **Template Integration:** Focuses on vNIC and vHBA templates, ensuring streamlined management and deployment within server environments.
- **Access Control:** Managed through privilege sets to ensure secure and reliable access to template catalog information.
- **Versioning and Updates:** Supports meta version updates, ensuring that override capabilities are current and aligned with system requirements.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_capability_template_catalog.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) An unique identifer for a capability descriptor. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
