---
subcategory: "workflow"
layout: "intersight"
page_title: "Intersight: intersight_workflow_service_item_attribute"
description: |-
        Service item attribute which represents all the artifacts created or related to this service item instance.

---

# Data Source: intersight_workflow_service_item_attribute
Service item attribute which represents all the artifacts created or related to this service item instance.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_workflow_service_item_attribute.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `datatype`:(string) Datatype, if any, backing the service item attribute definition. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Attribute name which is used in the attribute definition of the service item. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `type`:(string) Type of the service item attribute.* `None` - Default value if the service item attribute does not belong to any of the existing types.* `Configuration` - The service item attribute is a configuration from the designer or the end user.* `Inventory` - The service item attribute captures the inventory of the resource created by the service item deployment.* `Health` - The service item attribute describes the health of the resource created by the service item deployment.* `Output` - The service item attribute captures the artifact generated after performing an action on the service item. 
 
