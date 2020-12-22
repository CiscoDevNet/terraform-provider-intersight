---
subcategory: "fabric"
layout: "intersight"
page_title: "Intersight: intersight_fabric_switch_profile"
description: |-
  This specifies configuration policies for a managed network switch.
---

# Data Source: intersight_fabric_switch_profile
This specifies configuration policies for a managed network switch.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `action`:(string) User initiated action. Each profile type has its own supported actions. For HyperFlex cluster profile, the supported actions are -- Validate, Deploy, Continue, Retry, Abort, Unassign For server profile, the support actions are -- Deploy, Unassign. 
* `description`:(string) Description of the profile. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete profile. 
* `type`:(string) Defines the type of the profile. Accepted value is instance.* `instance` - The profile defines the configuration for a specific instance of a target. 
