---
subcategory: "partnerintegration"
layout: "intersight"
page_title: "Intersight: intersight_partnerintegration_model"
description: |-
        The Model object defines the endpoint model used for inventory collections.
        #### Purpose
        This provides the schema definition for the inventory endpoints, ensuring that the data structure is clearly defined and consistent for API consumption.
        #### Key Concepts
        - **Schema Definition:** Stores the endpoint model in YAML/JSON format.
        - **Endpoint Integration:** Defines the structure of the data exposed by the inventory service.

---

# Data Source: intersight_partnerintegration_model
The Model object defines the endpoint model used for inventory collections.
#### Purpose
This provides the schema definition for the inventory endpoints, ensuring that the data structure is clearly defined and consistent for API consumption.
#### Key Concepts
- **Schema Definition:** Stores the endpoint model in YAML/JSON format.
- **Endpoint Integration:** Defines the structure of the data exposed by the inventory service.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_partnerintegration_model.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Placeholder name for the model. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
