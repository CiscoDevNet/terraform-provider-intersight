---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_device_package_download_state"
description: |-
        The DevicePackageDownloadState object tracks the state of package downloads on HyperFlex devices. This provides detailed information about package installations, ensuring that devices have the necessary updates for optimal operation and health monitoring.
        #### Purpose
        DevicePackageDownloadState objects document the download and installation status of packages, verifying that devices are equipped with the latest tools for executing health checks. It is crucial for maintaining device readiness and system integrity.
        #### Key Concepts
        - **Installation Monitoring:** Tracks the progress and completion of package installations, ensuring that devices are updated with necessary software components.
        - **Checksum Validation:** Records package checksums to verify download integrity, protecting against corrupted or incomplete installations.
        - **Node-Level Detail:** Includes information on nodes where packages are installed, supporting targeted updates and maintenance activities.
        - **Device Integration:** Associated with specific HyperFlex devices, enabling device-centric package management and ensuring alignment with system requirements.

---

# Data Source: intersight_hyperflex_device_package_download_state
The DevicePackageDownloadState object tracks the state of package downloads on HyperFlex devices. This provides detailed information about package installations, ensuring that devices have the necessary updates for optimal operation and health monitoring.
#### Purpose
DevicePackageDownloadState objects document the download and installation status of packages, verifying that devices are equipped with the latest tools for executing health checks. It is crucial for maintaining device readiness and system integrity.
#### Key Concepts
- **Installation Monitoring:** Tracks the progress and completion of package installations, ensuring that devices are updated with necessary software components.
- **Checksum Validation:** Records package checksums to verify download integrity, protecting against corrupted or incomplete installations.
- **Node-Level Detail:** Includes information on nodes where packages are installed, supporting targeted updates and maintenance activities.
- **Device Integration:** Associated with specific HyperFlex devices, enabling device-centric package management and ensuring alignment with system requirements.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_device_package_download_state.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `checksum`:(string) Checksum of HyperFlex health check Debian package installed on the HyperFlex Device. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `hx_device_name`:(string) HyperFlex Device Name for which the package download state is tracked. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `timestamp`:(string) Timestamp of the last health check Debian package installation on the HyperFlex Device. 
 
