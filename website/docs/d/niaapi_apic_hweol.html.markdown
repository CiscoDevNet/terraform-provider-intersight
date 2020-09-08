
---
layout: "intersight"
page_title: "Intersight: intersight_niaapi_apic_hweol"
sidebar_current: "docs-intersight-data-source-niaapi-apic-hweol"
description: |-
The hardware end of life notice for APIC.
---

# Data Source: intersight_niaapi._apic_hweol
The hardware end of life notice for APIC.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `affected_pids`:(string) String contains the PID of hardwares affected by this notice, seperated by comma. 
* `announcement_date`:(string) When this notice is announced. 
* `announcement_date_epoch`:(int) Epoch time of Announcement Date. 
* `bulletin_no`:(string) The bulletinno of this hardware end of life notice. 
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `description`:(string) The description of this hardware end of life notice. 
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
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
