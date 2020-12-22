---
subcategory: "recovery"
layout: "intersight"
page_title: "Intersight: intersight_recovery_backup_profile"
description: |-
  Backup profile to initiate on-demand or scheduled backups at end points.
---

# Data Source: intersight_recovery_backup_profile
Backup profile to initiate on-demand or scheduled backups at end points.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `action`:(string) User initiated action. Each profile type has its own supported actions. For HyperFlex cluster profile, the supported actions are -- Validate, Deploy, Continue, Retry, Abort, Unassign For server profile, the support actions are -- Deploy, Unassign. 
* `description`:(string) Description of the profile. 
* `enabled`:(bool) Enables/Disables the schedule on the endpoint. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete profile. 
* `type`:(string) Defines the type of the profile. Accepted value is instance.* `instance` - The profile defines the configuration for a specific instance of a target. 
