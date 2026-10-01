---
subcategory: "capability"
layout: "intersight"
page_title: "Intersight: intersight_capability_switch_descriptor"
description: |-
        The SwitchDescriptor object uniquely identifies a switch/fabric-interconnect platform for capability catalog mapping.
        #### Purpose
        This enables consistent matching of discovered Fabric Interconnect inventory to catalog definitions, ensuring the system can retrieve the correct capability and constraint definitions for that hardware.
        #### Key Concepts
        - **Platform identification:** Provides stable keys for platform correlation.
        - **Catalog-backed validation:** Enables retrieving platform-specific capability/limit definitions.
        - **Revision-aware mapping:** Allows differentiation among revisions of similar models.
        - **Reusable catalog anchor:** Supports consistent references across multiple capability definitions.

---

# Data Source: intersight_capability_switch_descriptor
The SwitchDescriptor object uniquely identifies a switch/fabric-interconnect platform for capability catalog mapping.
#### Purpose
This enables consistent matching of discovered Fabric Interconnect inventory to catalog definitions, ensuring the system can retrieve the correct capability and constraint definitions for that hardware.
#### Key Concepts
- **Platform identification:** Provides stable keys for platform correlation.
- **Catalog-backed validation:** Enables retrieving platform-specific capability/limit definitions.
- **Revision-aware mapping:** Allows differentiation among revisions of similar models.
- **Reusable catalog anchor:** Supports consistent references across multiple capability definitions.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_capability_switch_descriptor.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Detailed information about the endpoint. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `expected_memory`:(int) The total expected memory for this hardware. 
* `is_avatar_ecmc`:(bool) Identifies whether Switch is part of Avatar series. 
* `is_ucsx_direct_switch`:(bool) Identifies whether Switch is part of UCSX Direct chassis. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) The model of the endpoint, for which this capability information is applicable. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `revision`:(string) Revision for the fabric interconnect. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `vendor`:(string) The vendor of the endpoint, for which this capability information is applicable. 
* `nr_version`:(string) The firmware or software version of the endpoint, for which this capability information is applicable. 
 
