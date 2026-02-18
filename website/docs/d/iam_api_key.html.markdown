---
subcategory: "iam"
layout: "intersight"
page_title: "Intersight: intersight_iam_api_key"
description: |-
        The ApiKey object is a critical component for secure authentication and authorization within the system. It allows clients to make API requests programmatically using a secure HTTP signature scheme.
        #### Purpose
        An ApiKey serves as a mechanism for authentication and authorization, enabling unattended clients, such as daemons, to interact with the API securely. This provides the means to verify the identity of a client and ensure they have the appropriate permissions to perform specific actions.
        #### Key Concepts
        - **Public Key Cryptography:** Utilizes public key cryptography to ensure secure authentication, with the server holding the public key and the client maintaining the confidentiality of the private key.
        - **Request Signing:** Involves signing HTTP requests to ensure integrity and authenticity, using cryptographic digest and signature mechanisms.
        - **Access Control:** Restricted to specific privilege sets, ensuring only authorized users can create, read, update, or delete API keys.
        - **Expiration and Rotation:** Supports expiration settings and key rotation, enhancing security by limiting the validity period of keys and facilitating the replacement of compromised or expired keys.

---

# Data Source: intersight_iam_api_key
The ApiKey object is a critical component for secure authentication and authorization within the system. It allows clients to make API requests programmatically using a secure HTTP signature scheme.
#### Purpose
An ApiKey serves as a mechanism for authentication and authorization, enabling unattended clients, such as daemons, to interact with the API securely. This provides the means to verify the identity of a client and ensure they have the appropriate permissions to perform specific actions.
#### Key Concepts
- **Public Key Cryptography:** Utilizes public key cryptography to ensure secure authentication, with the server holding the public key and the client maintaining the confidentiality of the private key.
- **Request Signing:** Involves signing HTTP requests to ensure integrity and authenticity, using cryptographic digest and signature mechanisms.
- **Access Control:** Restricted to specific privilege sets, ensuring only authorized users can create, read, update, or delete API keys.
- **Expiration and Rotation:** Supports expiration settings and key rotation, enhancing security by limiting the validity period of keys and facilitating the replacement of compromised or expired keys.
## Argument Reference
The results of this data source are stored in `results` property.
All objects matching the filter criteria are fetched through pagination.
To access the ith object of the results obtained, use `data.intersight_iam_api_key.<custom_name>.results[i].<propertyname>`.
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_moid`:(string) The Account ID for this managed object. 
* `admin_status`:(string) Used to trigger the enable or disable action on the API key. These actions change the status of an API key.* `enable` - Used to enable a disabled API key/App Registration. If the API key/App Registration is already expired, this action has no effect.* `disable` - Used to disable an active API key/App Registration. If the API key/App Registration is already expired, this action has no effect. 
* `create_time`:(string) The time when this managed object was created. 
* `domain_group_moid`:(string) The DomainGroup ID for this managed object. 
* `expiry_date_time`:(string) The expiration date of the API key which is set at the time of creation of the key. Its value can only be assigned a date that falls within the range determined by the maximum expiration time configured at the account level. The expiry date can be edited to be earlier or later, provided it stays within the designated expiry period. This period is determined by adding the 'startTime' property of the API key to the maximum expiry time configured at the account level. 
* `hash_algorithm`:(string) The cryptographic hash algorithm to calculate the message digest.* `SHA256` - The SHA-256 cryptographic hash, as defined by NIST in FIPS 180-4.* `SHA384` - The SHA-384 cryptographic hash, as defined by NIST in FIPS 180-4.* `SHA512` - The SHA-512 cryptographic hash, as defined by NIST in FIPS 180-4.* `SHA512_224` - The SHA-512/224 cryptographic hash, as defined by NIST in FIPS 180-4.* `SHA512_256` - The SHA-512/256 cryptographic hash, as defined by NIST in FIPS 180-4. 
* `is_never_expiring`:(bool) Used to mark the API key as a never-expiring API key. 
* `last_used_ip`:(string) The IP address from which the API key was last used. 
* `last_used_time`:(string) The time at which the API key was last used. It is updated every 24 hours. 
* `mod_time`:(string) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `oper_status`:(string) The current status of the API key that dictates the validity of the key.* `enabled` - An API key/App Registration having enabled status can be used for API invocation.* `disabled` - An API key/App Registration having disabled status cannot be used for API invocation.* `expired` - An API key/App Registration having expired status cannot be used for API invocation as the expiration date has passed. 
* `private_key`:(string) Holds the private key for the API key. 
* `purpose`:(string) The purpose of the API Key. 
* `shared_scope`:(string) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `signing_algorithm`:(string) The signing algorithm used by the client to authenticate API requests to Intersight.The signing algorithm must be compatible with the key generation specification.* `RSASSA-PKCS1-v1_5` - RSASSA-PKCS1-v1_5 is a RSA signature scheme specified in [RFC 8017](https://tools.ietf.org/html/rfc8017).RSASSA-PKCS1-v1_5 is included only for compatibility with existing applications.* `RSASSA-PSS` - RSASSA-PSS is a RSA signature scheme specified in [RFC 8017](https://tools.ietf.org/html/rfc8017).It combines the RSASP1 and RSAVP1 primitives with the EMSA-PSS encoding method.In the interest of increased robustness, RSASSA-PSS is required in new applications.* `Ed25519` - The Ed25519 signature algorithm, as specified in [RFC 8032](https://tools.ietf.org/html/rfc8032).Ed25519 is a public-key signature system with several attractive features, includingfast single-signature verification, very fast signing, fast key generation and high security level.* `Ecdsa` - The Elliptic Curve Digital Signature Standard (ECDSA), as defined by NIST in FIPS 186-4 and ANSI X9.62.The signature is encoded as a ASN.1 DER SEQUENCE with two INTEGERs (r and s), as defined in RFC3279.When using ECDSA signatures, configure the client to use the same signature encoding as specified on the server side.* `EcdsaP1363Format` - The Elliptic Curve Digital Signature Standard (ECDSA), as defined by NIST in FIPS 186-4 and ANSI X9.62.The signature is the raw concatenation of r and s, as defined in the ISO/IEC 7816-8 IEEE P.1363 standard.In that format, r and s are represented as unsigned, big endian numbers.Extra padding bytes (of value 0x00) is applied so that both r and s encodings have the same size.When using ECDSA signatures, configure the client to use the same signature encoding as specified on the server side. 
* `start_time`:(string) The timestamp at which an expiry date was first set on this API key. For expiring API keys, this field is same as the create time of the API key. For never-expiring API keys, this field is set initially to zero time value. If a never-expiry API key is later changed to have an expiration, the timestamp marking the start of this transition is recorded in this field. 
 
