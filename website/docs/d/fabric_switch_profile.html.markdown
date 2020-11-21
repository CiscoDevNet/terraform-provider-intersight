
---
layout: "intersight"
page_title: "Intersight: intersight_fabric_switch_profile"
sidebar_current: "docs-intersight-data-source-fabric-switch-profile"
description: |-
This specifies configuration policies for a managed network switch.
---

# Data Source: intersight_fabric._switch_profile
This specifies configuration policies for a managed network switch.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `action`:(string) User initiated action. Each profile type has its own supported actions. For HyperFlex cluster profile, the supported actions are -- Validate, Deploy, Continue, Retry, Abort, Unassign For server profile, the support actions are -- Deploy, Unassign. 
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `description`:(string) Description of the profile. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete profile. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `type`:(string) Defines the type of the profile. Accepted value is instance.* `instance` - The profile defines the configuration for a specific instance of a target. 
