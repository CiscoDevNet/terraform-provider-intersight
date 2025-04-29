---
subcategory: "fabric"
layout: "intersight"
page_title: "Intersight: intersight_fabric_mac_sec_policy"
description: |-
        A placeholder for MACSec (Media Access Control Security) configuration parameters, Primary/Fallback key chain and EAPol (Extensible Authentication Protocol over LAN) configurations.

---

# Resource: intersight_fabric_mac_sec_policy
A placeholder for MACSec (Media Access Control Security) configuration parameters, Primary/Fallback key chain and EAPol (Extensible Authentication Protocol over LAN) configurations.
## Argument Reference
The following arguments are supported:
* `account_moid`:(string)(ReadOnly) The Account ID for this managed object. 
* `ancestors`:(Array)(ReadOnly) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `cipher_suite`:(string) Cipher suite to be used for MACsec encryption.* `GCM-AES-XPN-256` - An extended Cipher Suite of GCM-AES-256 used in MACsec (Media Access Control Security) that incorporates Extended Packet Numbering (XPN) for enhanced security and scalability.* `GCM-AES-128` - This Cipher Suite employs the Advanced Encryption Standard (AES) with a 128-bit key in Galois/Counter Mode, offering both encryption and authentication.* `GCM-AES-256` - This Cipher Suite utilizes Advanced Encryption Standard (AES) with a 256-bit key in Galois/Counter Mode, offering a higher level of security compared to GCM-AES-128 due to the larger key size.* `GCM-AES-XPN-128` - An extended Cipher Suite of GCM-AES-128  used in MACsec (Media Access Control Security) that incorporates Extended Packet Numbering (XPN) to enhance security and scalability. 
* `confidentiality_offset`:(string) The MACsec confidentiality offset specifies the number of bytes starting from the frame header. MACsec encrypts only the bytes after the offset in a frame.* `CONF-OFFSET-0` - A value of 0 means the entire ethernet frame is encrypted.* `CONF-OFFSET-30` - The first 30 bytes of the ethernet frame are not encrypted, and the rest of the frame is encrypted.* `CONF-OFFSET-50` - The first 50 bytes of the ethernet frame are not encrypted, and the rest of the frame is encrypted. 
* `create_time`:(string)(ReadOnly) The time when this managed object was created. 
* `description`:(string) Description of the policy. 
* `domain_group_moid`:(string)(ReadOnly) The DomainGroup ID for this managed object. 
* `fallback_key_chain`:(HashMap) - Fallback keychain for managing an alternative set of security keys to be used when a secure session cannot be established using the primary keychain. 
This complex property has following sub-properties:
  + `name`:(string) Must be a maximum of 63 characters, without spacing. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `sec_keys`:(Array)
This complex property has following sub-properties:
    + `cryptographic_algorithm`:(string) The cryptographic algorithm that employs the cipher-based message authentication code (CMAC) mode of operation with advanced encryption standard (AES).* `AES_256_CMAC` - Uses the AES (Advanced Encryption Standard) algorithm with a 256-bit key to generate a CMAC.* `AES_128_CMAC` - Uses the AES (Advanced Encryption Standard) algorithm with a 128-bit key to generate a CMAC. 
    + `id`:(string) Must have an even number of hexadecimal characters (including 0-9 and A-F, only) with a length between 2 and 64 characters. For example, \ 10\ , \ 2000\ , \ ABCD1234\ . 
    + `is_octet_string_set`:(bool)(ReadOnly) Indicates whether the value of the 'octetString' property has been set. 
    + `key_type`:(string) The type of encryption used for the specified key.* `Type-0` - No encryption for the specified octetString.* `Type-6` - Proprietary advanced encryption standard for the specified octetString.* `Type-7` - Proprietary insecure encryption for the specified octetString. 
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
    + `octet_string`:(string) The key octet string is a shared secret used in cryptographic operations. The valid size and format of the octet string depend on the selected KeyCryptographicAlgorithm and KeyEncryptionType. It should start with the character 'J'. 
    + `send_lifetime_duration`:(int) The key lifetime duration in seconds after the start time. If a non-zero value is configured for the duration, the end time configuration for the key lifetime is ignored. 
    + `send_lifetime_end_time`:(string) The time of day and date when the key becomes inactive. 
    + `send_lifetime_infinite`:(bool) Indicates that the key remains active indefinitely after the specified start time. When this parameter is set, the end time and duration configurations for the key lifetime are ignored. 
    + `send_lifetime_start_time`:(string) The time of day and date when the key becomes active. 
    + `send_lifetime_time_zone`:(string) The time zone used for key lifetime configurations.* `UTC` - The Universal Time (UTC) for key lifetime configurations.* `Local` - The local time zone of the device for key lifetime configurations. 
    + `send_lifetime_unlimited`:(bool) Indicates that the key is always active. When this parameter is set, all other key lifetime configurations are ignored. 
* `include_icv_indicator`:(bool) Configures inclusion of the optional integrity check value (ICV) indicator as part of the transmitted MACsec key agreement protocol data unit (PDU). 
* `key_server_priority`:(int) The key server is selected by comparing key-server priority values during MACsec key agreement (MKA) message exchange between peer devices. Valid values range from 0 to 255. The lower the value, the higher the chance it will be selected as the key server. 
* `mac_sec_ea_pol`:(HashMap) - Extensible authentication protocol over LAN (EAPoL). MACsec transmits MACsec key agreement (MKA) protocol data units (PDUs) using EAPoL packets to establish a secure session. 
This complex property has following sub-properties:
  + `ea_pol_ethertype`:(string) Ethertype to use in extensible authentication protocol over LAN (EAPoL) frames for MACsec key agreement (MKA) protocol data units (PDUs). The range is between 0x600 - 0xffff. 
  + `ea_pol_mac_address`:(string) MAC address to use in extensible authentication protocol over LAN (EAPoL) for MACsec key agreement (MKA) protocol data units (PDUs). EAPol mac address should not be equal to all-zero (0000.0000.0000). 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `mod_time`:(string)(ReadOnly) The time when this managed object was last modified. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `organization`:(HashMap) - A reference to a organizationOrganization resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `owners`:
                (Array of schema.TypeString) -(ReadOnly)
* `parent`:(HashMap) -(ReadOnly) A reference to a moBaseMo resource.When the $expand query parameter is specified, the referenced resource is returned inline. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `permission_resources`:(Array)(ReadOnly) An array of relationships to moBaseMo resources. 
This complex property has following sub-properties:
  + `moid`:(string) The Moid of the referenced REST resource. 
  + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
  + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
* `primary_key_chain`:(HashMap) - Primary keychain for managing the default set of security keys for encryption and decryption. 
This complex property has following sub-properties:
  + `name`:(string) Must be a maximum of 63 characters, without spacing. 
  + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
  + `sec_keys`:(Array)
This complex property has following sub-properties:
    + `cryptographic_algorithm`:(string) The cryptographic algorithm that employs the cipher-based message authentication code (CMAC) mode of operation with advanced encryption standard (AES).* `AES_256_CMAC` - Uses the AES (Advanced Encryption Standard) algorithm with a 256-bit key to generate a CMAC.* `AES_128_CMAC` - Uses the AES (Advanced Encryption Standard) algorithm with a 128-bit key to generate a CMAC. 
    + `id`:(string) Must have an even number of hexadecimal characters (including 0-9 and A-F, only) with a length between 2 and 64 characters. For example, \ 10\ , \ 2000\ , \ ABCD1234\ . 
    + `is_octet_string_set`:(bool)(ReadOnly) Indicates whether the value of the 'octetString' property has been set. 
    + `key_type`:(string) The type of encryption used for the specified key.* `Type-0` - No encryption for the specified octetString.* `Type-6` - Proprietary advanced encryption standard for the specified octetString.* `Type-7` - Proprietary insecure encryption for the specified octetString. 
    + `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
    + `octet_string`:(string) The key octet string is a shared secret used in cryptographic operations. The valid size and format of the octet string depend on the selected KeyCryptographicAlgorithm and KeyEncryptionType. It should start with the character 'J'. 
    + `send_lifetime_duration`:(int) The key lifetime duration in seconds after the start time. If a non-zero value is configured for the duration, the end time configuration for the key lifetime is ignored. 
    + `send_lifetime_end_time`:(string) The time of day and date when the key becomes inactive. 
    + `send_lifetime_infinite`:(bool) Indicates that the key remains active indefinitely after the specified start time. When this parameter is set, the end time and duration configurations for the key lifetime are ignored. 
    + `send_lifetime_start_time`:(string) The time of day and date when the key becomes active. 
    + `send_lifetime_time_zone`:(string) The time zone used for key lifetime configurations.* `UTC` - The Universal Time (UTC) for key lifetime configurations.* `Local` - The local time zone of the device for key lifetime configurations. 
    + `send_lifetime_unlimited`:(bool) Indicates that the key is always active. When this parameter is set, all other key lifetime configurations are ignored. 
* `replay_window_size`:(int) Defines the size of the replay protection window. It determines the number of packets that can be received out of order without being considered replay attacks. 
* `sak_expiry_time`:(int) Time in seconds to force secure association key (SAK) rekey. Valid range is from 60 to 2592000 seconds when configured. When not configured, the SAK rekey interval is determined based on the interface speed. 
* `security_policy`:(string) The security policy specifies the level of MACsec enforcement on network traffic passing through a given interface.Should secure allows unencrypted traffic to flow until the MACsec key agreement (MKA) session is secured. After the MKA session is secured, the policy switches to only allow encrypted traffic to flow. Must secure imposes only MACsec encrypted traffic to flow. Traffic will be dropped, until the MKA session is not secured.* `Should-secure` - Should secure allows unencrypted traffic to flow until the MACsec key agreement (MKA) session is secured. After the MKA session is secured, the policy switches to only allow encrypted traffic to flow.* `Must-secure` - Must secure imposes only MACsec encrypted traffic to flow. Traffic will be dropped, until the MKA session is not secured. 
* `shared_scope`:(string)(ReadOnly) Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.Objects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. 
* `tags`:(Array)
This complex property has following sub-properties:
  + `key`:(string) The string representation of a tag key. 
  + `value`:(string) The string representation of a tag value. 
* `version_context`:(HashMap) -(ReadOnly) The versioning info for this managed object. 
This complex property has following sub-properties:
  + `interested_mos`:(Array)
This complex property has following sub-properties:
    + `moid`:(string) The Moid of the referenced REST resource. 
    + `object_type`:(string) The fully-qualified name of the remote type referred by this relationship. 
    + `selector`:(string) An OData $filter expression which describes the REST resource to be referenced. This field maybe set instead of 'moid' by clients.1. If 'moid' is set this field is ignored.1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of theresource matching the filter expression and populates it in the MoRef that is part of the objectinstance being inserted/updated to fulfill the REST request.An error is returned if the filter matches zero or more than one REST resource.An example filter string is: Serial eq '3AA8B7T11'. 
  + `marked_for_deletion`:(bool)(ReadOnly) The flag to indicate if snapshot is marked for deletion or not. If flag is set then snapshot will be removed after the successful deployment of the policy. 
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
`intersight_fabric_mac_sec_policy` can be imported using the Moid of the object, e.g.
```
$ terraform import intersight_fabric_mac_sec_policy.example 1234567890987654321abcde
``` 
