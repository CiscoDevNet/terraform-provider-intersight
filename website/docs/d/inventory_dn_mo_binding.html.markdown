
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
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `dn`:(string) The Distinguished Name for this object, used to uniquely identify this object. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `target_mo_id`:(string) The MO ID of the target MO for this particular Distinguished Name (dn). 
* `target_mo_type`:(string) The type of the target MO for this particular Distinguished Name (dn). 
