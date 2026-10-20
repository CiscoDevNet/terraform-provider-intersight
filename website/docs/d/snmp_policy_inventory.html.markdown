---
subcategory: "snmp"
layout: "intersight"
page_title: "Intersight: intersight_snmp_policy_inventory"
description: |-
        The SNMP Policy is a reusable policy for configuring Simple Network Management Protocol (SNMP) settings on endpoints like servers and chassis.
        #### Purpose
        The purpose of this policy is to enable and standardize device monitoring and management via SNMP. It allows administrators to configure the SNMP agent on bghendpoints, define access credentials for different SNMP versions, and specify destinations for SNMP traps, integrating the devices into existing network management systems.
        #### Key Concepts
        - **Multi-Version Support:** The policy allows for the independent configuration of SNMPv2c and SNMPv3, including enabling or disabling each version.
        - **User and Community Strings:** It supports the configuration of read-only community strings for SNMPv2c and detailed user-based security models (USM) for SNMPv3, including authentication and privacy protocols.
        - **Trap Configuration:** Administrators can define a list of trap destinations, specifying the recipient's address, port, and the SNMP version to use for the trap messages.
        - **Profile-Based Application:** The policy is attached to Server or Chassis Profiles to apply a consistent SNMP configuration across multiple devices.

---

# Data Source: intersight_snmp_policy_inventory
The SNMP Policy is a reusable policy for configuring Simple Network Management Protocol (SNMP) settings on endpoints like servers and chassis.
#### Purpose
The purpose of this policy is to enable and standardize device monitoring and management via SNMP. It allows administrators to configure the SNMP agent on bghendpoints, define access credentials for different SNMP versions, and specify destinations for SNMP traps, integrating the devices into existing network management systems.
#### Key Concepts
- **Multi-Version Support:** The policy allows for the independent configuration of SNMPv2c and SNMPv3, including enabling or disabling each version.
- **User and Community Strings:** It supports the configuration of read-only community strings for SNMPv2c and detailed user-based security models (USM) for SNMPv3, including authentication and privacy protocols.
- **Trap Configuration:** Administrators can define a list of trap destinations, specifying the recipient's address, port, and the SNMP version to use for the trap messages.
- **Profile-Based Application:** The policy is attached to Server or Chassis Profiles to apply a consistent SNMP configuration across multiple devices.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_snmp_policy_inventory.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `access_community_string`:(string) The default SNMPv1, SNMPv2c community name or SNMPv3 username to include on any trap messages sent to the SNMP host. The name can be 32 characters long. 
* `account_moid`:(string) The Account ID for this managed object. 
* `community_access`:(string) Controls access to the information in the inventory tables. Applicable only for SNMPv1 and SNMPv2c users.* `Disabled` - Blocks access to the information in the inventory tables.* `Limited` - Partial access to read the information in the inventory tables.* `Full` - Full access to read the information in the inventory tables. 
* `create_time`:(string) The time when this managed object was created. 
* `description`:(string) Description of the policy. 
* `device_mo_id`:(string) Device ID of the entity from where inventory is reported. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `enabled`:(bool) State of the SNMP Policy on the endpoint. If enabled, the endpoint sends SNMP traps to the designated host. 
* `engine_id`:(string) User-defined unique identification of the static engine. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the inventoried policy object. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `snmp_port`:(int) Port on which Cisco IMC SNMP agent runs. Enter a value between 1-65535. Reserved ports not allowed (22, 23, 80, 123, 389, 443, 623, 636, 2068, 3268, 3269). 
* `sys_contact`:(string) Contact person responsible for the SNMP implementation. Enter a string up to 64 characters, such as an email address or a name and telephone number. 
* `sys_location`:(string) Location of host on which the SNMP agent (server) runs. 
* `trap_community`:(string) SNMP community group used for sending SNMP trap to other devices. Valid only for SNMPv2c users. 
* `v2_enabled`:(bool) State of the SNMP v2c on the endpoint. If enabled, the endpoint sends SNMP v2c properties to the designated host. 
* `v3_enabled`:(bool) State of the SNMP v3 on the endpoint. If enabled, the endpoint sends SNMP v3 properties to the designated host. 
 
