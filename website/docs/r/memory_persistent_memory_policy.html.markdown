---
subcategory: "memory"
layout: "intersight"
page_title: "Intersight: intersight_memory_persistent_memory_policy"
description: |-
        The Persistent Memory policy defines the reusable Persistent Memory related configuration that can be applied on many servers. This policy allows the application of Persistent Memory Goals and creation of Persistent Memory Regions and Namespaces. The encryption of the Persistent Memory Modules can be enabled through this policy by providing a passphrase.

---

# Resource: intersight_memory_persistent_memory_policy
The Persistent Memory policy defines the reusable Persistent Memory related configuration that can be applied on many servers. This policy allows the application of Persistent Memory Goals and creation of Persistent Memory Regions and Namespaces. The encryption of the Persistent Memory Modules can be enabled through this policy by providing a passphrase.
## Usage Example
### Resource Creation

```hcl
resource "intersight_memory_persistent_memory_policy" "memory_persistent_memory_policy1" {
  name        = "memory_persistent_memory_policy1"
  description = "memory persistent memory policies"
  goals = [{
    object_type            = "memory.PersistentMemoryGoal"
    persistent_memory_type = "app-direct-non-interleaved"
    socket_id              = "All Sockets"
    additional_properties  = ""
    class_id               = "memory.PersistentMemoryGoal"
    memory_mode_percentage = 50
  }]
  local_security {
    object_type       = "memory.PersistentMemoryLocalSecurity"
    enabled           = true
    secure_passphrase = "ChangeMe"
  }
  logical_namespaces = [
    {
      additional_properties = ""
      class_id              = "memory.PersistentMemoryLocalSecurity"
      name                  = "logical_namespace_test"
      capacity              = 131072
      object_type           = "memory.PersistentMemoryLocalSecurity"
      mode                  = "block"
      socket_id             = 1
      socket_memory_id      = 6
    }
  ]
  management_mode   = "configured-from-intersight"
  retain_namespaces = false
  organization {
    object_type = "organization.Organization"
    moid        = var.organization
  }
}
```
## Argument Reference
The following arguments are supported:
* `account_moid`:(string)(ReadOnly) The Account ID for this managed object. 
* `ancestors`:(Array)(ReadOnly) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `create_time`:(string)(ReadOnly) The time when this managed object was created. 
* `description`:(string) Description of the policy. 
* `domain_group_moid`:(string)(ReadOnly) The DomainGroup ID for this managed object. 
* `goals`:(Array)
This complex property has following sub-properties:
  + `memory_mode_percentage`:(int) Volatile memory percentage. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `persistent_memory_type`:(string) Type of the Persistent Memory configuration where the Persistent Memory Modules are combined in an interleaved set or not.* `app-direct` - The App Direct interleaved Persistent Memory type.* `app-direct-non-interleaved` - The App Direct non-interleaved Persistent Memory type. 
  + `socket_id`:(string) CPU Socket ID to which this goal will be applied.* `All Sockets` - All the CPU socket IDs in a server. 
* `local_security`:(HashMap) - Local security for the Persistent Memory Modules on the server. 
This complex property has following sub-properties:
  + `enabled`:(bool) Persistent Memory Security state. 
  + `is_secure_passphrase_set`:(bool)(ReadOnly) Indicates whether the value of the 'securePassphrase' property has been set. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `secure_passphrase`:(string) Secure passphrase to be applied on the Persistent Memory Modules on the server. The allowed characters are a-z, A to Z, 0-9, and special characters =, \\u0021, &, \\#, $, %, +, ^, @, _, *, -. 
* `logical_namespaces`:(Array)
This complex property has following sub-properties:
  + `capacity`:(int) Capacity of this Namespace that is created or modified. 
  + `mode`:(string) Mode of this Namespace that is created or modified.* `raw` - The raw mode of Persistent Memory Namespace.* `block` - The block mode of Persistent Memory Namespace. 
  + `name`:(string) Name of this Namespace to be created on the server. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `socket_id`:(int) Socket ID of the region on which this Namespace has to be created or modified.* `1` - The first CPU socket in a server.* `2` - The second CPU socket in a server.* `3` - The third CPU socket in a server.* `4` - The fourth CPU socket in a server. 
  + `socket_memory_id`:(string) Socket Memory ID of the region on which this Namespace has to be created or modified.* `Not Applicable` - The socket memory ID is not applicable if app-direct persistent memory type is selected in the goal.* `2` - The second socket memory ID within a socket in a server.* `4` - The fourth socket memory ID within a socket in a server.* `6` - The sixth socket memory ID within a socket in a server.* `8` - The eighth socket memory ID within a socket in a server.* `10` - The tenth socket memory ID within a socket in a server.* `12` - The twelfth socket memory ID within a socket in a server. 
* `management_mode`:(string) Management Mode of the policy. This can be either Configured from Intersight or Configured from Operating System.* `configured-from-intersight` - The Persistent Memory Modules are configured from Intersight thorugh Persistent Memory policy.* `configured-from-operating-system` - The Persistent Memory Modules are configured from operating system thorugh OS tools. 
* `mod_time`:(string)(ReadOnly) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
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
* `profiles`:(Array) An array of relationships to policyAbstractConfigProfile resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `retain_namespaces`:(bool) Persistent Memory Namespaces to be retained or not. 
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
`intersight_memory_persistent_memory_policy` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_memory_persistent_memory_policy.example 1234567890987654321abcde
``` 
