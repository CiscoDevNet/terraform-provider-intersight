---
subcategory: "workflow"
layout: "intersight"
page_title: "Intersight: intersight_workflow_catalog_item_definition"
description: |-
        Catalog Item definition is a collection of Service items which are associated with workflow definition that can be used to deploy and manage service items.

---

# Data Source: intersight_workflow_catalog_item_definition
Catalog Item definition is a collection of Service items which are associated with workflow definition that can be used to deploy and manage service items.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_workflow_catalog_item_definition.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) The description for the catalog item which provides information on what are the pre-requisites to deploy the resource. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `label`:(string) A user friendly short name to identify the catalog item. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ) or an underscore (_). 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name for this catalog item definition. You can have multiple versions of the catalog item with the same name. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:) or an underscore (_). 
* `publish_status`:(string) Publish status of the catalog item.* `Draft` - The enum specifies the option as Draft which means the meta definition is being designed and tested.* `Published` - The enum specifies the option as Published which means the meta definition is ready for consumption.* `Archived` - The enum specifies the option as Archived which means the meta definition is archived and can no longer be consumed. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `support_status`:(string) The CatalogItem Support depicts the support status of catalog, the values will be any of Supported or Deprecated state. The user can create a Catalog Service Request if the support status is supported, if its Deprecated then it cannot be instantiated.* `Supported` - The definition is a supported version and there will be no changes to the mandatory inputs or outputs.* `Beta` - The definition is a Beta version and this version can under go changes until the version is marked supported.* `Deprecated` - The version of definition is deprecated and typically there will be a higher version of the same definition that has been added. 
* `user_id_or_email`:(string) The user identifier who created or updated the catalog item definition. 
* `nr_version`:(int) The version of the catalog item to support multiple versions. 
 
