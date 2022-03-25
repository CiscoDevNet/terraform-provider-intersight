---
subcategory: "iam"
layout: "intersight"
page_title: "Intersight: intersight_iam_api_key"
description: |-
        An API key is used to authenticate and authorize API requests sent by a client using the HTTP signature scheme. API keys can be used by unattended, daemon clients that need to send requests to Intersight programmatically. API keys are based on public key cryptography.
        To create an API key, the user must specify: 1. The purpose (description) of the API key, 2. The cryptographic hash algorithm, which is used to compute the digest of the body of HTTP requests, 3. The cryptographic parameters to generate a private/public key pair, e.g. RSA, ECDSA, EDDSA, key modulus, and 4. The signing algorithm, e.g. RSA PKCS v1.5, RSA PSS, ECDSA, EDDSA. The generated private key and public key are encoded in PEM format.
        The client owns the private key and is responsible for maintaining the confidentiality of the private key. The server holds the public key.
        The client must have a cryptographic provider compatible with the cryptographic parameters specified in the API key. For example, if you use the powershell SDK to write the client, make sure the appropriate cryptographic providers are installed on the local system. If you create an RSA key pair with modulus set to 2048, the client must support 2048-bit private keys. A maximum of 3 API keys per user is allowed.
        API keys are used to sign HTTP requests as follows: 1. A cryptographic digest of the body of the HTTP request is calculated using one of the supported cryptographic hash algorithms. 2. The value of the digest is base-64 encoded in the `Digest` HTTP header. 3. A signature is calculated as specified in the HTTP signature scheme, and the signature is added to the `Authorization` HTTP request header.
        All published Intersight SDKs support API keys.

---

# Resource: intersight_iam_api_key
An API key is used to authenticate and authorize API requests sent by a client using the HTTP signature scheme. API keys can be used by unattended, daemon clients that need to send requests to Intersight programmatically. API keys are based on public key cryptography.
To create an API key, the user must specify: 1. The purpose (description) of the API key, 2. The cryptographic hash algorithm, which is used to compute the digest of the body of HTTP requests, 3. The cryptographic parameters to generate a private/public key pair, e.g. RSA, ECDSA, EDDSA, key modulus, and 4. The signing algorithm, e.g. RSA PKCS v1.5, RSA PSS, ECDSA, EDDSA. The generated private key and public key are encoded in PEM format.
The client owns the private key and is responsible for maintaining the confidentiality of the private key. The server holds the public key.
The client must have a cryptographic provider compatible with the cryptographic parameters specified in the API key. For example, if you use the powershell SDK to write the client, make sure the appropriate cryptographic providers are installed on the local system. If you create an RSA key pair with modulus set to 2048, the client must support 2048-bit private keys. A maximum of 3 API keys per user is allowed.
API keys are used to sign HTTP requests as follows: 1. A cryptographic digest of the body of the HTTP request is calculated using one of the supported cryptographic hash algorithms. 2. The value of the digest is base-64 encoded in the `Digest` HTTP header. 3. A signature is calculated as specified in the HTTP signature scheme, and the signature is added to the `Authorization` HTTP request header.
All published Intersight SDKs support API keys.
## Usage Example
### Resource Creation

```hcl
resource "intersight_iam_api_key" "iam_api_key1" {
  hash_algorithm = "SHA256"
  key_spec = [{
    additional_properties = ""
    class_id              = "pkix.RsaAlgorithm"
    modulus               = 2048
    name                  = "RSA"
    object_type           = "pkix.RsaAlgorithm"
  }]
   parent = {
     moid        = var.iam_user
     object_type = "iam.User"
   }
   permission = {
     moid        = var.iam_permission
     object_type = "iam.Permission"
   }
   purpose           = "admin api"
   shared_scope      = ""
   signing_algorithm = "RSASSA-PKCS1-v1_5"
   user = {
     moid        = var.iam_user
     object_type = "iam.User"
   }
}

 variable "iam_permission" {
   type = string
   description = "value for iam_permission"
 }

 variable "iam_user" {
   type = string
   description = "value for iam_user"
 }

 variable "iam_role" {
   type = string
   description = "value for iam_role"
 }
```
## Argument Reference
The following arguments are supported:
* `account_moid`:(string)(ReadOnly) The Account ID for this managed object. 
* `ancestors`:(Array)(ReadOnly) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `create_time`:(string)(ReadOnly) The time when this managed object was created. 
* `domain_group_moid`:(string)(ReadOnly) The DomainGroup ID for this managed object. 
* `hash_algorithm`:(string) The cryptographic hash algorithm to calculate the message digest.* `SHA256` - The SHA-256 cryptographic hash, as defined by NIST in FIPS 180-4.* `SHA384` - The SHA-384 cryptographic hash, as defined by NIST in FIPS 180-4.* `SHA512` - The SHA-512 cryptographic hash, as defined by NIST in FIPS 180-4.* `SHA512_224` - The SHA-512/224 cryptographic hash, as defined by NIST in FIPS 180-4.* `SHA512_256` - The SHA-512/256 cryptographic hash, as defined by NIST in FIPS 180-4. 
* `key_spec`:(HashMap) - The key generation specification provides the algorithm and the parameters required for this algorithm to generate a private key, public key pair. Supported key generation schemes include RSA, ECDSA and Edwards-Curve Digital Signature Algorithm (EdDSA). 
This complex property has following sub-properties:
  + `additional_properties`:(JSON as string) - Additional Properties as per object type, can be added as JSON using `jsonencode()`. Allowed Types are: [pkix.EcdsaKeySpec](#pkixEcdsaKeySpec)
[pkix.EddsaKeySpec](#pkixEddsaKeySpec)
[pkix.RsaAlgorithm](#pkixRsaAlgorithm)
  + `name`:(string)(ReadOnly) Name of the key generation algorithm.* `RSA` - Key pairs should be generated by the RSA algorithm. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `mod_time`:(string)(ReadOnly) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `owners`:
                (Array of schema.TypeString) -(ReadOnly)
* `parent`:(HashMap) -(ReadOnly) A reference to a moBaseMo resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `permission`:(HashMap) -(ReadOnly) A reference to a iamPermission resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `permission_resources`:(Array)(ReadOnly) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `private_key`:(string) Holds the private key for the API key. 
* `purpose`:(string) The purpose of the API Key. 
* `shared_scope`:(string)(ReadOnly) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `signing_algorithm`:(string) The signing algorithm used by the client to authenticate API requests to Intersight.The signing algorithm must be compatible with the key generation specification.* `RSASSA-PKCS1-v1_5` - RSASSA-PKCS1-v1_5 is a RSA signature scheme specified in [RFC 8017](https://tools.ietf.org/html/rfc8017).RSASSA-PKCS1-v1_5 is included only for compatibility with existing applications.* `RSASSA-PSS` - RSASSA-PSS is a RSA signature scheme specified in [RFC 8017](https://tools.ietf.org/html/rfc8017).It combines the RSASP1 and RSAVP1 primitives with the EMSA-PSS encoding method.In the interest of increased robustness, RSASSA-PSS is required in new applications.* `Ed25519` - The Ed25519 signature algorithm, as specified in [RFC 8032](https://tools.ietf.org/html/rfc8032).Ed25519 is a public-key signature system with several attractive features, includingfast single-signature verification, very fast signing, fast key generation and high security level.* `Ecdsa` - The Elliptic Curve Digital Signature Standard (ECDSA), as defined by NIST in FIPS 186-4 and ANSI X9.62.The signature is encoded as a ASN.1 DER SEQUENCE with two INTEGERs (r and s), as defined in RFC3279.When using ECDSA signatures, configure the client to use the same signature encoding as specified on the server side.* `EcdsaP1363Format` - The Elliptic Curve Digital Signature Standard (ECDSA), as defined by NIST in FIPS 186-4 and ANSI X9.62.The signature is the raw concatenation of r and s, as defined in the ISO/IEC 7816-8 IEEE P.1363 standard.In that format, r and s are represented as unsigned, big endian numbers.Extra padding bytes (of value 0x00) is applied so that both r and s encodings have the same size.When using ECDSA signatures, configure the client to use the same signature encoding as specified on the server side. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `user`:(HashMap) -(ReadOnly) A reference to a iamUser resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `version_context`:(HashMap) -(ReadOnly) The versioning info for this managed object. 
This complex property has following sub-properties:
  + `interested_mos`:(Array)
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `ref_mo`:(HashMap) -(ReadOnly) A reference to the original Managed Object. 
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `timestamp`:(string)(ReadOnly) The time this versioned Managed Object was created. 
  + `nr_version`:(string)(ReadOnly) The version of the Managed Object, e.g. an incrementing number or a hash id. 
  + `version_type`:(string)(ReadOnly) Specifies type of version. Currently the only supported value is \ Configured\ that is used to keep track of snapshots of policies and profiles that are intendedto be configured to target endpoints.* `Modified` - Version created every time an object is modified.* `Configured` - Version created every time an object is configured to the service profile.* `Deployed` - Version created for objects related to a service profile when it is deployed. 


## Import
`intersight_iam_api_key` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_iam_api_key.example 1234567890987654321abcde
```
## Allowed Types in `AdditionalProperties`
 
### [pkix.EcdsaKeySpec](#argument-reference)
The key pair is generated using Elliptic Curve Digital Signature Algorithm (ECDSA), as defined in FIPS 186-4.
The ECDSA standard was originally developed for the American National Standards Institute by the Accredited
Standards Committee on Financial Services, X9.
ANS X9.62 defines methods for digital signature generation and verification using ECDSA. Specifications for
the generation of the domain parameters used during the generation and verification of digital signatures
are also included in ANS X9.62.
* `curve`:(string) A specific set of Elliptic Curve parameters, as recommended by NIST in FIPS 186-4.* `P256` - P256 returns a Curve which implements P-256, as defined in FIPS 186-4, section D.2.3.* `P224` - P224 returns a Curve which implements P-224, as defined in FIPS 186-4, section D.2.2.* `P384` - P384 returns a Curve which implements P-384, as defined in FIPS 186-4, section D.2.4.* `P521` - P521 returns a Curve which implements P-521, as defined in FIPS 186-4, section D.2.5. 

### [pkix.EddsaKeySpec](#argument-reference)
The key pair is generated using Edwards-Curve Digital Signature Algorithm (EdDSA).
The Edwards-curve Digital Signature Algorithm (EdDSA) is a variant of
Schnorr's signature system with (possibly twisted) Edwards curves.
* `algorithm`:(string) The EdDSA algorithm, as defined in RFC 8032.* `Ed25519` - The edwards25519 curve, as defined in RFC 8032 section 5.1.* `Ed25519ph` - The edwards25519 curve for the PureEdDSA variant.* `Ed25519ctx` - The edwards25519 curve for the HashEdDSA variant. 

### [pkix.RsaAlgorithm](#argument-reference)
The key pair is generated using the RSA algorithm and specified parameters.
* `modulus`:(int) The length of the RSA key, expressed in bits, for both public and private keys.* `2048` - A key length of 2048 bits.* `2560` - A key length of 2560 bits.* `3072` - A key length of 3072 bits.* `3584` - A key length of 3584 bits.* `4096` - A key length of 4096 bits. 
  
