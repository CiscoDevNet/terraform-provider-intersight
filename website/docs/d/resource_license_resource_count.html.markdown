
---
layout: "intersight"
page_title: "Intersight: intersight_resource_license_resource_count"
sidebar_current: "docs-intersight-data-source-resource-license-resource-count"
description: |-
LicenseResourceCount tracks the server count info for 3 different licensing tiers.
---

# Data Source: intersight_resource._license_resource_count
LicenseResourceCount tracks the server count info for 3 different licensing tiers.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `license_type`:(string) Type of licensing defined for this resource group. Used for licensing group.* `Base` - Base as a License type. It is default license type.* `Essential` - Essential as a License type.* `Standard` - Standard as a License type.* `Advantage` - Advantage as a License type.* `Premier` - Premier as a License type.* `IWO-Essential` - IWO-Essential as a License type.* `IWO-Advantage` - IWO-Advantage as a License type.* `IWO-Premier` - IWO-Premier as a License type. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `resource_count`:(int) The number of resource belongs to this licensing tier. 
