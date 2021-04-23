---
subcategory: "niaapi"
layout: "intersight"
page_title: "Intersight: intersight_niaapi_apic_hweol"
description: |-
  The hardware end of life notice for APIC.
---

# Data Source: intersight_niaapi_apic_hweol
The hardware end of life notice for APIC.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_niaapi_apic_hweol.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `affected_pids`:(string) String contains the PID of hardwares affected by this notice, seperated by comma. 
* `announcement_date`:(string) When this notice is announced. 
* `announcement_date_epoch`:(int) Epoch time of Announcement Date. 
* `bulletin_no`:(string) The bulletinno of this hardware end of life notice. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) The description of this hardware end of life notice. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `endof_new_service_attachment_date`:(string) Date time of end of new services attachment  . 
* `endof_new_service_attachment_date_epoch`:(int) Epoch time of New service attachment Date . 
* `endof_routine_failure_analysis_date`:(string) Date time of end of routinefailure analysis. 
* `endof_routine_failure_analysis_date_epoch`:(int) Epoch time of End of Routine Failure Analysis Date. 
* `endof_sale_date`:(string) When this hardware will end sale. 
* `endof_sale_date_epoch`:(int) Epoch time of End of Sale Date. 
* `endof_security_support`:(string) Date time of end of security support . 
* `endof_security_support_epoch`:(int) Epoch time of End of Security Support Date . 
* `endof_service_contract_renewal_date`:(string) Date time of end of service contract renew . 
* `endof_service_contract_renewal_date_epoch`:(int) Epoch time of End of Renewal service contract. 
* `endof_sw_maintenance_date`:(string) The date of end of maintainance. 
* `endof_sw_maintenance_date_epoch`:(int) Epoch time of End of maintenance Date. 
* `hardware_eol_url`:(string) Hardware end of sale URL link to the notice webpage. 
* `headline`:(string) The title of this hardware end of life notice. 
* `last_dateof_support`:(string) Date time of end of support . 
* `last_dateof_support_epoch`:(int) Epoch time of last date of support . 
* `last_ship_date`:(string) Date time of Lastship Date. 
* `last_ship_date_epoch`:(int) Epoch time of last ship Date. 
* `migration_url`:(string) The URL of this migration notice. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 