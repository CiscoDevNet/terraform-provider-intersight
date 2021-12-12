---
subcategory: "server"
layout: "intersight"
page_title: "Intersight: intersight_server_profile"
description: |-
  A profile specifying configuration settings for a physical server.
---

# Resource: intersight_server_profile
A profile specifying configuration settings for a physical server.
## Usage Example
### Resource Creation

```hcl
resource "intersight_server_profile" "server1" {
  name   = "server1"
  action = "No-op"
  tags {
    key   = "server"
    value = "demo"
  }
  organization {
    object_type = "organization.Organization"
    moid        = var.organization
  }
}
```
## Argument Reference
The following arguments are supported:
* `account_moid`:(string)(ReadOnly) The Account ID for this managed object. 
* `action`:(string) User initiated action. Each profile type has its own supported actions. For HyperFlex cluster profile, the supported actions are -- Validate, Deploy, Continue, Retry, Abort, Unassign For server profile, the support actions are -- Deploy, Unassign. 
* `action_params`:(Array)
This complex property has following sub-properties:
  + `name`:(string) The action parameter identifier. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `value`:(string) The action parameter value. 
* `ancestors`:(Array)(ReadOnly) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `assigned_server`:(HashMap) - A reference to a computePhysical resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `associated_server`:(HashMap) -(ReadOnly) A reference to a computePhysical resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `associated_server_pool`:(HashMap) - A reference to a resourcepoolPool resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `config_change_context`:(HashMap) -(ReadOnly) The configuration change state and results of the last change operation. 
This complex property has following sub-properties:
  + `config_change_error`:(string)(ReadOnly) Indicates reason for failure state of configChangeState. 
  + `config_change_state`:(string)(ReadOnly) Indicates a profile's configuration change state. Used for tracking pending-changes and out-of-synch states.* `Ok` - Config change state represents Validation for change/drift is successful or is not applicable.* `Validating-change` - Config change state represents policy changes are being validated. This state starts when policy is changed and becomes different from deployed changes (Pending-changes).* `Validating-drift` - Config change state represents endpoint configuration changes are being validated. This state starts when endpoint is changed and endpoint configuration becomes different from policy configured changes (Out-of-sync).* `Change-failed` - Config change state represents there is internal error in calculating policy change.* `Drift-failed` - Config change state represents there is internal error in calculating endpoint configuraion drift. 
  + `initial_config_context`:(HashMap) -(ReadOnly) Stores initial Configuration state. Used for reverting back to initial state of ConfigContext in case of validation failure. 
This complex property has following sub-properties:
    + `config_state`:(string)(ReadOnly) Indicates a profile's configuration deploying state. Values -- Assigned, Not-assigned, Associated, Pending-changes, Out-of-sync, Validating, Configuring, Failed. 
    + `config_type`:(string)(ReadOnly) The type of configuration running on the profile. Since profile deployments can configure multiple different settings, configType indicates which type of configuration is currently in progress. 
    + `control_action`:(string) System action to trigger the appropriate workflow. Values -- No_op, ConfigChange, Deploy, Unbind. 
    + `error_state`:(string) Indicates a profile's error state. Values -- Validation-error (Static validation error), Pre-config-error (Runtime validation error), Config-error (Runtime configuration error). 
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
    + `oper_state`:(string)(ReadOnly) Combined state (configState, and operational state of the associated physical resource) to indicate the current state of the profile. Values -- n/a, Power-off, Pending-changes, Configuring, Ok, Failed. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `config_change_details`:(Array)(ReadOnly) An array of relationships to serverConfigChangeDetail resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `config_changes`:(HashMap) - Pending configuration changes at the summary level. Detail changes are saved in configChangeDetails as a separate object. 
This complex property has following sub-properties:
  + `changes`:
                (Array of schema.TypeString) -
  + `disruptions`:
                (Array of schema.TypeString) -
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `config_context`:(HashMap) - The configuration state and results of the last configuration operation. 
This complex property has following sub-properties:
  + `config_state`:(string)(ReadOnly) Indicates a profile's configuration deploying state. Values -- Assigned, Not-assigned, Associated, Pending-changes, Out-of-sync, Validating, Configuring, Failed. 
  + `config_type`:(string)(ReadOnly) The type of configuration running on the profile. Since profile deployments can configure multiple different settings, configType indicates which type of configuration is currently in progress. 
  + `control_action`:(string) System action to trigger the appropriate workflow. Values -- No_op, ConfigChange, Deploy, Unbind. 
  + `error_state`:(string) Indicates a profile's error state. Values -- Validation-error (Static validation error), Pre-config-error (Runtime validation error), Config-error (Runtime configuration error). 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `oper_state`:(string)(ReadOnly) Combined state (configState, and operational state of the associated physical resource) to indicate the current state of the profile. Values -- n/a, Power-off, Pending-changes, Configuring, Ok, Failed. 
* `config_result`:(HashMap) -(ReadOnly) A reference to a serverConfigResult resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `create_time`:(string)(ReadOnly) The time when this managed object was created. 
* `description`:(string) Description of the profile. 
* `domain_group_moid`:(string)(ReadOnly) The DomainGroup ID for this managed object. 
* `is_pmc_deployed_secure_passphrase_set`:(bool)(ReadOnly) Indicates whether the value of the 'pmcDeployedSecurePassphrase' property has been set. 
* `leased_server`:(HashMap) - A reference to a computePhysical resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `mod_time`:(string)(ReadOnly) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the profile instance or profile template. 
* `organization`:(HashMap) - A reference to a organizationOrganization resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `owners`:
                (Array of schema.TypeString) -(ReadOnly)
* `parent`:(HashMap) -(ReadOnly) A reference to a moBaseMo resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `permission_resources`:(Array)(ReadOnly) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `pmc_deployed_secure_passphrase`:(string) Secure passphrase that is already deployed on all the Persistent Memory Modules on the server. This deployed passphrase is required during deploy of server profile if secure passphrase is changed or security is disabled in the attached persistent memory policy. 
* `policy_bucket`:(Array) An array of relationships to policyAbstractPolicy resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `resource_lease`:(HashMap) -(ReadOnly) A reference to a resourcepoolLease resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `running_workflows`:(Array)(ReadOnly) An array of relationships to workflowWorkflowInfo resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `server_assignment_mode`:(string) Source of the server assigned to the server profile. Values can be Static, Pool or None. Static is used if a server is attached directly to server profile. Pool is used if a resource pool is attached to server profile. None is used if no server or resource pool is attached to server profile.* `None` - No server is assigned to the server profile.* `Static` - Server is directly assigned to server profile using assign server.* `Pool` - Server is assigned from a resource pool. 
* `server_pool`:(HashMap) - A reference to a resourcepoolPool resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `shared_scope`:(string)(ReadOnly) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `src_template`:(HashMap) - A reference to a policyAbstractProfile resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `static_uuid_address`:(string) The UUID address for the server must include UUID prefix xxxxxxxx-xxxx-xxxx along with the UUID suffix of format xxxx-xxxxxxxxxxxx. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `target_platform`:(string) The platform for which the server profile is applicable. It can either be a server that is operating in standalone mode or which is attached to a Fabric Interconnect managed by Intersight.* `Standalone` - Servers which are operating in standalone mode i.e. not connected to a Fabric Interconnected.* `FIAttached` - Servers which are connected to a Fabric Interconnect that is managed by Intersight. 
* `type`:(string) Defines the type of the profile. Accepted values are instance or template.* `instance` - The profile defines the configuration for a specific instance of a target. 
* `uuid`:(string)(ReadOnly) The UUID address that is assigned to the server based on the UUID pool. 
* `uuid_address_type`:(string) UUID address allocation type selected to assign an UUID address for the server.* `NONE` - The user did not assign any UUID address.* `STATIC` - The user assigns a static UUID address.* `POOL` - The user selects a pool from which the address will be leased. 
* `uuid_lease`:(HashMap) -(ReadOnly) A reference to a uuidpoolUuidLease resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `uuid_pool`:(HashMap) - A reference to a uuidpoolPool resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `version_context`:(HashMap) -(ReadOnly) The versioning info for this managed object. 
This complex property has following sub-properties:
  + `interested_mos`:(Array)
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `ref_mo`:(HashMap) -(ReadOnly) A reference to the original Managed Object. 
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `timestamp`:(string)(ReadOnly) The time this versioned Managed Object was created. 
  + `nr_version`:(string)(ReadOnly) The version of the Managed Object, e.g. an incrementing number or a hash id. 
  + `version_type`:(string)(ReadOnly) Specifies type of version. Currently the only supported value is \ Configured\ that is used to keep track of snapshots of policies and profiles that are intendedto be configured to target endpoints.* `Modified` - Version created every time an object is modified.* `Configured` - Version created every time an object is configured to the service profile.* `Deployed` - Version created for objects related to a service profile when it is deployed. 

### Custom keywords
These are
* `wait_for_completion`:(bool) This model object can trigger workflows. Use this option to wait for all running workflows to reach a complete state. Default value is True i.e. wait.

## Import
`intersight_server_profile` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_server_profile.example 1234567890987654321abcde
``` 