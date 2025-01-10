---
subcategory: "workflow"
layout: "intersight"
page_title: "Intersight: intersight_workflow_service_item_definition"
description: |-
        Service Item definition is a collection of actions and associated workflow definition that can be used to deploy a service item.

---

# Data Source: intersight_workflow_service_item_definition
Service Item definition is a collection of actions and associated workflow definition that can be used to deploy a service item.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_workflow_service_item_definition.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `allow_multiple_service_item_instances`:(bool) Service item definition can declare that only one instance can be allowed within the customer account. 
* `create_time`:(string) The time when this managed object was created. 
* `create_user`:(string) The user identifier who created or cloned the service item definition. 
* `cvd_id`:(string) The Cisco Validated Design (CVD) Identifier that this service item provides. 
* `delete_instance_on_decommission`:(bool) The flag to indicate that service item instance will be deleted after the completion of decommission action. 
* `description`:(string) The description for this service item which provides information on what are the pre-requisites to deploy the service item and what features are supported on the service item. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `label`:(string) A user friendly short name to identify the service item. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ) or an underscore (_). 
* `license_entitlement`:(string) License entitlement required to run this service item.* `Base` - Base as a License type. It is default license type.* `Essential` - Essential as a License type.* `Standard` - Standard as a License type.* `Advantage` - Advantage as a License type.* `Premier` - Premier as a License type.* `IWO-Essential` - IWO-Essential as a License type.* `IWO-Advantage` - IWO-Advantage as a License type.* `IWO-Premier` - IWO-Premier as a License type.* `IKS-Advantage` - IKS-Advantage as a License type.* `INC-Premier-1GFixed` - Premier 1G Fixed license tier for Intersight Nexus Cloud.* `INC-Premier-10GFixed` - Premier 10G Fixed license tier for Intersight Nexus Cloud.* `INC-Premier-100GFixed` - Premier 100G Fixed license tier for Intersight Nexus Cloud.* `INC-Premier-Mod4Slot` - Premier Modular 4 slot license tier for Intersight Nexus Cloud.* `INC-Premier-Mod8Slot` - Premier Modular 8 slot license tier for Intersight Nexus Cloud.* `INC-Premier-D2OpsFixed` - Premier D2Ops fixed license tier for Intersight Nexus Cloud.* `INC-Premier-D2OpsMod` - Premier D2Ops modular license tier for Intersight Nexus Cloud.* `INC-Premier-CentralizedMod8Slot` - Premier modular license tier of switch type CentralizedMod8Slot for Intersight Nexus Cloud.* `INC-Premier-DistributedMod8Slot` - Premier modular license tier of switch type DistributedMod8Slot for Intersight Nexus Cloud.* `ERP-Advantage` - Advantage license tier for ERP workflows.* `IntersightTrial` - Virtual dummy license type to indicate trial. Used for UI display of trial mode Intersight tiers.* `IWOTrial` - Virtual dummy license type to indicate trial. Used for UI display of trial mode IKS tiers.* `IKSTrial` - Virtual dummy license type to indicate trial. Used for UI display of trial mode IWO tiers.* `INCTrial` - Virtual dummy license type to indicate trial. Used for UI display of trial mode Nexus tiers. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `mod_user`:(string) The user identifier who last updated the service item definition. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name for this service item definition. You can have multiple versions of the service item with the same name. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:) or an underscore (_). 
* `publish_status`:(string) The publish status of service item (Draft, Published, Archived).* `Draft` - The enum specifies the option as Draft which means the meta definition is being designed and tested.* `Published` - The enum specifies the option as Published which means the meta definition is ready for consumption.* `Archived` - The enum specifies the option as Archived which means the meta definition is archived and can no longer be consumed. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `status`:(string) State of service item considering the state of underlying service item actions definitions.* `Okay` - Deployment and other post-deployment actions are in valid state.* `Critical` - Deployment action is not in valid state.* `Warning` - Deployment action is in valid state, and one or more post-deployment actions are not in valid state. 
* `support_status`:(string) The service item can be marked as deprecated, supported or beta, the support status indicates that. When a new service item is introduced, it can be marked beta to indicate this is experimental and later moved to Supported status. When Service item is deprecated, it cannot be instantiated and used for a Catalog Item design.* `Supported` - The definition is a supported version and there will be no changes to the mandatory inputs or outputs.* `Beta` - The definition is a Beta version and this version can under go changes until the version is marked supported.* `Deprecated` - The version of definition is deprecated and typically there will be a higher version of the same definition that has been added. 
* `user_id_or_email`:(string) This attribute is deprecated and replaced by createUser and modUser. 
* `nr_version`:(int) The version of the service item to support multiple versions. 
 
