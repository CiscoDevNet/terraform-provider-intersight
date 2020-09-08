
---
layout: "intersight"
page_title: "Intersight: intersight_storage_sas_expander"
sidebar_current: "docs-intersight-data-source-storage-sas-expander"
description: |-
SAS Expander present in a server.
---

# Data Source: intersight_storage._sas_expander
SAS Expander present in a server.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `expander_id`:(int) Unique Identifier of the storage expander. 
* `model`:(string) This field identifies the model of the given component. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) The name  of the installed storage expander. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `oper_state`:(string) The operational state of the storage expander. 
* `operability`:(string) The operability status of the storage expander. 
* `presence`:(string) The availability of the storage expander. 
* `revision`:(string) This field identifies the revision of the given component. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `sas_address`:(string) The SAS address of the SAS expander. 
* `serial`:(string) This field identifies the serial of the given component. 
* `vendor`:(string) This field identifies the vendor of the given component. 
