# FirmwareEula

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "firmware.Eula"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "firmware.Eula"]
**Accepted** | Pointer to **bool** | Overall acceptance status for the account, both EULA and K9. | [optional] [readonly] 
**Content** | Pointer to **string** | Acceptance form content provided by cisco.com. | [optional] [readonly] 
**EulaAccepted** | Pointer to **bool** | EULA acceptance status for the account. | [optional] [readonly] 
**EulaContent** | Pointer to **string** | EULA acceptance form content provided by cisco.com. | [optional] [readonly] 
**K9Accepted** | Pointer to **bool** | K9 acceptance status for the account. | [optional] [readonly] 
**K9Content** | Pointer to **string** | K9 acceptance form content provided by cisco.com. | [optional] [readonly] 
**Account** | Pointer to [**NullableIamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 

## Methods

### NewFirmwareEula

`func NewFirmwareEula(classId string, objectType string, ) *FirmwareEula`

NewFirmwareEula instantiates a new FirmwareEula object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareEulaWithDefaults

`func NewFirmwareEulaWithDefaults() *FirmwareEula`

NewFirmwareEulaWithDefaults instantiates a new FirmwareEula object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FirmwareEula) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FirmwareEula) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FirmwareEula) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FirmwareEula) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FirmwareEula) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FirmwareEula) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAccepted

`func (o *FirmwareEula) GetAccepted() bool`

GetAccepted returns the Accepted field if non-nil, zero value otherwise.

### GetAcceptedOk

`func (o *FirmwareEula) GetAcceptedOk() (*bool, bool)`

GetAcceptedOk returns a tuple with the Accepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccepted

`func (o *FirmwareEula) SetAccepted(v bool)`

SetAccepted sets Accepted field to given value.

### HasAccepted

`func (o *FirmwareEula) HasAccepted() bool`

HasAccepted returns a boolean if a field has been set.

### GetContent

`func (o *FirmwareEula) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *FirmwareEula) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *FirmwareEula) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *FirmwareEula) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetEulaAccepted

`func (o *FirmwareEula) GetEulaAccepted() bool`

GetEulaAccepted returns the EulaAccepted field if non-nil, zero value otherwise.

### GetEulaAcceptedOk

`func (o *FirmwareEula) GetEulaAcceptedOk() (*bool, bool)`

GetEulaAcceptedOk returns a tuple with the EulaAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEulaAccepted

`func (o *FirmwareEula) SetEulaAccepted(v bool)`

SetEulaAccepted sets EulaAccepted field to given value.

### HasEulaAccepted

`func (o *FirmwareEula) HasEulaAccepted() bool`

HasEulaAccepted returns a boolean if a field has been set.

### GetEulaContent

`func (o *FirmwareEula) GetEulaContent() string`

GetEulaContent returns the EulaContent field if non-nil, zero value otherwise.

### GetEulaContentOk

`func (o *FirmwareEula) GetEulaContentOk() (*string, bool)`

GetEulaContentOk returns a tuple with the EulaContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEulaContent

`func (o *FirmwareEula) SetEulaContent(v string)`

SetEulaContent sets EulaContent field to given value.

### HasEulaContent

`func (o *FirmwareEula) HasEulaContent() bool`

HasEulaContent returns a boolean if a field has been set.

### GetK9Accepted

`func (o *FirmwareEula) GetK9Accepted() bool`

GetK9Accepted returns the K9Accepted field if non-nil, zero value otherwise.

### GetK9AcceptedOk

`func (o *FirmwareEula) GetK9AcceptedOk() (*bool, bool)`

GetK9AcceptedOk returns a tuple with the K9Accepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK9Accepted

`func (o *FirmwareEula) SetK9Accepted(v bool)`

SetK9Accepted sets K9Accepted field to given value.

### HasK9Accepted

`func (o *FirmwareEula) HasK9Accepted() bool`

HasK9Accepted returns a boolean if a field has been set.

### GetK9Content

`func (o *FirmwareEula) GetK9Content() string`

GetK9Content returns the K9Content field if non-nil, zero value otherwise.

### GetK9ContentOk

`func (o *FirmwareEula) GetK9ContentOk() (*string, bool)`

GetK9ContentOk returns a tuple with the K9Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK9Content

`func (o *FirmwareEula) SetK9Content(v string)`

SetK9Content sets K9Content field to given value.

### HasK9Content

`func (o *FirmwareEula) HasK9Content() bool`

HasK9Content returns a boolean if a field has been set.

### GetAccount

`func (o *FirmwareEula) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *FirmwareEula) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *FirmwareEula) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *FirmwareEula) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *FirmwareEula) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *FirmwareEula) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


