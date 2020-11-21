
---
layout: "intersight"
page_title: "Intersight: intersight_iaas_license_info"
sidebar_current: "docs-intersight-data-source-iaas-license-info"
description: |-
Describes about license info currently available in UCSD.
---

# Data Source: intersight_iaas._license_info
Describes about license info currently available in UCSD.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `license_expiration_date`:(string) UCS Director license expiration date. 
* `license_type`:(string) License type of UCSD whether it is EVAL/Permanent/Subscription.. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
