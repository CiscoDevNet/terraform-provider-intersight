
---
layout: "intersight"
page_title: "Intersight: intersight_inventory_dn_mo_binding"
sidebar_current: "docs-intersight-data-source-inventory-dn-mo-binding"
description: |-
DnMoBinding provides a binding between a Intersight MO and a UCSM MO which has a DN.
---

# Data Source: intersight_inventory._dn_mo_binding
DnMoBinding provides a binding between a Intersight MO and a UCSM MO which has a DN.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `dn`:(string) The Distinguished Name for this object, used to uniquely identify this object. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `target_mo_id`:(string) The MO ID of the target MO for this particular Distinguished Name (dn). 
* `target_mo_type`:(string) The type of the target MO for this particular Distinguished Name (dn). 
