---
subcategory: "oauth"
layout: "intersight"
page_title: "Intersight: intersight_oauth_access_token"
description: |-
        ### Overview
        The AccessToken object is a key component of the OAuth-based authentication system, designed to manage API access tokens associated with user accounts. This provides a secure mechanism for applications to access services on behalf of users, ensuring that permissions and access controls are upheld.
        #### Purpose
        The AccessToken object serves as a secure representation of authorization for specific applications to interact with designated services. It supports seamless and efficient access management while maintaining strict security protocols.
        #### Key Concepts
        - **Authorization Management:** - Represents the OAuth access token, which embodies the permissions granted to a specific application for accessing a service.
        - **Security:** - Ensures that access tokens are securely issued and managed, with support for refresh tokens to prolong access without user intervention.
        - **Access Control:** - The AccessToken object supports privilege sets to ensure that only authorized administrators can read or delete access tokens.
        - **Relationship Handling: Linked to the iam.Account object, representing the account responsible for managing the OAuth process.

---

# Data Source: intersight_oauth_access_token
### Overview
The AccessToken object is a key component of the OAuth-based authentication system, designed to manage API access tokens associated with user accounts. This provides a secure mechanism for applications to access services on behalf of users, ensuring that permissions and access controls are upheld.
#### Purpose
The AccessToken object serves as a secure representation of authorization for specific applications to interact with designated services. It supports seamless and efficient access management while maintaining strict security protocols.
#### Key Concepts
- **Authorization Management:** - Represents the OAuth access token, which embodies the permissions granted to a specific application for accessing a service.
- **Security:** - Ensures that access tokens are securely issued and managed, with support for refresh tokens to prolong access without user intervention.
- **Access Control:** - The AccessToken object supports privilege sets to ensure that only authorized administrators can read or delete access tokens.
- **Relationship Handling: Linked to the iam.Account object, representing the account responsible for managing the OAuth process.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_oauth_access_token.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `api_type`:(string) OAuth API type (e.g., Smart licensing API).* `Unknown` - Unknown serves as the default API type.* `SmartLicensing-API` - Smart licensing API type.* `CommerceEstimate-API` - Commerce Estimate API type. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `expiry`:(string) Access token expiration date and time. 
* `issuer`:(string) OAuth access token issuer. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `refresh_expiry`:(string) Refresh token expiration date and time. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `token_owner`:(string) MOID of the access token owner. 
 
