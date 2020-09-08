
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
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `license_expiration_date`:(string) UCS Director license expiration date. 
* `license_type`:(string) License type of UCSD whether it is EVAL/Permanent/Subscription.. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
