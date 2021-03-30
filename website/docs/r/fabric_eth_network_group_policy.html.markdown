---
subcategory: "fabric"
layout: "intersight"
page_title: "Intersight: intersight_fabric_eth_network_group_policy"
description: |-
  The allowed VLAN/s on an interface.
---

# Resource: intersight_fabric_eth_network_group_policy
The allowed VLAN/s on an interface.
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
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `vlan_settings`:(HashMap) - VLAN configuration for the virtual interface. 
This complex property has following sub-properties:
  + `allowed_vlans`:(string) Allowed VLAN IDs of the virtual interface. 
  + `native_vlan`:(int) Native VLAN ID of the virtual interface or the corresponding vethernet on the peer Fabric Interconnect to which the virtual interface is connected. If the native VLAN is not a part of the allowed VLANs, it will automatically be added to the list of allowed VLANs. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 


## Import
`intersight_fabric_eth_network_group_policy` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_fabric_eth_network_group_policy.example 1234567890987654321abcde
``` 