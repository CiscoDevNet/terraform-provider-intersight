---
subcategory: "appliance"
layout: "intersight"
page_title: "Intersight: intersight_appliance_certificate_setting"
description: |-
        The CertificateSetting object manages the certificate used by the Intersight Appliance for browser traffic. It ensures secure communication between the appliance and external interfaces.
        #### Purpose
        CertificateSetting facilitates the management of certificates used for secure browser interactions. This provides administrators with the ability to update and manage the appliance's certificates, ensuring compliance with security standards.
        #### Key Concepts
        - **Security Management:** Enables secure handling of browser certificates, supporting encrypted communication.
        - **Access Control:** Restricted to account administrators, ensuring that certificate management is conducted securely.
        - **Integration with Accounts:** Establishes relationships with accounts, aligning certificate settings with account management.

---

# Data Source: intersight_appliance_certificate_setting
The CertificateSetting object manages the certificate used by the Intersight Appliance for browser traffic. It ensures secure communication between the appliance and external interfaces.
#### Purpose
CertificateSetting facilitates the management of certificates used for secure browser interactions. This provides administrators with the ability to update and manage the appliance's certificates, ensuring compliance with security standards.
#### Key Concepts
- **Security Management:** Enables secure handling of browser certificates, supporting encrypted communication.
- **Access Control:** Restricted to account administrators, ensuring that certificate management is conducted securely.
- **Integration with Accounts:** Establishes relationships with accounts, aligning certificate settings with account management.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_appliance_certificate_setting.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
