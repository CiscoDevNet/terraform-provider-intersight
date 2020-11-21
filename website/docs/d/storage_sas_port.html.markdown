
---
layout: "intersight"
page_title: "Intersight: intersight_storage_sas_port"
sidebar_current: "docs-intersight-data-source-storage-sas-port"
description: |-
Sas Port details of the SAS endpoint.
---

# Data Source: intersight_storage._sas_port
Sas Port details of the SAS endpoint.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `address`:(string) The SAS Address assigned to storage port. 
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `disk_id`:(int) The unique disk identifier. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `end_point_id`:(int) The end-point Id assigned to storage port. 
* `link_description`:(string) The description for the link. 
* `link_speed`:(string) The link speed negotiated for communication. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
