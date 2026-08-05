---
subcategory: "mgmt"
layout: "intersight"
page_title: "Intersight: intersight_mgmt_config_restore_operation"
description: |-
        The ConfigRestoreOperation object represents an administrator-initiated configuration restore run. It tracks the end-to-end process of applying a previously captured backup to a target account, including progress, outcomes, and workflow coordination.
        #### Purpose
        ConfigRestoreOperation exists to orchestrate restoration of configuration from a chosen backup source, providing a top-level record that drives restore execution, tracks status, and reports import progress across account and organization scopes.
        #### Key Concepts
        - **End-to-End Restore Orchestration:** Coordinates staging, importing, and completion of a restore run, enabling customers to manage restoration as a single operation.
        - **Source Flexibility:** Designed to support restoring from different backup locations (e.g., local repository), depending on operational needs.
        - **Controlled Restore Behavior:** Supports configurable restore strategies and error-handling behavior so customers can align restores with change-management requirements.
        - **Progress and Outcome Visibility:** Aggregates import status reporting so customers can monitor completion and detect partial or incomplete outcomes.
        - **Workflow Integration:** Links to underlying workflow execution details to support operational auditability and troubleshooting.

---

# Resource: intersight_mgmt_config_restore_operation
The ConfigRestoreOperation object represents an administrator-initiated configuration restore run. It tracks the end-to-end process of applying a previously captured backup to a target account, including progress, outcomes, and workflow coordination.
#### Purpose
ConfigRestoreOperation exists to orchestrate restoration of configuration from a chosen backup source, providing a top-level record that drives restore execution, tracks status, and reports import progress across account and organization scopes.
#### Key Concepts
- **End-to-End Restore Orchestration:** Coordinates staging, importing, and completion of a restore run, enabling customers to manage restoration as a single operation.
- **Source Flexibility:** Designed to support restoring from different backup locations (e.g., local repository), depending on operational needs.
- **Controlled Restore Behavior:** Supports configurable restore strategies and error-handling behavior so customers can align restores with change-management requirements.
- **Progress and Outcome Visibility:** Aggregates import status reporting so customers can monitor completion and detect partial or incomplete outcomes.
- **Workflow Integration:** Links to underlying workflow execution details to support operational auditability and troubleshooting.
## Argument Reference
The following arguments are supported:
* `account`:(HashMap) - A reference to a iamAccount resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `account_moid`:(string)(ReadOnly) The Account ID for this managed object. 
* `acct_category_import_data`:(Array)
This complex property has following sub-properties:
  + `file_name`:(string)(ReadOnly) Name of the file that contains configurations for this category. 
  + `name`:(string)(ReadOnly) Name of the backup category. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `processed_types`:(Array)
This complex property has following sub-properties:
    + `child_types`:(Array)
This complex property has following sub-properties:
(Cyclic reference - this type contains a self-referencing property)
  + `filter`:(string)(ReadOnly) Filter used to obtain a list of objects of the given object type. 
  + `name`:(string)(ReadOnly) The type of the managed object. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `rest_path`:(string)(ReadOnly) The rest path for the MO type. 
  + `status`:(string)(ReadOnly) Import status for this category of objects.* `Scheduled` - Backup or restore process has been scheduled.* `NotInitiated` - Backup or restore process has not been initiated.* `InProgress` - Backup or restore process is in progress.* `Failed` - Backup or restore process has failed.* `Completed` - Backup or restore process has completed.* `NotComplete` - Backup or restore process has partially imported configuration. 
* `acct_import_jobs`:(Array)(ReadOnly) An array of relationships to bulkRequest resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `ancestors`:(Array)(ReadOnly) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `backup_account_moid`:(string)(ReadOnly) The moid of the account from which the backup was taken. 
* `backup_name`:(string)(ReadOnly) Name of the backup that needs to be restored. 
* `backup_source`:(string)(ReadOnly) Indicates whether the backup was created locally or imported.* `Local` - A local Intersight location.* `Uploaded` - A local location where the backup file is uploaded as mgmt.ConfigBackupFile MO. Intersight creates the backup from the uploaded location when this location type is set.* `Remote` - A remote location hosted in the user's datacenter. 
* `create_time`:(string)(ReadOnly) The time when this managed object was created. 
* `domain_group_moid`:(string)(ReadOnly) The DomainGroup ID for this managed object. 
* `end_time`:(string)(ReadOnly) End date and time of the restore process. 
* `ignore_secure_properties`:(bool) If set to true, secure properties will not be decrypted, and empty values will be used for these properties. As a result, if a secure property field is mandatory, the MO will be marked as incomplete. 
* `is_aes_key_set`:(bool)(ReadOnly) Indicates whether the value of the 'aesKey' property has been set. 
* `local_source`:(HashMap) - A reference to a mgmtConfigBackupInstance resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `mod_time`:(string)(ReadOnly) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) User provided identifier for the restore operation. 
* `options`:(HashMap) - Options to control the behavior of the restore operation. 
This complex property has following sub-properties:
  + `action_on_error`:(string) The action to be taken if an error is encountered during the restore operation.* `Continue` - The restore operation will continue on encountering an error.Any failure will be tagged as incomplete and the operation will continue importing other configurations.User shall review the partially imported configuration and fix later.* `Stop` - The restore operation will stop on encountering an error.Configurations that were changed as part of the restore operation till the error was encountered will remain in the system.* `Rollback` - Configurations that were changed as part of the restore operation till the error was encountered will be rolled back. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `orgs_to_restore`:
                (Array of schema.TypeString) -
  + `restore_action`:(string) Mechanism to be used for updating configurations as part of a restore operation.* `Merge` - Configurations in the backup will be merged with existing configurations in the target system.* `Replace` - Configurations in the target system will be replaced with the configurations in the backup.Configurations that are present only in the target system and not in the backup will be removed. 
* `org_import_jobs`:(Array)(ReadOnly) An array of relationships to bulkRequest resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `org_restores`:(Array)(ReadOnly) An array of relationships to mgmtOrgRestoreOperation resources. 
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
* `shared_scope`:(string)(ReadOnly) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `nr_source`:(string) The type of source from which the backup - local or remote - needs to be restored.* `Local` - A local Intersight location.* `Uploaded` - A local location where the backup file is uploaded as mgmt.ConfigBackupFile MO. Intersight creates the backup from the uploaded location when this location type is set.* `Remote` - A remote location hosted in the user's datacenter. 
* `start_time`:(string)(ReadOnly) Start date and time of the restore process. 
* `status`:(string)(ReadOnly) Status of the restore operation.* `Scheduled` - Backup or restore process has been scheduled.* `NotInitiated` - Backup or restore process has not been initiated.* `InProgress` - Backup or restore process is in progress.* `Failed` - Backup or restore process has failed.* `Completed` - Backup or restore process has completed.* `NotComplete` - Backup or restore process has partially imported configuration. 
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
* `workflow_info`:(HashMap) -(ReadOnly) A reference to a workflowWorkflowInfo resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 

### Custom keywords
These are
* `wait_for_completion`:(bool) This model object can trigger workflows. Use this option to wait for all running workflows to reach a complete state. Default value is True i.e. wait.

## Import
`intersight_mgmt_config_restore_operation` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_mgmt_config_restore_operation.example 1234567890987654321abcde
``` 
