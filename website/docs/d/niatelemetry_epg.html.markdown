
---
layout: "intersight"
page_title: "Intersight: intersight_niatelemetry_epg"
sidebar_current: "docs-intersight-data-source-niatelemetry-epg"
description: |-
Object is available at End Point Group scope. This currently applies only to the APIC environemt.
---

# Data Source: intersight_niatelemetry._epg
Object is available at End Point Group scope. This currently applies only to the APIC environemt.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `azure_pack_count`:(int) Azure Pack NAT with ASA feature usage. 
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `dn`:(string) Dn value for the End Point Groups present. 
* `epg_delimiter_count`:(int) EPG Delimiter scale where the delimiter value is present. 
* `fc_npv_count`:(int) Number of ports with FC path attribute of type FC. 
* `fcoe_count`:(string) Number of FCoE per End Point Group. 
* `fv_rs_dom_att_count`:(int) FvRsDomAtt scale per End Point Group with VMware configured. 
* `intra_epg_dvs_bm_count`:(int) Intra End Point Group Contract for Distributed Virtual Switch and BM feature usage. 
* `intra_epg_hyperv`:(string) Intra EPG Isolation for Hyper-V, enabled if pcEnfPref attribute is set to enforced. 
* `is_attr_based`:(string) Gets the state of End Point Groups with isAttrBasedEPg value as configured. 
* `microsegmentation`:(string) Gets the state of End Point Groups where microsegmentation is present. 
* `microsoft_useg_count`:(int) FvRsDomAtt scale per End Point Group with Microsoft configured. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name value for the End Point Groups present. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `orchsl_dev_vip_cfg_count`:(int) Simplified Service Graph Integration with Windows Azure Pack count scale. 
* `site_name`:(string) The Site name represents an APIC cluster. Service Engine can onboard multiple APIC clusters / sites. 
* `useg_hyperv_count`:(int) Logical Operators for attribute based microsegmentation in a hypervisor. 
