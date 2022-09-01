---
subcategory: "niatelemetry"
layout: "intersight"
page_title: "Intersight: intersight_niatelemetry_hcloud_details"
description: |-
        Inventory object available per device scope. This common object holds a device level information.

---

# Data Source: intersight_niatelemetry_hcloud_details
Inventory object available per device scope. This common object holds a device level information.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_niatelemetry_hcloud_details.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `dn`:(string) Dn for the inventories present. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `epg_count`:(int) Returns the total number of EPGs deployed. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `record_type`:(string) Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. 
* `record_version`:(string) Version of record being pushed. This determines what was the API version for data available from the device. 
* `router_count`:(int) Returns the total number of Cisco Cloud Routers deployed. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `site_name`:(string) Name of the APIC site from which this data is being collected. 
* `subnets_address`:(string) Returns the IP addresses of the subnets. 
* `subnets_count`:(int) Returns the total number of subnets deployed. 
* `transit_gateways_count`:(int) Returns the total number of Transit Gateways deployed. 
* `vpc_count`:(int) Returns the total number of VPCs deployed in Azure/AWS platforms. 
* `vpc_count_gcp`:(int) Returns the total number of VPCs deployed in GCP. 
 
