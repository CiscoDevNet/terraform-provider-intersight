---
subcategory: "equipment"
layout: "intersight"
page_title: "Intersight: intersight_equipment_tpm"
description: |-
        The equipment.Tpm object represents a Trusted Platform Module (TPM) security chip installed on a server's motherboard. A TPM is a hardware-based security component that provides secure storage for cryptographic keys, platform integrity measurements, and authentication.
        #### Purpose
        The Tpm object is used to inventory and report the status of the TPM on a server. It captures key information such as the TPM version, its activationStatus, and its administrative state (adminState). It is essential for security auditing and for enabling features that rely on hardware-based trust, such as secure boot and measured boot.
        #### Key Concepts
        - **Hardware Security Inventory:** Provides a detailed record of the server's TPM chip.
        - **Security State Reporting:** The activationStatus and ownership properties indicate the current security state of the TPM.
        - **Firmware and Versioning:** Captures the firmwareVersion and TPM specification version (e.g., 2.0).
        - **Server Component:** It is a child of a compute.Board, linking the TPM directly to the motherboard it is part of.

---

# Data Source: intersight_equipment_tpm
The equipment.Tpm object represents a Trusted Platform Module (TPM) security chip installed on a server's motherboard. A TPM is a hardware-based security component that provides secure storage for cryptographic keys, platform integrity measurements, and authentication.
#### Purpose
The Tpm object is used to inventory and report the status of the TPM on a server. It captures key information such as the TPM version, its activationStatus, and its administrative state (adminState). It is essential for security auditing and for enabling features that rely on hardware-based trust, such as secure boot and measured boot.
#### Key Concepts
- **Hardware Security Inventory:** Provides a detailed record of the server's TPM chip.
- **Security State Reporting:** The activationStatus and ownership properties indicate the current security state of the TPM.
- **Firmware and Versioning:** Captures the firmwareVersion and TPM specification version (e.g., "2.0").
- **Server Component:** It is a child of a compute.Board, linking the TPM directly to the motherboard it is part of.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_equipment_tpm.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `activation_status`:(string) Identifies the activation status of the TPM. 
* `admin_state`:(string) Identifies the admin configured state of the TPM. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `firmware_version`:(string) Firmware Version of the Trusted Platform Module. 
* `is_upgraded`:(bool) This field indicates the compute status of the catalog values for the associated component or hardware. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) This field displays the model number of the associated component or hardware. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `ownership`:(string) Identifies the ownership information of the TPM. 
* `presence`:(string) This field indicates the presence (equipped) or absence (absent) of the associated component or hardware. 
* `revision`:(string) This field displays the revised version of the associated component or hardware (if any). 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field displays the serial number of the associated component or hardware. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `tpm_id`:(int) Enter  the ID of the trusted platform module. 
* `vendor`:(string) This field displays the vendor information of the associated component or hardware. 
* `nr_version`:(string) Identifies the version of the Trusted Platform Module. 
 
