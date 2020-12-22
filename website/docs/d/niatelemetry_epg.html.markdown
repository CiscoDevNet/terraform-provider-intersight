---
subcategory: "niatelemetry"
layout: "intersight"
page_title: "Intersight: intersight_niatelemetry_epg"
description: |-
  Object is available at End Point Group scope. This currently applies only to the APIC environemt.
---

# Data Source: intersight_niatelemetry_epg
Object is available at End Point Group scope. This currently applies only to the APIC environemt.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `azure_pack_count`:(int) Azure Pack NAT with ASA feature usage. 
* `dn`:(string) Dn value for the End Point Groups present. 
* `epg_delimiter_count`:(int) Number of  objects with delimiter value present in EPG Delimiter attribute. 
* `fc_npv_count`:(int) Number of ports with FC path attribute of type FC. 
* `fcoe_count`:(string) Number of FCoE per End Point Group. 
* `fv_rs_dom_att_count`:(int) Number of FvRsDomAtt objects per End Point Group with VMware configuration. 
* `intra_epg_dvs_bm_count`:(int) Intra End Point Group Contract for Distributed Virtual Switch and BM feature usage. 
* `intra_epg_hyperv`:(string) Intra EPG Isolation for Hyper-V, enabled if pcEnfPref attribute is set to enforced. 
* `is_attr_based`:(string) Gets the state of End Point Groups with isAttrBasedEPg value as configured. 
* `microsegmentation`:(string) Gets the state of End Point Groups where microsegmentation is present. 
* `microsoft_useg_count`:(int) Number of FvRsDomAtt objects per End Point Group with Microsoft configuration. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name value for the End Point Groups present. 
* `orchsl_dev_vip_cfg_count`:(int) Number of objects with Simplified Service Graph Integration with Windows Azure Pack. 
* `record_type`:(string) Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. 
* `record_version`:(string) Version of record being pushed. This determines what was the API version for data available from the device. 
* `site_name`:(string) The Site name represents an APIC cluster. Service Engine can onboard multiple APIC clusters / sites. 
* `useg_hyperv_count`:(int) Logical Operators for attribute based microsegmentation in a hypervisor. 
