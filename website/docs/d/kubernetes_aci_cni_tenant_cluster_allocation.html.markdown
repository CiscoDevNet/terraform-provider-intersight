---
subcategory: "kubernetes"
layout: "intersight"
page_title: "Intersight: intersight_kubernetes_aci_cni_tenant_cluster_allocation"
description: |-
  Internally generated parameter allocations for a k8s cluster using a particular ACI CNI profile.
---

# Data Source: intersight_kubernetes_aci_cni_tenant_cluster_allocation
Internally generated parameter allocations for a k8s cluster using a particular ACI CNI profile.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `node_svc_ip_subnet`:(string) CIDR allocated for ACI node service IPs in this tenant cluster. 
* `pod_ip_subnet`:(string) CIDR allocated for pod IPs in this tenant cluster. 
* `vlan_end`:(string) End of VLAN range allocated to this tenant cluster. 
* `vlan_start`:(string) Start of VLAN range allocated to this tenant cluster. 
