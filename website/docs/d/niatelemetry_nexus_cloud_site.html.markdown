---
subcategory: "niatelemetry"
layout: "intersight"
page_title: "Intersight: intersight_niatelemetry_nexus_cloud_site"
description: |-
        Stores information of Nexus Cloud site devices.

---

# Data Source: intersight_niatelemetry_nexus_cloud_site
Stores information of Nexus Cloud site devices.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_niatelemetry_nexus_cloud_site.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `advisories`:(bool) Advisories setting status, based on license type. 
* `anomalies`:(bool) Anomalies setting status, based on license type. 
* `capacity_utilization`:(bool) Capacity utilization setting status, based on license type. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `license_type`:(string) Returns the license type of the device. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Returns the name of the Nexus Cloud site. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `site_type`:(string) Returns the type of the Nexus Cloud site. 
* `software_management`:(bool) Software management setting status, based on license type. 
* `uuid`:(string) Returns the uuid of the Nexus Cloud site. 
 
