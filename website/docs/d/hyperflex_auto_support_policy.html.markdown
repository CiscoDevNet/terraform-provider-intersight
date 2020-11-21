
---
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_auto_support_policy"
sidebar_current: "docs-intersight-data-source-hyperflex-auto-support-policy"
description: |-
A policy specifying the configuration to automatically generate support tickets with Cisco TAC.
---

# Data Source: intersight_hyperflex._auto_support_policy
A policy specifying the configuration to automatically generate support tickets with Cisco TAC.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `admin_state`:(bool) Enable or disable Auto Support. 
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `description`:(string) Description of the policy. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `service_ticket_receipient`:(string) The recipient email address for support tickets. 
