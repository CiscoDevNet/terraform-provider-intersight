---
subcategory: "appliance"
layout: "intersight"
page_title: "Intersight: intersight_appliance_cluster_info"
description: |-
        ClusterInfo model is used for a peer appliance to create a cluster with an existing, fully configured
        appliance.

---

# Data Source: intersight_appliance_cluster_info
ClusterInfo model is used for a peer appliance to create a cluster with an existing, fully configured
appliance.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_appliance_cluster_info.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `deployment_size`:(string) The deployment size of the node requiring to join cluster. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `gateway`:(string) Default gateway configured on the peer node. 
* `hostip`:(string) Publicly accessible IP of the peer node. 
* `hostname`:(string) Publicly accessible FQDN of the peer node. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `peerkey`:(string) The public key of peer host. 
* `responsekey`:(string) Public key returned to the client. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `status`:(string) The status of the cluster join process.* `Unknown` - The status of the appliance node is unknown.* `Operational` - The appliance node is operational.* `Impaired` - The appliance node is impaired.* `AttentionNeeded` - The appliance node needs attention.* `ReadyToJoin` - The node is ready to be added to a standalone Intersight Appliance to form a cluster.* `OutOfService` - The user has taken this node (part of a cluster) to out of service.* `ReadyForReplacement` - The cluster node is ready to be replaced.* `ReplacementInProgress` - The cluster node replacement is in progress.* `ReplacementFailed` - There was a failure during the cluster node replacement. 
* `subnetmask`:(string) Subnet Mask of the peer node. 
* `uuid`:(string) The UUID of the peer appliance. 
 
