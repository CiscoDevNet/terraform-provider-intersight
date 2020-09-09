
---
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_cluster_network_policy"
sidebar_current: "docs-intersight-resource-hyperflex-cluster-network-policy"
description: |-
  A policy specifying VLANs for management, VM migration, and VM traffic.
---

# Resource: intersight_hyperflex._cluster_network_policy
A policy specifying VLANs for management, VM migration, and VM traffic.
## Argument Reference
The following arguments are supported:
* `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
* `class_id`:(string)(Computed) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `cluster_profiles`:(Array) An array of relationships to hyperflexClusterProfile resources. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string)(Computed) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `description`:(string) Description of the policy. 
* `jumbo_frame`:(bool) Enable or disable jumbo frames. 
* `kvm_ip_range`:(Array with Maximum of one item) - The Out-of-band KVM IP range.Configures the service profiles to use IP addresses within this range for setting the KVM IP of a server. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string)(Computed) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
  + `end_addr`:(string) The end IPv4 address of the range. 
  + `gateway`:(string) The default gateway for the start and end IPv4 addresses. 
  + `netmask`:(string) The netmask specified in dot decimal notation.The start address, end address, and gateway must all be within the network specified by this netmask. 
  + `object_type`:(string) The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types. 
  + `start_addr`:(string) The start IPv4 address of the range. 
* `mac_prefix_range`:(Array with Maximum of one item) - The MAC address prefix range for configuring vNICs.Configures the service profiles to use MAC address prefixes within this range for setting the MAC address of server vNICs. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string)(Computed) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
  + `end_addr`:(string) The end MAC address prefix of a MAC address prefix range in the form of 00:25:B5:XX. 
  + `object_type`:(string) The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types. 
  + `start_addr`:(string) The start MAC address prefix of a MAC address prefix range in the form of 00:25:B5:XX. 
* `mgmt_vlan`:(Array with Maximum of one item) - The VLAN for the management network. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string)(Computed) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
  + `name`:(string) The name of the VLAN.The name can be from 1 to 32 characters long and can contain a combination of alphanumeric characters, underscores, and hyphens. 
  + `object_type`:(string) The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types. 
  + `vlan_id`:(int) The ID of the named VLAN. An ID of 0 means the traffic is untagged.The ID can be any number between 0 and 4095, inclusive. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `object_type`:(string)(Computed) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `organization`:(Array with Maximum of one item) - A reference to a organizationOrganization resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string)(Computed) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `uplink_speed`:(string) Link speed of the server adapter port to the upstream switch. When the policy is attached to a cluster profile with EDGE management platform, the uplink speed can be '1G' or '10G+'. Use '10G+' for link speeds of 10G or above. When the policy is attached to a cluster profile with Fabric Interconnect management platform, the uplink speed can be 'default' only.* `default` - Current default value set on the hardware platform.* `1G` - A link speed of 1 gigabit per second.* `10G` - A link speed of 10 gigabits per second or above. 
* `vm_migration_vlan`:(Array with Maximum of one item) - The VM migration VLAN.This VLAN is used for transfering VMs from one host to another during operations such a cluster upgrade. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string)(Computed) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
  + `name`:(string) The name of the VLAN.The name can be from 1 to 32 characters long and can contain a combination of alphanumeric characters, underscores, and hyphens. 
  + `object_type`:(string) The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types. 
  + `vlan_id`:(int) The ID of the named VLAN. An ID of 0 means the traffic is untagged.The ID can be any number between 0 and 4095, inclusive. 
* `vm_network_vlans`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string)(Computed) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
  + `name`:(string) The name of the VLAN.The name can be from 1 to 32 characters long and can contain a combination of alphanumeric characters, underscores, and hyphens. 
  + `object_type`:(string) The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types. 
  + `vlan_id`:(int) The ID of the named VLAN. An ID of 0 means the traffic is untagged.The ID can be any number between 0 and 4095, inclusive. 
