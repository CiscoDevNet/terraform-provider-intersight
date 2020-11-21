
---
layout: "intersight"
page_title: "Intersight: intersight_fabric_switch_cluster_profile"
sidebar_current: "docs-intersight-data-source-fabric-switch-cluster-profile"
description: |-
This specifies the configuration policies for a cluster of switches.
---

# Data Source: intersight_fabric._switch_cluster_profile
This specifies the configuration policies for a cluster of switches.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `description`:(string) Description of the profile. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete profile. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `switch_profiles_count`:(int) Number of switch profiles that are part of this cluster profile. 
* `type`:(string) Defines the type of the profile. Accepted value is instance.* `instance` - The profile defines the configuration for a specific instance of a target. 
