---
subcategory: "server"
layout: "intersight"
page_title: "Intersight: intersight_server_profile_template"
description: |-
        The ProfileTemplate object is an essential element in configuration management, providing a templated approach to server profile creation and updates. This simplifies the management of server profiles by allowing standardized configurations to be reused across multiple instances.
        #### Purpose
        ProfileTemplate serves as the blueprint for creating server profiles, enabling efficient configuration management and consistency across deployments. It allows bulk operations such as cloning and updating profiles based on template changes.
        #### Key Concepts
        - **Template-Based Management:** Offers a structured approach to server configuration by establishing baseline settings that can be propagated to derived profiles.
        - **Bulk Operations:** Supports asynchronous operations such as cloning and updating, facilitating large-scale management of server profiles.
        - **Version Control:** Ensures template modifications are safely propagated to profiles without compromising existing configurations.
        - **Access Governance:** Privilege sets govern access to the template, maintaining secure control over server configuration processes.
        A profile template specifying configuration settings for a physical server.

---

# Resource: intersight_server_profile_template
The ProfileTemplate object is an essential element in configuration management, providing a templated approach to server profile creation and updates. This simplifies the management of server profiles by allowing standardized configurations to be reused across multiple instances.   
#### Purpose  
ProfileTemplate serves as the blueprint for creating server profiles, enabling efficient configuration management and consistency across deployments. It allows bulk operations such as cloning and updating profiles based on template changes.  
#### Key Concepts
- **Template-Based Management:** Offers a structured approach to server configuration by establishing baseline settings that can be propagated to derived profiles. 
- **Bulk Operations:** Supports asynchronous operations such as cloning and updating, facilitating large-scale management of server profiles. 
- **Version Control:** Ensures template modifications are safely propagated to profiles without compromising existing configurations.
- **Access Governance:** Privilege sets govern access to the template, maintaining secure control over server configuration processes. 
A profile template specifying configuration settings for a physical server.
## Usage Example
### Resource Creation

```hcl
resource "intersight_server_profile_template" "template1" {
  name            = "server_profile_template1"
  description     = "demo server profile template"
  target_platform = "FIAttached"
   organization {
     object_type = "organization.Organization"
     moid        = var.organization
   }
  # the following policy_bucket statements map different policies to this
  # template -- the object_type shows the policy type
   policy_bucket {
     moid        = var.precision_policy
     object_type = "boot.PrecisionPolicy"
   }
   policy_bucket {
     moid = var.access_policy
     object_type = "access.Policy"
   }
   policy_bucket {
     moid = var.sol_policy
     object_type = "sol.Policy"
   }
}
variable "organization" {
   type = string
   description = "<value for organization>"
 }

variable "sol_policy"{
  type = string
  description = "Moid of sol.Policy"
}

variable "access_policy"{
  type = string
  description = "Moid of access.Policy"
}

variable "precision_policy"{
  type = string
  description = "Moid of boot.PrecisionPolicy"
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
* `config_context`:(HashMap) - The configuration state and results of the last configuration operation. 
This complex property has following sub-properties:
  + `config_state`:(string)(ReadOnly) Indicates a profile's configuration deploying state. Values -- Assigned, Not-assigned, Associated, Pending-changes, Out-of-sync, Validating, Configuring, Failed. 
  + `config_state_summary`:(string)(ReadOnly) Indicates a profile's configuration deploying state. Values -- Assigned, Not-assigned, Associated, InConsistent, Validating, Configuring, Failed, Activating, UnConfiguring.* `None` - The default state is none.* `Not-assigned` - Server is not assigned to the profile.* `Assigned` - Server is assigned to the profile and the configurations are not yet deployed.* `Preparing` - Preparing to deploy the configuration.* `Validating` - Profile validation in progress.* `Evaluating` - Policy edit configuration change evaluation in progress.* `Configuring` - Profile deploy operation is in progress.* `UnConfiguring` - Server is unassigned and config cleanup is in progress.* `Analyzing` - Profile changes are being analyzed.* `Activating` - Configuration is being activated at the endpoint.* `Inconsistent` - Profile is inconsistent with the endpoint configuration.* `Associated` - The profile configuration has been applied to the endpoint and no inconsistencies have been detected.* `Failed` - The last action on the profile has failed.* `Not-complete` - Config import operation on the profile is not complete.* `Waiting-for-resource` - Waiting for the resource to be allocated for the profile.* `Partially-deployed` - The profile configuration has been applied on a subset of endpoints. 
  + `config_type`:(string)(ReadOnly) The type of configuration running on the profile. Since profile deployments can configure multiple different settings, configType indicates which type of configuration is currently in progress. 
  + `control_action`:(string) System action to trigger the appropriate workflow. Values -- No_op, ConfigChange, Deploy, Unbind. 
  + `error_state`:(string) Indicates a profile's error state. Values -- Validation-error (Static validation error), Pre-config-error (Runtime validation error), Config-error (Runtime configuration error). 
  + `inconsistency_reason`:
                (Array of schema.TypeString) -
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `oper_state`:(string)(ReadOnly) Combined state (configState, and operational state of the associated physical resource) to indicate the current state of the profile. Values -- n/a, Power-off, Pending-changes, Configuring, Ok, Failed. 
* `config_result`:(HashMap) -(ReadOnly) A reference to a serverConfigResult resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `create_time`:(string)(ReadOnly) The time when this managed object was created. 
* `deployed_policies`:
                (Array of schema.TypeString) -
* `description`:(string) Description of the profile. 
* `domain_group_moid`:(string)(ReadOnly) The DomainGroup ID for this managed object. 
* `enable_override`:(bool) When enabled, the configuration of the derived instances may override the template configuration. 
* `management_mode`:(string)(ReadOnly) The management mode of the server.* `IntersightStandalone` - Intersight Standalone mode of operation.* `Intersight` - Intersight managed mode of operation. 
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
* `partially_deployed_policies`:(Array)
This complex property has following sub-properties:
  + `end_point_context`:(string)(ReadOnly) Information about the endpoint to which it is applied.* `Server` - Configuration is applied to a server context.* `FI` - Configuration is applied to a Fabric Identifier (FI) context.* `IOM` - Configuration is applied to an Input/Output Module (IOM) context. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `policy`:(string)(ReadOnly) The name of the policy for which entry is created. 
* `permission_resources`:(Array)(ReadOnly) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `policy_bucket`:(Array) An array of relationships to policyAbstractPolicy resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `policy_change_details`:(Array)
This complex property has following sub-properties:
  + `changes`:
                (Array of schema.TypeString) -
  + `config_change_context`:(HashMap) - Context information on the change. 
This complex property has following sub-properties:
    + `dependent_policy_list`:
                (Array of schema.TypeString) -
    + `entity_data`:(JSON as string)(ReadOnly) The data of the object present in config result context. 
    + `entity_moid`:(string) The Moid of the object present in config result context. 
    + `entity_name`:(string) The name of the object present in config result context. 
    + `entity_type`:(string) The type of the object present in config result context. 
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
    + `parent_moid`:(string) The Moid of the parent object present in config result context. 
    + `parent_policy_object_type`:(string) The type of the policy object associated with the profile. 
    + `parent_type`:(string) The type of the parent object present in config result context. 
  + `config_change_flag`:(string) Config change flag to differentiate Pending-changes and Config-drift.* `Pending-changes` - Config change flag represents changes are due to not deployed changes from Intersight.* `Drift-changes` - Config change flag represents changes are due to endpoint configuration changes. 
  + `disruptions`:
                (Array of schema.TypeString) -
  + `message`:(string) Detailed description of the config change. 
  + `mod_status`:(string) Modification status of the mo in this config change.* `None` - The 'none' operation/state.Indicates a configuration profile has been deployed, and the desired configuration matches the actual device configuration.* `Created` - The 'create' operation/state.Indicates a configuration profile has been created and associated with a device, but the configuration specified in the profilehas not been applied yet. For example, this could happen when the user creates a server profile and has not deployed the profile yet.* `Modified` - The 'update' operation/state.Indicates some of the desired configuration changes specified in a profile have not been been applied to the associated device.This happens when the user has made changes to a profile and has not deployed the changes yet, or when the workflow to applythe configuration changes has not completed successfully.* `Deleted` - The 'delete' operation/state.Indicates a configuration profile has been been disassociated from a device and the user has not undeployed these changes yet. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `removed_policies`:
                (Array of schema.TypeString) -
* `reported_policy_changes`:(Array)
This complex property has following sub-properties:
  + `change_id`:(string)(ReadOnly) The change evaluation identifier for which the change is reported. 
  + `change_status`:(string)(ReadOnly) The status of policy change evaluation which has been reported.* `Initiated` - The status when policy change evaluation is triggered for a policy.* `Reported` - The status when policy change evaluation is reported for a policy. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `policy_type`:(string)(ReadOnly) The type of policy for which the change has been reported. 
* `scheduled_actions`:(Array)
This complex property has following sub-properties:
  + `action`:(string) Name of the action to be performed on the profile. 
  + `action_qualifier`:(HashMap) - Qualifiers to control the action being triggered. Action qualifiers are to be specified based on the type of disruptions that an action is to be restricted to. For example, trigger the 'Deploy' action to only perform network and management plane configurations. 
This complex property has following sub-properties:
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `proceed_on_reboot`:(bool) ProceedOnReboot can be used to acknowledge server reboot while triggering deploy/activate. 
* `shared_scope`:(string)(ReadOnly) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `src_template`:(HashMap) - A reference to a policyAbstractProfile resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `ancestor_definitions`:(Array)
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `definition`:(HashMap) -(ReadOnly) The definition is a reference to the tag definition object.The tag definition object contains the properties of the tag such as name, type, and description. 
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `key`:(string) The string representation of a tag key. 
  + `propagated`:(bool)(ReadOnly) Propagated is a boolean flag that indicates whether the tag is propagated to the related managed objects. 
  + `sys_tag`:(bool)(ReadOnly) Specifies whether the tag is user-defined or owned by the system. 
  + `type`:(string)(ReadOnly) An enum type that defines the type of tag. Supported values are 'pathtag' and 'keyvalue'.* `KeyValue` - KeyValue type of tag. Key is required for these tags. Value is optional.* `PathTag` - Key contain path information. Value is not present for these tags. The path is created by using the '/' character as a delimiter.For example, if the tag is \ A/B/C\ , then \ A\  is the parent tag, \ B\  is the child tag of \ A\  and \ C\  is the child tag of \ B\ . 
  + `value`:(string) The string representation of a tag value. 
* `target_platform`:(string) The platform for which the server profile is applicable. It can either be a server that is operating in standalone mode or which is attached to a Fabric Interconnect managed by Intersight.* `Standalone` - Servers which are operating in standalone mode i.e. not connected to a Fabric Interconnected.* `FIAttached` - Servers which are connected to a Fabric Interconnect that is managed by Intersight.* `UnifiedEdgeServer` - Unified Edge sleds that is managed by Intersight. 
* `template_actions`:(Array)
This complex property has following sub-properties:
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `params`:(Array)
This complex property has following sub-properties:
    + `name`:(string) The action parameter identifier. The supported values are SyncType and SyncTimer for the template sync action.* `None` - The default parameter that implies that no action parameter is required for the template action.* `SyncType` - The parameter that describes the type of sync action such as SyncOne or SyncFailed supported on any template or derived object.* `SyncTimer` - The parameter for the initial delay in seconds after which the sync action must be executed. The supported range is from 0 to 60 seconds.* `OverriddenList` - The parameter applicable in attach operation indicating the configurations that must override the template configurations. 
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
    + `value`:(string) The action parameter value is based on the action parameter type. Supported action parameters and their values are-a) Name - SyncType, Supported Values - SyncFailed, SyncOne.b) Name - SyncTimer, Supported Values - 0 to 60 seconds.c) Name - OverriddenList, Supported Values - Comma Separated list of overridable configurations. 
  + `type`:(string) The action type to be executed.* `Sync` - The action to merge values from the template to its derived objects.* `Deploy` - The action to execute deploy action on all the objects derived from the template that is mainly applicable for the various profile types.* `Detach` - The action to detach the current derived object from its attached template.* `Attach` - The action to attach the current object to the specified template. 
* `type`:(string) Defines the type of the profile. Accepted values are instance or template.* `instance` - The profile defines the configuration for a specific instance of a target. 
* `update_status`:(string)(ReadOnly) The template sync status with all derived objects.* `None` - The Enum value represents that the object is not attached to any template.* `OK` - The Enum value represents that the object values are in sync with attached template.* `Scheduled` - The Enum value represents that the object sync from attached template is scheduled from template.* `InProgress` - The Enum value represents that the object sync with the attached template is in progress.* `OutOfSync` - The Enum value represents that the object values are not in sync with attached template. 
* `usage`:(int)(ReadOnly) The count of the server profiles derived from the template. 
* `usage_count`:(int)(ReadOnly) The number of objects derived from a Template MO instance. 
* `uuid_address_type`:(string) UUID address allocation type selected to assign an UUID address for the server.* `NONE` - The user did not assign any UUID address.* `STATIC` - The user assigns a static UUID address.* `POOL` - The user selects a pool from which the address will be leased. 
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
  + `marked_for_deletion`:(bool)(ReadOnly) The flag to indicate if snapshot is marked for deletion or not. If flag is set then snapshot will be removed after the successful deployment of the policy. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `ref_mo`:(HashMap) -(ReadOnly) A reference to the original Managed Object. 
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `timestamp`:(string)(ReadOnly) The time this versioned Managed Object was created. 
  + `nr_version`:(string)(ReadOnly) The version of the Managed Object, e.g. an incrementing number or a hash id. 
  + `version_type`:(string)(ReadOnly) Specifies type of version. Currently the only supported value is \ Configured\ that is used to keep track of snapshots of policies and profiles that are intendedto be configured to target endpoints.* `Modified` - Version created every time an object is modified.* `Configured` - Version created every time an object is configured to the service profile.* `Deployed` - Version created for objects related to a service profile when it is deployed. 


## Import
`intersight_server_profile_template` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_server_profile_template.example 1234567890987654321abcde
``` 
