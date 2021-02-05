---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_ext_fc_storage_policy"
description: |-
  A policy specifying external storage connectivity information via Fabric attached FC storage.
---

# Resource: intersight_hyperflex_ext_fc_storage_policy
A policy specifying external storage connectivity information via Fabric attached FC storage.
## Argument Reference
The following arguments are supported:
* `admin_state`:(bool) Enables or disables external FC storage configuration. 
* `cluster_profiles`:(Array) An array of relationships to hyperflexClusterProfile resources. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `description`:(string) Description of the policy. 
* `exta_traffic`:(HashMap) - VSAN for the primary Fabric Interconnect external FC storage traffic. 
This complex property has following sub-properties:
  + `name`:(string) The name of the VSAN.The name can be from 1 to 32 characters long and can contain a combination of alphanumeric characters, underscores, and hyphens. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `vsan_id`:(int) The ID of the named VSAN.The ID can be any number between 1 and 4093, inclusive. 
* `extb_traffic`:(HashMap) - VSAN for the secondary Fabric Interconnect external FC storage traffic. 
This complex property has following sub-properties:
  + `name`:(string) The name of the VSAN.The name can be from 1 to 32 characters long and can contain a combination of alphanumeric characters, underscores, and hyphens. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `vsan_id`:(int) The ID of the named VSAN.The ID can be any number between 1 and 4093, inclusive. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `organization`:(HashMap) - A reference to a organizationOrganization resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `wwxn_prefix_range`:(HashMap) - The range of WWxN addresses to use for the FC storage configuration. 
This complex property has following sub-properties:
  + `end_addr`:(string) The end WWxN prefix of a WWPN/WWNN range in the form of 20:00:00:25:B5:XX. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `start_addr`:(string) The start WWxN prefix of a WWPN/WWNN range in the form of 20:00:00:25:B5:XX. 


## Import
`intersight_hyperflex_ext_fc_storage_policy` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_hyperflex_ext_fc_storage_policy.example 1234567890987654321abcde
``` 