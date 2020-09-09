
---
layout: "intersight"
page_title: "Intersight: intersight_equipment_tpm"
sidebar_current: "docs-intersight-data-source-equipment-tpm"
description: |-
TPM security chip on server board.
---

# Data Source: intersight_equipment._tpm
TPM security chip on server board.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `activation_status`:(string) Identifies the activation status of the TPM. 
* `admin_state`:(string) Identifies the admin configured state of the TPM. 
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `model`:(string) This field identifies the model of the given component. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `ownership`:(string) Identifies the ownership information of the TPM. 
* `presence`:(string) Identifies the presence of the trusted platform module. 
* `revision`:(string) This field identifies the revision of the given component. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field identifies the serial of the given component. 
* `tpm_id`:(int) Enter  the ID of the trusted platform module. 
* `vendor`:(string) This field identifies the vendor of the given component. 
* `version`:(string) Identifies the revision of the Trusted Platform Module. 
