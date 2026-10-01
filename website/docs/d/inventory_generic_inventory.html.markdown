---
subcategory: "inventory"
layout: "intersight"
page_title: "Intersight: intersight_inventory_generic_inventory"
description: |-
        GenericInventories represent arbitrary inventory expressed as key/value pairs (for example, OS tools inventory from UCSM). They are used when the inventory schema is not modeled as dedicated typed objects.
        #### Purpose
        Provide a flexible mechanism to store and query unstructured or semi-structured inventory data as key/value entries.
        #### Key Concepts
        - **Key/value inventory:** Represents inventory items where the structure is best expressed as a key and value.
        - **Schema flexibility:** Allows capturing inventory not worth modeling as strongly typed managed objects.
        - **Holder association:** Typically grouped under a GenericInventoryHolder for an endpoint context.

---

# Data Source: intersight_inventory_generic_inventory
GenericInventories represent arbitrary inventory expressed as key/value pairs (for example, OS tools inventory from UCSM). They are used when the inventory schema is not modeled as dedicated typed objects.
#### Purpose
Provide a flexible mechanism to store and query unstructured or semi-structured inventory data as key/value entries.
#### Key Concepts
- **Key/value inventory:** Represents inventory items where the structure is best expressed as a key and value.
- **Schema flexibility:** Allows capturing inventory not worth modeling as strongly typed managed objects.
- **Holder association:** Typically grouped under a GenericInventoryHolder for an endpoint context.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_inventory_generic_inventory.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `key`:(string) Key of inventory data for Generic Inventory data set. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `type`:(string) Type of inventory data for Generic Inventory data set. 
* `value`:(string) Value of inventory data for Generic Inventory data set. 
 
