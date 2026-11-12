---
subcategory: "iam"
layout: "intersight"
page_title: "Intersight: intersight_iam_certificate_request"
description: |-
        CertificateRequests model certificate signing request (CSR) workflows within an account. A CertificateRequest captures the subject identity and SANs needed to generate a PEM-formatted CSR, optionally supporting a self-signed path.
        #### Purpose
        Enable account-scoped certificate lifecycle operations by generating CSRs that can be submitted to a certificate authority (or self-signed by the appliance where supported).
        #### Key Concepts
        - **CSR generation:** Produces a PEM CSR from provided subject and SAN inputs.
        - **Self-signed option:** Supports conditional behavior for self-signed vs CA-signed flows.
        - **Account-scoped security artifact:** Inherits permissions from the owning account and is managed under account governance.
        - **Certificate pairing:** Typically links forward to a Certificate object that represents the signed certificate returned from a CA.

---

# Data Source: intersight_iam_certificate_request
CertificateRequests model certificate signing request (CSR) workflows within an account. A CertificateRequest captures the subject identity and SANs needed to generate a PEM-formatted CSR, optionally supporting a self-signed path.
#### Purpose
Enable account-scoped certificate lifecycle operations by generating CSRs that can be submitted to a certificate authority (or self-signed by the appliance where supported).
#### Key Concepts
- **CSR generation:** Produces a PEM CSR from provided subject and SAN inputs.
- **Self-signed option:** Supports conditional behavior for self-signed vs CA-signed flows.
- **Account-scoped security artifact:** Inherits permissions from the owning account and is managed under account governance.
- **Certificate pairing:** Typically links forward to a Certificate object that represents the signed certificate returned from a CA.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iam_certificate_request.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `email_address`:(string) User input email address, an optional part of the subject of the certificate request. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the certificate request. 
* `request`:(string) Generated certificate signing request (CSR) in PEM format. 
* `self_signed`:(bool) Whether the user wants the generated CSR to be self-signed by the appliance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
