---
subcategory: "capability"
layout: "intersight"
page_title: "Intersight: intersight_capability_psu_descriptor"
description: |-
        The capability.PsuDescriptor object is a hardware descriptor that uniquely identifies a Power Supply Unit (PSU) within the capability catalog. It serves as a reference record for specific models of PSUs.
        #### Purpose
        The purpose of this object is to maintain a catalog of known PSU models. By capturing key identifiers like vendor, model, version, and revision, it allows the system to validate the PSUs installed in a server or chassis. This ensures that the components are recognized, supported, and compatible with the equipment they are installed in.
        #### Key Concepts
        - **Hardware Cataloging:** Provides a structured entry for a specific PSU model in the system's hardware catalog.
        - **Unique Identification:** Uses a composite key of vendor, model, version, and revision to uniquely identify a PSU type.
        - **Component Validation:** Enables the system to verify that the PSUs discovered in a device match the known, supported hardware definitions.
        - **Extensible Framework:** As a HardwareDescriptor, it is part of an extensible catalog that can be updated with new PSU models as they become available.

---

# Data Source: intersight_capability_psu_descriptor
The capability.PsuDescriptor object is a hardware descriptor that uniquely identifies a Power Supply Unit (PSU) within the capability catalog. It serves as a reference record for specific models of PSUs.
#### Purpose
The purpose of this object is to maintain a catalog of known PSU models. By capturing key identifiers like vendor, model, version, and revision, it allows the system to validate the PSUs installed in a server or chassis. This ensures that the components are recognized, supported, and compatible with the equipment they are installed in.
#### Key Concepts
- **Hardware Cataloging:** Provides a structured entry for a specific PSU model in the system's hardware catalog.
- **Unique Identification:** Uses a composite key of vendor, model, version, and revision to uniquely identify a PSU type.
- **Component Validation:** Enables the system to verify that the PSUs discovered in a device match the known, supported hardware definitions.
- **Extensible Framework:** As a HardwareDescriptor, it is part of an extensible catalog that can be updated with new PSU models as they become available.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_capability_psu_descriptor.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Detailed information about the endpoint. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) The model of the endpoint, for which this capability information is applicable. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `revision`:(string) Revision for the power supply. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `vendor`:(string) The vendor of the endpoint, for which this capability information is applicable. 
* `nr_version`:(string) The firmware or software version of the endpoint, for which this capability information is applicable. 
 
