---
subcategory: "macpool"
layout: "intersight"
page_title: "Intersight: intersight_macpool_lease"
description: |-
        Leases represent individual MAC addresses allocated from a MAC pool (dynamic allocation) or obtained via static assignment within the MAC universe. A lease is the “in-use” record that ties a MAC address to an owning entity such as a server profile.
        #### Purpose
        Track allocation of MAC addresses, enforce uniqueness, support reserved-identity allocation, and provide an ownership link between a MAC address and the entity consuming it.
        #### Key Concepts
        - **Allocated identity record**: `macAddress` is the leased MAC and is create-only, ensuring stable identity once issued.
        - **Reservation-aware allocation**: `reservation` can be provided to allocate an already-reserved identity and carry reservation details/conditions.
        - **Preferred MAC behavior (dynamic only)**: `preferredMacAddress` allows best-effort selection during dynamic requests; if unavailable/out of range/reserved/leased, the next available MAC is allocated.
        - **Migration support (dynamic only)**: When used with a migrate behavior (not shown here but referenced), an existing lease can be replaced.
        - **Ownership linkage**: `assignedToEntity` links the lease to the consuming managed object (e.g., a server profile), enabling traceability and cleanup workflows.
        - **Universe containment**: `universe` ties the lease to the account-wide MAC bookkeeping container (cascade on universe deletion).
        - **Pool selection at creation**: `pool` is create-only, reflecting that the lease is requested/issued against a specific pool.

---

# Data Source: intersight_macpool_lease
Leases represent individual MAC addresses allocated from a MAC pool (dynamic allocation) or obtained via static assignment within the MAC universe. A lease is the “in-use” record that ties a MAC address to an owning entity such as a server profile.
#### Purpose
Track allocation of MAC addresses, enforce uniqueness, support reserved-identity allocation, and provide an ownership link between a MAC address and the entity consuming it.
#### Key Concepts
- **Allocated identity record**: `macAddress` is the leased MAC and is create-only, ensuring stable identity once issued.
- **Reservation-aware allocation**: `reservation` can be provided to allocate an already-reserved identity and carry reservation details/conditions.
- **Preferred MAC behavior (dynamic only)**: `preferredMacAddress` allows best-effort selection during dynamic requests; if unavailable/out of range/reserved/leased, the next available MAC is allocated.
- **Migration support (dynamic only)**: When used with a migrate behavior (not shown here but referenced), an existing lease can be replaced.
- **Ownership linkage**: `assignedToEntity` links the lease to the consuming managed object (e.g., a server profile), enabling traceability and cleanup workflows.
- **Universe containment**: `universe` ties the lease to the account-wide MAC bookkeeping container (cascade on universe deletion).
- **Pool selection at creation**: `pool` is create-only, reflecting that the lease is requested/issued against a specific pool.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_macpool_lease.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `allocation_type`:(string) Type of the lease allocation either static or dynamic (i.e via pool).* `dynamic` - Identifiers to be allocated by system.* `static` - Identifiers are assigned by the user. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `has_duplicate`:(bool) HasDuplicate represents if there are other pools in which this id exists. 
* `mac_address`:(string) MAC address allocated for pool-based allocation. 
* `migrate`:(bool) The migration capability is applicable only for dynamic lease requests and it works in conjunction with  preferred ID. If there is an existing dynamic or static lease that matches the preferred ID, that existing  lease will be migrated to the current pool. That means the existing lease will be deleted and a new lease  will be created in the pool. If there is a reservation exists that matches with preferred ID, that  reservation will be kept as is and next available ID from the pool will be leased. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `preferred_mac_address`:(string) The preferred MAC address can be specified only for dynamic lease requests. Intersight will make its best  effort to allocate that MAC address if it is available in the pool. If the specified preferred MAC address  is not in the range of the pool or if it is already leased or reserved, then the next available MAC address  from the pool will be leased. Since this feature is specific to dynamic lease requests only, static lease  request will fail if it specifies the preferred MAC address property. When the preferred MAC address  property is specified in conjunction with 'migrate' property, existing static or dynamic lease will be  replaced by the new lease. Migration also supported only for dynamic lease requests. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
