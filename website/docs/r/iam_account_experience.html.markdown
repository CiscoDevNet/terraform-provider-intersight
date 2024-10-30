---
subcategory: "iam"
layout: "intersight"
page_title: "Intersight: intersight_iam_account_experience"
description: |-
        The beta features enabled for the specified account.

---

# Resource: intersight_iam_account_experience
The beta features enabled for the specified account.
## Usage Example
### Resource Creation

```hcl
resource "intersight_iam_account_experience" "iam_account_experience1" {
  account {
    object_type = "iam.Account"
    moid        = var.iam_account
  }
  features {
    class_id    = "iam.FeatureDefinition"
    feature     = "UnrecognizedValue"
    object_type = "iam.FeatureDefinition"
  }
  parent {
    moid        = var.iam_account
    object_type = "iam.Account"
  }
}

variable "iam_account" {
  type        = string
  description = "value for iam_account"
}
```
## Argument Reference
The following arguments are supported:
* `account`:(HashMap) - A reference to a iamAccount resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `account_moid`:(string)(ReadOnly) The Account ID for this managed object. 
* `ancestors`:(Array)(ReadOnly) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `create_time`:(string)(ReadOnly) The time when this managed object was created. 
* `domain_group_moid`:(string)(ReadOnly) The DomainGroup ID for this managed object. 
* `features`:(Array)
This complex property has following sub-properties:
  + `feature`:(string) The beta feature that will be enabled for specific account.* `IWO` - Intersight Workflow Optimizer.* `Hitachi` - Support to claim Hitachi Storage arrays using the Intersight Orchestrator framework.* `KubernetesExtension` - Extension to the IKS and Adopted Clusters.* `NetAppIO` - Support to claim NetApp Storage arrays as IO targets.* `IvsPublicCloud` - Enables virtualization service for public clouds.* `TerraformCloud` - Enables an ability to create Terraform Cloud.* `WashingtonEFT` - Support for EFT customers to use Washington firmware images for upgrades.* `Solutions` - Support for managing solutions.* `IksBm` - Enables Intersight Kubernetes Service on Baremetal server.* `NexusCloud` - Enables Nexus Cloud services functionality.* `NexusCloudTrial` - Enables Nexus Cloud trial period.* `NexusCloudUpgradeAssist` - Enables Nexus Cloud upgrade assist.* `NexusCloudSustainability` - Enables Nexus Cloud sustainability.* `PlatformUIRefresh` - Enables platform refreshed UI with new service launcher.* `IksGpu` - Enables GPU support for Intersight Kubernetes Service.* `IwoAppServiceVerticalScaling` - Enables vertical Scaling of App Service Plans.* `IwoDataExporter` - Enables IWO Data Exporter component.* `IwoMigrate` - Enables IWO data Migration.* `NexusCloudTechPreviewGold` - Enable Nexus Cloud Preview of stable features, available for public consumption.* `NexusCloudTechPreviewSilver` - Enable Nexus Cloud Preview of beta features. This feature set is intended for consumption by internal audiences.* `NexusCloudTechPreviewBronze` - Enable Nexus Cloud Preview of features still in development. This feature set is intended for consumption by internal audiences.* `DisconnectedTargetAlarm` - Raise an alarm when a Target is disconnected from Intersight. Intersight is unable to manage disconnected Targets.* `AsAService` - Enable AsAService Preview of beta features. This feature set is intended for consumption by selective audiences.* `EMEA` - Enable all avaialble features on Intersight EMEA region.* `CrossPlatformNavigation` - Enable Cross-Platform Navigation on UI.* `WorkflowsPreview` - Enable Workflows preview for this account.* `ManualKEKSupport` - Enable support for Manual key to enhance storage controller security for this account.* `FirmwareConsolidationEFT` - Enable usage of firmware images transitioned from CCO to Intersight. This feature set is intended for consumption by internal audiences. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `mod_time`:(string)(ReadOnly) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
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
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
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


## Import
`intersight_iam_account_experience` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_iam_account_experience.example 1234567890987654321abcde
``` 
