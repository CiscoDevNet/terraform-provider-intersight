---
subcategory: "oauth"
layout: "intersight"
page_title: "Intersight: intersight_oauth_authorization"
description: |-
        ### Overview
        The Authorization object underpins the OAuth2 authorization process, enabling user consent for external interactions on behalf of their account. It supports secure and efficient management of user credentials for accessing external services.
        #### Purpose
        Authorization serves as the conduit for user consent, allowing Intersight to interact with external software repositories securely on behalf of users, particularly in support of OAuth2 resource owner grant types.
        #### Key Concepts
        - **User Consent:** - Facilitates secure consent mechanisms for external interactions, ensuring user credentials are protected.
        - **Security Measures:** - Incorporates encryption and secure handling of user credentials, safeguarding access and interactions.
        - **Privileged Access:** - Ensures that only authorized administrators can create, update, or read authorizations, maintaining strict access controls.
        - **Account Integration:** - Associated with the iam.Account object, providing a structured approach to managing user authorization processes.

---

# Data Source: intersight_oauth_authorization
### Overview
The Authorization object underpins the OAuth2 authorization process, enabling user consent for external interactions on behalf of their account. It supports secure and efficient management of user credentials for accessing external services.
#### Purpose
Authorization serves as the conduit for user consent, allowing Intersight to interact with external software repositories securely on behalf of users, particularly in support of OAuth2 resource owner grant types.
#### Key Concepts
- **User Consent:** - Facilitates secure consent mechanisms for external interactions, ensuring user credentials are protected.
- **Security Measures:** - Incorporates encryption and secure handling of user credentials, safeguarding access and interactions.
- **Privileged Access:** - Ensures that only authorized administrators can create, update, or read authorizations, maintaining strict access controls.
- **Account Integration:** - Associated with the iam.Account object, providing a structured approach to managing user authorization processes.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_oauth_authorization.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `api_type`:(string) OAuth API type (e.g., Smart licensing API).* `Unknown` - Unknown serves as the default API type.* `SmartLicensing-API` - Smart licensing API type.* `CommerceEstimate-API` - Commerce Estimate API type. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `is_password_set`:(bool) Indicates whether the value of the 'password' property has been set. 
* `is_user_id_set`:(bool) Indicates whether the value of the 'userId' property has been set. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `password`:(string) Password used by Intersight to create OAuth2 tokens for interacting with the external repository on behalf of the user account. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `user_id`:(string) Username used by Intersight to create OAuth2 tokens for interacting with the external repository on behalf of the user account. 
 
