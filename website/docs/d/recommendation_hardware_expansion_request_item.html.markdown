---
subcategory: "recommendation"
layout: "intersight"
page_title: "Intersight: intersight_recommendation_hardware_expansion_request_item"
description: |-
        Entity representing the user request for expansion of each hardware item.

---

# Data Source: intersight_recommendation_hardware_expansion_request_item
Entity representing the user request for expansion of each hardware item.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_recommendation_hardware_expansion_request_item.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `item_type`:(string) Type of hardware item for which the capacity increase is requested by the user. For example, CPU.* `None` - The Enum value None represents that no value was set for the hardware type.* `CPU` - The Enum value CPU represents that the hardware type requested for expansion is a physical CPU.* `Memory` - The Enum value Memory represents that the hardware type requested for expansion is a memory unit.* `Storage` - The Enum value Storage represents that the hardware type requested for expansion is a storage unit, ie, storage drives. 
* `max_value`:(float) The maximum value allowed for expansion for the hardware type on the referred registered device. 
* `max_value_unit`:(string) Unit type for the maximum value of the resource. For example, TB, GB, MB.* `TB` - The Enum value TB represents that the measurement unit is in terabytes.* `MB` - The Enum value MB represents that the measurement unit is in megabytes.* `GB` - The Enum value GB represents that the measurement unit is in gigabytes.* `MHz` - The Enum value MHz represents that the measurement unit is in megahertz.* `GHz` - The Enum value GHz represents that the measurement unit is in gigahertz.* `Percentage` - The Enum value Percentage represents that the expansion request is in the percentage of resource increase. For example, a 20% increase in CPU capacity. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `unit_type`:(string) Unit type for the expansion request, i.e., if the increase is requested as a raw value in TB, GB, etc., or in percentage increase.* `TB` - The Enum value TB represents that the measurement unit is in terabytes.* `MB` - The Enum value MB represents that the measurement unit is in megabytes.* `GB` - The Enum value GB represents that the measurement unit is in gigabytes.* `MHz` - The Enum value MHz represents that the measurement unit is in megahertz.* `GHz` - The Enum value GHz represents that the measurement unit is in gigahertz.* `Percentage` - The Enum value Percentage represents that the expansion request is in the percentage of resource increase. For example, a 20% increase in CPU capacity. 
* `value`:(float) Value of the expansion request which can be absolute value or percentage increase. 
 
