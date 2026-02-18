---
subcategory: "niaapi"
layout: "intersight"
page_title: "Intersight: intersight_niaapi_nia_metadata"
description: |-
        The NiaMetadata object is a key element in metadata management, providing structured information about the latest metadata files available for download. It aids users in accessing up-to-date metadata efficiently.
        #### Purpose
        NiaMetadata offers comprehensive details about metadata packages, including versioning and checksum information. It is essential for users to access and verify the integrity of metadata files.
        #### Key Concepts
        - **Metadata Management:** Facilitates the download and verification of metadata files, ensuring users have access to accurate information.
        - **Versioning and Integrity:** Provides version numbers and checksums for metadata files, supporting integrity verification.
        - **Structured Access:** Ensures users can efficiently access and utilize metadata in their systems.

---

# Data Source: intersight_niaapi_nia_metadata
The NiaMetadata object is a key element in metadata management, providing structured information about the latest metadata files available for download. It aids users in accessing up-to-date metadata efficiently.
#### Purpose
NiaMetadata offers comprehensive details about metadata packages, including versioning and checksum information. It is essential for users to access and verify the integrity of metadata files.
#### Key Concepts
- **Metadata Management:** Facilitates the download and verification of metadata files, ensuring users have access to accurate information.
- **Versioning and Integrity:** Provides version numbers and checksums for metadata files, supporting integrity verification.
- **Structured Access:** Ensures users can efficiently access and utilize metadata in their systems.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_niaapi_nia_metadata.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `date`:(string) The date when this package is generated. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `metadata_chksum`:(string) Chksum used to check the integrity of the Metadata file downloaded. 
* `metadata_filename`:(string) The Filename of this Metadata package. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `nr_version`:(int) The version number of the Metadata package. 
 
