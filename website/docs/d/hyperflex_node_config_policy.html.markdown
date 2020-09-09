
---
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_node_config_policy"
sidebar_current: "docs-intersight-data-source-hyperflex-node-config-policy"
description: |-
A policy specifying node details such as management and storage data IP ranges. For HyperFlex Edge, storage data IP range is pre-defined.
---

# Data Source: intersight_hyperflex._node_config_policy
A policy specifying node details such as management and storage data IP ranges. For HyperFlex Edge, storage data IP range is pre-defined.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `description`:(string) Description of the policy. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `node_name_prefix`:(string) The node name prefix that is used to automatically generate the default hostname for each server.A dash (-) will be appended to the prefix followed by the node number to form a hostname.This default naming scheme can be manually overridden in the node configuration.The maximum length of a prefix is 60, must only contain alphanumeric characters or dash (-), and muststart with an alphanumeric character. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
