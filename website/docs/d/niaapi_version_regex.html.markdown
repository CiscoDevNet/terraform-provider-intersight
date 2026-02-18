---
subcategory: "niaapi"
layout: "intersight"
page_title: "Intersight: intersight_niaapi_version_regex"
description: |-
        The VersionRegex object is integral to software version management, providing regular expression patterns for recognizing version strings. It is crucial for parsing and managing software versions efficiently.
        #### Purpose
        VersionRegex provides regular expression patterns to parse software version strings, ensuring users can accurately identify and manage software releases. It supports efficient version control and management.
        #### Key Concepts
        - **Regex Patterns:** Offers regular expressions for parsing version strings, aiding users in managing software versions.
        - **Version Identification:** Facilitates accurate recognition of software releases, supporting effective version control.
        - **System Integration:** Ensures compatibility with systems for seamless version parsing and management.

---

# Data Source: intersight_niaapi_version_regex
The VersionRegex object is integral to software version management, providing regular expression patterns for recognizing version strings. It is crucial for parsing and managing software versions efficiently.
#### Purpose
VersionRegex provides regular expression patterns to parse software version strings, ensuring users can accurately identify and manage software releases. It supports efficient version control and management.
#### Key Concepts
- **Regex Patterns:** Offers regular expressions for parsing version strings, aiding users in managing software versions.
- **Version Identification:** Facilitates accurate recognition of software releases, supporting effective version control.
- **System Integration:** Ensures compatibility with systems for seamless version parsing and management.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_niaapi_version_regex.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `nr_version`:(string) Version number for the Version Regex data, also used as identity. 
 
