---
subcategory: "sdcard"
layout: "intersight"
page_title: "Intersight: intersight_sdcard_policy"
description: |-
  Policy for configuring SD Card settings on endpoint.
---

# Resource: intersight_sdcard_policy
Policy for configuring SD Card settings on endpoint.
## Argument Reference
The following arguments are supported:
* `description`:(string) Description of the policy. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `organization`:(HashMap) - A reference to a organizationOrganization resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `partitions`:(Array)
This complex property has following sub-properties:
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `type`:(string) This specifies the type of the partition. Allowed values are OS, Utility.* `OS` - This partition contains virtual drives where user can install operating system.* `Utility` - This partition contains virtual drives for utilities such as SCU, HUU, Drivers and Diagnostics. 
  + `virtual_drives`:(Array)
This complex property has following sub-properties:
    + `additional_properties`:(JSON) - Additional Properties as per object type, can be added as JSON using `jsonencode()`. Allowed Types are: [sdcard.Diagnostics](#sdcardDiagnostics)
[sdcard.Drivers](#sdcardDrivers)
[sdcard.HostUpgradeUtility](#sdcardHostUpgradeUtility)
[sdcard.OperatingSystem](#sdcardOperatingSystem)
[sdcard.ServerConfigurationUtility](#sdcardServerConfigurationUtility)
[sdcard.UserPartition](#sdcardUserPartition)
    + `enable`:(bool) Enable the respective virtual drive to be available to the host. 
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `profiles`:(Array) An array of relationships to policyAbstractConfigProfile resources. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 

## Usage Example
### Resource Creation
```hcl
resource "intersight_sdcard_policy" "sdcard1" {
  name        = "sdcard1"
  description = "test policy"
  organization {
    object_type = "organization.Organization"
    moid = var.organization
  }
  partitions {
    type        = "OS"
    object_type = "sdcard.Partition"

    virtual_drives {
      enable      = true
      object_type = "sdcard.OperatingSystem"
      additional_properties = jsonencode({
        Name = "Hypervisor"
      })
    }
  }

  profiles {
    moid        = intersight_server_profile.server1.id
    object_type = "server.Profile"
  }
}
```

## Import
`intersight_sdcard_policy` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_sdcard_policy.example 1234567890987654321abcde
``` 