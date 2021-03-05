---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_node_config_policy"
description: |-
  A policy specifying node details such as management and storage data IP ranges. For HyperFlex Edge, storage data IP range is pre-defined.
---

# Data Source: intersight_hyperflex_node_config_policy
A policy specifying node details such as management and storage data IP ranges. For HyperFlex Edge, storage data IP range is pre-defined.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_node_config_policy.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string) Description of the policy. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `node_name_prefix`:(string) The node name prefix that is used to automatically generate the default hostname for each server.A dash (-) will be appended to the prefix followed by the node number to form a hostname.This default naming scheme can be manually overridden in the node configuration.The maximum length of a prefix is 60, must only contain alphanumeric characters or dash (-), and muststart with an alphanumeric character. 
 