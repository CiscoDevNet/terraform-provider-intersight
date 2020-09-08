
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
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `license_type`:(string) Type of licensing defined for this resource group. Used for licensing group.* `Base` - Base as a License type. It is default license type.* `Essential` - Essential as a License type.* `Standard` - Standard as a License type.* `Advantage` - Advantage as a License type.* `Premier` - Premier as a License type. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `resource_count`:(int) The number of resource belongs to this licensing tier. 
