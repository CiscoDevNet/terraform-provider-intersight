---
subcategory: "inventory"
layout: "intersight"
page_title: "Intersight: intersight_inventory_dn_mo_binding"
description: |-
        DnMoBindings provide a mapping between an Intersight managed object and a UCSM managed object that is uniquely identified by a Distinguished Name (DN). This binding enables correlation between Intersight MO identities and UCSM DN-based identities for the same underlying entity.
        #### Purpose
        Enable reliable cross-system correlation by binding UCSM DN identifiers to the corresponding Intersight target MO identifiers and types.
        #### Key Concepts
        - **DN-based identity mapping:** Stores the UCSM Distinguished Name (`dn`) as the key identifier for the binding.
        - **Target MO reference:** Captures the Intersight target MO identity (`targetMoId`) and classification (`targetMoType`) associated with that DN.
        - **Device context:** Optionally associated with a `registeredDevice` to scope bindings to the correct UCSM/endpoint context, with cascade cleanup on device deletion.
        - **Read-only reference data:** Exposed via READ for lookup/correlation workflows rather than as a mutable configuration object.

---

# Data Source: intersight_inventory_dn_mo_binding
DnMoBindings provide a mapping between an Intersight managed object and a UCSM managed object that is uniquely identified by a Distinguished Name (DN). This binding enables correlation between Intersight MO identities and UCSM DN-based identities for the same underlying entity.
#### Purpose
Enable reliable cross-system correlation by binding UCSM DN identifiers to the corresponding Intersight target MO identifiers and types.
#### Key Concepts
- **DN-based identity mapping:** Stores the UCSM Distinguished Name (`dn`) as the key identifier for the binding.
- **Target MO reference:** Captures the Intersight target MO identity (`targetMoId`) and classification (`targetMoType`) associated with that DN.
- **Device context:** Optionally associated with a `registeredDevice` to scope bindings to the correct UCSM/endpoint context, with cascade cleanup on device deletion.
- **Read-only reference data:** Exposed via READ for lookup/correlation workflows rather than as a mutable configuration object.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_inventory_dn_mo_binding.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `dn`:(string) The Distinguished Name for this object, used to uniquely identify this object. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `target_mo_id`:(string) The MO ID of the target MO for this particular Distinguished Name (dn). 
* `target_mo_type`:(string) The type of the target MO for this particular Distinguished Name (dn). 
 
