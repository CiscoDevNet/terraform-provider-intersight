# IamApiKeyAllOf

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
**SigningAlgorithm** | Pointer to **string** | The signing algorithm used by the client to authenticate API requests to Intersight. The signing algorithm must be compatible with the key generation specification. * &#x60;RSASSA-PKCS1-v1_5&#x60; - RSASSA-PKCS1-v1_5 is a RSA signature scheme specified in [RFC 8017](https://tools.ietf.org/html/rfc8017).RSASSA-PKCS1-v1_5 is included only for compatibility with existing applications. * &#x60;RSASSA-PSS&#x60; - RSASSA-PSS is a RSA signature scheme specified in [RFC 8017](https://tools.ietf.org/html/rfc8017).It combines the RSASP1 and RSAVP1 primitives with the EMSA-PSS encoding method.In the interest of increased robustness, RSASSA-PSS is required in new applications. * &#x60;Ed25519&#x60; - The Ed25519 signature algorithm, as specified in [RFC 8032](https://tools.ietf.org/html/rfc8032).Ed25519 is a public-key signature system with several attractive features, includingfast single-signature verification, very fast signing, fast key generation and high security level. * &#x60;Ecdsa&#x60; - The Elliptic Curve Digital Signature Standard (ECDSA), as defined by NIST in FIPS 186-4 and ANSI X9.62.The signature is encoded as a ASN.1 DER SEQUENCE with two INTEGERs (r and s), as defined in RFC3279.When using ECDSA signatures, configure the client to use the same signature encoding as specified on the server side. * &#x60;EcdsaP1363Format&#x60; - The Elliptic Curve Digital Signature Standard (ECDSA), as defined by NIST in FIPS 186-4 and ANSI X9.62.The signature is the raw concatenation of r and s, as defined in the ISO/IEC 7816-8 IEEE P.1363 standard.In that format, r and s are represented as unsigned, big endian numbers.Extra padding bytes (of value 0x00) is applied so that both r and s encodings have the same size.When using ECDSA signatures, configure the client to use the same signature encoding as specified on the server side. | [optional] [default to "RSASSA-PKCS1-v1_5"]
**StartTime** | Pointer to **time.Time** | The timestamp at which an expiry date was first set on this API key. For expiring API keys, this field is same as the create time of the API key. For never-expiring API keys, this field is set initially to zero time value. If a never-expiry API key is later changed to have an expiration, the timestamp marking the start of this transition is recorded in this field. | [optional] [readonly] 
**Permission** | Pointer to [**IamPermissionRelationship**](IamPermissionRelationship.md) |  | [optional] 
**User** | Pointer to [**IamUserRelationship**](IamUserRelationship.md) |  | [optional] 

## Methods

### NewIamApiKeyAllOf

`func NewIamApiKeyAllOf(classId string, objectType string, ) *IamApiKeyAllOf`

NewIamApiKeyAllOf instantiates a new IamApiKeyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamApiKeyAllOfWithDefaults

`func NewIamApiKeyAllOfWithDefaults() *IamApiKeyAllOf`

NewIamApiKeyAllOfWithDefaults instantiates a new IamApiKeyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamApiKeyAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamApiKeyAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamApiKeyAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamApiKeyAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamApiKeyAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamApiKeyAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdminStatus

`func (o *IamApiKeyAllOf) GetAdminStatus() string`

GetAdminStatus returns the AdminStatus field if non-nil, zero value otherwise.

### GetAdminStatusOk

`func (o *IamApiKeyAllOf) GetAdminStatusOk() (*string, bool)`

GetAdminStatusOk returns a tuple with the AdminStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminStatus

`func (o *IamApiKeyAllOf) SetAdminStatus(v string)`

SetAdminStatus sets AdminStatus field to given value.

### HasAdminStatus

`func (o *IamApiKeyAllOf) HasAdminStatus() bool`

HasAdminStatus returns a boolean if a field has been set.

### GetExpiryDateTime

`func (o *IamApiKeyAllOf) GetExpiryDateTime() time.Time`

GetExpiryDateTime returns the ExpiryDateTime field if non-nil, zero value otherwise.

### GetExpiryDateTimeOk

`func (o *IamApiKeyAllOf) GetExpiryDateTimeOk() (*time.Time, bool)`

GetExpiryDateTimeOk returns a tuple with the ExpiryDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDateTime

`func (o *IamApiKeyAllOf) SetExpiryDateTime(v time.Time)`

SetExpiryDateTime sets ExpiryDateTime field to given value.

### HasExpiryDateTime

`func (o *IamApiKeyAllOf) HasExpiryDateTime() bool`

HasExpiryDateTime returns a boolean if a field has been set.

### GetHashAlgorithm

`func (o *IamApiKeyAllOf) GetHashAlgorithm() string`

GetHashAlgorithm returns the HashAlgorithm field if non-nil, zero value otherwise.

### GetHashAlgorithmOk

`func (o *IamApiKeyAllOf) GetHashAlgorithmOk() (*string, bool)`

GetHashAlgorithmOk returns a tuple with the HashAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashAlgorithm

`func (o *IamApiKeyAllOf) SetHashAlgorithm(v string)`

SetHashAlgorithm sets HashAlgorithm field to given value.

### HasHashAlgorithm

`func (o *IamApiKeyAllOf) HasHashAlgorithm() bool`

HasHashAlgorithm returns a boolean if a field has been set.

### GetIsNeverExpiring

`func (o *IamApiKeyAllOf) GetIsNeverExpiring() bool`

GetIsNeverExpiring returns the IsNeverExpiring field if non-nil, zero value otherwise.

### GetIsNeverExpiringOk

`func (o *IamApiKeyAllOf) GetIsNeverExpiringOk() (*bool, bool)`

GetIsNeverExpiringOk returns a tuple with the IsNeverExpiring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNeverExpiring

`func (o *IamApiKeyAllOf) SetIsNeverExpiring(v bool)`

SetIsNeverExpiring sets IsNeverExpiring field to given value.

### HasIsNeverExpiring

`func (o *IamApiKeyAllOf) HasIsNeverExpiring() bool`

HasIsNeverExpiring returns a boolean if a field has been set.

### GetKeySpec

`func (o *IamApiKeyAllOf) GetKeySpec() PkixKeyGenerationSpec`

GetKeySpec returns the KeySpec field if non-nil, zero value otherwise.

### GetKeySpecOk

`func (o *IamApiKeyAllOf) GetKeySpecOk() (*PkixKeyGenerationSpec, bool)`

GetKeySpecOk returns a tuple with the KeySpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySpec

`func (o *IamApiKeyAllOf) SetKeySpec(v PkixKeyGenerationSpec)`

SetKeySpec sets KeySpec field to given value.

### HasKeySpec

`func (o *IamApiKeyAllOf) HasKeySpec() bool`

HasKeySpec returns a boolean if a field has been set.

### SetKeySpecNil

`func (o *IamApiKeyAllOf) SetKeySpecNil(b bool)`

 SetKeySpecNil sets the value for KeySpec to be an explicit nil

### UnsetKeySpec
`func (o *IamApiKeyAllOf) UnsetKeySpec()`

UnsetKeySpec ensures that no value is present for KeySpec, not even an explicit nil
### GetLastUsedIp

`func (o *IamApiKeyAllOf) GetLastUsedIp() string`

GetLastUsedIp returns the LastUsedIp field if non-nil, zero value otherwise.

### GetLastUsedIpOk

`func (o *IamApiKeyAllOf) GetLastUsedIpOk() (*string, bool)`

GetLastUsedIpOk returns a tuple with the LastUsedIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedIp

`func (o *IamApiKeyAllOf) SetLastUsedIp(v string)`

SetLastUsedIp sets LastUsedIp field to given value.

### HasLastUsedIp

`func (o *IamApiKeyAllOf) HasLastUsedIp() bool`

HasLastUsedIp returns a boolean if a field has been set.

### GetLastUsedTime

`func (o *IamApiKeyAllOf) GetLastUsedTime() time.Time`

GetLastUsedTime returns the LastUsedTime field if non-nil, zero value otherwise.

### GetLastUsedTimeOk

`func (o *IamApiKeyAllOf) GetLastUsedTimeOk() (*time.Time, bool)`

GetLastUsedTimeOk returns a tuple with the LastUsedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedTime

`func (o *IamApiKeyAllOf) SetLastUsedTime(v time.Time)`

SetLastUsedTime sets LastUsedTime field to given value.

### HasLastUsedTime

`func (o *IamApiKeyAllOf) HasLastUsedTime() bool`

HasLastUsedTime returns a boolean if a field has been set.

### GetOperStatus

`func (o *IamApiKeyAllOf) GetOperStatus() string`

GetOperStatus returns the OperStatus field if non-nil, zero value otherwise.

### GetOperStatusOk

`func (o *IamApiKeyAllOf) GetOperStatusOk() (*string, bool)`

GetOperStatusOk returns a tuple with the OperStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperStatus

`func (o *IamApiKeyAllOf) SetOperStatus(v string)`

SetOperStatus sets OperStatus field to given value.

### HasOperStatus

`func (o *IamApiKeyAllOf) HasOperStatus() bool`

HasOperStatus returns a boolean if a field has been set.

### GetPrivateKey

`func (o *IamApiKeyAllOf) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *IamApiKeyAllOf) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *IamApiKeyAllOf) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *IamApiKeyAllOf) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetPurpose

`func (o *IamApiKeyAllOf) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *IamApiKeyAllOf) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *IamApiKeyAllOf) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *IamApiKeyAllOf) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### GetSigningAlgorithm

`func (o *IamApiKeyAllOf) GetSigningAlgorithm() string`

GetSigningAlgorithm returns the SigningAlgorithm field if non-nil, zero value otherwise.

### GetSigningAlgorithmOk

`func (o *IamApiKeyAllOf) GetSigningAlgorithmOk() (*string, bool)`

GetSigningAlgorithmOk returns a tuple with the SigningAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningAlgorithm

`func (o *IamApiKeyAllOf) SetSigningAlgorithm(v string)`

SetSigningAlgorithm sets SigningAlgorithm field to given value.

### HasSigningAlgorithm

`func (o *IamApiKeyAllOf) HasSigningAlgorithm() bool`

HasSigningAlgorithm returns a boolean if a field has been set.

### GetStartTime

`func (o *IamApiKeyAllOf) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *IamApiKeyAllOf) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *IamApiKeyAllOf) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *IamApiKeyAllOf) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetPermission

`func (o *IamApiKeyAllOf) GetPermission() IamPermissionRelationship`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *IamApiKeyAllOf) GetPermissionOk() (*IamPermissionRelationship, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *IamApiKeyAllOf) SetPermission(v IamPermissionRelationship)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *IamApiKeyAllOf) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetUser

`func (o *IamApiKeyAllOf) GetUser() IamUserRelationship`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *IamApiKeyAllOf) GetUserOk() (*IamUserRelationship, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *IamApiKeyAllOf) SetUser(v IamUserRelationship)`

SetUser sets User field to given value.

### HasUser

`func (o *IamApiKeyAllOf) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


