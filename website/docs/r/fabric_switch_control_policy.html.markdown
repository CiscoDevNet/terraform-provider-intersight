---
subcategory: "fabric"
layout: "intersight"
page_title: "Intersight: intersight_fabric_switch_control_policy"
description: |-
  A policy to configure the Switching Mode, Port VLAN Optimization, MAC Aging Time, Reserved VLAN Range of the FI.
---

# Resource: intersight_fabric_switch_control_policy
A policy to configure the Switching Mode, Port VLAN Optimization, MAC Aging Time, Reserved VLAN Range of the FI.
## Argument Reference
The following arguments are supported:
* `description`:(string) Description of the policy. 
* `mac_aging_settings`:(HashMap) - This specifies the MAC aging option and time settings. 
This complex property has following sub-properties:
  + `mac_aging_option`:(string) This specifies one of the option to configure the MAC address aging time.* `Default` - This option sets the default MAC address aging time to 14500 seconds for End Host mode.* `Custom` - This option allows the the user to configure the MAC address aging time on the switch. For Switch Model UCS-FI-6454 or higher, the valid range is 120 to 918000 seconds and the switch will set the lower multiple of 5 of the given time.* `Never` - This option disables the MAC address aging process and never allows the MAC address entries to get removed from the table. 
  + `mac_aging_time`:(int) Define the MAC address aging time in seconds. This field is valid when the \ Custom\  MAC address aging option is selected. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `organization`:(HashMap) - A reference to a organizationOrganization resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `profiles`:(Array) An array of relationships to fabricSwitchProfile resources. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `vlan_port_optimization_enabled`:(bool) To enable or disable the VLAN port count optimization. 


## Import
`intersight_fabric_switch_control_policy` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_fabric_switch_control_policy.example 1234567890987654321abcde
``` 