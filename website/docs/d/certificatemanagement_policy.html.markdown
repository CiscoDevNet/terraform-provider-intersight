---
subcategory: "certificatemanagement"
layout: "intersight"
page_title: "Intersight: intersight_certificatemanagement_policy"
description: |-
        The Certificate Management Policy is a reusable policy for managing security certificates on endpoints such as servers and Fabric Interconnects. This provides a centralized way to deploy custom certificates and their corresponding private keys.
        #### Purpose
        The main purpose of this policy is to replace the default, self-signed certificates on management controllers (like a server's CIMC) with custom certificates signed by a trusted Certificate Authority (CA). This enhances security by ensuring that management interfaces are trusted and that communication is encrypted. It also supports the deployment of Root CA certificates for validating external services.
        #### Key Concepts
        - **Centralized Certificate Deployment:** Allows administrators to manage and deploy X.509 certificates and private keys from a single policy to multiple endpoints.
        - **Reusable and Profile-Based:** As a policy object, it can be attached to Server or Chassis Profiles to ensure consistent certificate configuration across managed devices.
        - **Supports Multiple Certificate Types:** The policy can manage different types of certificates, including IMC certificates (for the management controller itself) and Root CA certificates (for trusting external services like LDAP).
        - **Secure Credential Handling:** The private key associated with a certificate is a write-only, encrypted property, ensuring it is handled securely within the system.

---

# Data Source: intersight_certificatemanagement_policy
The Certificate Management Policy is a reusable policy for managing security certificates on endpoints such as servers and Fabric Interconnects. This provides a centralized way to deploy custom certificates and their corresponding private keys.
#### Purpose
The main purpose of this policy is to replace the default, self-signed certificates on management controllers (like a server's CIMC) with custom certificates signed by a trusted Certificate Authority (CA). This enhances security by ensuring that management interfaces are trusted and that communication is encrypted. It also supports the deployment of Root CA certificates for validating external services.
#### Key Concepts
- **Centralized Certificate Deployment:** Allows administrators to manage and deploy X.509 certificates and private keys from a single policy to multiple endpoints.
- **Reusable and Profile-Based:** As a policy object, it can be attached to Server or Chassis Profiles to ensure consistent certificate configuration across managed devices.
- **Supports Multiple Certificate Types:** The policy can manage different types of certificates, including IMC certificates (for the management controller itself) and Root CA certificates (for trusting external services like LDAP).
- **Secure Credential Handling:** The private key associated with a certificate is a write-only, encrypted property, ensuring it is handled securely within the system.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_certificatemanagement_policy.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the policy. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
