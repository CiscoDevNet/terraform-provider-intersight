
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
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `email_address`:(string) User input email address, an optional part of the subject of the certificate request. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the certificate request. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `request`:(string) Generated certificate signing request (CSR) in PEM format. 
* `self_signed`:(bool) Whether the user wants the generated CSR to be self-signed by the appliance. 
