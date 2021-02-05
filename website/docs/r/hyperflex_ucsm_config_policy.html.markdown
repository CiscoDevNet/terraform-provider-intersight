---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_ucsm_config_policy"
description: |-
  A policy specifying UCS Manager settings such as service profile org, KVM IP addresses, and MAC prefix for server configuration in Fabric Interconnect environment.
---

# Resource: intersight_hyperflex_ucsm_config_policy
A policy specifying UCS Manager settings such as service profile org, KVM IP addresses, and MAC prefix for server configuration in Fabric Interconnect environment.
## Argument Reference
The following arguments are supported:
* `cluster_profiles`:(Array) An array of relationships to hyperflexClusterProfile resources. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `description`:(string) Description of the policy. 
* `kvm_ip_range`:(HashMap) - The Out-of-band KVM IP range.Configures the service profiles to use IP addresses within this range for setting the KVM IP of a server. 
This complex property has following sub-properties:
  + `end_addr`:(string) The end IPv4 address of the range. 
  + `gateway`:(string) The default gateway for the start and end IPv4 addresses. 
  + `netmask`:(string) The netmask specified in dot decimal notation.The start address, end address, and gateway must all be within the network specified by this netmask. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `start_addr`:(string) The start IPv4 address of the range. 
* `mac_prefix_range`:(HashMap) - The MAC address prefix range for configuring vNICs.Configures the service profiles to use MAC address prefixes within this range for setting the MAC address of server vNICs. 
This complex property has following sub-properties:
  + `end_addr`:(string) The end MAC address prefix of a MAC address prefix range in the form of 00:25:B5:XX. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `start_addr`:(string) The start MAC address prefix of a MAC address prefix range in the form of 00:25:B5:XX. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `organization`:(HashMap) - A reference to a organizationOrganization resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `server_firmware_version`:(string) The server firmware bundle version used for server components such as CIMC, adapters, BIOS, etc. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 


## Import
`intersight_hyperflex_ucsm_config_policy` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_hyperflex_ucsm_config_policy.example 1234567890987654321abcde
``` 