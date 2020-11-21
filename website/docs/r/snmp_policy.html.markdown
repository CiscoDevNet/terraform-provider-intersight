
---
layout: "intersight"
page_title: "Intersight: intersight_snmp_policy"
sidebar_current: "docs-intersight-resource-snmp-policy"
description: |-
  Policy to configure SNMP settings on endpoint.
---

# Resource: intersight_snmp._policy
Policy to configure SNMP settings on endpoint.
## Argument Reference
The following arguments are supported:
* `access_community_string`:(string) The default SNMPv1, SNMPv2c community name or SNMPv3 username to include on any trap messages sent to the SNMP host. The name can be 18 characters long. 
* `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `community_access`:(string) Controls access to the information in the inventory tables. Applicable only for SNMPv1 and SNMPv2c users.* `Disabled` - Blocks access to the information in the inventory tables.* `Limited` - Partial access to read the information in the inventory tables.* `Full` - Full access to read the information in the inventory tables. 
* `description`:(string) Description of the policy. 
* `enabled`:(bool) State of the SNMP Policy on the endpoint. If enabled, the endpoint sends SNMP traps to the designated host. 
* `engine_id`:(string) User-defined unique identification of the static engine. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `organization`:(Array with Maximum of one item) - A reference to a organizationOrganization resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `profiles`:(Array) An array of relationships to policyAbstractConfigProfile resources. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `snmp_port`:(int) Port on which Cisco IMC SNMP agent runs. Enter a value between 1-65535. Reserved ports not allowed (22, 23, 80, 123, 389, 443, 623, 636, 2068, 3268, 3269). 
* `snmp_traps`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
  + `destination`:(string) Address to which the SNMP trap information is sent. 
  + `enabled`:(bool) Enables/disables the trap on the server If enabled, trap is active on the server. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
  + `port`:(int) Port used by the server to communicate with trap destination. Enter a value between 1-65535. Reserved ports not allowed (22, 23, 80, 123, 389, 443, 623, 636, 2068, 3268, 3269). 
  + `type`:(string) Type of trap which decides whether to receive a notification when a trap is received at the destination.* `Trap` - Do not receive notifications when trap is sent to the destination.* `Inform` - Receive notifications when trap is sent to the destination. This option is valid only for V2 users. 
  + `user`:(string) SNMP user for the trap. Applicable only to SNMPv3. 
  + `version`:(string) SNMP version used for the trap.* `V3` - SNMP v3 trap version notifications.* `V2` - SNMP v2 trap version notifications. 
* `snmp_users`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `auth_password`:(string) Authorization password for the user. 
  + `auth_type`:(string) Authorization protocol for authenticating the user.* `NA` - Authentication protocol is not applicable.* `MD5` - MD5 protocol is used to authenticate SNMP user.* `SHA` - SHA protocol is used to authenticate SNMP user. 
  + `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
  + `is_auth_password_set`:(bool)(Computed) Indicates whether the value of the 'authPassword' property has been set. 
  + `is_privacy_password_set`:(bool)(Computed) Indicates whether the value of the 'privacyPassword' property has been set. 
  + `name`:(string) SNMP username. Must have a minimum of 1 and and a maximum of 31 characters. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
  + `privacy_password`:(string) Privacy password for the user. 
  + `privacy_type`:(string) Privacy protocol for the user.* `NA` - Privacy protocol is not applicable.* `DES` - DES privacy protocol is used for SNMP user.* `AES` - AES privacy protocol is used for SNMP user. 
  + `security_level`:(string) Security mechanism used for communication between agent and manager.* `AuthPriv` - The user requires both an authorization password and a privacy password.* `NoAuthNoPriv` - The user does not require an authorization or privacy password.* `AuthNoPriv` - The user requires an authorization password but not a privacy password. 
* `sys_contact`:(string) Contact person responsible for the SNMP implementation. Enter a string up to 64 characters, such as an email address or a name and telephone number. 
* `sys_location`:(string) Location of host on which the SNMP agent (server) runs. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `trap_community`:(string) SNMP community group used for sending SNMP trap to other devices. Valid only for SNMPv2c users. 
