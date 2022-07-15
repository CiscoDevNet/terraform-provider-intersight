---
subcategory: "convergedinfra"
layout: "intersight"
page_title: "Intersight: intersight_convergedinfra_pod"
description: |-
        A pod is unit of deployment of converged infrastructure. Contains inventory information related to the health, HCL, storage,
        nodes, etc. of the pod.

---

# Data Source: intersight_convergedinfra_pod
A pod is unit of deployment of converged infrastructure. Contains inventory information related to the health, HCL, storage,
nodes, etc. of the pod.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_convergedinfra_pod.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `deployment_type`:(string) The deployment type for this integrated system.* `FlexPodInfra` - The deployment type for a pod is of Infrastructure.* `FlexPodNG` - The deployment type for a pod is of Nextgen type. 
* `description`:(string) Description of the pod. A short note about the nature or purpose of the pod. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `interop_status`:(string) The interoperability status for this solution pod.* `NotEvaluated` - The interoperability compliance for the component has not be checked.* `Approved` - The component is valid as per the interoperability compliance check.* `NotApproved` - The component is not valid as per the interoperability compliance check.* `Incomplete` - The interoperability compliance check could not be completed for the component due to incomplete data. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the pod. Concrete pod will be created with this name. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `type`:(string) Defines the type of the pod.* `FlexPod` - Pod type is FlexPod, an integrated infrastructure solution developed by Cisco and NetApp.* `FlashStack` - Pod type is FlashStack, an integrated infrastructure solution developed by Cisco and Pure Storage. 
 
