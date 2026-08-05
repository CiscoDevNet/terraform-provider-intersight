---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_cluster_replication_network_policy"
description: |-
        The ClusterReplicationNetworkPolicy object is a fundamental element within the HyperFlex system, designed to specify the replication network parameters for a cluster. It ensures efficient data replication across network environments, facilitating seamless data synchronization and redundancy.
        #### Purpose
        ClusterReplicationNetworkPolicy provides the framework for managing replication network configurations within a HyperFlex Cluster. It allows for the detailed specification of network settings, ensuring reliable and optimized data replication across different nodes and clusters.
        #### Key Concepts
        - **Network Configuration:** Defines essential network parameters, such as VLANs, IP ranges, and bandwidth, ensuring robust and efficient data replication.
        - **Optimization:** Supports bandwidth management and MTU settings, allowing for tailored network configurations that enhance replication performance.
        - **Scalability:** Accommodates multiple IP ranges and replication settings, supporting the dynamic expansion and modification of network configurations.
        - **Interoperability:** Integrates with cluster profiles and policies, ensuring coherent network management aligned with overall cluster operations.

---

# Data Source: intersight_hyperflex_cluster_replication_network_policy
The ClusterReplicationNetworkPolicy object is a fundamental element within the HyperFlex system, designed to specify the replication network parameters for a cluster. It ensures efficient data replication across network environments, facilitating seamless data synchronization and redundancy. 
#### Purpose  
ClusterReplicationNetworkPolicy provides the framework for managing replication network configurations within a HyperFlex Cluster. It allows for the detailed specification of network settings, ensuring reliable and optimized data replication across different nodes and clusters.  
#### Key Concepts  
- **Network Configuration:** Defines essential network parameters, such as VLANs, IP ranges, and bandwidth, ensuring robust and efficient data replication. 
- **Optimization:** Supports bandwidth management and MTU settings, allowing for tailored network configurations that enhance replication performance. 
- **Scalability:** Accommodates multiple IP ranges and replication settings, supporting the dynamic expansion and modification of network configurations. 
- **Interoperability:** Integrates with cluster profiles and policies, ensuring coherent network management aligned with overall cluster operations.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_cluster_replication_network_policy.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the policy. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `replication_bandwidth_mbps`:(int) Bandwidth for the Replication network in Mbps. 
* `replication_mtu`:(int) MTU for the Replication network. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
