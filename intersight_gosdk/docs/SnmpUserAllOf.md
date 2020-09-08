# SnmpUserAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthPassword** | Pointer to **string** | Authorization password for the user. | [optional] 
**AuthType** | Pointer to **string** | Authorization protocol for authenticating the user. * &#x60;NA&#x60; - Authentication protocol is not applicable. * &#x60;MD5&#x60; - MD5 protocol is used to authenticate SNMP user. * &#x60;SHA&#x60; - SHA protocol is used to authenticate SNMP user. | [optional] [default to "NA"]
**IsAuthPasswordSet** | Pointer to **bool** | Indicates whether the value of the &#39;authPassword&#39; property has been set. | [optional] [readonly] 
**IsPrivacyPasswordSet** | Pointer to **bool** | Indicates whether the value of the &#39;privacyPassword&#39; property has been set. | [optional] [readonly] 
**Name** | Pointer to **string** | SNMP username. Must have a minimum of 1 and and a maximum of 31 characters. | [optional] 
**PrivacyPassword** | Pointer to **string** | Privacy password for the user. | [optional] 
**PrivacyType** | Pointer to **string** | Privacy protocol for the user. * &#x60;NA&#x60; - Privacy protocol is not applicable. * &#x60;DES&#x60; - DES privacy protocol is used for SNMP user. * &#x60;AES&#x60; - AES privacy protocol is used for SNMP user. | [optional] [default to "NA"]
**SecurityLevel** | Pointer to **string** | Security mechanism used for communication between agent and manager. * &#x60;AuthPriv&#x60; - The user requires both an authorization password and a privacy password. * &#x60;NoAuthNoPriv&#x60; - The user does not require an authorization or privacy password. * &#x60;AuthNoPriv&#x60; - The user requires an authorization password but not a privacy password. | [optional] [default to "AuthPriv"]

## Methods

### NewSnmpUserAllOf

`func NewSnmpUserAllOf() *SnmpUserAllOf`

NewSnmpUserAllOf instantiates a new SnmpUserAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnmpUserAllOfWithDefaults

`func NewSnmpUserAllOfWithDefaults() *SnmpUserAllOf`

NewSnmpUserAllOfWithDefaults instantiates a new SnmpUserAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthPassword

`func (o *SnmpUserAllOf) GetAuthPassword() string`

GetAuthPassword returns the AuthPassword field if non-nil, zero value otherwise.

### GetAuthPasswordOk

`func (o *SnmpUserAllOf) GetAuthPasswordOk() (*string, bool)`

GetAuthPasswordOk returns a tuple with the AuthPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPassword

`func (o *SnmpUserAllOf) SetAuthPassword(v string)`

SetAuthPassword sets AuthPassword field to given value.

### HasAuthPassword

`func (o *SnmpUserAllOf) HasAuthPassword() bool`

HasAuthPassword returns a boolean if a field has been set.

### GetAuthType

`func (o *SnmpUserAllOf) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *SnmpUserAllOf) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *SnmpUserAllOf) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *SnmpUserAllOf) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetIsAuthPasswordSet

`func (o *SnmpUserAllOf) GetIsAuthPasswordSet() bool`

GetIsAuthPasswordSet returns the IsAuthPasswordSet field if non-nil, zero value otherwise.

### GetIsAuthPasswordSetOk

`func (o *SnmpUserAllOf) GetIsAuthPasswordSetOk() (*bool, bool)`

GetIsAuthPasswordSetOk returns a tuple with the IsAuthPasswordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAuthPasswordSet

`func (o *SnmpUserAllOf) SetIsAuthPasswordSet(v bool)`

SetIsAuthPasswordSet sets IsAuthPasswordSet field to given value.

### HasIsAuthPasswordSet

`func (o *SnmpUserAllOf) HasIsAuthPasswordSet() bool`

HasIsAuthPasswordSet returns a boolean if a field has been set.

### GetIsPrivacyPasswordSet

`func (o *SnmpUserAllOf) GetIsPrivacyPasswordSet() bool`

GetIsPrivacyPasswordSet returns the IsPrivacyPasswordSet field if non-nil, zero value otherwise.

### GetIsPrivacyPasswordSetOk

`func (o *SnmpUserAllOf) GetIsPrivacyPasswordSetOk() (*bool, bool)`

GetIsPrivacyPasswordSetOk returns a tuple with the IsPrivacyPasswordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivacyPasswordSet

`func (o *SnmpUserAllOf) SetIsPrivacyPasswordSet(v bool)`

SetIsPrivacyPasswordSet sets IsPrivacyPasswordSet field to given value.

### HasIsPrivacyPasswordSet

`func (o *SnmpUserAllOf) HasIsPrivacyPasswordSet() bool`

HasIsPrivacyPasswordSet returns a boolean if a field has been set.

### GetName

`func (o *SnmpUserAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SnmpUserAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SnmpUserAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SnmpUserAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrivacyPassword

`func (o *SnmpUserAllOf) GetPrivacyPassword() string`

GetPrivacyPassword returns the PrivacyPassword field if non-nil, zero value otherwise.

### GetPrivacyPasswordOk

`func (o *SnmpUserAllOf) GetPrivacyPasswordOk() (*string, bool)`

GetPrivacyPasswordOk returns a tuple with the PrivacyPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyPassword

`func (o *SnmpUserAllOf) SetPrivacyPassword(v string)`

SetPrivacyPassword sets PrivacyPassword field to given value.

### HasPrivacyPassword

`func (o *SnmpUserAllOf) HasPrivacyPassword() bool`

HasPrivacyPassword returns a boolean if a field has been set.

### GetPrivacyType

`func (o *SnmpUserAllOf) GetPrivacyType() string`

GetPrivacyType returns the PrivacyType field if non-nil, zero value otherwise.

### GetPrivacyTypeOk

`func (o *SnmpUserAllOf) GetPrivacyTypeOk() (*string, bool)`

GetPrivacyTypeOk returns a tuple with the PrivacyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyType

`func (o *SnmpUserAllOf) SetPrivacyType(v string)`

SetPrivacyType sets PrivacyType field to given value.

### HasPrivacyType

`func (o *SnmpUserAllOf) HasPrivacyType() bool`

HasPrivacyType returns a boolean if a field has been set.

### GetSecurityLevel

`func (o *SnmpUserAllOf) GetSecurityLevel() string`

GetSecurityLevel returns the SecurityLevel field if non-nil, zero value otherwise.

### GetSecurityLevelOk

`func (o *SnmpUserAllOf) GetSecurityLevelOk() (*string, bool)`

GetSecurityLevelOk returns a tuple with the SecurityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityLevel

`func (o *SnmpUserAllOf) SetSecurityLevel(v string)`

SetSecurityLevel sets SecurityLevel field to given value.

### HasSecurityLevel

`func (o *SnmpUserAllOf) HasSecurityLevel() bool`

HasSecurityLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


