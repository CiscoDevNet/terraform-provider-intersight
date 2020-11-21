
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
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `description`:(string) Description of the profile. 
* `hxdp_data_ip`:(string) IP address for storage data network (Controller VM interface). 
* `hxdp_mgmt_ip`:(string) IP address for HyperFlex management network. 
* `hypervisor_data_ip`:(string) IP address for storage data network (Hypervisor interface). 
* `hypervisor_mgmt_ip`:(string) IP address for Hypervisor management network. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete profile. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `type`:(string) Defines the type of the profile. Accepted value is instance.* `instance` - The profile defines the configuration for a specific instance of a target. 
