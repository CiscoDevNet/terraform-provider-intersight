---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_health_check_definition"
description: |-
        HyperFlex health check definition metadata.

---

# Data Source: intersight_hyperflex_health_check_definition
HyperFlex health check definition metadata.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_health_check_definition.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `category`:(string) Category that the health check belongs to. 
* `common_causes`:(string) Static information detailing the common causes for the health check failure. 
* `configuration`:(string) Execution configuration fo the health check script. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the health check definition. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `health_impact`:(string) Static information detailing the health impact of the health check failure. 
* `internal_name`:(string) Internal name of the health check definition. 
* `minimum_hyper_flex_version`:(string) Minimum HyperFlex version that the check is supported on. Defaults to version 3.0.1. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the health check definition. 
* `reference`:(string) Static information containing additional reference for the health check. 
* `resolution`:(string) Static information detailing the possible remediation actions that can be taken to remedy the health check failure. 
* `script_execution_mode`:(string) Execution mode of the health script on the HyperFlex cluster.* `ON_DEMAND` - Execute the health check on-demand.* `SCHEDULED` - Execute the health check on a scheduled interval. 
* `script_execution_on_compute_nodes`:(bool) Indicates if the script needs to be executed on HyperFlex compute nodes. |Typically, scripts are only executed on the storage Nodes. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `supported_hypervisor_type`:(string) Hypervisor type that the Health Check is supported on (All, if it is hypervisor agnostic).* `All` - The Health Check is hypervisor-agnostic.* `ESXi` - The Health Check is supported only on Vmware ESXi hypervisor of any version.* `HyperV` - The Health Check is supported only on Microsoft HyperV hypervisor. 
* `target_execution_type`:(string) Indicates whether the health check is executed only on the leader node, or on all nodes in the HyperFlex cluster.* `EXECUTE_ON_LEADER_NODE` - Execute the health check script only on the HyperFlex cluster's leader node.* `EXECUTE_ON_ALL_NODES` - Execute health check on all nodes and aggregate the results.* `EXECUTE_ON_ALL_NODES_AND_AGGREGATE` - Execute the health check on all Nodes and perform custom aggregation.* `EXECUTE_ON_CURRENT_NODE` - The HyperFlex health check is executed on the node which receives the request. 
* `timeout`:(int) Health check script execution timeout. 
 
