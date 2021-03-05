---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_ext_fc_storage_policy"
description: |-
  A policy specifying external storage connectivity information via Fabric attached FC storage.
---

# Data Source: intersight_hyperflex_ext_fc_storage_policy
A policy specifying external storage connectivity information via Fabric attached FC storage.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_ext_fc_storage_policy.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `admin_state`:(bool) Enables or disables external FC storage configuration. 
* `description`:(string) Description of the policy. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
 