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
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_niatelemetry_epg.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `azure_pack_count`:(int) Azure Pack NAT with ASA feature usage. 
* `create_time`:(string) The time when this managed object was created. 
* `dn`:(string) Dn value for the End Point Groups present. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `epg_delimiter_count`:(int) Number of  objects with delimiter value present in EPG Delimiter attribute. 
* `fc_npv_count`:(int) Number of ports with FC path attribute of type FC. 
* `fcoe_count`:(string) Number of FCoE per End Point Group. 
* `fv_rs_dom_att_count`:(int) Number of FvRsDomAtt objects per End Point Group with VMware configuration. 
* `intra_epg_dvs_bm_count`:(int) Intra End Point Group Contract for Distributed Virtual Switch and BM feature usage. 
* `intra_epg_hyperv`:(string) Intra EPG Isolation for Hyper-V, enabled if pcEnfPref attribute is set to enforced. 
* `is_attr_based`:(string) Gets the state of End Point Groups with isAttrBasedEPg value as configured. 
* `microsegmentation`:(string) Gets the state of End Point Groups where microsegmentation is present. 
* `microsoft_useg_count`:(int) Number of FvRsDomAtt objects per End Point Group with Microsoft configuration. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name value for the End Point Groups present. 
* `orchsl_dev_vip_cfg_count`:(int) Number of objects with Simplified Service Graph Integration with Windows Azure Pack. 
* `record_type`:(string) Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. 
* `record_version`:(string) Version of record being pushed. This determines what was the API version for data available from the device. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `site_name`:(string) The Site name represents an APIC cluster. Service Engine can onboard multiple APIC clusters / sites. 
* `static_path_attachment_count`:(int) Number of static path Attachments. 
* `useg_hyperv_count`:(int) Logical Operators for attribute based microsegmentation in a hypervisor. 
 