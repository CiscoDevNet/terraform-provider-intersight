---
subcategory: "capability"
layout: "intersight"
page_title: "Intersight: intersight_capability_io_card_descriptor"
description: |-
        IoCardDescriptors are capability-catalog hardware descriptors that uniquely identify an IO card module and describe key connectivity and behavior characteristics needed for platform compatibility and topology modeling. Beyond basic vendor/model/version/revision identity, they capture host-port counts, bifurcation details, UIF connectivity, native-speed behavior, and policy applicability.
        #### Purpose
        Provide a canonical, capability-aware catalog entry for IO card modules so the platform can determine how a given IO card should be modeled, connected, and constrained in supported configurations.
        #### Key Concepts
        - **Extended hardware identity**: Identity includes `vendor`, `model`, `version`, `revision`, `numHifPorts`, `uifConnectivity`, and `section`, reflecting that connectivity characteristics are part of what makes an IO card variant unique.
        - **Host interface topology**: `numHifPorts` describes host-interface port count per blade, which is important for connectivity planning and validation.
        - **Port mapping/bifurcation behavior**: `bifPortNum` and `uifConnectivity` model how uplink/UIF ports relate to IOM ports and how the card’s connectivity is structured.
        - **Native-speed behavior**: `nativeSpeedMasterPortNum` and `nativeHifPortChannelRequired` describe how native-speed configurations should be anchored and whether host port-channeling is required.
        - **Platform-specific classification**: `isUcsxDirectIoCard` indicates whether the IO card belongs to a UCS X Direct chassis context.
        - **Policy applicability constraints**: `unsupportedPolicies` lists policies that do not apply to the IO card, enabling validation and UI/workflow guardrails.
        - **Section-scoped catalog control**: Inherits permissions from `section`, with CRUD reserved for `CapabilityCatalog Administrator`.

---

# Data Source: intersight_capability_io_card_descriptor
IoCardDescriptors are capability-catalog hardware descriptors that uniquely identify an IO card module and describe key connectivity and behavior characteristics needed for platform compatibility and topology modeling. Beyond basic vendor/model/version/revision identity, they capture host-port counts, bifurcation details, UIF connectivity, native-speed behavior, and policy applicability.
#### Purpose
Provide a canonical, capability-aware catalog entry for IO card modules so the platform can determine how a given IO card should be modeled, connected, and constrained in supported configurations.
#### Key Concepts
- **Extended hardware identity**: Identity includes `vendor`, `model`, `version`, `revision`, `numHifPorts`, `uifConnectivity`, and `section`, reflecting that connectivity characteristics are part of what makes an IO card variant unique.
- **Host interface topology**: `numHifPorts` describes host-interface port count per blade, which is important for connectivity planning and validation.
- **Port mapping/bifurcation behavior**: `bifPortNum` and `uifConnectivity` model how uplink/UIF ports relate to IOM ports and how the card’s connectivity is structured.
- **Native-speed behavior**: `nativeSpeedMasterPortNum` and `nativeHifPortChannelRequired` describe how native-speed configurations should be anchored and whether host port-channeling is required.
- **Platform-specific classification**: `isUcsxDirectIoCard` indicates whether the IO card belongs to a UCS X Direct chassis context.
- **Policy applicability constraints**: `unsupportedPolicies` lists policies that do not apply to the IO card, enabling validation and UI/workflow guardrails.
- **Section-scoped catalog control**: Inherits permissions from `section`, with CRUD reserved for `CapabilityCatalog Administrator`.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_capability_io_card_descriptor.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `bif_port_num`:(int) Identifies the bif port number for the iocard module. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Detailed information about the endpoint. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `is_ucsx_direct_io_card`:(bool) Identifies whether the iocard module is a part of the UCSX Direct chassis. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) The model of the endpoint, for which this capability information is applicable. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `native_hif_port_channel_required`:(bool) Identifies whether host port-channel is required to be configured for the iocard module. 
* `native_speed_master_port_num`:(int) Primary port number for native speed configuration for the iocard module. 
* `num_hif_ports`:(int) Number of hif ports per blade for the iocard module. 
* `revision`:(string) Revision for the iocard module. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `uif_connectivity`:(string) Connectivity information between UIF Uplink ports and IOM ports.* `inline` - UIF uplink ports and IOM ports are connected inline.* `cross-connected` - UIF uplink ports and IOM ports are cross-connected, a case in washington chassis. 
* `vendor`:(string) The vendor of the endpoint, for which this capability information is applicable. 
* `nr_version`:(string) The firmware or software version of the endpoint, for which this capability information is applicable. 
 
