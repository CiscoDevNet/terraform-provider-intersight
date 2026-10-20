---
subcategory: "niatelemetry"
layout: "intersight"
page_title: "Intersight: intersight_niatelemetry_aaa_ldap_provider_details"
description: |-
        The AaaLdapProviderDetails object captures LDAP provider information for authentication in APIC environments. It streamlines integration with external LDAP systems, supporting secure and scalable identity management.
        #### Purpose
        AaaLdapProviderDetails documents configuration and operational parameters for LDAP providers, enabling APIC to leverage directory-based authentication.
        #### Key Concepts
        - **Directory Integration**: Supports LDAP authentication for APIC.
        - **Provider Configuration**: Captures settings such as base DN, port, and root DN.
        - **Operational Security**: Enhances secure access and identity management.

---

# Data Source: intersight_niatelemetry_aaa_ldap_provider_details
The AaaLdapProviderDetails object captures LDAP provider information for authentication in APIC environments. It streamlines integration with external LDAP systems, supporting secure and scalable identity management.
#### Purpose
AaaLdapProviderDetails documents configuration and operational parameters for LDAP providers, enabling APIC to leverage directory-based authentication.
#### Key Concepts
- **Directory Integration**: Supports LDAP authentication for APIC.
- **Provider Configuration**: Captures settings such as base DN, port, and root DN.
- **Operational Security**: Enhances secure access and identity management.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_niatelemetry_aaa_ldap_provider_details.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `base_dn`:(string) Base dn of the AAA ldap provider in APIC. 
* `create_time`:(string) The time when this managed object was created. 
* `dn`:(string) Dn of the AAA ldap provider in APIC. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `port`:(string) Port of the AAA ldap provider in APIC. 
* `record_type`:(string) Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. 
* `record_version`:(string) Version of record being pushed. This determines what was the API version for data available from the device. 
* `root_dn`:(string) Root dn of the AAA ldap provider in APIC. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `site_name`:(string) Name of the APIC site from which this data is being collected. 
 
