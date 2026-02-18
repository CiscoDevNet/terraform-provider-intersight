---
subcategory: "capability"
layout: "intersight"
page_title: "Intersight: intersight_capability_port_group_aggregation_def"
description: |-
        The PortGroupAggregationDef defines how port-groups on FEX/IOCARD modules can be aggregated.
        #### Purpose
        A PortGroupAggregationDef serves as the blueprint for defining port group aggregation capabilities, supporting efficient network management and configuration.
        #### Key Concepts
        - **Aggregation Capability:** Specifies the aggregation capabilities associated with port groups, enhancing network efficiency and resource management.
        - **Module Integration:** Focuses on FEX/IOCARD module port groups, ensuring seamless integration and configuration within network systems.
        - **Access Management:** Utilizes privilege sets to ensure secure access and management of port group capabilities, maintaining system integrity.

---

# Data Source: intersight_capability_port_group_aggregation_def
The PortGroupAggregationDef defines how port-groups on FEX/IOCARD modules can be aggregated.
#### Purpose
A PortGroupAggregationDef serves as the blueprint for defining port group aggregation capabilities, supporting efficient network management and configuration. 
#### Key Concepts
- **Aggregation Capability:** Specifies the aggregation capabilities associated with port groups, enhancing network efficiency and resource management.
- **Module Integration:** Focuses on FEX/IOCARD module port groups, ensuring seamless integration and configuration within network systems.
- **Access Management:** Utilizes privilege sets to ensure secure access and management of port group capabilities, maintaining system integrity.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_capability_port_group_aggregation_def.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `aggregation_cap`:(string) Aggregation capability for port group. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `hw40_g_port_group_cap`:(bool) Indicates support for 40G port group capability. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) An unique identifer for a capability descriptor. 
* `pgtype`:(string) The type of port group for which this capability is defined. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
