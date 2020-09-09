
---
layout: "intersight"
page_title: "Intersight: intersight_niatelemetry_tenant"
sidebar_current: "docs-intersight-data-source-niatelemetry-tenant"
description: |-
Object is available at Tenant scope. This currently applies only to the APIC environemt.
---

# Data Source: intersight_niatelemetry._tenant
Object is available at Tenant scope. This currently applies only to the APIC environemt.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `bfd_if_pol_count`:(int) Bidirectional Forwarding Detection bfdIfPol Model Object count scale. 
* `bfd_ifp_count`:(int) Bidirectional Forwarding Detection Interface Policy count scale. 
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `dn`:(string) Dn for the tenants present. 
* `fhs_bd_pol_count`:(int) First hop security count scale. Checks for presence of IP source gaurd, dynamic arp inspection. 
* `fv_rs_bd_to_fhs_count`:(int) First hop security count scale. Checks for presence of IP source gaurd, dynamic arp inspection. 
* `fv_site_connp_count`:(int) Multi-Site scale for fvSiteConnp Model Object. 
* `l3_multicast_count`:(int) Number of layer 3 multicasts. 
* `l3_multicast_ctx_count`:(int) Number of layer 3 multicast CtxP. 
* `l3_multicast_if_count`:(int) Number of layer 3 multicast IfP. 
* `l3out_count`:(int) L3 out scale for the tenants present. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `site_name`:(string) The Site name represents an APIC cluster. Service Engine can onboard multiple APIC clusters / sites. 
* `ssm`:(string) SSM property feature usage. 
* `trace_route_ep_count`:(int) Number of ITrace route endpoint per tenant. 
* `trace_route_ep_ext_count`:(int) Number of ITrace endpoint external routes per tenant. 
* `trace_route_ext_ep_count`:(int) Number of ITrace external endpoint routes per tenant. 
* `trace_route_ext_ext_count`:(int) Number of ITrace external routes per tenant. 
* `vns_abs_graph_count`:(int) L4 to L7 Services graph count scale. 
* `vns_backup_pol_count`:(int) Policy Based Routing standby Node count scale. Checks the Policy Based Routing backup policy. 
* `vns_redirect_dest_count`:(int) Policy Based Routing standby Node count scale. Policy based redirect requires a destination to redirect traffic. 
* `vns_svc_redirect_pol_count`:(int) Policy Based Routing and Policy Based Service Insertion count scale. Includes without L4-L7 package. 
* `vrf_count`:(int) Vrf scale count per tenant. 
