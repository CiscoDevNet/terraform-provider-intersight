
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
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `description`:(string) Description of the policy. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `node_name_prefix`:(string) The node name prefix that is used to automatically generate the default hostname for each server.A dash (-) will be appended to the prefix followed by the node number to form a hostname.This default naming scheme can be manually overridden in the node configuration.The maximum length of a prefix is 60, must only contain alphanumeric characters or dash (-), and muststart with an alphanumeric character. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
