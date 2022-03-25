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
* `cvd_id`:(string) The Cisco Validated Design (CVD) Identifier that this service item provides. 
* `delete_instance_on_decommission`:(bool) The flag to indicate that service item instance will be deleted after the completion of decommission action. 
* `description`:(string) The description for this service item which provides information on what are the pre-requisites to deploy the service item and what features are supported on the service item. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `label`:(string) A user friendly short name to identify the service item. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ) or an underscore (_). 
* `license_entitlement`:(string) License entitlement required to run this service item.* `Base` - Base as a License type. It is default license type.* `Essential` - Essential as a License type.* `Standard` - Standard as a License type.* `Advantage` - Advantage as a License type.* `Premier` - Premier as a License type.* `IWO-Essential` - IWO-Essential as a License type.* `IWO-Advantage` - IWO-Advantage as a License type.* `IWO-Premier` - IWO-Premier as a License type.* `IKS-Advantage` - IKS-Advantage as a License type. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name for this service item definition. You can have multiple versions of the service item with the same name. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:) or an underscore (_). 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `nr_version`:(int) The version of the service item to support multiple versions. 
 
