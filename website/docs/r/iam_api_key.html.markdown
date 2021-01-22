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
## Argument Reference
The following arguments are supported:
* `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `hash_algorithm`:(string) The cryptographic hash algorithm to calculate the message digest.* `SHA256` - The SHA-256 cryptographic hash, as defined by NIST in FIPS 180-4.* `SHA384` - The SHA-384 cryptographic hash, as defined by NIST in FIPS 180-4.* `SHA512` - The SHA-512 cryptographic hash, as defined by NIST in FIPS 180-4.* `SHA512_224` - The SHA-512/224 cryptographic hash, as defined by NIST in FIPS 180-4.* `SHA512_256` - The SHA-512/256 cryptographic hash, as defined by NIST in FIPS 180-4. 
* `key_spec`:(Array with Maximum of one item) - The key generation specification provides the algorithm and the parameters required for this algorithm to generate a private key, public key pair. Supported key generation schemes include RSA, ECDSA and Edwards-Curve Digital Signature Algorithm (EdDSA). 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
  + `name`:(string)(Computed) Name of the key generation algorithm.* `RSA` - Key pairs should be generated by the RSA algorithm. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `permission`:(Array with Maximum of one item) -(Computed) A reference to a iamPermission resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `private_key`:(string) Holds the private key for the API key. 
* `purpose`:(string) The purpose of the API Key. 
* `signing_algorithm`:(string) The signing algorithm used by the client to authenticate API requests to Intersight.The signing algorithm must be compatible with the key generation specification.* `RSASSA-PKCS1-v1_5` - RSASSA-PKCS1-v1_5 is a RSA signature scheme specified in [RFC 8017](https://tools.ietf.org/html/rfc8017).RSASSA-PKCS1-v1_5 is included only for compatibility with existing applications.* `RSASSA-PSS` - RSASSA-PSS is a RSA signature scheme specified in [RFC 8017](https://tools.ietf.org/html/rfc8017).It combines the RSASP1 and RSAVP1 primitives with the EMSA-PSS encoding method.In the interest of increased robustness, RSASSA-PSS is required in new applications.* `Ed25519` - The Ed25519 signature algorithm, as specified in [RFC 8032](https://tools.ietf.org/html/rfc8032).Ed25519 is a public-key signature system with several attractive features, includingfast single-signature verification, very fast signing, fast key generation and high security level.* `Ecdsa` - The Elliptic Curve Digital Signature Standard (ECDSA), as defined by NIST in FIPS 186-4 and ANSI X9.62.The signature is encoded as a ASN.1 DER SEQUENCE with two INTEGERs (r and s), as defined in RFC3279.When using ECDSA signatures, configure the client to use the same signature encoding as specified on the server side.* `EcdsaP1363Format` - The Elliptic Curve Digital Signature Standard (ECDSA), as defined by NIST in FIPS 186-4 and ANSI X9.62.The signature is the raw concatenation of r and s, as defined in the ISO/IEC 7816-8 IEEE P.1363 standard.In that format, r and s are represented as unsigned, big endian numbers.Extra padding bytes (of value 0x00) is applied so that both r and s encodings have the same size.When using ECDSA signatures, configure the client to use the same signature encoding as specified on the server side. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `user`:(Array with Maximum of one item) -(Computed) A reference to a iamUser resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `additional_properties`:
(Array with Maximum of one item) - Add additional properties in json format inside `jsonencode()` for this object.
  + `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
  + `moid`:(string)(Computed) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string)(Computed) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 


## Import
`intersight_iam_api_key` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_iam_api_key.example 1234567890987654321abcde
```