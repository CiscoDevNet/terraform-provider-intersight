---
subcategory: "connectorpack"
layout: "intersight"
page_title: "Intersight: intersight_connectorpack_connector_pack_upgrade"
description: |-
  Used to download or install connector packs on the target device.
---

# Resource: intersight_connectorpack_connector_pack_upgrade
Used to download or install connector packs on the target device.
## Argument Reference
The following arguments are supported:
* `connector_pack_op_type`:(string) The type of operation to be performed on UCS Director.* `Install` - Installs the requisite connector packs on UCS Director.* `Push` - Pushes the requisite connector packs to UCS Director. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `ucsd_info`:(HashMap) - A reference to a iaasUcsdInfo resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `workflow`:(HashMap) -(Computed) A reference to a workflowWorkflowInfo resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 

### Custom keywords
These are
* `wait_for_completion`:(bool) This model object can trigger workflows. Use this option to wait for all running workflows to reach a complete state. Default value is True i.e. wait.

## Import
`intersight_connectorpack_connector_pack_upgrade` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_connectorpack_connector_pack_upgrade.example 1234567890987654321abcde
``` 