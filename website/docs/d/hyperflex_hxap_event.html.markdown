---
subcategory: "hyperflex"
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_hxap_event"
description: |-
  Events associated with HyperFlex Application Platform compute cluster.
---

# Data Source: intersight_hyperflex_hxap_event
Events associated with HyperFlex Application Platform compute cluster.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_hyperflex_hxap_event.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `first_time`:(string) First timestamp of the event occurrence. 
* `identity`:(string) Internally generated identity (UUID) of this event. 
* `last_time`:(string) Last timestamp of the event occurrence. 
* `message`:(string) Full description of the event. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `owner_name`:(string) Name of the owner with which event is associated. 
* `owner_type`:(string) Type of the object with which event is associated (Host, Cluster, VM).* `Unknown` - Value is Unknown if the target type is unidentified.* `Cluster` - Cluster refers to HyperFlex AP Cluster.* `Host` - Host refers to server node which is part of HyperFlex AP Cluster.* `VM` - VM refers to Virtual machine available on a HyperFlex AP Node.* `Disk` - Disk refers to Virtual Disk available on a HyperFlex AP Cluster. 
* `owner_uuid`:(string) UUID of the owner with which event is associated. 
* `severity`:(string) Severity level of the event (Info/Warning/Critical).* `None` - The Enum value None represents that there is no severity.* `Info` - The Enum value Info represents the Informational level of severity.* `Critical` - The Enum value Critical represents the Critical level of severity.* `Warning` - The Enum value Warning represents the Warning level of severity.* `Cleared` - The Enum value Cleared represents that the alarm severity has been cleared. 
 