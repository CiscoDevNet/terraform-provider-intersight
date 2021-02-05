---
subcategory: "vnic"
layout: "intersight"
page_title: "Intersight: intersight_vnic_eth_network_policy"
description: |-
  An Ethernet Network policy determines if the port can carry single VLAN (Access) or multiple VLANs (Trunk) traffic. You can specify the VLAN to be associated with an Ethernet packet if no tag is found.
---

# Resource: intersight_vnic_eth_network_policy
An Ethernet Network policy determines if the port can carry single VLAN (Access) or multiple VLANs (Trunk) traffic. You can specify the VLAN to be associated with an Ethernet packet if no tag is found.
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
* `target_platform`:(string) The platform for which the server profile is applicable. It can either be a server that is operating in standalone mode or which is attached to a Fabric Interconnect managed by Intersight.* `Standalone` - Servers which are operating in standalone mode i.e. not connected to a Fabric Interconnected.* `FIAttached` - Servers which are connected to a Fabric Interconnect that is managed by Intersight. 
* `vlan_settings`:(HashMap) - VLAN configuration for the virtual interface. 
This complex property has following sub-properties:
  + `allowed_vlans`:(string) Allowed VLAN IDs of the virtual interface. 
  + `default_vlan`:(int) Native VLAN ID of the virtual interface or the corresponding vethernet on the peer Fabric Interconnect to which the virtual interface is connected. Setting the ID to 0 will not associate any native VLAN to the traffic on the virtual interface. 
  + `mode`:(string) Option to determine if the port can carry single VLAN (Access) or multiple VLANs (Trunk) traffic.* `ACCESS` - An access port carries traffic only for a single VLAN on the interface.* `TRUNK` - A trunk port can have two or more VLANs configured on the interface. It can carry traffic for several VLANs simultaneously. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 

## Usage Example
### Resource Creation
```hcl
resource "intersight_vnic_eth_network_policy" "v_eth_network1" {
  name = "v_eth_network1"
  organization {
    object_type = "organization.Organization"
    moid = var.organization
  }
  vlan_settings {
    object_type = "vnic.VlanSettings"
    default_vlan = 1
    mode = "ACCESS"
  }
}
```

## Import
`intersight_vnic_eth_network_policy` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_vnic_eth_network_policy.example 1234567890987654321abcde
``` 