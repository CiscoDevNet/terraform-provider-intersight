---
subcategory: "capability"
layout: "intersight"
page_title: "Intersight: intersight_capability_fex_capability_def"
description: |-
        The FexCapabilityDef object describes capability flags for a Fabric Extender (FEX) platform in the capability catalog.
        #### Purpose
        This captures feature support (for example, whether certain port-level configurations are supported) so the system can validate and tailor configuration workflows for the specific FEX platform.
        #### Key Concepts
        - **Feature discovery:** Declares what a FEX platform supports (capability flags).
        - **Validation enablement:** Drives platform-aware checks before applying configuration.
        - **Platform-specific behavior:** Helps avoid applying unsupported settings on certain FEX models.
        - **Catalog governance:** Provides a single source of truth for FEX capability differences.

---

# Data Source: intersight_capability_fex_capability_def
The FexCapabilityDef object describes capability flags for a Fabric Extender (FEX) platform in the capability catalog.
#### Purpose
This captures feature support (for example, whether certain port-level configurations are supported) so the system can validate and tailor configuration workflows for the specific FEX platform.
#### Key Concepts
- **Feature discovery:** Declares what a FEX platform supports (capability flags).
- **Validation enablement:** Drives platform-aware checks before applying configuration.
- **Platform-specific behavior:** Helps avoid applying unsupported settings on certain FEX models.
- **Catalog governance:** Provides a single source of truth for FEX capability differences.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_capability_fex_capability_def.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `fec_config_on_hif_port_supported`:(bool) FEC config on HIF port for Fabric Extender. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) An unique identifer for a capability descriptor. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
