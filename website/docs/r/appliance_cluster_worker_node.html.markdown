---
subcategory: "appliance"
layout: "intersight"
page_title: "Intersight: intersight_appliance_cluster_worker_node"
description: |-
        ClusterWorkerNode is an object that tracks the Intersight Appliance's process for adding a worker node.

---

# Resource: intersight_appliance_cluster_worker_node
ClusterWorkerNode is an object that tracks the Intersight Appliance's process for adding a worker node.
## Argument Reference
The following arguments are supported:
* `account`:(HashMap) -(ReadOnly) A reference to a iamAccount resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `account_moid`:(string)(ReadOnly) The Account ID for this managed object. 
* `ancestors`:(Array)(ReadOnly) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `completed_phases`:(Array)
This complex property has following sub-properties:
  + `current_node`:(int)(ReadOnly) Node number install is running on. 
  + `current_subphase`:(int)(ReadOnly) Subphase of the phase currently running. 
  + `elapsed_time`:(int)(ReadOnly) Elapsed time in seconds to complete the current install phase. 
  + `end_time`:(string)(ReadOnly) End date of the software install phase. 
  + `failed`:(bool)(ReadOnly) Indicates if the install phase has failed or not. 
  + `message`:(string)(ReadOnly) Status message set during the install phase. 
  + `name`:(string)(ReadOnly) Name of the install phase.* `Backup` - Backup the currently running node.* `EnableBootstrap` - Disable echo and enable nginx on the currently running node.* `UpdateAnsibleHosts` - Update /etc/ansible/hosts on each node.* `UpdateNetworkConfig` - Update Network config for node IP change scenario.* `SyncImages` - Sync images and manifest to each node.* `ConfigureEtcd` - Configure etcd on each node.* `DeployServices` - Deploy kubernetes services from node 1.* `InstallEquinox` - Configure and install equinox on each node.* `Validate` - Validate equinox is running in primary/secondary mode.* `CheckApplianceClusterState` - Check the appliance cluster status.* `Success` - Upgrade completed successfully.* `Fail` - Indicates that the upgrade process has failed.* `Cancel` - Indicates that the upgrade was canceled by the Intersight Appliance. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `pending_nodes`:
                (Array of schema.TypeInt) -
  + `start_time`:(string)(ReadOnly) Start date of the software install phase. 
  + `status`:(string)(ReadOnly) Status of the install phase. 
* `create_time`:(string)(ReadOnly) The time when this managed object was created. 
* `current_phase`:(HashMap) -(ReadOnly) Current phase of the Intersight Appliance's software install. 
This complex property has following sub-properties:
  + `current_node`:(int)(ReadOnly) Node number install is running on. 
  + `current_subphase`:(int)(ReadOnly) Subphase of the phase currently running. 
  + `elapsed_time`:(int)(ReadOnly) Elapsed time in seconds to complete the current install phase. 
  + `end_time`:(string)(ReadOnly) End date of the software install phase. 
  + `failed`:(bool)(ReadOnly) Indicates if the install phase has failed or not. 
  + `message`:(string)(ReadOnly) Status message set during the install phase. 
  + `name`:(string)(ReadOnly) Name of the install phase.* `Backup` - Backup the currently running node.* `EnableBootstrap` - Disable echo and enable nginx on the currently running node.* `UpdateAnsibleHosts` - Update /etc/ansible/hosts on each node.* `UpdateNetworkConfig` - Update Network config for node IP change scenario.* `SyncImages` - Sync images and manifest to each node.* `ConfigureEtcd` - Configure etcd on each node.* `DeployServices` - Deploy kubernetes services from node 1.* `InstallEquinox` - Configure and install equinox on each node.* `Validate` - Validate equinox is running in primary/secondary mode.* `CheckApplianceClusterState` - Check the appliance cluster status.* `Success` - Upgrade completed successfully.* `Fail` - Indicates that the upgrade process has failed.* `Cancel` - Indicates that the upgrade was canceled by the Intersight Appliance. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `pending_nodes`:
                (Array of schema.TypeInt) -
  + `start_time`:(string)(ReadOnly) Start date of the software install phase. 
  + `status`:(string)(ReadOnly) Status of the install phase. 
* `domain_group_moid`:(string)(ReadOnly) The DomainGroup ID for this managed object. 
* `elapsed_time`:(int)(ReadOnly) Elapsed time in seconds during the software install. 
* `end_time`:(string)(ReadOnly) End date of the software install. 
* `error_code`:(int)(ReadOnly) Error code for Intersight Appliance's software install. In case of failure - this code will help decide if software install can be retried. 
* `hostname`:(string) Hostname of the worker node being added. 
* `messages`:
                (Array of schema.TypeString) -
* `mod_time`:(string)(ReadOnly) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `node_info`:(Array)
This complex property has following sub-properties:
  + `gateway`:(string)(ReadOnly) Gateway ip address of the cluster node. 
  + `hostname`:(string)(ReadOnly) Hostname of the cluster node. 
  + `ip_address`:(string)(ReadOnly) Ip address of the cluster node. 
  + `netmask`:(string)(ReadOnly) Netmask of the cluster node. 
  + `node_id`:(int)(ReadOnly) Id number of the cluster node. 
  + `node_moid`:(string)(ReadOnly) Moid of the corresponding appliance.ClusterInfo or appliance.NodeInfo mo. 
  + `node_type`:(string)(ReadOnly) The node type of Intersight Virtual Appliance.* `standalone` - Single Node Intersight Virtual Appliance.* `management` - Management node type when Intersight Virtual Appliance is running as management-worker deployment.* `hamanagement` - Management node type when Intersight Virtual Appliance is running as multi node HA deployment.* `metrics` - Metrics node when Intersight Virtual Appliance is running management-metrics node. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `status`:(string)(ReadOnly) Status of the cluster node.* `Unknown` - The status of the appliance node is unknown.* `Operational` - The appliance node is operational.* `Impaired` - The appliance node is impaired.* `AttentionNeeded` - The appliance node needs attention.* `ReadyToJoin` - The node is ready to be added to a standalone Intersight Appliance to form a cluster.* `OutOfService` - The user has taken this node (part of a cluster) to out of service.* `ReadyForReplacement` - The cluster node is ready to be replaced.* `ReplacementInProgress` - The cluster node replacement is in progress.* `ReplacementFailed` - There was a failure during the cluster node replacement.* `WorkerNodeInstInProgress` - The worker node installation is in progress.* `WorkerNodeInstSuccess` - The worker node installation succeeded.* `WorkerNodeInstFailed` - The worker node installation failed. 
* `owners`:
                (Array of schema.TypeString) -(ReadOnly)
* `parent`:(HashMap) -(ReadOnly) A reference to a moBaseMo resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `permission_resources`:(Array)(ReadOnly) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `remote_dns`:(string) Round robin DNS address, which should be able to resolve the hostnames of all the nodes in the cluster. 
* `reuse_node`:(bool) Indicates if the worker node being added is being reused. 
* `session_id`:(string)(ReadOnly) Session Moid for the user session. 
* `shared_scope`:(string)(ReadOnly) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `start_time`:(string) Start date of the software install. UI can modify startTime to re-schedule an install. 
* `status`:(string)(ReadOnly) Status of the Intersight Appliance's software install.* `NotReady` - Cluster is not ready. Install cannot be triggered.* `Ready` - Cluster is ready. Install can be triggered.* `InProgress` - Install is currently in progress.* `Success` - Install was run and succeeded.* `Fail` - Install was run and failed. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `ancestor_definitions`:(Array)
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `definition`:(HashMap) -(ReadOnly) The definition is a reference to the tag definition object.The tag definition object contains the properties of the tag such as name, type, and description. 
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `key`:(string) The string representation of a tag key. 
  + `propagated`:(bool)(ReadOnly) Propagated is a boolean flag that indicates whether the tag is propagated to the related managed objects. 
  + `type`:(string)(ReadOnly) An enum type that defines the type of tag. Supported values are 'pathtag' and 'keyvalue'.* `KeyValue` - KeyValue type of tag. Key is required for these tags. Value is optional.* `PathTag` - Key contain path information. Value is not present for these tags. The path is created by using the '/' character as a delimiter.For example, if the tag is \ A/B/C\ , then \ A\  is the parent tag, \ B\  is the child tag of \ A\  and \ C\  is the child tag of \ B\ . 
  + `value`:(string) The string representation of a tag value. 
* `total_nodes`:(int)(ReadOnly) Total number of nodes in the system. 
* `total_phases`:(int)(ReadOnly) TotalPhase represents the total number of the install phases for one install. 
* `version_context`:(HashMap) -(ReadOnly) The versioning info for this managed object. 
This complex property has following sub-properties:
  + `interested_mos`:(Array)
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `marked_for_deletion`:(bool)(ReadOnly) The flag to indicate if snapshot is marked for deletion or not. If flag is set then snapshot will be removed after the successful deployment of the policy. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `ref_mo`:(HashMap) -(ReadOnly) A reference to the original Managed Object. 
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `timestamp`:(string)(ReadOnly) The time this versioned Managed Object was created. 
  + `nr_version`:(string)(ReadOnly) The version of the Managed Object, e.g. an incrementing number or a hash id. 
  + `version_type`:(string)(ReadOnly) Specifies type of version. Currently the only supported value is \ Configured\ that is used to keep track of snapshots of policies and profiles that are intendedto be configured to target endpoints.* `Modified` - Version created every time an object is modified.* `Configured` - Version created every time an object is configured to the service profile.* `Deployed` - Version created for objects related to a service profile when it is deployed. 
* `vip`:(HashMap) -(ReadOnly) VIP hostname and ip address of the onprem cluster. 
This complex property has following sub-properties:
  + `gateway`:(string)(ReadOnly) Gateway ip address of the cluster node. 
  + `hostname`:(string)(ReadOnly) Hostname of the cluster node. 
  + `ip_address`:(string)(ReadOnly) Ip address of the cluster node. 
  + `netmask`:(string)(ReadOnly) Netmask of the cluster node. 
  + `node_id`:(int)(ReadOnly) Id number of the cluster node. 
  + `node_moid`:(string)(ReadOnly) Moid of the corresponding appliance.ClusterInfo or appliance.NodeInfo mo. 
  + `node_type`:(string)(ReadOnly) The node type of Intersight Virtual Appliance.* `standalone` - Single Node Intersight Virtual Appliance.* `management` - Management node type when Intersight Virtual Appliance is running as management-worker deployment.* `hamanagement` - Management node type when Intersight Virtual Appliance is running as multi node HA deployment.* `metrics` - Metrics node when Intersight Virtual Appliance is running management-metrics node. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `status`:(string)(ReadOnly) Status of the cluster node.* `Unknown` - The status of the appliance node is unknown.* `Operational` - The appliance node is operational.* `Impaired` - The appliance node is impaired.* `AttentionNeeded` - The appliance node needs attention.* `ReadyToJoin` - The node is ready to be added to a standalone Intersight Appliance to form a cluster.* `OutOfService` - The user has taken this node (part of a cluster) to out of service.* `ReadyForReplacement` - The cluster node is ready to be replaced.* `ReplacementInProgress` - The cluster node replacement is in progress.* `ReplacementFailed` - There was a failure during the cluster node replacement.* `WorkerNodeInstInProgress` - The worker node installation is in progress.* `WorkerNodeInstSuccess` - The worker node installation succeeded.* `WorkerNodeInstFailed` - The worker node installation failed. 


## Import
`intersight_appliance_cluster_worker_node` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_appliance_cluster_worker_node.example 1234567890987654321abcde
``` 
