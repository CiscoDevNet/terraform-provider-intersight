---
subcategory: "niatelemetry"
layout: "intersight"
page_title: "Intersight: intersight_niatelemetry_fault"
description: |-
        The Faults object provides detailed information about faults occurring within a fabric scope.
        #### Purpose
        It serves as a centralized repository for fault data, allowing administrators to identify, analyze, and remediate issues within the fabric.
        #### Key Concepts
        - **Fault Reporting:** Captures fault codes, severity, cause, and creation time.
        - **Operational Visibility:** Provides the necessary context to diagnose and resolve fabric-level incidents.

---

# Data Source: intersight_niatelemetry_fault
The Faults object provides detailed information about faults occurring within a fabric scope.
#### Purpose
It serves as a centralized repository for fault data, allowing administrators to identify, analyze, and remediate issues within the fabric.
#### Key Concepts
- **Fault Reporting:** Captures fault codes, severity, cause, and creation time.
- **Operational Visibility:** Provides the necessary context to diagnose and resolve fabric-level incidents.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_niatelemetry_fault.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `cause`:(string) Cause of the fault present. 
* `code`:(string) Code of the fault present. 
* `create_time`:(string) The time when this managed object was created. 
* `created_time`:(string) Created time of the fault present. 
* `description`:(string) Description of the fault present. 
* `dn`:(string) Dn value for the fault present. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `record_type`:(string) Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. 
* `record_version`:(string) Version of record being pushed. This determines what was the API version for data available from the device. 
* `severity`:(string) Severity of the fault present. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `site_name`:(string) The Site name represents an APIC cluster. Service Engine can onboard multiple APIC clusters/sites. 
* `type`:(string) Type of the fault present. 
 
