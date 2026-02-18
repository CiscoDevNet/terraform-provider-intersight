---
subcategory: "resource"
layout: "intersight"
page_title: "Intersight: intersight_resource_license_resource_count"
description: |-
        The LicenseResourceCount object tracks the number of resources associated with different licensing tiers, aiding in licensing management and resource allocation.
        #### Purpose
        LicenseResourceCount provides a structured approach to managing licensing information within the system, ensuring that resources are accurately accounted for and associated with appropriate licensing tiers. It supports efficient resource tracking and licensing compliance.
        #### Key Concepts
        - **Licensing Management:** Facilitates the tracking and management of resources based on licensing tiers, aiding in compliance and resource allocation.
        - **Resource Tracking:** Provides accurate counts of resources within each licensing tier, supporting efficient resource management.
        - **Integration:** Seamlessly integrates with resource groups to ensure that licensing information is properly managed and utilized.

---

# Data Source: intersight_resource_license_resource_count
The LicenseResourceCount object tracks the number of resources associated with different licensing tiers, aiding in licensing management and resource allocation.
#### Purpose
LicenseResourceCount provides a structured approach to managing licensing information within the system, ensuring that resources are accurately accounted for and associated with appropriate licensing tiers. It supports efficient resource tracking and licensing compliance.
#### Key Concepts
- **Licensing Management:** Facilitates the tracking and management of resources based on licensing tiers, aiding in compliance and resource allocation.
- **Resource Tracking:** Provides accurate counts of resources within each licensing tier, supporting efficient resource management.
- **Integration:** Seamlessly integrates with resource groups to ensure that licensing information is properly managed and utilized.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_resource_license_resource_count.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `license_type`:(string) Type of licensing defined for this resource group. Used for licensing group.* `Base` - Base as a License type. It is default license type.* `Essential` - Essential as a License type.* `Standard` - Standard as a License type.* `Advantage` - Advantage as a License type.* `Premier` - Premier as a License type.* `IWO-Essential` - IWO-Essential as a License type.* `IWO-Advantage` - IWO-Advantage as a License type.* `IWO-Premier` - IWO-Premier as a License type.* `IKS-Advantage` - IKS-Advantage as a License type.* `INC-Premier-1GFixed` - Premier 1G Fixed license tier for Intersight Nexus Cloud.* `INC-Premier-10GFixed` - Premier 10G Fixed license tier for Intersight Nexus Cloud.* `INC-Premier-100GFixed` - Premier 100G Fixed license tier for Intersight Nexus Cloud.* `INC-Premier-Mod4Slot` - Premier Modular 4 slot license tier for Intersight Nexus Cloud.* `INC-Premier-Mod8Slot` - Premier Modular 8 slot license tier for Intersight Nexus Cloud.* `INC-Premier-D2OpsFixed` - Premier D2Ops fixed license tier for Intersight Nexus Cloud.* `INC-Premier-D2OpsMod` - Premier D2Ops modular license tier for Intersight Nexus Cloud.* `INC-Premier-CentralizedMod8Slot` - Premier modular license tier of switch type CentralizedMod8Slot for Intersight Nexus Cloud.* `INC-Premier-DistributedMod8Slot` - Premier modular license tier of switch type DistributedMod8Slot for Intersight Nexus Cloud.* `ERP-Advantage` - Advantage license tier for ERP workflows.* `IntersightTrial` - Virtual dummy license type to indicate trial. Used for UI display of trial mode Intersight tiers.* `IWOTrial` - Virtual dummy license type to indicate trial. Used for UI display of trial mode IKS tiers.* `IKSTrial` - Virtual dummy license type to indicate trial. Used for UI display of trial mode IWO tiers.* `INCTrial` - Virtual dummy license type to indicate trial. Used for UI display of trial mode Nexus tiers. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `resource_count`:(int) The number of resource belongs to this licensing tier. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
