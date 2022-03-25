---
subcategory: "firmware"
layout: "intersight"
page_title: "Intersight: intersight_firmware_upgrade_status"
description: |-
        The status for the upgrade operation to include the status for the download and upgrade stages.

---

# Data Source: intersight_firmware_upgrade_status
The status for the upgrade operation to include the status for the download and upgrade stages.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_firmware_upgrade_status.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `download_error`:(string) Any error encountered. Set to empty when download is in progress or completed. 
* `download_message`:(string) The message from the endpoint during the download. 
* `download_percentage`:(int) The percentage of the image downloaded in the endpoint. 
* `download_progress`:(int) The download progress of the file represented as a percentage between 0% and 100%. If progress reporting is not possible a value of -1 is sent. 
* `download_retries`:(int) The number of retries the plugin attempted before succeeding or failing the download. 
* `download_stage`:(string) The image download stages. Example:downloading, flashing. 
* `ep_power_status`:(string) The server power status after the upgrade request is submitted in the endpoint.* `none` - Server power status is none.* `powered on` - Server power status is powered on.* `powered off` - Server power status is powered off. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `overall_error`:(string) The reason for the operation failure. 
* `overall_percentage`:(int) The overall percentage of the operation. 
* `overallstatus`:(string) The overall status of the operation.* `none` - Upgrade stage is no upgrade stage.* `started` - Upgrade stage is started.* `prepare initiating` - Upgrade configuration is being prepared.* `prepare initiated` - Upgrade configuration is initiated.* `prepared` - Upgrade configuration is prepared.* `download initiating` - Upgrade stage is download initiating.* `download initiated` - Upgrade stage is download initiated.* `downloading` - Upgrade stage is downloading.* `downloaded` - Upgrade stage is downloaded.* `upgrade initiating on fabric A` - Upgrade stage is in upgrade initiating when upgrade is being started in endopint.* `upgrade initiated on fabric A` - Upgrade stage is in upgrade initiated when the upgrade has started in endpoint.* `upgrading fabric A` - Upgrade stage is in upgrading when the upgrade requires reboot to complete.* `rebooting fabric A` - Upgrade is in rebooting when the endpoint is being rebooted.* `upgraded fabric A` - Upgrade stage is in upgraded when the corresponding endpoint has completed.* `upgrade initiating on fabric B` - Upgrade stage is in upgrade initiating when upgrade is being started in endopint.* `upgrade initiated on fabric B` - Upgrade stage is in upgrade initiated when upgrade has started in endpoint.* `upgrading fabric B` - Upgrade stage is in upgrading when the upgrade requires reboot to complete.* `rebooting fabric B` - Upgrade is in rebooting when the endpoint is being rebooted.* `upgraded fabric B` - Upgrade stage is in upgraded when the corresponding endpoint has completed.* `upgrade initiating` - Upgrade stage is upgrade initiating.* `upgrade initiated` - Upgrade stage is upgrade initiated.* `upgrading` - Upgrade stage is upgrading.* `oob images staging` - Out-of-band component images staging.* `oob images staged` - Out-of-band component images staged.* `rebooting` - Upgrade is rebooting the endpoint.* `upgraded` - Upgrade stage is upgraded.* `success` - Upgrade stage is success.* `failed` - Upgrade stage is upgrade failed.* `terminated` - Upgrade stage is terminated.* `pending` - Upgrade stage is pending.* `ReadyForCache` - The image is ready to be cached into the Intersight Appliance.* `Caching` - The image will be cached into Intersight Appliance or an endpoint cache.* `Cached` - The image has been cached into the Intersight Appliance or endpoint cache.* `CachingFailed` - The image caching into the Intersight Appliance failed or endpoint cache. 
* `pending_type`:(string) Pending reason for the upgrade waiting.* `none` - Upgrade pending reason is none.* `pending for next reboot` - Upgrade pending reason is pending for next reboot. 
* `sd_card_download_error`:(string) The error message from the endpoint during the SD card download. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
