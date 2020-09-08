
---
layout: "intersight"
page_title: "Intersight: intersight_iam_certificate_request"
sidebar_current: "docs-intersight-data-source-iam-certificate-request"
description: |-
The information required to generate a certificate signing request (CSR),
which is a block of encoded text that is given to a Certificate Authority when applying for an SSL Certificate.
---

# Data Source: intersight_iam._certificate_request
The information required to generate a certificate signing request (CSR),
which is a block of encoded text that is given to a Certificate Authority when applying for an SSL Certificate.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `email_address`:(string) User input email address, an optional part of the subject of the certificate request. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the certificate request. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `request`:(string) Generated certificate signing request (CSR) in PEM format. 
* `self_signed`:(bool) Whether the user wants the generated CSR to be self-signed by the appliance. 
