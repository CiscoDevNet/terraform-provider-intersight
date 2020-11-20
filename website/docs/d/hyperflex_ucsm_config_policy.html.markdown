
---
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_ucsm_config_policy"
sidebar_current: "docs-intersight-data-source-hyperflex-ucsm-config-policy"
description: |-
A policy specifying UCS Manager settings such as service profile org, KVM IP addresses, and MAC prefix for server configuration in Fabric Interconnect environment.
---

# Data Source: intersight_hyperflex._ucsm_config_policy
A policy specifying UCS Manager settings such as service profile org, KVM IP addresses, and MAC prefix for server configuration in Fabric Interconnect environment.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `description`:(string) Description of the policy. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `server_firmware_version`:(string) The server firmware bundle version used for server components such as CIMC, adapters, BIOS, etc. 
