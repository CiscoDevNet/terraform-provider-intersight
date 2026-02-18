---
subcategory: "license"
layout: "intersight"
page_title: "Intersight: intersight_license_license_info_view"
description: |-
        The LicenseInfoView object provides a custom view for license information, designed to be displayed by user interfaces, enhancing user experience and information accessibility.
        #### Purpose
        LicenseInfoView serves as a presentation layer for license information, offering structured views that enhance the user's ability to access and understand licensing data.
        #### Key Concepts
        - **Custom View Creation:** Facilitates the creation of tailored views for license information, ensuring data is presented in an accessible and user-friendly manner.
        - **UI Integration:** Designed for integration with user interfaces, providing seamless access to licensing data.
        - **Information Accessibility:** Enhances the ability of users to access and interpret license information, supporting informed decision-making.
        - **Organizational Resource:** Operates within shared organizational contexts, ensuring consistent presentation across different accounts and user roles.

---

# Data Source: intersight_license_license_info_view
The LicenseInfoView object provides a custom view for license information, designed to be displayed by user interfaces, enhancing user experience and information accessibility.
#### Purpose
LicenseInfoView serves as a presentation layer for license information, offering structured views that enhance the user's ability to access and understand licensing data.
#### Key Concepts
- **Custom View Creation:** Facilitates the creation of tailored views for license information, ensuring data is presented in an accessible and user-friendly manner.
- **UI Integration:** Designed for integration with user interfaces, providing seamless access to licensing data.
- **Information Accessibility:** Enhances the ability of users to access and interpret license information, supporting informed decision-making.
- **Organizational Resource:** Operates within shared organizational contexts, ensuring consistent presentation across different accounts and user roles.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_license_license_info_view.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
