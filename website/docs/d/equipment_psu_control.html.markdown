---
subcategory: "equipment"
layout: "intersight"
page_title: "Intersight: intersight_equipment_psu_control"
description: |-
        The equipment.PsuControl object represents the overall power state and redundancy configuration of the Power Supply Units (PSUs) within a chassis.
        #### Purpose
        The primary function of this object is to provide a consolidated view of the chassis's power redundancy status. It reports on the collective state of the PSUs, including the overall redundancy mode (redundancy), the input and output power states, and any system-level power-related health issues.
        #### Key Concepts
        - **System-Level Power State:** Aggregates the status of all PSUs in a chassis into a single management object.
        - **Redundancy Monitoring:** The redundancy and clusterState properties are key indicators of the chassis's power fault tolerance (e.g., N+1, non-redundant).
        - **Health Aggregation:** The operReason property provides details on chassis-wide power issues, such as a loss of redundancy (PsuRedundancyLostCritical).
        - **Chassis-Specific:** It is a child of an equipment.Chassis, providing a summary of the power subsystem for that specific enclosure.

---

# Data Source: intersight_equipment_psu_control
The equipment.PsuControl object represents the overall power state and redundancy configuration of the Power Supply Units (PSUs) within a chassis.
#### Purpose
The primary function of this object is to provide a consolidated view of the chassis's power redundancy status. It reports on the collective state of the PSUs, including the overall redundancy mode (redundancy), the input and output power states, and any system-level power-related health issues.
#### Key Concepts
- **System-Level Power State:** Aggregates the status of all PSUs in a chassis into a single management object.
- **Redundancy Monitoring:** The redundancy and clusterState properties are key indicators of the chassis's power fault tolerance (e.g., N+1, non-redundant).
- **Health Aggregation:** The operReason property provides details on chassis-wide power issues, such as a loss of redundancy (PsuRedundancyLostCritical).
- **Chassis-Specific:** It is a child of an equipment.Chassis, providing a summary of the power subsystem for that specific enclosure.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_equipment_psu_control.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `cluster_state`:(string) This field identifies the cluster state of the psu redundancy. 
* `create_time`:(string) The time when this managed object was created. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `input_power_state`:(string) This field identifies the input power state of the psus. 
* `is_upgraded`:(bool) This field indicates the compute status of the catalog values for the associated component or hardware. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) This field displays the model number of the associated component or hardware. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) This field identifies the name of psu control object. 
* `oper_qualifier`:(string) This field identifies the operational qualifier for the psu redundancy. 
* `oper_state`:(string) This field identifies the operational state of the psu redundancy. 
* `output_power_state`:(string) This field identifies the output power state of the psus. 
* `presence`:(string) This field indicates the presence (equipped) or absence (absent) of the associated component or hardware. 
* `redundancy`:(string) This field identifies the redundancy state of the psus. 
* `revision`:(string) This field displays the revised version of the associated component or hardware (if any). 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field displays the serial number of the associated component or hardware. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `vendor`:(string) This field displays the vendor information of the associated component or hardware. 
 
