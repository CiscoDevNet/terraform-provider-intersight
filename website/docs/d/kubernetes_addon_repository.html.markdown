---
subcategory: "kubernetes"
layout: "intersight"
page_title: "Intersight: intersight_kubernetes_addon_repository"
description: |-
        Docker registry or helm repository which hosts helm charts and docker images.

---

# Data Source: intersight_kubernetes_addon_repository
Docker registry or helm repository which hosts helm charts and docker images.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_kubernetes_addon_repository.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `insecure_skip_verification`:(bool) Allow connecting to http registries or https registries which do not have certificate signed by a well known CA. 
* `is_password_set`:(bool) Indicates whether the value of the 'password' property has been set. 
* `is_token_set`:(bool) Indicates whether the value of the 'token' property has been set. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the addon repository or registry. 
* `repository_url`:(string) URL for the repository where the addon is hosted. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `username`:(string) Username to authenticate to the addon registry. 
 
