---
subcategory: "chassis"
layout: "intersight"
page_title: "Intersight: intersight_chassis_iom_profile"
description: |-
  A profile specifying configuration settings for IOM.
---

# Data Source: intersight_chassis_iom_profile
A profile specifying configuration settings for IOM.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `action`:(string) User initiated action. Each profile type has its own supported actions. For HyperFlex cluster profile, the supported actions are -- Validate, Deploy, Continue, Retry, Abort, Unassign For server profile, the support actions are -- Deploy, Unassign. 
* `description`:(string) Description of the profile. 
* `iom_entity`:(string) IOM in chassis for which IOM profile is applicable. or which is attached to a Fabric Interconnect managed by Intersight.* `IOMA` - IOM on left side of chassis.* `IOMB` - IOM on right side of chassis. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete profile. 
* `type`:(string) Defines the type of the profile. Accepted value is instance.* `instance` - The profile defines the configuration for a specific instance of a target. 
