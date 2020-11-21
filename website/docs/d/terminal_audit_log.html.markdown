
---
layout: "intersight"
page_title: "Intersight: intersight_terminal_audit_log"
sidebar_current: "docs-intersight-data-source-terminal-audit-log"
description: |-
Audit log of remote terminal user sessions.
---

# Data Source: intersight_terminal._audit_log
Audit log of remote terminal user sessions.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `end_time`:(string) The time the terminal was closed. If terminal has not closed, value is zero time. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `start_time`:(string) The time the terminal session was opened. 
