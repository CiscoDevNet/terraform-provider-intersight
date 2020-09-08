
---
layout: "intersight"
page_title: "Intersight: intersight_equipment_shared_io_module"
sidebar_current: "docs-intersight-data-source-equipment-shared-io-module"
description: |-
I/O Controller present inside SIOC to provide data path from S-series server to FI.
---

# Data Source: intersight_equipment._shared_io_module
I/O Controller present inside SIOC to provide data path from S-series server to FI.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `config_state`:(string) This field identifies the configuration state for this SIOM Unit. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `discovery`:(string) This field identifies the discovery state of SIOM. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `mac_of_shared_iom_aside`:(string) This field identifies the MAC of IOM-A side. 
* `mac_of_shared_iom_bside`:(string) This field identifies the MAC of IOM-B side. 
* `model`:(string) This field identifies the model of the given component. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `oper_state`:(string) This field identifies the SIOM operational state. 
* `part_number`:(string) This field identifies the Part Number for this SIOM Unit. 
* `reachability`:(string) This field identifies the reachability to FI-A and B side. 
* `revision`:(string) This field identifies the revision of the given component. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field identifies the serial of the given component. 
* `usr_lbl`:(string) User label configured for the SIOM. 
* `vendor`:(string) This field identifies the vendor of the given component. 
* `vid`:(string) This field identifies the vendor id for this SIOM Unit. 
