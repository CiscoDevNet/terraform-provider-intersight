---
subcategory: "vnic"
layout: "intersight"
page_title: "Intersight: intersight_vnic_iscsi_boot_policy"
description: |-
  Configuration parameters to enable a server to boot its operating system from an iSCSI target machine located remotely over a network.
---

# Resource: intersight_vnic_iscsi_boot_policy
Configuration parameters to enable a server to boot its operating system from an iSCSI target machine located remotely over a network.
## Argument Reference
The following arguments are supported:
* `auto_targetvendor_name`:(string) Auto target interface that is represented via the Initiator name or the DHCP vendor ID. The vendor ID can be up to 32 alphanumeric characters. 
* `chap`:(HashMap) - CHAP authentication parameters for iSCSI Target. 
This complex property has following sub-properties:
  + `is_password_set`:(bool)(Computed) Indicates whether the value of the 'password' property has been set. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `password`:(string) Password of Initiator/Target Interface. Enter between 12 and 16 characters, including special characters except spaces, tabs, line breaks. 
  + `user_id`:(string) User Id of Initiator/Target Interface. Enter between 1 and 128 characters, spaces, or special characters. 
* `description`:(string) Description of the policy. 
* `initiator_ip_pool`:(HashMap) - A reference to a ippoolPool resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `initiator_ip_source`:(string) Source Type of Initiator IP Address - Auto/Static/Pool.* `None` - Type defines that property is not applicable for an interface.* `Auto` - The system selects an interface automatically - DHCP.* `Static` - Type represents that static information or properties are associated to an interface.* `Pool` - Type defines that property value will be fetched from an associated pool. 
* `initiator_static_ip_v4_address`:(string) Static IP address provided for iSCSI Initiator. 
* `initiator_static_ip_v4_config`:(HashMap) - IPV4 configurations such as Netmask, Gateway and DNS for iSCSI Initiator. 
This complex property has following sub-properties:
  + `gateway`:(string) IP address of the default IPv4 gateway. 
  + `netmask`:(string) A subnet mask is a 32-bit number that masks an IP address and divides the IP address into network address and host address. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `primary_dns`:(string) IP Address of the primary Domain Name System (DNS) server. 
  + `secondary_dns`:(string) IP Address of the secondary Domain Name System (DNS) server. 
* `iscsi_adapter_policy`:(HashMap) - A reference to a vnicIscsiAdapterPolicy resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `mutual_chap`:(HashMap) - Mutual CHAP authentication parameters for iSCSI Initiator. Two-way CHAP mechanism. 
This complex property has following sub-properties:
  + `is_password_set`:(bool)(Computed) Indicates whether the value of the 'password' property has been set. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `password`:(string) Password of Initiator/Target Interface. Enter between 12 and 16 characters, including special characters except spaces, tabs, line breaks. 
  + `user_id`:(string) User Id of Initiator/Target Interface. Enter between 1 and 128 characters, spaces, or special characters. 
* `name`:(string) Name of the concrete policy. 
* `organization`:(HashMap) - A reference to a organizationOrganization resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `primary_target_policy`:(HashMap) - A reference to a vnicIscsiStaticTargetPolicy resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `secondary_target_policy`:(HashMap) - A reference to a vnicIscsiStaticTargetPolicy resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `target_source_type`:(string) Source Type of Targets - Auto/Static.* `None` - Type defines that property is not applicable for an interface.* `Auto` - The system selects an interface automatically - DHCP.* `Static` - Type represents that static information or properties are associated to an interface.* `Pool` - Type defines that property value will be fetched from an associated pool. 


## Import
`intersight_vnic_iscsi_boot_policy` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_vnic_iscsi_boot_policy.example 1234567890987654321abcde
``` 