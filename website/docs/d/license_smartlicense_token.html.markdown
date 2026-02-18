---
subcategory: "license"
layout: "intersight"
page_title: "Intersight: intersight_license_smartlicense_token"
description: |-
        The SmartlicenseToken object manages the collection of license registration tokens, ensuring secure and efficient handling of licensing credentials.
        #### Purpose
        SmartlicenseToken serves as the repository for license registration tokens, providing mechanisms for secure storage and retrieval of licensing credentials.
        #### Key Concepts
        - **Token Management:** Centralizes storage and management of license registration tokens, ensuring secure handling of credentials.
        - **Credential Security:** Focuses on the security aspects of token management, ensuring tokens are protected and accessible only to authorized users.
        - **Organizational Resource:** Operates within shared organizational contexts, allowing consistent application across different accounts and user roles.
        - **Integration:** Designed to integrate seamlessly with licensing processes, supporting efficient registration and authorization flows.

---

# Data Source: intersight_license_smartlicense_token
The SmartlicenseToken object manages the collection of license registration tokens, ensuring secure and efficient handling of licensing credentials.
#### Purpose
SmartlicenseToken serves as the repository for license registration tokens, providing mechanisms for secure storage and retrieval of licensing credentials.
#### Key Concepts
- **Token Management:** Centralizes storage and management of license registration tokens, ensuring secure handling of credentials.
- **Credential Security:** Focuses on the security aspects of token management, ensuring tokens are protected and accessible only to authorized users.
- **Organizational Resource:** Operates within shared organizational contexts, allowing consistent application across different accounts and user roles.
- **Integration:** Designed to integrate seamlessly with licensing processes, supporting efficient registration and authorization flows.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_license_smartlicense_token.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `token`:(string) Smart license registration token. 
 
