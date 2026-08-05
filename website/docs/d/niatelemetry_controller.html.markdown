---
subcategory: "niatelemetry"
layout: "intersight"
page_title: "Intersight: intersight_niatelemetry_controller"
description: |-
        A Controller managed by a nexus dashboard instance.

---

# Data Source: intersight_niatelemetry_controller
A Controller managed by a nexus dashboard instance.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_niatelemetry_controller.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `anomaly_status`:(string) The anomaly status of the managed Controller. 
* `compliance_status`:(string) The compliance status of the managed Controller. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `fabric_type`:(string) The type of fabric the managed switch belongs to. 
* `hardware_conformance`:(string) The hardware conformance status of the managed Controller. 
* `hardware_end_of_vulnerability_support_date`:(string) The hardware end of vulnerability support date of the managed Controller. 
* `hardware_last_date_of_support`:(string) The hardware last date of support of the managed Controller. 
* `hostname`:(string) The hostname of the managed Controller. 
* `in_band_ip_v4_address`:(string) The inband IPv4 address of the managed Controller. 
* `in_band_ip_v6_address`:(string) The inband IPv6 address of the managed Controller. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `model`:(string) The model of the managed Controller. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `operating_system`:(string) The operating system of the managed Controller. 
* `operating_system_version`:(string) The OS version of the managed Controller. 
* `out_of_band_ip_v4_address`:(string) The out-of-band IPv4 address of the managed Controller. 
* `out_of_band_ip_v6_address`:(string) The out-of-band IPv6 address of the managed Controller. 
* `role`:(string) The role of the managed Controller. 
* `serial`:(string) The serial number of the managed Controller. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `site`:(string) The name of the managed fabric. 
* `software_conformance`:(string) The software conformance status of the managed Controller. 
* `software_end_of_vulnerability_support_date`:(string) The software end of vulnerability support date of the managed Controller. 
* `software_last_date_of_support`:(string) The software last date of support of the managed Controller. 
* `software_version`:(string) The software version of the managed Controller. 
* `strong_asset_id`:(string) The DBID of the managed Controller. 
* `telemetry_enabled`:(bool) Indicates whether telemetry is enabled for the fabric. 
 
