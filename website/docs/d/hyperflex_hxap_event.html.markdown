
---
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_hxap_event"
sidebar_current: "docs-intersight-data-source-hyperflex-hxap-event"
description: |-
Events associated with HyperFlex Application Platform compute cluster.
---

# Data Source: intersight_hyperflex._hxap_event
Events associated with HyperFlex Application Platform compute cluster.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `first_time`:(string) First timestamp of the event occurrence. 
* `identity`:(string) Internally generated identity (UUID) of this event. 
* `last_time`:(string) Last timestamp of the event occurrence. 
* `message`:(string) Full description of the event. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `owner_name`:(string) Name of the owner with which event is associated. 
* `owner_type`:(string) Type of the object with which event is associated (Host, Cluster, VM).* `Unknown` - Value is Unknown if the target type is unidentified.* `Cluster` - Cluster refers to HyperFlex AP Cluster.* `Host` - Host refers to server node which is part of HyperFlex AP Cluster.* `VM` - VM refers to Virtual machine available on a HyperFlex AP Node.* `Disk` - Disk refers to Virtual Disk available on a HyperFlex AP Cluster. 
* `owner_uuid`:(string) UUID of the owner with which event is associated. 
* `severity`:(string) Severity level of the event (Info/Warning/Critical).* `None` - The Enum value None represents that there is no severity.* `Info` - The Enum value Info represents the Informational level of severity.* `Critical` - The Enum value Critical represents the Critical level of severity.* `Warning` - The Enum value Warning represents the Warning level of severity.* `Cleared` - The Enum value Cleared represents that the alarm severity has been cleared. 
