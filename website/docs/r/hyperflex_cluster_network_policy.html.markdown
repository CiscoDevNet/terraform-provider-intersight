---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_cluster_network_policy"
description: |-
  A policy specifying VLANs for management, VM migration, and VM traffic.
---

# Resource: intersight_hyperflex_cluster_network_policy
A policy specifying VLANs for management, VM migration, and VM traffic.
## Usage Example
### Resource Creation

```hcl
resource "intersight_hyperflex_cluster_network_policy" "hyperflex_cluster_network_policy1" {
  mgmt_vlan {
    name    = "hx-inband-mgmt"
    vlan_id = 301
  }
  jumbo_frame  = true
  uplink_speed = "10G"
  mac_prefix_range {
    end_addr   = "00:25:B5:D5"
    start_addr = "00:25:B5:D5"
  }
  organization {
    object_type = "organization.Organization"
    moid        = var.organization_organization
  }
  name = "hyperflex_cluster_network_policy1"
}
```
## Argument Reference
The following arguments are supported:
* `account_moid`:(string)(Computed) The Account ID for this managed object. 
* `ancestors`:(Array)(Computed) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `cluster_profiles`:(Array) An array of relationships to hyperflexClusterProfile resources. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `create_time`:(string)(Computed) The time when this managed object was created. 
* `description`:(string) Description of the policy. 
* `domain_group_moid`:(string)(Computed) The DomainGroup ID for this managed object. 
* `jumbo_frame`:(bool) Enable or disable jumbo frames. 
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
* `mgmt_vlan`:(HashMap) - The VLAN for the management network. 
This complex property has following sub-properties:
  + `name`:(string) The name of the VLAN.The name can be from 1 to 32 characters long and can contain a combination of alphanumeric characters, underscores, and hyphens. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `vlan_id`:(int) The ID of the named VLAN. An ID of 0 means the traffic is untagged.The ID can be any number between 0 and 4095, inclusive. 
* `mod_time`:(string)(Computed) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `organization`:(HashMap) - A reference to a organizationOrganization resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `owners`:
                (Array of schema.TypeString) -(Computed)
* `parent`:(HashMap) -(Computed) A reference to a moBaseMo resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `permission_resources`:(Array)(Computed) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `shared_scope`:(string)(Computed) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `uplink_speed`:(string) Link speed of the server adapter port to the upstream switch. When the policy is attached to a cluster profile with EDGE management platform, the uplink speed can be '1G' or '10G+'. Use '10G+' for link speeds of 10G or above. When the policy is attached to a cluster profile with Fabric Interconnect management platform, the uplink speed can be 'default' only.* `default` - Current default value set on the hardware platform.* `1G` - A link speed of 1 gigabit per second.* `10G` - A link speed of 10 gigabits per second or above. 
* `version_context`:(HashMap) -(Computed) The versioning info for this managed object. 
This complex property has following sub-properties:
  + `interested_mos`:(Array)
This complex property has following sub-properties:
    + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `ref_mo`:(HashMap) -(Computed) A reference to the original Managed Object. 
This complex property has following sub-properties:
    + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `timestamp`:(string)(Computed) The time this versioned Managed Object was created. 
  + `nr_version`:(string)(Computed) The version of the Managed Object, e.g. an incrementing number or a hash id. 
  + `version_type`:(string)(Computed) Specifies type of version. Currently the only supported value is \ Configured\ that is used to keep track of snapshots of policies and profiles that are intendedto be configured to target endpoints.* `Modified` - Version created every time an object is modified.* `Configured` - Version created every time an object is configured to the service profile.* `Deployed` - Version created for objects related to a service profile when it is deployed. 
* `vm_migration_vlan`:(HashMap) - The VM migration VLAN.This VLAN is used for transfering VMs from one host to another during operations such a cluster upgrade.For HyperFlex Application Platform clusters, this VLAN is also used for hypervisor control traffic such asnode to node communication, pod-to-pod communication, etc. 
This complex property has following sub-properties:
  + `name`:(string) The name of the VLAN.The name can be from 1 to 32 characters long and can contain a combination of alphanumeric characters, underscores, and hyphens. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `vlan_id`:(int) The ID of the named VLAN. An ID of 0 means the traffic is untagged.The ID can be any number between 0 and 4095, inclusive. 
* `vm_network_vlans`:(Array)
This complex property has following sub-properties:
  + `name`:(string) The name of the VLAN.The name can be from 1 to 32 characters long and can contain a combination of alphanumeric characters, underscores, and hyphens. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `vlan_id`:(int) The ID of the named VLAN. An ID of 0 means the traffic is untagged.The ID can be any number between 0 and 4095, inclusive. 


## Import
`intersight_hyperflex_cluster_network_policy` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_hyperflex_cluster_network_policy.example 1234567890987654321abcde
``` 