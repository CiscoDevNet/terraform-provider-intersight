---
subcategory: "iam"
layout: "intersight"
page_title: "Intersight: intersight_iam_account_experience"
description: |-
  The beta features enabled for the specified account.
---

# Resource: intersight_iam_account_experience
The beta features enabled for the specified account.
## Argument Reference
The following arguments are supported:
* `account`:(HashMap) - A reference to a iamAccount resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `features`:(Array)
This complex property has following sub-properties:
  + `feature`:(string) The beta feature that will be enabled for specific account.* `IWO` - Intersight Workflow Optimizer.* `Hitachi` - Support to claim Hitachi Storage arrays using the Intersight Orchestrator framework.* `Kubernetes` - Enables ability to create and manage Kubernetes clusters.* `NetAppIO` - Support to claim NetApp Storage arrays as IO targets.* `IvsPublicCloud` - Enables virtualization service for public clouds.* `TerraformCloud` - Enables an ability to create Terraform Cloud. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 


## Import
`intersight_iam_account_experience` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_iam_account_experience.example 1234567890987654321abcde
``` 