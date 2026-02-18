---
subcategory: "capability"
layout: "intersight"
page_title: "Intersight: intersight_capability_cimc_firmware_descriptor"
description: |-
        The capability.CimcFirmwareDescriptor is a hardware descriptor object that defines server capabilities based on its Cisco Integrated Management Controller (CIMC) firmware version. It is used to determine feature support for specific firmware releases.
        #### Purpose
        The main purpose of this object is to create a catalog that maps CIMC firmware versions to specific features. This allows the management system to dynamically determine if a server's firmware supports functionalities like UUID-based identification or local user password management, ensuring that operations are only attempted on compatible firmware.
        #### Key Concepts
        - **Firmware-Based Capability Mapping:** Links server capabilities directly to the vendor, model, and version of its CIMC firmware.
        - **Feature Support Gates:** Provides minimum firmware versions (uuidSupportedVer, localUserPswdSupportedVer) required for certain features to be available.
        - **Integration Logic:** Informs the system about specific integration details, such as whether the server uses an endpoint proxy (adapterEpProxyEnabled) to communicate with its adapters.
        - **Extensible Catalog:** As a HardwareDescriptor, it is part of an extensible catalog that can be updated to reflect new firmware releases and their corresponding capabilities.

---

# Data Source: intersight_capability_cimc_firmware_descriptor
The capability.CimcFirmwareDescriptor is a hardware descriptor object that defines server capabilities based on its Cisco Integrated Management Controller (CIMC) firmware version. It is used to determine feature support for specific firmware releases.
#### Purpose
The main purpose of this object is to create a catalog that maps CIMC firmware versions to specific features. This allows the management system to dynamically determine if a server's firmware supports functionalities like UUID-based identification or local user password management, ensuring that operations are only attempted on compatible firmware.
#### Key Concepts
- **Firmware-Based Capability Mapping:** Links server capabilities directly to the vendor, model, and version of its CIMC firmware.
- **Feature Support Gates:** Provides minimum firmware versions (uuidSupportedVer, localUserPswdSupportedVer) required for certain features to be available.
- **Integration Logic:** Informs the system about specific integration details, such as whether the server uses an endpoint proxy (adapterEpProxyEnabled) to communicate with its adapters.
- **Extensible Catalog:** As a HardwareDescriptor, it is part of an extensible catalog that can be updated to reflect new firmware releases and their corresponding capabilities.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_capability_cimc_firmware_descriptor.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `adapter_ep_proxy_enabled`:(bool) Indicates whether the server uses ep proxy to communicate with the adapter. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Detailed information about the endpoint. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `local_user_pswd_supported_ver`:(string) Minimum server firmware version for local users password properties feature support. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) The model of the endpoint, for which this capability information is applicable. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `revision`:(string) Revision information for the server. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `uuid_supported_ver`:(string) Minimum server firmware version for UUID feature support. 
* `vendor`:(string) The vendor of the endpoint, for which this capability information is applicable. 
* `nr_version`:(string) The firmware or software version of the endpoint, for which this capability information is applicable. 
 
