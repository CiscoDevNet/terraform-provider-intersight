---
subcategory: "equipment"
layout: "intersight"
page_title: "Intersight: intersight_equipment_shared_io_module"
description: |-
  I/O Controller present inside SIOC to provide data path from S-series server to FI.
---

# Data Source: intersight_equipment_shared_io_module
I/O Controller present inside SIOC to provide data path from S-series server to FI.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_equipment_shared_io_module.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `config_state`:(string) This field identifies the configuration state for this SIOM Unit. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `discovery`:(string) This field identifies the discovery state of SIOM. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mac_of_shared_iom_aside`:(string) This field identifies the MAC of IOM-A side. 
* `mac_of_shared_iom_bside`:(string) This field identifies the MAC of IOM-B side. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) This field identifies the model of the given component. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `oper_state`:(string) This field identifies the SIOM operational state. 
* `part_number`:(string) This field identifies the Part Number for this SIOM Unit. 
* `presence`:(string) This field identifies the presence (equipped) or absence of the given component. 
* `reachability`:(string) This field identifies the reachability to FI-A and B side. 
* `revision`:(string) This field identifies the revision of the given component. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field identifies the serial of the given component. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `usr_lbl`:(string) User label configured for the SIOM. 
* `vendor`:(string) This field identifies the vendor of the given component. 
* `vid`:(string) This field identifies the vendor id for this SIOM Unit. 
 