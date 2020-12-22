---
subcategory: "kubernetes"
layout: "intersight"
page_title: "Intersight: intersight_kubernetes_addon_repository"
description: |-
  Docker registry or helm repository which hosts helm charts and docker images.
---

# Resource: intersight_kubernetes_addon_repository
Docker registry or helm repository which hosts helm charts and docker images.
## Argument Reference
The following arguments are supported:
* `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
* `ca_cert`:(Array with Maximum of one item) - CA certificate for the addon registry if it is not a well known CA. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
  + `issuer`:(Array with Maximum of one item) -(Computed) The X.509 distinguished name of the issuer of this certificate. 
This complex property has following sub-properties:
    + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
    + `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
    + `common_name`:(string)(Computed) A required component that identifies a person or an object. 
    + `country`:
                (Array of schema.TypeString) -
    + `locality`:
                (Array of schema.TypeString) -
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
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
  + `subject`:(Array with Maximum of one item) -(Computed) The X.509 distinguished name of the subject of this certificate. 
This complex property has following sub-properties:
    + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
    + `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
    + `common_name`:(string)(Computed) A required component that identifies a person or an object. 
    + `country`:
                (Array of schema.TypeString) -
    + `locality`:
                (Array of schema.TypeString) -
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
    + `organization`:
                (Array of schema.TypeString) -
    + `organizational_unit`:
                (Array of schema.TypeString) -
    + `state`:
                (Array of schema.TypeString) -
* `catalog`:(Array with Maximum of one item) - A reference to a workflowCatalog resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `insecure_skip_verification`:(bool) Allow connecting to http registries or https registries which do not have certificate signed by a well known CA. 
* `is_password_set`:(bool)(Computed) Indicates whether the value of the 'password' property has been set. 
* `is_token_set`:(bool)(Computed) Indicates whether the value of the 'token' property has been set. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the addon repository or registry. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `registered_device`:(Array with Maximum of one item) - A reference to a assetDeviceRegistration resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `repository_url`:(string) URL for the repository where the addon is hosted. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `username`:(string) Username to authenticate to the addon registry. 


## Import
`intersight_kubernetes_addon_repository` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_kubernetes_addon_repository.example 1234567890987654321abcde
```