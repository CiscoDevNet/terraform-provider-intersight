---
subcategory: "ipmioverlan"
layout: "intersight"
page_title: "Intersight: intersight_ipmioverlan_policy_inventory"
description: |-
        Intelligent Platform Management Interface Over LAN Policy.

---

# Data Source: intersight_ipmioverlan_policy_inventory
Intelligent Platform Management Interface Over LAN Policy.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_ipmioverlan_policy_inventory.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the policy. 
* `device_mo_id`:(string) Device ID of the entity from where inventory is reported. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `enabled`:(bool) State of the IPMI Over LAN service on the endpoint. 
* `encryption_key`:(string) The encryption key to use for IPMI communication. It should have an even number of hexadecimal characters and not exceed 40 characters. Use “00” to disable encryption key use. This configuration is supported by all Standalone C-Series servers. FI-attached C-Series servers with firmware at minimum of 4.2.3a support this configuration. B/X-Series servers with firmware at minimum of 5.1.0.x support this configuration. IPMI commands using this key should append zeroes to the key to achieve a length of 40 characters. 
* `is_encryption_key_set`:(bool) Indicates whether the value of the 'encryptionKey' property has been set. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the inventoried policy object. 
* `privilege`:(string) The highest privilege level that can be assigned to an IPMI session on a server. This configuration is supported by all Standalone C-Series servers. FI-attached C-Series servers with firmware at minimum of 4.2.3a support this configuration. B/X-Series servers with firmware at minimum of 5.1.0.x support this configuration. Privilege level user is not supported for B/X-Series servers.* `admin` - Privilege to perform all actions available through IPMI.* `user` - Privilege to perform some functions through IPMI but restriction on performing administrative tasks.* `read-only` - Privilege to view information throught IPMI but restriction on making any changes. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
 
