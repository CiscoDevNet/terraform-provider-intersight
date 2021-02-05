---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_cluster_replication_network_policy_deployment"
description: |-
  Record of HyperFlex Cluster replication network policy deployment.
---

# Resource: intersight_hyperflex_cluster_replication_network_policy_deployment
Record of HyperFlex Cluster replication network policy deployment.
## Argument Reference
The following arguments are supported:
* `cluster`:(HashMap) - A reference to a hyperflexCluster resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `cluster_uuid`:(string)(Computed) Uuid of the HyperFlex cluster. 
* `description`:(string)(Computed) Description from corresponding ClusterReplicationNetworkPolicy. 
* `discovered`:(bool) True if record created by discovery on HyperFlex cluster. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string)(Computed) Name from corresponding ClusterReplicationNetworkPolicy. 
* `organization`:(HashMap) - A reference to a organizationOrganization resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `policy_moid`:(string)(Computed) Deployed network policy moid. 
* `profile_moid`:(string)(Computed) Deployed cluster profile moid. 
* `replication_bandwidth_mbps`:(int)(Computed) Bandwidth for the Replication network in Mbps. 
* `replication_ipranges`:(Array)
This complex property has following sub-properties:
  + `end_addr`:(string) The end IPv4 address of the range. 
  + `gateway`:(string) The default gateway for the start and end IPv4 addresses. 
  + `netmask`:(string) The netmask specified in dot decimal notation.The start address, end address, and gateway must all be within the network specified by this netmask. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `start_addr`:(string) The start IPv4 address of the range. 
* `replication_mtu`:(int)(Computed) MTU for the Replication network. 
* `replication_vlan`:(HashMap) -(Computed) VLAN for the Replication network. 
This complex property has following sub-properties:
  + `name`:(string) The name of the VLAN.The name can be from 1 to 32 characters long and can contain a combination of alphanumeric characters, underscores, and hyphens. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `vlan_id`:(int) The ID of the named VLAN. An ID of 0 means the traffic is untagged.The ID can be any number between 0 and 4095, inclusive. 
* `request_id`:(string)(Computed) Unique request ID allowing retry of the same logical request following a transient communication failure. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 


## Import
`intersight_hyperflex_cluster_replication_network_policy_deployment` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_hyperflex_cluster_replication_network_policy_deployment.example 1234567890987654321abcde
``` 