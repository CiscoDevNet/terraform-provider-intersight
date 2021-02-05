---
subcategory: "chassis"
layout: "intersight"
page_title: "Intersight: intersight_chassis_config_import"
description: |-
  Configuration import action will import the existing configuration from chassis and generate Intersight policies profile from it. At end of successful import, chassis will be assigned to the generated profile which has policies associated with it. No chassis profile or policies will be generated if configuration import fails.
---

# Resource: intersight_chassis_config_import
Configuration import action will import the existing configuration from chassis and generate Intersight policies profile from it. At end of successful import, chassis will be assigned to the generated profile which has policies associated with it. No chassis profile or policies will be generated if configuration import fails.
## Argument Reference
The following arguments are supported:
* `chassis`:(HashMap) - A reference to a equipmentChassis resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `chassis_profile`:(HashMap) -(Computed) A reference to a chassisProfile resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `description`:(string) Description of the imported profile. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `organization`:(HashMap) - A reference to a organizationOrganization resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `policy_prefix`:(string) Policy prefix for the policies of the imported chassis profile. 
* `policy_types`:
                (Array of schema.TypeString) -
* `profile_name`:(string) Profile name for the imported chassis profile. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 


## Import
`intersight_chassis_config_import` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_chassis_config_import.example 1234567890987654321abcde
``` 