---
subcategory: "firmware"
layout: "intersight"
page_title: "Intersight: intersight_firmware_switch_upgrade"
description: |-
        The SwitchUpgrade object is designed for upgrading firmware on Fabric Interconnects, streamlining the process of downloading and applying firmware images from cloud repositories or network shares.
        #### Purpose
        SwitchUpgrade simplifies the firmware upgrade process for Fabric Interconnects, enhancing network performance and security by facilitating timely updates. This provides options for both direct download and network share-based upgrades for UCSM devices. IMM Fabric Interconnect supports direct download.
        #### Key Concepts
        - **Fabric Interconnect Optimization:** Tailored to address the specific needs of Fabric Interconnects, ensuring efficient and effective firmware upgrades.
        - **Evacuation Management:** Includes features to manage fabric evacuation during upgrades, enhancing operational continuity and minimizing disruptions.
        - **Flexible Deployment:** Supports both direct download from Cisco sources and network share-based upgrades, catering to varied infrastructure configurations.
        - **Security and Compliance:** Implements rigorous access and licensing controls, ensuring upgrades are performed securely and in compliance with organizational policies.

---

# Data Source: intersight_firmware_switch_upgrade
The SwitchUpgrade object is designed for upgrading firmware on Fabric Interconnects, streamlining the process of downloading and applying firmware images from cloud repositories or network shares.
#### Purpose
SwitchUpgrade simplifies the firmware upgrade process for Fabric Interconnects, enhancing network performance and security by facilitating timely updates. This provides options for both direct download and network share-based upgrades for UCSM devices. IMM Fabric Interconnect supports direct download.
#### Key Concepts
- **Fabric Interconnect Optimization:** Tailored to address the specific needs of Fabric Interconnects, ensuring efficient and effective firmware upgrades.
- **Evacuation Management:** Includes features to manage fabric evacuation during upgrades, enhancing operational continuity and minimizing disruptions.
- **Flexible Deployment:** Supports both direct download from Cisco sources and network share-based upgrades, catering to varied infrastructure configurations.
- **Security and Compliance:** Implements rigorous access and licensing controls, ensuring upgrades are performed securely and in compliance with organizational policies.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_firmware_switch_upgrade.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `enable_fabric_evacuation`:(bool) The flag to enable or disable fabric evacuation during the switch firmware upgrade. In case of IMM, it is mandatory to have the Fabric Interconnects associated with domain profile for fabric evacuation to happen. 
* `enable_pdb_fpga_upgrade`:(bool) The flag to enable or disable firmware upgrade functionality for the PDB FPGA. 
* `enable_psu_upgrade`:(bool) The flag to enable or disable firmware upgrade functionality for the Power Supply Unit (PSU). 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `skip_estimate_impact`:(bool) User has the option to skip the estimate impact calculation. 
* `skip_wait_for_io_path_connectivity`:(bool) The flag to enable or disable the option to wait for IO paths connectivity during the switch firmware upgrade. 
* `status`:(string) Status of the upgrade operation.* `NONE` - Upgrade status is not populated.* `IN_PROGRESS` - The upgrade is in progress.* `SUCCESSFUL` - The upgrade successfully completed.* `FAILED` - The upgrade shows failed status.* `TERMINATED` - The upgrade has been terminated. 
* `switch_name`:(string) Name of the Fabric Interconnect on which the firmware upgrade operation is performed. 
* `upgrade_type`:(string) Desired upgrade mode to choose either direct download based upgrade or network share upgrade.* `direct_upgrade` - Upgrade mode is direct download.* `network_upgrade` - Upgrade mode is network upgrade. 
* `wait_time_out`:(int) Specifies a timeout period, in minutes, before the firmware upgrade begins. The valid range is -1 to 1000. A value of -1 requires manual user acknowledgment to proceed, 0 starts the upgrade immediately, and values from 1 to 1000 wait the specified number of minutes before starting. The upgrade will automatically begin once the timeout expires, but it can also be initiated manually at any time before the timeout ends. If no value is specified, manual user acknowledgment is required, equivalent to -1. 
 
