---
subcategory: "server"
layout: "intersight"
page_title: "Intersight: intersight_server_profile"
description: |-
  A profile specifying configuration settings for a physical server.
---

# Data Source: intersight_server_profile
A profile specifying configuration settings for a physical server.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `action`:(string) User initiated action. Each profile type has its own supported actions. For HyperFlex cluster profile, the supported actions are -- Validate, Deploy, Continue, Retry, Abort, Unassign For server profile, the support actions are -- Deploy, Unassign. 
* `description`:(string) Description of the profile. 
* `is_pmc_deployed_secure_passphrase_set`:(bool) Indicates whether the value of the 'pmcDeployedSecurePassphrase' property has been set. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete profile. 
* `pmc_deployed_secure_passphrase`:(string) Secure passphrase that is already deployed on all the Persistent Memory Modules on the server. This deployed passphrase is required during deploy of server profile if secure passphrase is changed or security is disabled in the attached persistent memory policy. 
* `target_platform`:(string) The platform for which the server profile is applicable. It can either be a server that is operating in standalone mode or which is attached to a Fabric Interconnect managed by Intersight.* `Standalone` - Servers which are operating in standalone mode i.e. not connected to a Fabric Interconnected.* `FIAttached` - Servers which are connected to a Fabric Interconnect that is managed by Intersight. 
* `type`:(string) Defines the type of the profile. Accepted value is instance.* `instance` - The profile defines the configuration for a specific instance of a target. 
