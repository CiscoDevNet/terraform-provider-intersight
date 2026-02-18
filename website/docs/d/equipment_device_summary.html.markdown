---
subcategory: "equipment"
layout: "intersight"
page_title: "Intersight: intersight_equipment_device_summary"
description: |-
        The equipment.DeviceSummary object provides a simplified, aggregated view of an inventory component, such as a server or chassis. It is designed to be a lightweight object for use in summary lists and high-level dashboards.
        #### Purpose
        The main purpose of DeviceSummary is to offer a consolidated and easy-to-query representation of a piece of equipment. It abstracts away the complexity of different device types (e.g., compute.Blade vs. compute.RackUnit) and presents a consistent set of key identifiers like model, serial, and dn. This simplifies the development of user interfaces that display lists of mixed hardware.
        #### Key Concepts
        - **Aggregated View:** Combines the most common and essential properties of an inventory object into a single, flat structure.
        - **Unified Model:** Provides a consistent representation for different types of underlying hardware, simplifying data consumption.
        - **Performance-Oriented:** Designed to be lightweight and efficiently queried, making it suitable for high-level summary views.
        - **Source Traceability:** The sourceObjectType and parent relationships allow for drilling down from the summary view to the full, detailed inventory object.

---

# Data Source: intersight_equipment_device_summary
The equipment.DeviceSummary object provides a simplified, aggregated view of an inventory component, such as a server or chassis. It is designed to be a lightweight object for use in summary lists and high-level dashboards.
#### Purpose
The main purpose of DeviceSummary is to offer a consolidated and easy-to-query representation of a piece of equipment. It abstracts away the complexity of different device types (e.g., compute.Blade vs. compute.RackUnit) and presents a consistent set of key identifiers like model, serial, and dn. This simplifies the development of user interfaces that display lists of mixed hardware.
#### Key Concepts
- **Aggregated View:** Combines the most common and essential properties of an inventory object into a single, flat structure.
- **Unified Model:** Provides a consistent representation for different types of underlying hardware, simplifying data consumption.
- **Performance-Oriented:** Designed to be lightweight and efficiently queried, making it suitable for high-level summary views.
- **Source Traceability:** The sourceObjectType and parent relationships allow for drilling down from the summary view to the full, detailed inventory object.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_equipment_device_summary.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `dn`:(string) The distinguished name that unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) This field identifies the model of the given component. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `serial`:(string) This field identifies the serial number of the given component. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `source_object_type`:(string) The source object type of the given component. 
 
