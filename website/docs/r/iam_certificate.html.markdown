---
subcategory: "iam"
layout: "intersight"
page_title: "Intersight: intersight_iam_certificate"
description: |-
  Holds a certificate, signed by a CAcert.
---

# Resource: intersight_iam_certificate
Holds a certificate, signed by a CAcert.
## Argument Reference
The following arguments are supported:
* `certificate`:(HashMap) - User-input pem-encoded certificate, signed by a CAcert. 
This complex property has following sub-properties:
  + `issuer`:(HashMap) -(Computed) The X.509 distinguished name of the issuer of this certificate. 
This complex property has following sub-properties:
    + `common_name`:(string)(Computed) A required component that identifies a person or an object. 
    + `country`:
                (Array of schema.TypeString) -
    + `locality`:
                (Array of schema.TypeString) -
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
    + `organization`:
                (Array of schema.TypeString) -
    + `organizational_unit`:
                (Array of schema.TypeString) -
    + `state`:
                (Array of schema.TypeString) -
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `pem_certificate`:(string) The base64 encoded certificate in PEM format. 
  + `sha256_fingerprint`:(string)(Computed) The computed SHA-256 fingerprint of the certificate. Equivalent to 'openssl x509 -fingerprint -sha256'. 
  + `signature_algorithm`:(string)(Computed) Signature algorithm, as specified in [RFC 5280](https://tools.ietf.org/html/rfc5280). 
  + `subject`:(HashMap) -(Computed) The X.509 distinguished name of the subject of this certificate. 
This complex property has following sub-properties:
    + `common_name`:(string)(Computed) A required component that identifies a person or an object. 
    + `country`:
                (Array of schema.TypeString) -
    + `locality`:
                (Array of schema.TypeString) -
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
    + `organization`:
                (Array of schema.TypeString) -
    + `organizational_unit`:
                (Array of schema.TypeString) -
    + `state`:
                (Array of schema.TypeString) -
* `certificate_request`:(HashMap) - A reference to a iamCertificateRequest resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `status`:(string)(Computed) Status of the certificate.* `PendingValidation` - The certificate has not been validated.* `Valid` - The certificate is valid.* `Invalid` - Ther certificate is invalid. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 


## Import
`intersight_iam_certificate` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_iam_certificate.example 1234567890987654321abcde
``` 