---
subcategory: "workflow"
layout: "intersight"
page_title: "Intersight: intersight_workflow_catalog_service_request"
description: |-
        Catalog Service Request is one instance of a catalog item based on a catalog item definition.

---

# Data Source: intersight_workflow_catalog_service_request
Catalog Service Request is one instance of a catalog item based on a catalog item definition.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_workflow_catalog_service_request.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) The description for this catalog service request which provides information on request from the user. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `end_time`:(string) The time at which the service request completed or stopped. 
* `label`:(string) A user friendly short name to identify the catalog service request instance. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ) or an underscore (_). 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) A name of the catalog service request instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `status`:(string) Status of the catalog service request instance which determines the actions that are allowed on this instance.* `NotCreated` - The service item is not yet created and it is in a draft mode. A service item instance can be deleted in this state.* `InProgress` - An action is in progress and until that action has reached a final state, another action cannot be started.* `Failed` - The last action on the service item instance failed and corrective measures need to be taken to bring the service item instance back to valid state.* `Okay` - The last action on the service item instance completed and the service item instance is in Okay state.* `Decommissioned` - The service item is decommissioned and can be safely deleted. A service item instance in any other state after it has been created cannot be deleted until it has been decommissioned. 
* `user_id_or_email`:(string) The user identifier who invoked the request to create the catalog service request. 
 
