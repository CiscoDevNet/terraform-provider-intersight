---
subcategory: "equipment"
layout: "intersight"
page_title: "Intersight: intersight_equipment_psu_control"
description: |-
  This represents the power states of an equipment.
---

# Data Source: intersight_equipment_psu_control
This represents the power states of an equipment.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_equipment_psu_control.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `cluster_state`:(string) This field identifies the cluster state of the psu redundancy. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `input_power_state`:(string) This field identifies the input power state of the psus. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) This field identifies the model of the given component. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) This field identifies the name of psu control object. 
* `oper_qualifier`:(string) This field identifies the operational qualifier for the psu redundancy. 
* `oper_state`:(string) This field identifies the operational state of the psu redundancy. 
* `output_power_state`:(string) This field identifies the output power state of the psus. 
* `presence`:(string) This field identifies the presence (equipped) or absence of the given component. 
* `redundancy`:(string) This field identifies the redundancy state of the psus. 
* `revision`:(string) This field identifies the revision of the given component. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field identifies the serial of the given component. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `vendor`:(string) This field identifies the vendor of the given component. 
 