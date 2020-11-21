
---
layout: "intersight"
page_title: "Intersight: intersight_firmware_chassis_upgrade"
sidebar_current: "docs-intersight-data-source-firmware-chassis-upgrade"
description: |-
Firmware upgrade operation for chassis that downloads the image located at Cisco/appliance/user provided HTTP repository or uses the image from a network share and upgrade. Direct download is used for upgrade that uses the image from a Cisco repository or an appliance repository. Network share is used for upgrade that use the image from a network share from your data center.
---

# Data Source: intersight_firmware._chassis_upgrade
Firmware upgrade operation for chassis that downloads the image located at Cisco/appliance/user provided HTTP repository or uses the image from a network share and upgrade. Direct download is used for upgrade that uses the image from a Cisco repository or an appliance repository. Network share is used for upgrade that use the image from a network share from your data center.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `skip_estimate_impact`:(bool) User has the option to skip the estimate impact calculation. 
* `status`:(string) Status of the upgrade operation.* `NONE` - Upgrade status is not populated.* `IN_PROGRESS` - The upgrade is in progress.* `SUCCESSFUL` - The upgrade successfully completed.* `FAILED` - The upgrade shows failed status.* `TERMINATED` - The upgrade has been terminated. 
* `upgrade_type`:(string) Desired upgrade mode to choose either direct download based upgrade or network share upgrade.* `direct_upgrade` - Upgrade mode is direct download.* `network_upgrade` - Upgrade mode is network upgrade. 
