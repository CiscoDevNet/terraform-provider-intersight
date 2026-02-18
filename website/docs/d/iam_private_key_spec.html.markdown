---
subcategory: "iam"
layout: "intersight"
page_title: "Intersight: intersight_iam_private_key_spec"
description: |-
        PrivateKeySpec defines parameters for generating a private key that will be used to sign a certificate signing request (CSR).
        #### Purpose
        It standardizes the key generation process for certificate workflows.
        #### Key Concepts
        - **CSR Workflow:** Couples key generation with a certificate request.
        - **Algorithm Selection:** Encapsulates key algorithm and parameters.
        - **Scoped Inheritance:** Governed by the associated certificate request.
        - **Security Posture:** Ensures consistent, policy-aligned key creation.

---

# Data Source: intersight_iam_private_key_spec
PrivateKeySpec defines parameters for generating a private key that will be used to sign a certificate signing request (CSR).
#### Purpose
It standardizes the key generation process for certificate workflows.
#### Key Concepts
- **CSR Workflow:** Couples key generation with a certificate request.
- **Algorithm Selection:** Encapsulates key algorithm and parameters.
- **Scoped Inheritance:** Governed by the associated certificate request.
- **Security Posture:** Ensures consistent, policy-aligned key creation.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iam_private_key_spec.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
