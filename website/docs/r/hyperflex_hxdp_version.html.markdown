---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_hxdp_version"
description: |-
  A HyperFlex Data Platform version.
---

# Resource: intersight_hyperflex_hxdp_version
A HyperFlex Data Platform version.
## Argument Reference
The following arguments are supported:
* `app_catalog`:(HashMap) - A reference to a hyperflexAppCatalog resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `nr_version`:(string) The HyperFlex Data Platform version. 


## Import
`intersight_hyperflex_hxdp_version` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_hyperflex_hxdp_version.example 1234567890987654321abcde
``` 