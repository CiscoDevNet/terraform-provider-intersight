---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_health"
description: |-
  The data health status and ability of the HyperFlex storage cluster to tolerate failures.
---

# Data Source: intersight_hyperflex_health
The data health status and ability of the HyperFlex storage cluster to tolerate failures.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_health.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `arbitration_service_state`:(string) The status of the HyperFlex cluster's connection to the Intersight arbitration service. The arbitration service state is only applicable to 2-node edge clusters.* `NOT_AVAILABLE` - The cluster does not require a connection to the arbitration service.* `UNKNOWN` - The cluster's connection state to the arbitration service cannot be determined.* `ONLINE` - The cluster is connected to the arbitration service.* `OFFLINE` - The cluster is disconnected from the arbitration service. 
* `create_time`:(string) The time when this managed object was created. 
* `data_replication_compliance`:(string) The HyperFlex cluster's compliance to the configured replication factor. It indicates that the compliance has degraded if the number of copies of data is reduced.* `UNKNOWN` - The replication compliance of the HyperFlex cluster is not known.* `COMPLIANT` - The HyperFlex cluster is compliant with the replication policy. All data on the cluster is replicated according to the configured replication factor.* `NON_COMPLIANT` - The HyperFlex cluster is not compliant with the replication policy. Some data on the cluster is not replicated in accordance with the configured replication factor. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `state`:(string) The operational status of the HyperFlex cluster.* `UNKNOWN` - The operational status of the cluster cannot be determined.* `ONLINE` - The HyperFlex cluster is online and is performing IO operations.* `OFFLINE` - The HyperFlex cluster is offline and is not ready to perform IO operations.* `ENOSPACE` - The HyperFlex cluster is out of available storage capacity and cannot perform write transactions.* `READONLY` - The HyperFlex cluster is not accepting write transactions, but can still display static cluster information. 
* `uuid`:(string) The unique identifier for the cluster. 
* `zk_health`:(string) The health status of the HyperFlex cluster's zookeeper ensemble.* `NOT_AVAILABLE` - The operational status of the ZK ensemble is not provided by the HyperFlex cluster.* `UNKNOWN` - The operational status of the ZK ensemble cannot be determined.* `ONLINE` - The ZK ensemble is online and operational.* `OFFLINE` - The ZK ensemble is offline and not operational. 
 