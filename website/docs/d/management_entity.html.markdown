
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
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `cluster_link_state`:(string) Cluster link state between the Fabric Interconnects. 
* `cluster_readiness`:(string) Cluster readiness of the Fabric Interconnect. 
* `cluster_state`:(string) Cluster state of the Fabric Interconnect. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `entity_id`:(string) Identity of the Fabric Interconnect - A/B. 
* `leadership`:(string) Role (Primary / Subordinate) of the Fabric Interconnect. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
