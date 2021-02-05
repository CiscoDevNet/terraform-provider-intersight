---
subcategory: "vnic"
layout: "intersight"
page_title: "Intersight: intersight_vnic_iscsi_adapter_policy"
description: |-
  Set of iSCSI properties that govern the host-side behavior of the adapter.
---

# Resource: intersight_vnic_iscsi_adapter_policy
Set of iSCSI properties that govern the host-side behavior of the adapter.
## Argument Reference
The following arguments are supported:
* `connection_time_out`:(int) The number of seconds to wait until Cisco UCS assumes that the initial login has failed and the iSCSI adapter is unavailable. 
* `description`:(string) Description of the policy. 
* `dhcp_timeout`:(int) The number of seconds to wait before the initiator assumes that the DHCP server is unavailable. 
* `lun_busy_retry_count`:(int) The number of times to retry the connection in case of a failure during iSCSI LUN discovery. 
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


## Import
`intersight_vnic_iscsi_adapter_policy` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_vnic_iscsi_adapter_policy.example 1234567890987654321abcde
``` 