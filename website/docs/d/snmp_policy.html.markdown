---
subcategory: "snmp"
layout: "intersight"
page_title: "Intersight: intersight_snmp_policy"
description: |-
        Policy to configure SNMP settings on endpoint.

---

# Data Source: intersight_snmp_policy
Policy to configure SNMP settings on endpoint.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_snmp_policy.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `access_community_string`:(string) The default SNMPv1, SNMPv2c community name or SNMPv3 username to include on any trap messages sent to the SNMP host. The name can be 18 characters long. 
* `account_moid`:(string) The Account ID for this managed object. 
* `community_access`:(string) Controls access to the information in the inventory tables. Applicable only for SNMPv1 and SNMPv2c users.* `Disabled` - Blocks access to the information in the inventory tables.* `Limited` - Partial access to read the information in the inventory tables.* `Full` - Full access to read the information in the inventory tables. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the policy. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `enabled`:(bool) State of the SNMP Policy on the endpoint. If enabled, the endpoint sends SNMP traps to the designated host. 
* `engine_id`:(string) User-defined unique identification of the static engine. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `snmp_port`:(int) Port on which Cisco IMC SNMP agent runs. Enter a value between 1-65535. Reserved ports not allowed (22, 23, 80, 123, 389, 443, 623, 636, 2068, 3268, 3269). 
* `sys_contact`:(string) Contact person responsible for the SNMP implementation. Enter a string up to 64 characters, such as an email address or a name and telephone number. 
* `sys_location`:(string) Location of host on which the SNMP agent (server) runs. 
* `trap_community`:(string) SNMP community group used for sending SNMP trap to other devices. Valid only for SNMPv2c users. 
* `v2_enabled`:(bool) State of the SNMP v2c on the endpoint. If enabled, the endpoint sends SNMP v2c properties to the designated host. 
* `v3_enabled`:(bool) State of the SNMP v3 on the endpoint. If enabled, the endpoint sends SNMP v3 properties to the designated host. 
 
