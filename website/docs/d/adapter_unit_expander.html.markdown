---
subcategory: "adapter"
layout: "intersight"
page_title: "Intersight: intersight_adapter_unit_expander"
description: |-
        The adapter.UnitExpander object represents a physical extension card for a network adapter within a server. It is used for inventory and to identify expander cards that augment the functionality or port count of a primary adapter.
        #### Purpose
        The UnitExpander object serves to inventory adapter extension cards, which are not full-featured adapters themselves but provide additional physical connectivity. It captures key hardware identifiers to ensure the component can be correctly identified and managed within the system.
        #### Key Concepts
        - **Hardware Identification:** Provides essential hardware details such as partNumber and vid (virtual ID) for unique identification.
        - **Hierarchical Relationship:** It is directly associated with a parent adapter.Unit object, clearly defining its role as an extension of a primary adapter.
        - **Simplified Model:** Focuses purely on the identity of the expander hardware, distinct from the more complex adapter.Unit which models the full adapter.

---

# Data Source: intersight_adapter_unit_expander
The adapter.UnitExpander object represents a physical extension card for a network adapter within a server. It is used for inventory and to identify expander cards that augment the functionality or port count of a primary adapter.
#### Purpose
The UnitExpander object serves to inventory adapter extension cards, which are not full-featured adapters themselves but provide additional physical connectivity. It captures key hardware identifiers to ensure the component can be correctly identified and managed within the system.
#### Key Concepts
- **Hardware Identification:** Provides essential hardware details such as partNumber and vid (virtual ID) for unique identification.
- **Hierarchical Relationship:** It is directly associated with a parent adapter.Unit object, clearly defining its role as an extension of a primary adapter.
- **Simplified Model:** Focuses purely on the identity of the expander hardware, distinct from the more complex adapter.Unit which models the full adapter.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_adapter_unit_expander.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `is_upgraded`:(bool) This field indicates the compute status of the catalog values for the associated component or hardware. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) This field displays the model number of the associated component or hardware. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `part_number`:(string) This field identifies the partNumber of the given component. 
* `presence`:(string) This field indicates the presence (equipped) or absence (absent) of the associated component or hardware. 
* `revision`:(string) This field displays the revised version of the associated component or hardware (if any). 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field displays the serial number of the associated component or hardware. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `vendor`:(string) This field displays the vendor information of the associated component or hardware. 
* `vid`:(string) This field identifies the virtual id of the given component. 
 
