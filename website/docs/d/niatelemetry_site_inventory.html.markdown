---
subcategory: "niatelemetry"
layout: "intersight"
page_title: "Intersight: intersight_niatelemetry_site_inventory"
description: |-
  Details of sites onboarded onto an ND.
---

# Data Source: intersight_niatelemetry_site_inventory
Details of sites onboarded onto an ND.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_niatelemetry_site_inventory.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `firmware_version`:(string) Version of the specified site. 
* `install_type`:(string) Fine-grained type DCNM either SAN or LAN. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the APIC / DCNM site onboarded. 
* `nexus_dashboard`:(string) Name of ND on which site has been onboarded. 
* `nodes`:(int) Number of nodes the site contains. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `type`:(string) Type of site onboarded either APIC or DCNM. 
 