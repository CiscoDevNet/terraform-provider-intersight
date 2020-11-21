
---
layout: "intersight"
page_title: "Intersight: intersight_iam_certificate"
sidebar_current: "docs-intersight-data-source-iam-certificate"
description: |-
Holds a certificate, signed by a CAcert.
---

# Data Source: intersight_iam._certificate
Holds a certificate, signed by a CAcert.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `status`:(string) Status of the certificate.* `PendingValidation` - The certificate has not been validated.* `Valid` - The certificate is valid.* `Invalid` - Ther certificate is invalid. 
