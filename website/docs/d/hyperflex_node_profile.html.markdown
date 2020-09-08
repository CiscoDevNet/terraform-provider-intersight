
---
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_node_profile"
sidebar_current: "docs-intersight-data-source-hyperflex-node-profile"
description: |-
A configuration profile per node in the HyperFlex cluster.
It defines node settings such as IP address configuration for hypervisor management network, storage data network, HyperFlex management network, and the assigned physical server.
---

# Data Source: intersight_hyperflex._node_profile
A configuration profile per node in the HyperFlex cluster.
It defines node settings such as IP address configuration for hypervisor management network, storage data network, HyperFlex management network, and the assigned physical server.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `description`:(string) Description of the profile. 
* `hxdp_data_ip`:(string) IP address for storage data network (Controller VM interface). 
* `hxdp_mgmt_ip`:(string) IP address for HyperFlex management network. 
* `hypervisor_data_ip`:(string) IP address for storage data network (Hypervisor interface). 
* `hypervisor_mgmt_ip`:(string) IP address for Hypervisor management network. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete profile. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `type`:(string) Defines the type of the profile. Accepted value is instance.* `instance` - The profile defines the configuration for a specific instance of a target. 
