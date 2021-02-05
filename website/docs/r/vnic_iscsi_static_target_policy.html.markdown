---
subcategory: "vnic"
layout: "intersight"
page_title: "Intersight: intersight_vnic_iscsi_static_target_policy"
description: |-
  Configuration parameters that defines the reachability of iSCSI Target portal.
---

# Resource: intersight_vnic_iscsi_static_target_policy
Configuration parameters that defines the reachability of iSCSI Target portal.
## Argument Reference
The following arguments are supported:
* `description`:(string) Description of the policy. 
* `ip_address`:(string) The IPv4 address assigned to the iSCSI target. 
* `lun`:(HashMap) - The LUN parameters associated with an iSCSI target. 
This complex property has following sub-properties:
  + `bootable`:(bool) Specifies LUN is bootable. 
  + `lun_id`:(int) The Identifier of the LUN. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `organization`:(HashMap) - A reference to a organizationOrganization resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `port`:(int) The port associated with the iSCSI target. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `target_name`:(string) Qualified Name (IQN) or Extended Unique Identifier (EUI) name of the iSCSI target. 


## Import
`intersight_vnic_iscsi_static_target_policy` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_vnic_iscsi_static_target_policy.example 1234567890987654321abcde
``` 