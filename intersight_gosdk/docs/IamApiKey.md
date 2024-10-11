# IamApiKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.ApiKey"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.ApiKey"]
**AdminStatus** | Pointer to **string** | Used to trigger the enable or disable action on the API key. These actions change the status of an API key. * &#x60;enable&#x60; - Used to enable a disabled API key/App Registration. If the API key/App Registration is already expired, this action has no effect. * &#x60;disable&#x60; - Used to disable an active API key/App Registration. If the API key/App Registration is already expired, this action has no effect. | [optional] [default to "enable"]
**ExpiryDateTime** | Pointer to **time.Time** | The expiration date of the API key which is set at the time of creation of the key. Its value can only be assigned a date that falls within the range determined by the maximum expiration time configured at the account level. The expiry date can be edited to be earlier or later, provided it stays within the designated expiry period. This period is determined by adding the &#39;startTime&#39; property of the API key to the maximum expiry time configured at the account level. | [optional] 
**HashAlgorithm** | Pointer to **string** | The cryptographic hash algorithm to calculate the message digest. * &#x60;SHA256&#x60; - The SHA-256 cryptographic hash, as defined by NIST in FIPS 180-4. * &#x60;SHA384&#x60; - The SHA-384 cryptographic hash, as defined by NIST in FIPS 180-4. * &#x60;SHA512&#x60; - The SHA-512 cryptographic hash, as defined by NIST in FIPS 180-4. * &#x60;SHA512_224&#x60; - The SHA-512/224 cryptographic hash, as defined by NIST in FIPS 180-4. * &#x60;SHA512_256&#x60; - The SHA-512/256 cryptographic hash, as defined by NIST in FIPS 180-4. | [optional] [default to "SHA256"]
**IsNeverExpiring** | Pointer to **bool** | Used to mark the API key as a never-expiring API key. | [optional] [default to false]
**KeySpec** | Pointer to [**NullablePkixKeyGenerationSpec**](PkixKeyGenerationSpec.md) |  | [optional] 
**LastUsedIp** | Pointer to **string** | The IP address from which the API key was last used. | [optional] [readonly] 
**LastUsedTime** | Pointer to **time.Time** | The time at which the API key was last used. It is updated every 24 hours. | [optional] [readonly] 
**OperStatus** | Pointer to **string** | The current status of the API key that dictates the validity of the key. * &#x60;enabled&#x60; - An API key/App Registration having enabled status can be used for API invocation. * &#x60;disabled&#x60; - An API key/App Registration having disabled status cannot be used for API invocation. * &#x60;expired&#x60; - An API key/App Registration having expired status cannot be used for API invocation as the expiration date has passed. | [optional] [readonly] [default to "enabled"]
**PrivateKey** | Pointer to **string** | Holds the private key for the API key. | [optional] 
**Purpose** | Pointer to **string** | The purpose of the API Key. | [optional] 
**Scope** | Pointer to [**NullableIamSwitchScopePermissions**](IamSwitchScopePermissions.md) |  | [optional] 
**SigningAlgorithm** | Pointer to **string** | The signing algorithm used by the client to authenticate API requests to Intersight. The signing algorithm must be compatible with the key generation specification. * &#x60;RSASSA-PKCS1-v1_5&#x60; - RSASSA-PKCS1-v1_5 is a RSA signature scheme specified in [RFC 8017](https://tools.ietf.org/html/rfc8017).RSASSA-PKCS1-v1_5 is included only for compatibility with existing applications. * &#x60;RSASSA-PSS&#x60; - RSASSA-PSS is a RSA signature scheme specified in [RFC 8017](https://tools.ietf.org/html/rfc8017).It combines the RSASP1 and RSAVP1 primitives with the EMSA-PSS encoding method.In the interest of increased robustness, RSASSA-PSS is required in new applications. * &#x60;Ed25519&#x60; - The Ed25519 signature algorithm, as specified in [RFC 8032](https://tools.ietf.org/html/rfc8032).Ed25519 is a public-key signature system with several attractive features, includingfast single-signature verification, very fast signing, fast key generation and high security level. * &#x60;Ecdsa&#x60; - The Elliptic Curve Digital Signature Standard (ECDSA), as defined by NIST in FIPS 186-4 and ANSI X9.62.The signature is encoded as a ASN.1 DER SEQUENCE with two INTEGERs (r and s), as defined in RFC3279.When using ECDSA signatures, configure the client to use the same signature encoding as specified on the server side. * &#x60;EcdsaP1363Format&#x60; - The Elliptic Curve Digital Signature Standard (ECDSA), as defined by NIST in FIPS 186-4 and ANSI X9.62.The signature is the raw concatenation of r and s, as defined in the ISO/IEC 7816-8 IEEE P.1363 standard.In that format, r and s are represented as unsigned, big endian numbers.Extra padding bytes (of value 0x00) is applied so that both r and s encodings have the same size.When using ECDSA signatures, configure the client to use the same signature encoding as specified on the server side. | [optional] [default to "RSASSA-PKCS1-v1_5"]
**StartTime** | Pointer to **time.Time** | The timestamp at which an expiry date was first set on this API key. For expiring API keys, this field is same as the create time of the API key. For never-expiring API keys, this field is set initially to zero time value. If a never-expiry API key is later changed to have an expiration, the timestamp marking the start of this transition is recorded in this field. | [optional] [readonly] 
**Permission** | Pointer to [**NullableIamPermissionRelationship**](IamPermissionRelationship.md) |  | [optional] 
**User** | Pointer to [**NullableIamUserRelationship**](IamUserRelationship.md) |  | [optional] 

## Methods

### NewIamApiKey

`func NewIamApiKey(classId string, objectType string, ) *IamApiKey`

NewIamApiKey instantiates a new IamApiKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamApiKeyWithDefaults

`func NewIamApiKeyWithDefaults() *IamApiKey`

NewIamApiKeyWithDefaults instantiates a new IamApiKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamApiKey) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamApiKey) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamApiKey) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamApiKey) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamApiKey) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamApiKey) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdminStatus

`func (o *IamApiKey) GetAdminStatus() string`

GetAdminStatus returns the AdminStatus field if non-nil, zero value otherwise.

### GetAdminStatusOk

`func (o *IamApiKey) GetAdminStatusOk() (*string, bool)`

GetAdminStatusOk returns a tuple with the AdminStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminStatus

`func (o *IamApiKey) SetAdminStatus(v string)`

SetAdminStatus sets AdminStatus field to given value.

### HasAdminStatus

`func (o *IamApiKey) HasAdminStatus() bool`

HasAdminStatus returns a boolean if a field has been set.

### GetExpiryDateTime

`func (o *IamApiKey) GetExpiryDateTime() time.Time`

GetExpiryDateTime returns the ExpiryDateTime field if non-nil, zero value otherwise.

### GetExpiryDateTimeOk

`func (o *IamApiKey) GetExpiryDateTimeOk() (*time.Time, bool)`

GetExpiryDateTimeOk returns a tuple with the ExpiryDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDateTime

`func (o *IamApiKey) SetExpiryDateTime(v time.Time)`

SetExpiryDateTime sets ExpiryDateTime field to given value.

### HasExpiryDateTime

`func (o *IamApiKey) HasExpiryDateTime() bool`

HasExpiryDateTime returns a boolean if a field has been set.

### GetHashAlgorithm

`func (o *IamApiKey) GetHashAlgorithm() string`

GetHashAlgorithm returns the HashAlgorithm field if non-nil, zero value otherwise.

### GetHashAlgorithmOk

`func (o *IamApiKey) GetHashAlgorithmOk() (*string, bool)`

GetHashAlgorithmOk returns a tuple with the HashAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashAlgorithm

`func (o *IamApiKey) SetHashAlgorithm(v string)`

SetHashAlgorithm sets HashAlgorithm field to given value.

### HasHashAlgorithm

`func (o *IamApiKey) HasHashAlgorithm() bool`

HasHashAlgorithm returns a boolean if a field has been set.

### GetIsNeverExpiring

`func (o *IamApiKey) GetIsNeverExpiring() bool`

GetIsNeverExpiring returns the IsNeverExpiring field if non-nil, zero value otherwise.

### GetIsNeverExpiringOk

`func (o *IamApiKey) GetIsNeverExpiringOk() (*bool, bool)`

GetIsNeverExpiringOk returns a tuple with the IsNeverExpiring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNeverExpiring

`func (o *IamApiKey) SetIsNeverExpiring(v bool)`

SetIsNeverExpiring sets IsNeverExpiring field to given value.

### HasIsNeverExpiring

`func (o *IamApiKey) HasIsNeverExpiring() bool`

HasIsNeverExpiring returns a boolean if a field has been set.

### GetKeySpec

`func (o *IamApiKey) GetKeySpec() PkixKeyGenerationSpec`

GetKeySpec returns the KeySpec field if non-nil, zero value otherwise.

### GetKeySpecOk

`func (o *IamApiKey) GetKeySpecOk() (*PkixKeyGenerationSpec, bool)`

GetKeySpecOk returns a tuple with the KeySpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySpec

`func (o *IamApiKey) SetKeySpec(v PkixKeyGenerationSpec)`

SetKeySpec sets KeySpec field to given value.

### HasKeySpec

`func (o *IamApiKey) HasKeySpec() bool`

HasKeySpec returns a boolean if a field has been set.

### SetKeySpecNil

`func (o *IamApiKey) SetKeySpecNil(b bool)`

 SetKeySpecNil sets the value for KeySpec to be an explicit nil

### UnsetKeySpec
`func (o *IamApiKey) UnsetKeySpec()`

UnsetKeySpec ensures that no value is present for KeySpec, not even an explicit nil
### GetLastUsedIp

`func (o *IamApiKey) GetLastUsedIp() string`

GetLastUsedIp returns the LastUsedIp field if non-nil, zero value otherwise.

### GetLastUsedIpOk

`func (o *IamApiKey) GetLastUsedIpOk() (*string, bool)`

GetLastUsedIpOk returns a tuple with the LastUsedIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedIp

`func (o *IamApiKey) SetLastUsedIp(v string)`

SetLastUsedIp sets LastUsedIp field to given value.

### HasLastUsedIp

`func (o *IamApiKey) HasLastUsedIp() bool`

HasLastUsedIp returns a boolean if a field has been set.

### GetLastUsedTime

`func (o *IamApiKey) GetLastUsedTime() time.Time`

GetLastUsedTime returns the LastUsedTime field if non-nil, zero value otherwise.

### GetLastUsedTimeOk

`func (o *IamApiKey) GetLastUsedTimeOk() (*time.Time, bool)`

GetLastUsedTimeOk returns a tuple with the LastUsedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedTime

`func (o *IamApiKey) SetLastUsedTime(v time.Time)`

SetLastUsedTime sets LastUsedTime field to given value.

### HasLastUsedTime

`func (o *IamApiKey) HasLastUsedTime() bool`

HasLastUsedTime returns a boolean if a field has been set.

### GetOperStatus

`func (o *IamApiKey) GetOperStatus() string`

GetOperStatus returns the OperStatus field if non-nil, zero value otherwise.

### GetOperStatusOk

`func (o *IamApiKey) GetOperStatusOk() (*string, bool)`

GetOperStatusOk returns a tuple with the OperStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperStatus

`func (o *IamApiKey) SetOperStatus(v string)`

SetOperStatus sets OperStatus field to given value.

### HasOperStatus

`func (o *IamApiKey) HasOperStatus() bool`

HasOperStatus returns a boolean if a field has been set.

### GetPrivateKey

`func (o *IamApiKey) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *IamApiKey) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *IamApiKey) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *IamApiKey) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetPurpose

`func (o *IamApiKey) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *IamApiKey) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *IamApiKey) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *IamApiKey) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### GetScope

`func (o *IamApiKey) GetScope() IamSwitchScopePermissions`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *IamApiKey) GetScopeOk() (*IamSwitchScopePermissions, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *IamApiKey) SetScope(v IamSwitchScopePermissions)`

SetScope sets Scope field to given value.

### HasScope

`func (o *IamApiKey) HasScope() bool`

HasScope returns a boolean if a field has been set.

### SetScopeNil

`func (o *IamApiKey) SetScopeNil(b bool)`

 SetScopeNil sets the value for Scope to be an explicit nil

### UnsetScope
`func (o *IamApiKey) UnsetScope()`

UnsetScope ensures that no value is present for Scope, not even an explicit nil
### GetSigningAlgorithm

`func (o *IamApiKey) GetSigningAlgorithm() string`

GetSigningAlgorithm returns the SigningAlgorithm field if non-nil, zero value otherwise.

### GetSigningAlgorithmOk

`func (o *IamApiKey) GetSigningAlgorithmOk() (*string, bool)`

GetSigningAlgorithmOk returns a tuple with the SigningAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningAlgorithm

`func (o *IamApiKey) SetSigningAlgorithm(v string)`

SetSigningAlgorithm sets SigningAlgorithm field to given value.

### HasSigningAlgorithm

`func (o *IamApiKey) HasSigningAlgorithm() bool`

HasSigningAlgorithm returns a boolean if a field has been set.

### GetStartTime

`func (o *IamApiKey) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *IamApiKey) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *IamApiKey) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *IamApiKey) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetPermission

`func (o *IamApiKey) GetPermission() IamPermissionRelationship`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *IamApiKey) GetPermissionOk() (*IamPermissionRelationship, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *IamApiKey) SetPermission(v IamPermissionRelationship)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *IamApiKey) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### SetPermissionNil

`func (o *IamApiKey) SetPermissionNil(b bool)`

 SetPermissionNil sets the value for Permission to be an explicit nil

### UnsetPermission
`func (o *IamApiKey) UnsetPermission()`

UnsetPermission ensures that no value is present for Permission, not even an explicit nil
### GetUser

`func (o *IamApiKey) GetUser() IamUserRelationship`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *IamApiKey) GetUserOk() (*IamUserRelationship, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *IamApiKey) SetUser(v IamUserRelationship)`

SetUser sets User field to given value.

### HasUser

`func (o *IamApiKey) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *IamApiKey) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *IamApiKey) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


