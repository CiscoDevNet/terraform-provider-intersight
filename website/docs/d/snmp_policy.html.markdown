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
The following arguments can be used to get data of already created objects in Intersight appliance:
* `access_community_string`:(string) The default SNMPv1, SNMPv2c community name or SNMPv3 username to include on any trap messages sent to the SNMP host. The name can be 18 characters long. 
* `community_access`:(string) Controls access to the information in the inventory tables. Applicable only for SNMPv1 and SNMPv2c users.* `Disabled` - Blocks access to the information in the inventory tables.* `Limited` - Partial access to read the information in the inventory tables.* `Full` - Full access to read the information in the inventory tables. 
* `description`:(string) Description of the policy. 
* `enabled`:(bool) State of the SNMP Policy on the endpoint. If enabled, the endpoint sends SNMP traps to the designated host. 
* `engine_id`:(string) User-defined unique identification of the static engine. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `snmp_port`:(int) Port on which Cisco IMC SNMP agent runs. Enter a value between 1-65535. 
* `sys_contact`:(string) Contact person responsible for the SNMP implementation. Enter a string up to 64 characters, such as an email address or a name and telephone number. 
* `sys_location`:(string) Location of host on which the SNMP agent (server) runs. 
* `trap_community`:(string) SNMP community group used for sending SNMP trap to other devices. Valid only for SNMPv2c users. 
