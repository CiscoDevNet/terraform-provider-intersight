---
subcategory: "recommendation"
layout: "intersight"
page_title: "Intersight: intersight_recommendation_software_item"
description: |-
        Entity representing the software recommendation for the managed end point.

---

# Data Source: intersight_recommendation_software_item
Entity representing the software recommendation for the managed end point.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_recommendation_software_item.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `message`:(string) The user visible message which informs user of the type of software recommendation. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name of the physical device recommended. 
* `personality`:(string) The personality of the physical device recommended. 
* `recommendation_type`:(string) The software-recommendation type, for example, HXDP version, HyperV or ESXi version, etc.* `None` - The Enum value None represents the default software recommendation value.* `HXDPVersion` - The Enum value HXDPVersion represents that the software recommendation is to upgrade the HyperFlex Data Platform build version.* `NodeRatioLicense` - The Enum value NodeRatioLicense represents that the software recommendation is to upgrade the HyperFlex Data Platform license for using 1:2 converged to compute node ratio limits.* `DCNoFILicense` - The Enum value DCNoFILicense represents that the software recommendation is to upgrade the HyperFlex Data Platform license for using DC-No-FI limits.* `LAZExistingStatus` - The Enum value LAZExistingStatus represents that the software recommendation indicates that the HyperFlex cluster might have LAZ enabled.* `LAZNewStatus` - The Enum value LAZNewStatus represents that the software recommendation is to enable LAZ with expansion on the HyperFlex Cluster.* `EVCStatus` - The Enum value EVCStatus represents that the software recommendation is to enable Enhanced VMotion on the HypeFlex Cluster. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `type`:(string) The type of physical device recommended.* `None` - The Enum value None represents that no value was set for the item type.* `Disk` - The Enum value Disk represents that the item type recommended is a Physical Disk.* `Node` - The Enum value Node represents that the item type recommended is a Storage Node.* `CPU` - The Enum value CPU represents that the item type recommended is a Processor.* `Memory` - The Enum value Memory represents that the item type recommended is a memory unit.* `Cluster` - The Enum value Cluster represents that the item type recommended is a HyperFlex Cluster. 
 
