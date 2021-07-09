---
subcategory: "niatelemetry"
layout: "intersight"
page_title: "Intersight: intersight_niatelemetry_apic_core_file_details"
description: |-
  Object to capture Core File details in APIC.
---

# Data Source: intersight_niatelemetry_apic_core_file_details
Object to capture Core File details in APIC.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_niatelemetry_apic_core_file_details.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `annotation`:(string) Annotation of the Core file in APIC. 
* `child_action`:(string) Child action of the Core file in APIC. 
* `collection_time`:(string) Collection Time of the Core file in APIC. 
* `create_time`:(string) The time when this managed object was created. 
* `data_type`:(string) Data type of the Core file in APIC. 
* `dn`:(string) Dn for the Core file in the inventory. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `export_file_uri`:(string) Export file URI of the Core file in APIC. 
* `export_status`:(string) Export status of the Core file in APIC. 
* `export_status_str`:(int) Export status str of the Core file in APIC. 
* `export_tech_sup_file_uri`:(string) Export tech Sup file URI of the Core file in APIC. 
* `exported_to_controller`:(string) Return if file is exported To Controller or not in APIC. 
* `ext_mngd_by`:(string) Ext Mngd By of the Core file in APIC. 
* `file_size`:(int) File size of the Core file in APIC. 
* `host_name`:(string) Host Name of the Core file in APIC. 
* `lc_own`:(string) Lc owner of the Core file in APIC. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `mod_ts`:(int) Mod Ts of the Core file in APIC. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `node_id`:(string) Node Id of the Core file in APIC. 
* `pol_name`:(string) Pol Name of the Core file in APIC. 
* `record_type`:(string) Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. 
* `record_version`:(string) Version of record being pushed. This determines what was the API version for data available from the device. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `site_name`:(string) Name of the APIC site from which this data is being collected. 
* `status`:(string) Status of the Core file in APIC. 
* `uid`:(string) UId of the Core file in the APIC. 
* `userdom`:(string) User dom of the Core file in APIC. 
 