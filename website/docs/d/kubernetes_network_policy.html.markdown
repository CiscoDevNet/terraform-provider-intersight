---
subcategory: "kubernetes"
layout: "intersight"
page_title: "Intersight: intersight_kubernetes_network_policy"
description: |-
  A policy specifying the CIDR for internal networks in a Kubernetes cluster like POD network, and Service network.
---

# Data Source: intersight_kubernetes_network_policy
A policy specifying the CIDR for internal networks in a Kubernetes cluster like POD network, and Service network.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `cni_type`:(string) Supported CNI type. Currently we only support Calico.* `Calico` - Calico CNI plugin as described in https://github.com/projectcalico/cni-plugin.* `Aci` - Cisco ACI Container Network Interface plugin. 
* `description`:(string) Description of the policy. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `pod_network_cidr`:(string) CIDR block to allocate pod network IP addresses from. 
* `service_cidr`:(string) CIDR block to allocate cluster service IP addresses from. 
