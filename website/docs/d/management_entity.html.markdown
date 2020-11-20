
---
layout: "intersight"
page_title: "Intersight: intersight_management_entity"
sidebar_current: "docs-intersight-data-source-management-entity"
description: |-
Logical representation that captures the role of each Fabric Interconnect in UCS Manager.
---

# Data Source: intersight_management._entity
Logical representation that captures the role of each Fabric Interconnect in UCS Manager.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `cluster_link_state`:(string) Cluster link state between the Fabric Interconnects. 
* `cluster_readiness`:(string) Cluster readiness of the Fabric Interconnect. 
* `cluster_state`:(string) Cluster state of the Fabric Interconnect. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `entity_id`:(string) Identity of the Fabric Interconnect - A/B. 
* `leadership`:(string) Role (Primary / Subordinate) of the Fabric Interconnect. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
