---
subcategory: "servicenow"
layout: "intersight"
page_title: "Intersight: intersight_servicenow_incident_doc"
description: |-
        The IncidentDoc object is part of the ServiceNow service, providing a schema for incident-related documentation. It serves as a foundational element for structuring and organizing incident data.
        #### Purpose
        IncidentDoc provides the structure necessary for recording and managing incident data, ensuring consistent documentation and integration across systems. It supports the standardization of incident information, facilitating efficient data handling.
        #### Key Concepts
        - **Schema Definition:** Establishes a clear schema for incident data, supporting standardized data entry and management.
        - **Integration:** Facilitates integration with other systems and processes, ensuring seamless interaction and data exchange.
        - **Consistency:** Promotes consistent documentation practices, enhancing data reliability and accessibility.
        - **Flexibility:** Allows customization to accommodate specific organizational needs and requirements.

---

# Data Source: intersight_servicenow_incident_doc
The IncidentDoc object is part of the ServiceNow service, providing a schema for incident-related documentation. It serves as a foundational element for structuring and organizing incident data.
#### Purpose
IncidentDoc provides the structure necessary for recording and managing incident data, ensuring consistent documentation and integration across systems. It supports the standardization of incident information, facilitating efficient data handling.
#### Key Concepts
- **Schema Definition:** Establishes a clear schema for incident data, supporting standardized data entry and management.
- **Integration:** Facilitates integration with other systems and processes, ensuring seamless interaction and data exchange.
- **Consistency:** Promotes consistent documentation practices, enhancing data reliability and accessibility.
- **Flexibility:** Allows customization to accommodate specific organizational needs and requirements.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_servicenow_incident_doc.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `example_value`:(string) Example value for Incident property. 
* `internal_type`:(string) Internal type for Incident property. 
* `label`:(string) Label for Incident property. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `property_name`:(string) Name for Incident property. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
