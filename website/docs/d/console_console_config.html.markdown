---
subcategory: "console"
layout: "intersight"
page_title: "Intersight: intersight_console_console_config"
description: |-
        Concrete class for terminal Console.

---

# Data Source: intersight_console_console_config
Concrete class for terminal Console.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_console_console_config.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `data_bits`:(int) The databits numbers in console. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `parity`:(string) The parity bit of the console.* `none` - If Parity is none, parity checking is not performed and the parity bit is not transmitted.* `odd` - If Parity is odd, the number of mark bits (1s) in the data is counted, and the parity bit is asserted or unasserted to obtain an odd number of mark bits.* `even` - If Parity is even, the number of mark bits in the data is counted, and the parity bit is asserted or unasserted to obtain an even number of mark bits.* `mark` - If Parity is mark, the parity bit is asserted.* `space` - If Parity is space, the parity bit is unasserted. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `speed`:(int) The speed of the console. 
* `stop_bits`:(int) The stopbits of async line in console. 
 
