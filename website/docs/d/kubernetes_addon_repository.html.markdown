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
To access the ith object of the results obtained, use `data.intersight_kubernetes_addon_repository.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `insecure_skip_verification`:(bool) Allow connecting to http registries or https registries which do not have certificate signed by a well known CA. 
* `is_password_set`:(bool) Indicates whether the value of the 'password' property has been set. 
* `is_token_set`:(bool) Indicates whether the value of the 'token' property has been set. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the addon repository or registry. 
* `repository_url`:(string) URL for the repository where the addon is hosted. 
* `username`:(string) Username to authenticate to the addon registry. 
 