---
subcategory: "crd"
layout: "intersight"
page_title: "Intersight: intersight_crd_custom_resource"
description: |-
        Custom Kubernetes resource launcher service.

---

# Data Source: intersight_crd_custom_resource
Custom Kubernetes resource launcher service.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_crd_custom_resource.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `dc_launcher`:(string) The type of custom resource in Kubernetes. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `image`:(string) The docker image URL for the cloud DC. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) A string property called 'name', used as part of a uniqueness constraint. Refer to the 'identity' specification in this MO definition. 
* `namespace`:(string) Namespace for launching the deployment associated with the custom resource. 
* `port`:(int) Port used for the cloud DC. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `target_id`:(string) Target Id for the cloud DC. 
* `target_moid`:(string) Target MOID for the cloud DC. 
* `target_type`:(string) Target type for the cloud DC. 
 
