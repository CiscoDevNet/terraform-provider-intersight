---
subcategory: "kubernetes"
layout: "intersight"
page_title: "Intersight: intersight_kubernetes_aci_cni_tenant_cluster_allocation"
description: |-
  Internally generated parameter allocations for a k8s cluster using a particular ACI CNI profile.
---

# Resource: intersight_kubernetes_aci_cni_tenant_cluster_allocation
Internally generated parameter allocations for a k8s cluster using a particular ACI CNI profile.
## Argument Reference
The following arguments are supported:
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `node_svc_ip_subnet`:(string)(Computed) CIDR allocated for ACI node service IPs in this tenant cluster. 
* `organization`:(HashMap) - A reference to a organizationOrganization resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `pod_ip_subnet`:(string)(Computed) CIDR allocated for pod IPs in this tenant cluster. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `vlan_end`:(string)(Computed) End of VLAN range allocated to this tenant cluster. 
* `vlan_start`:(string)(Computed) Start of VLAN range allocated to this tenant cluster. 


## Import
`intersight_kubernetes_aci_cni_tenant_cluster_allocation` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_kubernetes_aci_cni_tenant_cluster_allocation.example 1234567890987654321abcde
``` 