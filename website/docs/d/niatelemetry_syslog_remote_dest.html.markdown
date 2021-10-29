---
subcategory: "niatelemetry"
layout: "intersight"
page_title: "Intersight: intersight_niatelemetry_syslog_remote_dest"
description: |-
  Object to capture Syslog remote dest details.
---

# Data Source: intersight_niatelemetry_syslog_remote_dest
Object to capture Syslog remote dest details.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_niatelemetry_syslog_remote_dest.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `admin_state`:(string) Admin state of Syslog remote dest in APIC. 
* `common_policy`:(string) Parent common policy for syslog src in APIC. 
* `create_time`:(string) The time when this managed object was created. 
* `dn`:(string) Dn of the Syslog remote dest in APIC. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `host`:(string) Host of Syslog remote dest in APIC. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `record_type`:(string) Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. 
* `record_version`:(string) Version of record being pushed. This determines what was the API version for data available from the device. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `site_name`:(string) Name of the APIC site from which this data is being collected. 
* `syslog_rs_dest_grp`:(string) Dn of sys log src dest grp in APIC. 
* `syslog_src`:(string) Dn of parent syslog src for the syslog dest grp in APIC. 
 