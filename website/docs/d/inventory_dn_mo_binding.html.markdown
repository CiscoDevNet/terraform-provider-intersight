---
subcategory: "inventory"
layout: "intersight"
page_title: "Intersight: intersight_inventory_dn_mo_binding"
description: |-
  DnMoBinding provides a binding between a Intersight MO and a UCSM MO which has a DN.
---

# Data Source: intersight_inventory_dn_mo_binding
DnMoBinding provides a binding between a Intersight MO and a UCSM MO which has a DN.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `dn`:(string) The Distinguished Name for this object, used to uniquely identify this object. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `target_mo_id`:(string) The MO ID of the target MO for this particular Distinguished Name (dn). 
* `target_mo_type`:(string) The type of the target MO for this particular Distinguished Name (dn). 
