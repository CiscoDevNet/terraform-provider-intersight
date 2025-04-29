# EtherMacsecOperData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "ether.MacsecOperData"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "ether.MacsecOperData"]
**AuthMode** | Pointer to **string** | The authentication mode used for MACsec encryption. | [optional] [readonly] 
**CipherSuite** | Pointer to **string** | Cipher suite to be used for MACsec encryption. | [optional] [readonly] 
**ConfidentialityOffset** | Pointer to **string** | The MACsec confidentiality offset specifies the number of bytes starting from the frame header. MACsec encrypts only the bytes after the offset in a frame. | [optional] [readonly] 
**KeyServer** | Pointer to **string** | The value indicates that the device is acting as a key server, responsible for distributing encryption keys to other devices in the MACsec-enabled session. | [optional] [readonly] 
**SessionState** | Pointer to **string** | The state of the MACsec session. | [optional] [readonly] 
**StateReason** | Pointer to **string** | The reason for the MACsec session state. | [optional] [readonly] 

## Methods

### NewEtherMacsecOperData

`func NewEtherMacsecOperData(classId string, objectType string, ) *EtherMacsecOperData`

NewEtherMacsecOperData instantiates a new EtherMacsecOperData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEtherMacsecOperDataWithDefaults

`func NewEtherMacsecOperDataWithDefaults() *EtherMacsecOperData`

NewEtherMacsecOperDataWithDefaults instantiates a new EtherMacsecOperData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *EtherMacsecOperData) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *EtherMacsecOperData) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *EtherMacsecOperData) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *EtherMacsecOperData) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EtherMacsecOperData) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EtherMacsecOperData) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAuthMode

`func (o *EtherMacsecOperData) GetAuthMode() string`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *EtherMacsecOperData) GetAuthModeOk() (*string, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *EtherMacsecOperData) SetAuthMode(v string)`

SetAuthMode sets AuthMode field to given value.

### HasAuthMode

`func (o *EtherMacsecOperData) HasAuthMode() bool`

HasAuthMode returns a boolean if a field has been set.

### GetCipherSuite

`func (o *EtherMacsecOperData) GetCipherSuite() string`

GetCipherSuite returns the CipherSuite field if non-nil, zero value otherwise.

### GetCipherSuiteOk

`func (o *EtherMacsecOperData) GetCipherSuiteOk() (*string, bool)`

GetCipherSuiteOk returns a tuple with the CipherSuite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCipherSuite

`func (o *EtherMacsecOperData) SetCipherSuite(v string)`

SetCipherSuite sets CipherSuite field to given value.

### HasCipherSuite

`func (o *EtherMacsecOperData) HasCipherSuite() bool`

HasCipherSuite returns a boolean if a field has been set.

### GetConfidentialityOffset

`func (o *EtherMacsecOperData) GetConfidentialityOffset() string`

GetConfidentialityOffset returns the ConfidentialityOffset field if non-nil, zero value otherwise.

### GetConfidentialityOffsetOk

`func (o *EtherMacsecOperData) GetConfidentialityOffsetOk() (*string, bool)`

GetConfidentialityOffsetOk returns a tuple with the ConfidentialityOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidentialityOffset

`func (o *EtherMacsecOperData) SetConfidentialityOffset(v string)`

SetConfidentialityOffset sets ConfidentialityOffset field to given value.

### HasConfidentialityOffset

`func (o *EtherMacsecOperData) HasConfidentialityOffset() bool`

HasConfidentialityOffset returns a boolean if a field has been set.

### GetKeyServer

`func (o *EtherMacsecOperData) GetKeyServer() string`

GetKeyServer returns the KeyServer field if non-nil, zero value otherwise.

### GetKeyServerOk

`func (o *EtherMacsecOperData) GetKeyServerOk() (*string, bool)`

GetKeyServerOk returns a tuple with the KeyServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyServer

`func (o *EtherMacsecOperData) SetKeyServer(v string)`

SetKeyServer sets KeyServer field to given value.

### HasKeyServer

`func (o *EtherMacsecOperData) HasKeyServer() bool`

HasKeyServer returns a boolean if a field has been set.

### GetSessionState

`func (o *EtherMacsecOperData) GetSessionState() string`

GetSessionState returns the SessionState field if non-nil, zero value otherwise.

### GetSessionStateOk

`func (o *EtherMacsecOperData) GetSessionStateOk() (*string, bool)`

GetSessionStateOk returns a tuple with the SessionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionState

`func (o *EtherMacsecOperData) SetSessionState(v string)`

SetSessionState sets SessionState field to given value.

### HasSessionState

`func (o *EtherMacsecOperData) HasSessionState() bool`

HasSessionState returns a boolean if a field has been set.

### GetStateReason

`func (o *EtherMacsecOperData) GetStateReason() string`

GetStateReason returns the StateReason field if non-nil, zero value otherwise.

### GetStateReasonOk

`func (o *EtherMacsecOperData) GetStateReasonOk() (*string, bool)`

GetStateReasonOk returns a tuple with the StateReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateReason

`func (o *EtherMacsecOperData) SetStateReason(v string)`

SetStateReason sets StateReason field to given value.

### HasStateReason

`func (o *EtherMacsecOperData) HasStateReason() bool`

HasStateReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


