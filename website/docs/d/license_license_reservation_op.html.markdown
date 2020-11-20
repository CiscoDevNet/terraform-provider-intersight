
---
layout: "intersight"
page_title: "Intersight: intersight_license_license_reservation_op"
sidebar_current: "docs-intersight-data-source-license-license-reservation-op"
description: |-
Customer operation object to request reservation code.
---

# Data Source: intersight_license._license_reservation_op
Customer operation object to request reservation code.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `auth_code`:(string) Revervation code used to install the license. 
* `auth_code_installed`:(bool) Flag to indicate whether authorization code is installed. 
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `confirm_code`:(string) Confirm code used to complete the license update on smart license account. 
* `generate_request_code`:(bool) Trigger the generation of request code for specific license reservation. 
* `generate_return_code`:(bool) Trigger the generation of return code for specific license reservation. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `request_code`:(string) Revervation code used to generate authorization code from CSSM. 
* `return_code`:(string) Return code used to return the reserved license to smart license account. 
