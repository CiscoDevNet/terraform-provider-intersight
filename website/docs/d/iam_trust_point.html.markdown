---
subcategory: "iam"
layout: "intersight"
page_title: "Intersight: intersight_iam_trust_point"
description: |-
        The TrustPoint object represents a trusted source of certificates within the system, enabling secure communication and authentication.
        #### Purpose
        A TrustPoint serves as a repository for certificates, affirming the identity of trusted entities and facilitating secure interactions with external systems.
        #### Key Concepts
        - **Certificate Management:** Stores and manages X.509 certificates, supporting both root and intermediate certificates for authentication purposes.
        - **Security Assurance:** Provides a basis for secure communication by validating the identity of external entities through trusted certificates.
        - **Account Association:** Associates certificates with specific accounts, ensuring they are available for authentication within the account's scope.
        - **Certificate Import:** Allows the import of third-party certificates, enhancing security through integration with external certificate authorities.

---

# Data Source: intersight_iam_trust_point
The TrustPoint object represents a trusted source of certificates within the system, enabling secure communication and authentication.
#### Purpose
A TrustPoint serves as a repository for certificates, affirming the identity of trusted entities and facilitating secure interactions with external systems.
#### Key Concepts
- **Certificate Management:** Stores and manages X.509 certificates, supporting both root and intermediate certificates for authentication purposes.
- **Security Assurance:** Provides a basis for secure communication by validating the identity of external entities through trusted certificates.
- **Account Association:** Associates certificates with specific accounts, ensuring they are available for authentication within the account's scope.
- **Certificate Import:** Allows the import of third-party certificates, enhancing security through integration with external certificate authorities.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iam_trust_point.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `chain`:(string) The certificate information for this trusted point. The certificate must be in Base64 encoded X.509 (CER) format. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
