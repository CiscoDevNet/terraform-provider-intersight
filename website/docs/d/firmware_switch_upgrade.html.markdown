---
subcategory: "firmware"
layout: "intersight"
page_title: "Intersight: intersight_firmware_switch_upgrade"
description: |-
  Firmware upgrade operation for Fabric Interconnects that downloads the image located at Cisco/appliance/user provided HTTP repository or uses the image from a network share and upgrade. Direct download is used for upgrade that uses the image from a Cisco repository or an appliance repository. Network share is used for upgrade that use the image from a network share from your data center.
---

# Data Source: intersight_firmware_switch_upgrade
Firmware upgrade operation for Fabric Interconnects that downloads the image located at Cisco/appliance/user provided HTTP repository or uses the image from a network share and upgrade. Direct download is used for upgrade that uses the image from a Cisco repository or an appliance repository. Network share is used for upgrade that use the image from a network share from your data center.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_firmware_switch_upgrade.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `enable_fabric_evacuation`:(bool) The flag to enable or disable fabric evacuation during the switch firmware upgrade. In case of IMM, it is mandatory to have the Fabric Interconnects associated with domain profile for fabric evacuation to happen. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `skip_estimate_impact`:(bool) User has the option to skip the estimate impact calculation. 
* `status`:(string) Status of the upgrade operation.* `NONE` - Upgrade status is not populated.* `IN_PROGRESS` - The upgrade is in progress.* `SUCCESSFUL` - The upgrade successfully completed.* `FAILED` - The upgrade shows failed status.* `TERMINATED` - The upgrade has been terminated. 
* `upgrade_type`:(string) Desired upgrade mode to choose either direct download based upgrade or network share upgrade.* `direct_upgrade` - Upgrade mode is direct download.* `network_upgrade` - Upgrade mode is network upgrade. 
 