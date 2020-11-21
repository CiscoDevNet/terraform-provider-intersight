
---
layout: "intersight"
page_title: "Intersight: intersight_appliance_external_syslog_setting"
sidebar_current: "docs-intersight-data-source-appliance-external-syslog-setting"
description: |-
Configure External Syslog Server.
---

# Data Source: intersight_appliance._external_syslog_setting
Configure External Syslog Server.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `enabled`:(bool) Enable or disable External Syslog Server. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `port`:(int) External Syslog Server Port for connection establishment. 
* `server`:(string) External Syslog Server Address, can be IP address or hostname. 
