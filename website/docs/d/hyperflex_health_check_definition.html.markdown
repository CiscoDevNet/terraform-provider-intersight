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
The following arguments can be used to get data of already created objects in Intersight appliance:
* `category`:(string) Category that the health check belongs to. 
* `common_causes`:(string) Static information detailing the common causes for the health check failure. 
* `configuration`:(string) Execution configuration fo the health check script. 
* `description`:(string) Description of the health check definition. 
* `health_impact`:(string) Static information detailing the health impact of the health check failure. 
* `internal_name`:(string) Internal name of the health check definition. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the health check definition. 
* `reference`:(string) Static information containing additional reference for the health check. 
* `resolution`:(string) Static information detailing the possible remediation actions that can be taken to remedy the health check failure. 
* `script_execution_mode`:(string) Execution mode of the health script on the HyperFlex cluster.* `ON_DEMAND` - Execute the health check on-demand.* `SCHEDULED` - Execute the health check on a scheduled interval. 
* `script_execution_on_compute_nodes`:(bool) Indicates if the script needs to be executed on HyperFlex compute nodes. |Typically, scripts are only executed on the storage Nodes. 
* `target_execution_type`:(string) Indicates whether the health check is executed only on the leader node, or on all nodes in the HyperFlex cluster.* `EXECUTE_ON_LEADER_NODE` - Execute the health check script only on the HyperFlex cluster's leader node.* `EXECUTE_ON_ALL_NODES` - Execute health check on all nodes and aggregate the results.* `EXECUTE_ON_ALL_NODES_AND_AGGREGATE` - Execute the health check on all Nodes and perform custom aggregation. 
* `timeout`:(int) Health check script execution timeout. 
