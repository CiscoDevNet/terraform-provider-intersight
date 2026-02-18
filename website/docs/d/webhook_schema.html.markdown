---
subcategory: "webhook"
layout: "intersight"
page_title: "Intersight: intersight_webhook_schema"
description: |-
        The Schema object is integral to event management, providing the framework that defines the structure of events. It ensures consistency and validity in the way events are described and handled across the system.
        #### Purpose
        A Schema outlines the format and content of events, acting as a template for event generation. It is essential for maintaining uniformity and accuracy in event data processing and interpretation.
        #### Key Concepts
        - **Event Definition:** Specifies the event type and source, ensuring each schema represents a unique event scenario within the system.
        - **Data Integrity:** Enforces constraints on event data, guaranteeing that events adhere to predefined standards and formats.
        - **Access Control:** Utilizes privilege sets to regulate schema creation, updates, and deletion, protecting the integrity of event definitions.

---

# Data Source: intersight_webhook_schema
The Schema object is integral to event management, providing the framework that defines the structure of events. It ensures consistency and validity in the way events are described and handled across the system.
#### Purpose
A Schema outlines the format and content of events, acting as a template for event generation. It is essential for maintaining uniformity and accuracy in event data processing and interpretation.
#### Key Concepts
- **Event Definition:** Specifies the event type and source, ensuring each schema represents a unique event scenario within the system.
- **Data Integrity:** Enforces constraints on event data, guaranteeing that events adhere to predefined standards and formats.
- **Access Control:** Utilizes privilege sets to regulate schema creation, updates, and deletion, protecting the integrity of event definitions.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_webhook_schema.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `data`:(string) The actual JSON schema data. The data must be a valid JSON schema string. 
* `description`:(string) The description of the schema. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `event_type`:(string) The event type of the schema. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `nr_source`:(string) The source of the schema. 
 
