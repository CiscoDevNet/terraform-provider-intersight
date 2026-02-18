---
subcategory: "niaapi"
layout: "intersight"
page_title: "Intersight: intersight_niaapi_sn_validator_metadata"
description: |-
        The SnValidatorMetadata object is crucial for metadata management, offering structured information about Serial Metadata files. It supports users in accessing and verifying metadata efficiently.
        #### Purpose
        SnValidatorMetadata provides comprehensive details about Serial Metadata files, including versioning and checksum information. It is essential for users to verify the integrity of metadata files.
        #### Key Concepts
        - **Metadata Management:** Facilitates the download and verification of Serial Metadata files.
        - **Versioning and Integrity:** Provides version numbers and checksums for metadata files.
        - **Structured Access:** Ensures users can efficiently access Serial Metadata in their systems.

---

# Data Source: intersight_niaapi_sn_validator_metadata
The SnValidatorMetadata object is crucial for metadata management, offering structured information about Serial Metadata files. It supports users in accessing and verifying metadata efficiently.
#### Purpose
SnValidatorMetadata provides comprehensive details about Serial Metadata files, including versioning and checksum information. It is essential for users to verify the integrity of metadata files.
#### Key Concepts
- **Metadata Management:** Facilitates the download and verification of Serial Metadata files.
- **Versioning and Integrity:** Provides version numbers and checksums for metadata files.
- **Structured Access:** Ensures users can efficiently access Serial Metadata in their systems.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_niaapi_sn_validator_metadata.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `checksum`:(string) Checksum of SnValidatorData file. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `file_name`:(string) The filename of sn metadata file. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `nr_version`:(int) The version number of the SerialNumber Metadata. 
 
