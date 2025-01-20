---
subcategory: "license"
layout: "intersight"
page_title: "Intersight: intersight_license_account_license_data"
description: |-
        License information for an account.

---

# Data Source: intersight_license_account_license_data
License information for an account.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_license_account_license_data.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_id`:(string) Root user's ID of the account. 
* `account_moid`:(string) The Account ID for this managed object. 
* `agent_data`:(string) Agent trusted store data. 
* `auth_expire_time`:(string) Authorization expiration time. 
* `auth_initial_time`:(string) Intial authorization time. 
* `auth_next_time`:(string) Next time for the authorization. 
* `category`:(string) Account license data category name. 
* `create_time`:(string) The time when this managed object was created. 
* `default_license_type`:(string) Default license tier set by the user.* `Base` - Base as a License type. It is default license type.* `Essential` - Essential as a License type.* `Standard` - Standard as a License type.* `Advantage` - Advantage as a License type.* `Premier` - Premier as a License type.* `IWO-Essential` - IWO-Essential as a License type.* `IWO-Advantage` - IWO-Advantage as a License type.* `IWO-Premier` - IWO-Premier as a License type.* `IKS-Advantage` - IKS-Advantage as a License type.* `INC-Premier-1GFixed` - Premier 1G Fixed license tier for Intersight Nexus Cloud.* `INC-Premier-10GFixed` - Premier 10G Fixed license tier for Intersight Nexus Cloud.* `INC-Premier-100GFixed` - Premier 100G Fixed license tier for Intersight Nexus Cloud.* `INC-Premier-Mod4Slot` - Premier Modular 4 slot license tier for Intersight Nexus Cloud.* `INC-Premier-Mod8Slot` - Premier Modular 8 slot license tier for Intersight Nexus Cloud.* `INC-Premier-D2OpsFixed` - Premier D2Ops fixed license tier for Intersight Nexus Cloud.* `INC-Premier-D2OpsMod` - Premier D2Ops modular license tier for Intersight Nexus Cloud.* `INC-Premier-CentralizedMod8Slot` - Premier modular license tier of switch type CentralizedMod8Slot for Intersight Nexus Cloud.* `INC-Premier-DistributedMod8Slot` - Premier modular license tier of switch type DistributedMod8Slot for Intersight Nexus Cloud.* `ERP-Advantage` - Advantage license tier for ERP workflows.* `IntersightTrial` - Virtual dummy license type to indicate trial. Used for UI display of trial mode Intersight tiers.* `IWOTrial` - Virtual dummy license type to indicate trial. Used for UI display of trial mode IKS tiers.* `IKSTrial` - Virtual dummy license type to indicate trial. Used for UI display of trial mode IWO tiers.* `INCTrial` - Virtual dummy license type to indicate trial. Used for UI display of trial mode Nexus tiers. 
* `default_license_type_newer_models`:(string) Default license tier for newer model M7+ servers set by the user.* `Base` - Base as a License type. It is default license type.* `Essential` - Essential as a License type.* `Standard` - Standard as a License type.* `Advantage` - Advantage as a License type.* `Premier` - Premier as a License type.* `IWO-Essential` - IWO-Essential as a License type.* `IWO-Advantage` - IWO-Advantage as a License type.* `IWO-Premier` - IWO-Premier as a License type.* `IKS-Advantage` - IKS-Advantage as a License type.* `INC-Premier-1GFixed` - Premier 1G Fixed license tier for Intersight Nexus Cloud.* `INC-Premier-10GFixed` - Premier 10G Fixed license tier for Intersight Nexus Cloud.* `INC-Premier-100GFixed` - Premier 100G Fixed license tier for Intersight Nexus Cloud.* `INC-Premier-Mod4Slot` - Premier Modular 4 slot license tier for Intersight Nexus Cloud.* `INC-Premier-Mod8Slot` - Premier Modular 8 slot license tier for Intersight Nexus Cloud.* `INC-Premier-D2OpsFixed` - Premier D2Ops fixed license tier for Intersight Nexus Cloud.* `INC-Premier-D2OpsMod` - Premier D2Ops modular license tier for Intersight Nexus Cloud.* `INC-Premier-CentralizedMod8Slot` - Premier modular license tier of switch type CentralizedMod8Slot for Intersight Nexus Cloud.* `INC-Premier-DistributedMod8Slot` - Premier modular license tier of switch type DistributedMod8Slot for Intersight Nexus Cloud.* `ERP-Advantage` - Advantage license tier for ERP workflows.* `IntersightTrial` - Virtual dummy license type to indicate trial. Used for UI display of trial mode Intersight tiers.* `IWOTrial` - Virtual dummy license type to indicate trial. Used for UI display of trial mode IKS tiers.* `IKSTrial` - Virtual dummy license type to indicate trial. Used for UI display of trial mode IWO tiers.* `INCTrial` - Virtual dummy license type to indicate trial. Used for UI display of trial mode Nexus tiers. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `error_desc`:(string) The detailed error message when there is any error related to license sync of this account. 
* `group`:(string) Account license data group name. 
* `highest_compliant_license_tier`:(string) The highest license tier which is in compliant of this account.* `Base` - Base as a License type. It is default license type.* `Essential` - Essential as a License type.* `Standard` - Standard as a License type.* `Advantage` - Advantage as a License type.* `Premier` - Premier as a License type.* `IWO-Essential` - IWO-Essential as a License type.* `IWO-Advantage` - IWO-Advantage as a License type.* `IWO-Premier` - IWO-Premier as a License type.* `IKS-Advantage` - IKS-Advantage as a License type.* `INC-Premier-1GFixed` - Premier 1G Fixed license tier for Intersight Nexus Cloud.* `INC-Premier-10GFixed` - Premier 10G Fixed license tier for Intersight Nexus Cloud.* `INC-Premier-100GFixed` - Premier 100G Fixed license tier for Intersight Nexus Cloud.* `INC-Premier-Mod4Slot` - Premier Modular 4 slot license tier for Intersight Nexus Cloud.* `INC-Premier-Mod8Slot` - Premier Modular 8 slot license tier for Intersight Nexus Cloud.* `INC-Premier-D2OpsFixed` - Premier D2Ops fixed license tier for Intersight Nexus Cloud.* `INC-Premier-D2OpsMod` - Premier D2Ops modular license tier for Intersight Nexus Cloud.* `INC-Premier-CentralizedMod8Slot` - Premier modular license tier of switch type CentralizedMod8Slot for Intersight Nexus Cloud.* `INC-Premier-DistributedMod8Slot` - Premier modular license tier of switch type DistributedMod8Slot for Intersight Nexus Cloud.* `ERP-Advantage` - Advantage license tier for ERP workflows.* `IntersightTrial` - Virtual dummy license type to indicate trial. Used for UI display of trial mode Intersight tiers.* `IWOTrial` - Virtual dummy license type to indicate trial. Used for UI display of trial mode IKS tiers.* `IKSTrial` - Virtual dummy license type to indicate trial. Used for UI display of trial mode IWO tiers.* `INCTrial` - Virtual dummy license type to indicate trial. Used for UI display of trial mode Nexus tiers. 
* `last_cssm_sync`:(string) Specifies last sync time with CSSM. 
* `last_renew`:(string) Specifies last certificate renew time with SA. 
* `last_sync`:(string) Specifies last sync time with SA. 
* `last_updated_time`:(string) Record's last update datetime. 
* `license_state`:(string) Aggregrated mode for the agent. 
* `license_tech_support_info`:(string) Tech-support info of a smart-agent. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `register_expire_time`:(string) Registration exipiration time. 
* `register_initial_time`:(string) Initial time of registration. 
* `register_next_time`:(string) Next time for the license registration. 
* `registration_status`:(string) Registration status of a smart-agent. 
* `renew_failure_string`:(string) License renewal failure message. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `smart_account`:(string) Name of the smart account. 
* `smart_account_domain`:(string) Domain Name of the smart account. 
* `smart_api_enabled`:(bool) Indicate whether API integration is enabled. 
* `smart_api_sync_status`:(string) The detailed error message when there is any smart API sync error related to this account. 
* `sync_status`:(string) Current sync status for the account. 
* `virtual_account`:(string) Name of the virtual account. 
 
