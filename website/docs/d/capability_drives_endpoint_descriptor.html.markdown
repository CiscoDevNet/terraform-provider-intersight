---
subcategory: "capability"
layout: "intersight"
page_title: "Intersight: intersight_capability_drives_endpoint_descriptor"
description: |-
        The capability.DrivesEndpointDescriptor object is a descriptor used to uniquely identify a storage drive (e.g., HDD, SSD) within the capability catalog. It serves as a centralized inventory record for all supported drive models.
        #### Purpose
        The main function of this object is to create and maintain a definitive catalog of storage drives. It stores essential identification attributes like vendor, model, pid (Product ID), and partNumber. This catalog is used by the system to validate drive compatibility with server platforms, verify hardware configurations, and provide detailed inventory reports.
        #### Key Concepts
        - **Storage Drive Catalog:** Provides a structured system for managing information about various drive models.
        - **Unique Identification:** Uniquely identifies each drive type using a combination of vendor, model, and pid.
        - **Alias Support:** The aliasModel property allows for mapping multiple model names to a single descriptor, accommodating vendor rebranding or minor variations.
        - **Platform Compatibility:** The supportedPlatformsPids list links each drive to the specific server platforms it is supported on, enabling automated compatibility checks.

---

# Data Source: intersight_capability_drives_endpoint_descriptor
The capability.DrivesEndpointDescriptor object is a descriptor used to uniquely identify a storage drive (e.g., HDD, SSD) within the capability catalog. It serves as a centralized inventory record for all supported drive models.
#### Purpose
The main function of this object is to create and maintain a definitive catalog of storage drives. It stores essential identification attributes like vendor, model, pid (Product ID), and partNumber. This catalog is used by the system to validate drive compatibility with server platforms, verify hardware configurations, and provide detailed inventory reports.
#### Key Concepts
- **Storage Drive Catalog:** Provides a structured system for managing information about various drive models.
- **Unique Identification:** Uniquely identifies each drive type using a combination of vendor, model, and pid.
- **Alias Support:** The aliasModel property allows for mapping multiple model names to a single descriptor, accommodating vendor rebranding or minor variations.
- **Platform Compatibility:** The supportedPlatformsPids list links each drive to the specific server platforms it is supported on, enabling automated compatibility checks.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_capability_drives_endpoint_descriptor.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `alias_model`:(string) This field is to provide alias model of the drive. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) This field is to provide description of the drive. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) This field is to provide model of the drive. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `part_number`:(string) This field is to provide partNumber of the drive. 
* `pid`:(string) This field is to provide pid of the drive. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `vendor`:(string) This field is to provide vendor of the drive. 
 
