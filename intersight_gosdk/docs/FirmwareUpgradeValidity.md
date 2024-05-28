# FirmwareUpgradeValidity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "firmware.UpgradeValidity"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "firmware.UpgradeValidity"]
**ErrorMessage** | Pointer to **string** | The error string returned while checking for a target device&#39;s validity for firmware upgrade. | [optional] [readonly] 
**IsValid** | Pointer to **bool** | This flag denotes whether the target device is a valid target for firmware upgrade. | [optional] [readonly] 
**Server** | Pointer to [**NullableComputePhysicalRelationship**](ComputePhysicalRelationship.md) |  | [optional] 

## Methods

### NewFirmwareUpgradeValidity

`func NewFirmwareUpgradeValidity(classId string, objectType string, ) *FirmwareUpgradeValidity`

NewFirmwareUpgradeValidity instantiates a new FirmwareUpgradeValidity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareUpgradeValidityWithDefaults

`func NewFirmwareUpgradeValidityWithDefaults() *FirmwareUpgradeValidity`

NewFirmwareUpgradeValidityWithDefaults instantiates a new FirmwareUpgradeValidity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FirmwareUpgradeValidity) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FirmwareUpgradeValidity) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FirmwareUpgradeValidity) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FirmwareUpgradeValidity) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FirmwareUpgradeValidity) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FirmwareUpgradeValidity) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetErrorMessage

`func (o *FirmwareUpgradeValidity) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *FirmwareUpgradeValidity) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *FirmwareUpgradeValidity) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *FirmwareUpgradeValidity) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetIsValid

`func (o *FirmwareUpgradeValidity) GetIsValid() bool`

GetIsValid returns the IsValid field if non-nil, zero value otherwise.

### GetIsValidOk

`func (o *FirmwareUpgradeValidity) GetIsValidOk() (*bool, bool)`

GetIsValidOk returns a tuple with the IsValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsValid

`func (o *FirmwareUpgradeValidity) SetIsValid(v bool)`

SetIsValid sets IsValid field to given value.

### HasIsValid

`func (o *FirmwareUpgradeValidity) HasIsValid() bool`

HasIsValid returns a boolean if a field has been set.

### GetServer

`func (o *FirmwareUpgradeValidity) GetServer() ComputePhysicalRelationship`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *FirmwareUpgradeValidity) GetServerOk() (*ComputePhysicalRelationship, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *FirmwareUpgradeValidity) SetServer(v ComputePhysicalRelationship)`

SetServer sets Server field to given value.

### HasServer

`func (o *FirmwareUpgradeValidity) HasServer() bool`

HasServer returns a boolean if a field has been set.

### SetServerNil

`func (o *FirmwareUpgradeValidity) SetServerNil(b bool)`

 SetServerNil sets the value for Server to be an explicit nil

### UnsetServer
`func (o *FirmwareUpgradeValidity) UnsetServer()`

UnsetServer ensures that no value is present for Server, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


