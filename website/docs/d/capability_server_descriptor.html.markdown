---
subcategory: "capability"
layout: "intersight"
page_title: "Intersight: intersight_capability_server_descriptor"
description: |-
        The capability.ServerDescriptor object is a hardware descriptor that uniquely identifies an Intersight Managed Mode (IMM) server model and defines its specific capabilities and characteristics.
        #### Purpose
        The ServerDescriptor object serves as a central point in the capability catalog for defining model-specific server attributes. It allows the system to understand the unique features, limitations, and default configurations of a particular server model, such as UCSB-B200-M5. It is crucial for applying correct policies, enabling or disabling features, and ensuring proper management.
        #### Key Concepts
        - **Model-Specific Capabilities:** Defines attributes that are specific to a server model, such as whether the NCSI side-band interface is enabled (isNcsiEnabled) or the default PCIe slot for an MLOM adapter (mlomAdapterPcieSlotNumber).
        - **Policy Applicability:** The unsupportedPolicies property lists policies that are not applicable to a specific server model, preventing misconfiguration.
        - **Form Factor Identification:** Specifies the server's physical form factor (serverFormFactor), such as blade or rack which influences management logic.
        - **Unique Identification:** Uniquely identifies a server capability profile by its vendor and model.

---

# Data Source: intersight_capability_server_descriptor
The capability.ServerDescriptor object is a hardware descriptor that uniquely identifies an Intersight Managed Mode (IMM) server model and defines its specific capabilities and characteristics.
#### Purpose
The ServerDescriptor object serves as a central point in the capability catalog for defining model-specific server attributes. It allows the system to understand the unique features, limitations, and default configurations of a particular server model, such as UCSB-B200-M5. It is crucial for applying correct policies, enabling or disabling features, and ensuring proper management.
#### Key Concepts
- **Model-Specific Capabilities:** Defines attributes that are specific to a server model, such as whether the NCSI side-band interface is enabled (isNcsiEnabled) or the default PCIe slot for an MLOM adapter (mlomAdapterPcieSlotNumber).
- **Policy Applicability:** The unsupportedPolicies property lists policies that are not applicable to a specific server model, preventing misconfiguration.
- **Form Factor Identification:** Specifies the server's physical form factor (serverFormFactor), such as "blade" or "rack" which influences management logic.
- **Unique Identification:** Uniquely identifies a server capability profile by its vendor and model.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_capability_server_descriptor.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Detailed information about the endpoint. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `is_ncsi_enabled`:(bool) Indicates whether the CIMC to VIC side-band interface is enabled on the server. 
* `is_ppl_enabled`:(bool) Indicates Processor Package Power Limit for the server. 
* `mlom_adapter_pcie_slot_number`:(int) Indicates PCIe Slot numerical value for each Server model MLOM slot. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) The model of the endpoint, for which this capability information is applicable. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `server_form_factor`:(string) The form factor (blade/rack/etc) of the server.* `unknown` - The form factor of the server is unknown.* `blade` - Blade server form factor.* `rack` - Rack unit server form factor. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `vendor`:(string) The vendor of the endpoint, for which this capability information is applicable. 
* `nr_version`:(string) The firmware or software version of the endpoint, for which this capability information is applicable. 
 
