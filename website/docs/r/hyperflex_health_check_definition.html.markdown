---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_health_check_definition"
description: |-
  HyperFlex health check definition metadata.
---

# Resource: intersight_hyperflex_health_check_definition
HyperFlex health check definition metadata.
## Usage Example
### Resource Creation

```hcl
resource "intersight_hyperflex_health_check_definition" "hyperflex_health_check_definition1" {
  category                          = "Cluster Resiliency Check"
  common_causes                     = "Persistent drive on one of the nodes is in deny-list or unavailable."
  description                       = "Number of Persistent Drive Failures Tolerable"
  internal_name                     = "NumberOfPersistentDriveFailuresTolerable"
  name                              = "Number of Persistent Drive Failures Tolerable"
  resolution                        = "Check the status of the disk in HX Connect."
  script_execution_mode             = "ON_DEMAND"
  script_execution_on_compute_nodes = false
  target_execution_type             = "EXECUTE_ON_LEADER_NODE"
  timeout                           = 0
}
```
## Argument Reference
The following arguments are supported:
* `account_moid`:(string)(Computed) The Account ID for this managed object. 
* `ancestors`:(Array)(Computed) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `category`:(string) Category that the health check belongs to. 
* `common_causes`:(string) Static information detailing the common causes for the health check failure. 
* `configuration`:(string) Execution configuration fo the health check script. 
* `create_time`:(string)(Computed) The time when this managed object was created. 
* `default_health_check_script_info`:(HashMap) -(Computed) Default version Script info. 
This complex property has following sub-properties:
  + `aggregate_script_name`:(string)(Computed) Health check aggregate script that runs in the HyperFlex Leader Node. |It aggregates the output of all HyperFlex nodes and provides the health check result. 
  + `hyperflex_version`:(string)(Computed) HyperFlex Data Platform version running on the target device. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `script_execute_location`:(string)(Computed) Location of the health check script's execution on the HyperFlex device. 
  + `script_name`:(string)(Computed) Name of the health check script to be executed. 
  + `nr_version`:(string)(Computed) Version of the health check script associated with the health check definition. 
* `description`:(string) Description of the health check definition. 
* `domain_group_moid`:(string)(Computed) The DomainGroup ID for this managed object. 
* `health_check_script_infos`:(Array)
This complex property has following sub-properties:
  + `aggregate_script_name`:(string)(Computed) Health check aggregate script that runs in the HyperFlex Leader Node. |It aggregates the output of all HyperFlex nodes and provides the health check result. 
  + `hyperflex_version`:(string)(Computed) HyperFlex Data Platform version running on the target device. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `script_execute_location`:(string)(Computed) Location of the health check script's execution on the HyperFlex device. 
  + `script_name`:(string)(Computed) Name of the health check script to be executed. 
  + `nr_version`:(string)(Computed) Version of the health check script associated with the health check definition. 
* `health_impact`:(string) Static information detailing the health impact of the health check failure. 
* `internal_name`:(string) Internal name of the health check definition. 
* `mod_time`:(string)(Computed) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the health check definition. 
* `owners`:
                (Array of schema.TypeString) -(Computed)
* `parent`:(HashMap) -(Computed) A reference to a moBaseMo resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `permission_resources`:(Array)(Computed) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `reference`:(string) Static information containing additional reference for the health check. 
* `resolution`:(string) Static information detailing the possible remediation actions that can be taken to remedy the health check failure. 
* `script_execution_mode`:(string) Execution mode of the health script on the HyperFlex cluster.* `ON_DEMAND` - Execute the health check on-demand.* `SCHEDULED` - Execute the health check on a scheduled interval. 
* `script_execution_on_compute_nodes`:(bool) Indicates if the script needs to be executed on HyperFlex compute nodes. |Typically, scripts are only executed on the storage Nodes. 
* `shared_scope`:(string)(Computed) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `target_execution_type`:(string) Indicates whether the health check is executed only on the leader node, or on all nodes in the HyperFlex cluster.* `EXECUTE_ON_LEADER_NODE` - Execute the health check script only on the HyperFlex cluster's leader node.* `EXECUTE_ON_ALL_NODES` - Execute health check on all nodes and aggregate the results.* `EXECUTE_ON_ALL_NODES_AND_AGGREGATE` - Execute the health check on all Nodes and perform custom aggregation. 
* `timeout`:(int) Health check script execution timeout. 
* `unsupported_hyper_flex_versions`:
                (Array of schema.TypeString) -
* `version_context`:(HashMap) -(Computed) The versioning info for this managed object. 
This complex property has following sub-properties:
  + `interested_mos`:(Array)
This complex property has following sub-properties:
    + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `ref_mo`:(HashMap) -(Computed) A reference to the original Managed Object. 
This complex property has following sub-properties:
    + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `timestamp`:(string)(Computed) The time this versioned Managed Object was created. 
  + `nr_version`:(string)(Computed) The version of the Managed Object, e.g. an incrementing number or a hash id. 
  + `version_type`:(string)(Computed) Specifies type of version. Currently the only supported value is \ Configured\ that is used to keep track of snapshots of policies and profiles that are intendedto be configured to target endpoints.* `Modified` - Version created every time an object is modified.* `Configured` - Version created every time an object is configured to the service profile.* `Deployed` - Version created for objects related to a service profile when it is deployed. 


## Import
`intersight_hyperflex_health_check_definition` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_hyperflex_health_check_definition.example 1234567890987654321abcde
``` 