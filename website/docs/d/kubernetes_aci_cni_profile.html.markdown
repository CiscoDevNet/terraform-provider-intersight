---
subcategory: "kubernetes"
layout: "intersight"
page_title: "Intersight: intersight_kubernetes_aci_cni_profile"
description: |-
        Configuration for an ACI CNI profile.

---

# Data Source: intersight_kubernetes_aci_cni_profile
Configuration for an ACI CNI profile.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_kubernetes_aci_cni_profile.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `aaep_name`:(string) Name of ACI AAEP (Attachable Access Entity Profile) to be used for all Kubernetes clusters using this policy. 
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the profile. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `ext_svc_dyn_subnet_start`:(string) Start of range of IP subnets for external services with dynamic IP allocation for use by Kubernetes clusters using this ACI CNI policy. 
* `ext_svc_static_subnet_start`:(string) Start of range of IP subnets for external services with static IP allocation for use by Kubernetes clusters using this ACI CNI policy. 
* `infra_vlan_id`:(int) Value of ACI infrastructuere VLAN ID for the ACI fabric. 
* `l3_out_network_name`:(string) Name of ACI L3Out network to be used for all Kubernetes clusters using this policy. 
* `l3_out_policy_name`:(string) Name of ACI L3Out policy to be used for all Kubernetes clusters using this policy. 
* `l3_out_tenant`:(string) Tenant in ACI used by this L3Out and Common VRF. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the profile instance or profile template. 
* `nested_vmm_domain`:(string) VMM domain within which Kubernetes clusters using this policy are nested. 
* `node_svc_subnet_start`:(string) Start of range of ACI Node Service IP subnets to use by Kubernetes clusters using this ACI CNI policy This is used for the service graph which is used for ACI PBR based load balancing. 
* `node_vlan_range_end`:(int) Ending value of VLAN range used to assign Node VLAN Ids for each Kubernetes cluster using this policy. 
* `node_vlan_range_start`:(int) Starting value of VLAN range used to assign Node VLAN Ids for each Kubernetes cluster using this policy. 
* `number_of_kubernetes_clusters`:(int) Number of k8s clusters currently using this ACI CNI profile. 
* `opflex_multicast_address_range`:(string) Range of IP Multicast addresses to be used by the Opflex protocol for Kubernetes clusters using this policy. 
* `pod_subnet_start`:(string) Start of range of Kubernetes pod IP subnets to use by Kubernetes clusters using this ACI CNI policy This should be a /8 IP subnet so that multiple /16 subnets can be assigned for pod subnets of Kubernetes clusters using this profile. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `svc_subnet_start`:(string) Start of range of Kubernetes Service IP subnets to use by Kubernetes clusters using this ACI CNI policy Currently this is fixed internally and read-only. 
* `type`:(string) Defines the type of the profile. Accepted values are instance or template.* `instance` - The profile defines the configuration for a specific instance of a target. 
* `vrf`:(string) VRF (Virtual Routing and Forwarding) domain to be used within ACI fabric by all k8s clusters using this policy. 
 
