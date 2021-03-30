---
subcategory: "kubernetes"
layout: "intersight"
page_title: "Intersight: intersight_kubernetes_cluster_profile"
description: |-
  Cluster profile specifies the config profile for a Kubernetes cluster. It also depicts operations to control the life cycle of a Kubernetes cluster.
---

# Data Source: intersight_kubernetes_cluster_profile
Cluster profile specifies the config profile for a Kubernetes cluster. It also depicts operations to control the life cycle of a Kubernetes cluster.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_kubernetes_cluster_profile.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `action`:(string) User initiated action. Each profile type has its own supported actions. For HyperFlex cluster profile, the supported actions are -- Validate, Deploy, Continue, Retry, Abort, Unassign For server profile, the support actions are -- Deploy, Unassign. 
* `description`:(string) Description of the profile. 
* `managed_mode`:(string) Management mode for the cluster. In some cases Intersight kubernetes service is not requiredto provision and manage the management entities and endpoints (for e.g. EKS). In most other casesit will be required to provision and manage these entities and endpoints.* `Provided` - Cluster management entities and endpoints are provided by the infrastructure platform.* `Managed` - Cluster management entities and endpoints are provisioned and managed by Intersight kubernetes service. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete profile. 
* `status`:(string) Status of the Kubernetes cluster and its nodes.* `Undeployed` - The cluster is undeployed.* `Configuring` - The cluster is being configured.* `Deploying` - The cluster is being deployed.* `Undeploying` - The cluster is being undeployed.* `DeployFailedTerminal` - The cluster deployment failed terminally and can not be recovered.* `DeployFailed` - The cluster deployment failed.* `Upgrading` - The cluster is being upgraded.* `Deleting` - The cluster is being deleted.* `DeleteFailed` - The cluster delete failed.* `Ready` - The cluster is ready for use.* `Active` - The cluster is being active.* `Shutdown` - All the nodes in the cluster are powered off.* `Terminated` - The cluster is terminated.* `Deployed` - The cluster is deployed. The cluster may not yet be ready for use.* `UndeployFailed` - The cluster undeploy action failed.* `NotReady` - The cluster is created and some nodes are not ready. 
* `type`:(string) Defines the type of the profile. Accepted value is instance.* `instance` - The profile defines the configuration for a specific instance of a target. 
 