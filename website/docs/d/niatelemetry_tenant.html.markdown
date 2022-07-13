---
subcategory: "niatelemetry"
layout: "intersight"
page_title: "Intersight: intersight_niatelemetry_tenant"
description: |-
        Object is available at Tenant scope.

---

# Data Source: intersight_niatelemetry_tenant
Object is available at Tenant scope.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_niatelemetry_tenant.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `bfd_if_pol_count`:(int) Number of Bidirectional Forwarding Detection bfdIfPol Model Objects. 
* `bfd_ifp_count`:(int) Number of objects with Bidirectional Forwarding Detection Interface Policy. 
* `create_time`:(string) The time when this managed object was created. 
* `dhcp_rs_prov_count`:(int) Number of tenants with Dynamic Host Configuration Protocol enabled. 
* `dn`:(string) Dn for the tenants present. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `fhs_bd_pol_count`:(int) Number of objects with First hop security. Checks for presence of IP source gaurd, dynamic arp inspection. 
* `fv_ap_count`:(int) Number of application profiles per tenant. 
* `fv_bd_count`:(int) Number of bridge domains with hardware proxy enabled per tenant. 
* `fv_bd_subnet_count`:(int) Number of bridge domain subnets per tenant. 
* `fv_bdno_arp_count`:(int) Number of bridge domains with ARP flodding disabled per tenant. 
* `fv_cep_count`:(int) Count of number of endpoints per tenant. 
* `fv_rs_bd_to_fhs_count`:(int) Number of objects with First hop security. Checks for presence of IP source gaurd, dynamic arp inspection. 
* `fv_rs_bd_to_out_count`:(int) Number of bridge domains connected to layer 3 out per tenant. 
* `fv_site_connp_count`:(int) Number of Multi-Site objects. 
* `fv_subnet_count`:(int) Number of subnets per tenant. 
* `ip_static_route_count`:(int) Number of IP static routes per tenant. 
* `l3_multicast_count`:(int) Number of layer 3 multicasts. 
* `l3_multicast_ctx_count`:(int) Number of layer 3 multicast CtxP. 
* `l3_multicast_if_count`:(int) Number of layer 3 multicast IfP. 
* `l3out_count`:(int) Number of L3 out objects for the tenants present. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `qos_custom_pol_count`:(int) Number of Quality Of Service Custom Policy. 
* `record_type`:(string) Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. 
* `record_version`:(string) Version of record being pushed. This determines what was the API version for data available from the device. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `site_name`:(string) The Site name represents an APIC cluster. Service Engine can onboard multiple APIC clusters / sites. 
* `ssm`:(string) SSM property feature usage. 
* `ssm_count`:(int) Number of context-level ssm translate policies per tenant. 
* `tcam_opt_count`:(int) Number of TCAM optimization enabled per tenant. 
* `tenant_health`:(int) Health of each tenant in the APIC. 
* `trace_route_ep_count`:(int) Number of ITrace route endpoint per tenant. 
* `trace_route_ep_ext_count`:(int) Number of ITrace endpoint external routes per tenant. 
* `trace_route_ext_ep_count`:(int) Number of ITrace external endpoint routes per tenant. 
* `trace_route_ext_ext_count`:(int) Number of ITrace external routes per tenant. 
* `vns_abs_graph_count`:(int) Number of objects with L4 to L7 Services graph. 
* `vns_backup_pol_count`:(int) Number of objects with Policy Based Routing standby Node. Checks the Policy Based Routing backup policy. 
* `vns_redirect_dest_count`:(int) Number of objects with Policy Based Routing standby Node. Policy based redirect requires a destination to redirect traffic. 
* `vns_svc_redirect_pol_count`:(int) Number of Policy Based Routing and Policy Based Service Insertion objects. Includes without L4-L7 package. 
* `vrf_count`:(int) Number of Vrfs per tenant. 
* `vz_br_cp_count`:(int) Number of Zoning Policy per tenant. 
* `vz_rt_cons_count`:(int) Number of Client Contract between End Point Groups per tenant. 
* `vz_rt_prov_count`:(int) Number of Client Contract between End Point Groups per tenant. 
 
