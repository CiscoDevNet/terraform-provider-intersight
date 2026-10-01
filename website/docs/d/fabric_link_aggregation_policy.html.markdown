---
subcategory: "fabric"
layout: "intersight"
page_title: "Intersight: intersight_fabric_link_aggregation_policy"
description: |-
        The LinkAggregationPolicy object defines link aggregation behavior (including LACP-related intent) used for port-channels and aggregated links.
        #### Purpose
        LinkAggregationPolicy enables consistent configuration of aggregated links by defining LACP behavior and link aggregation controls in a reusable policy form, supporting predictable port-channel behavior across deployments.
        #### Key Concepts
        - **Port-channel behavior definition:** Encodes how aggregated links should negotiate and operate.
        - **Reusable policy intent:** Allows consistent LAG configuration across many port-channels and profiles.
        - **Operational stability:** Helps standardize aggregation behavior to reduce misconfiguration risk.
        - **Integration with port roles:** Typically consumed by port-channel role objects that define port-channel membership.

---

# Data Source: intersight_fabric_link_aggregation_policy
The LinkAggregationPolicy object defines link aggregation behavior (including LACP-related intent) used for port-channels and aggregated links.
#### Purpose
LinkAggregationPolicy enables consistent configuration of aggregated links by defining LACP behavior and link aggregation controls in a reusable policy form, supporting predictable port-channel behavior across deployments.
#### Key Concepts
- **Port-channel behavior definition:** Encodes how aggregated links should negotiate and operate.
- **Reusable policy intent:** Allows consistent LAG configuration across many port-channels and profiles.
- **Operational stability:** Helps standardize aggregation behavior to reduce misconfiguration risk.
- **Integration with port roles:** Typically consumed by port-channel role objects that define port-channel membership.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_fabric_link_aggregation_policy.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the policy. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `lacp_rate`:(string) Configures the LACP control-packet rate. Fast sends a packet every second and Normal sends one every 30 seconds.* `normal` - Sends LACP control packets once every 30 seconds.* `fast` - Sends LACP control packets once every second. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `suspend_individual`:(bool) Flag tells the switch whether to suspend the port if it didn’t receive LACP PDU. 
 
