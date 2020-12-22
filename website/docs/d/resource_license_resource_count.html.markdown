---
subcategory: "resource"
layout: "intersight"
page_title: "Intersight: intersight_resource_license_resource_count"
description: |-
  LicenseResourceCount tracks the server count info for 3 different licensing tiers.
---

# Data Source: intersight_resource_license_resource_count
LicenseResourceCount tracks the server count info for 3 different licensing tiers.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `license_type`:(string) Type of licensing defined for this resource group. Used for licensing group.* `Base` - Base as a License type. It is default license type.* `Essential` - Essential as a License type.* `Standard` - Standard as a License type.* `Advantage` - Advantage as a License type.* `Premier` - Premier as a License type.* `IWO-Essential` - IWO-Essential as a License type.* `IWO-Advantage` - IWO-Advantage as a License type.* `IWO-Premier` - IWO-Premier as a License type. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `resource_count`:(int) The number of resource belongs to this licensing tier. 
