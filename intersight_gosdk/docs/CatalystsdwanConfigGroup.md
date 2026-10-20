# CatalystsdwanConfigGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "catalystsdwan.ConfigGroup"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "catalystsdwan.ConfigGroup"]
**ConfigGroupId** | Pointer to **string** | UUID for the Catalyst SDWAN config group. | [optional] 
**Devices** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** | The Catalyst SDWAN config group name. | [optional] 
**NumberOfDevices** | Pointer to **int64** | The Catalyst SDWAN config group number of devices. | [optional] 
**Solution** | Pointer to **string** | The Catalyst SDWAN config group solution. | [optional] 

## Methods

### NewCatalystsdwanConfigGroup

`func NewCatalystsdwanConfigGroup(classId string, objectType string, ) *CatalystsdwanConfigGroup`

NewCatalystsdwanConfigGroup instantiates a new CatalystsdwanConfigGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalystsdwanConfigGroupWithDefaults

`func NewCatalystsdwanConfigGroupWithDefaults() *CatalystsdwanConfigGroup`

NewCatalystsdwanConfigGroupWithDefaults instantiates a new CatalystsdwanConfigGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CatalystsdwanConfigGroup) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CatalystsdwanConfigGroup) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CatalystsdwanConfigGroup) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CatalystsdwanConfigGroup) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CatalystsdwanConfigGroup) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CatalystsdwanConfigGroup) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConfigGroupId

`func (o *CatalystsdwanConfigGroup) GetConfigGroupId() string`

GetConfigGroupId returns the ConfigGroupId field if non-nil, zero value otherwise.

### GetConfigGroupIdOk

`func (o *CatalystsdwanConfigGroup) GetConfigGroupIdOk() (*string, bool)`

GetConfigGroupIdOk returns a tuple with the ConfigGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigGroupId

`func (o *CatalystsdwanConfigGroup) SetConfigGroupId(v string)`

SetConfigGroupId sets ConfigGroupId field to given value.

### HasConfigGroupId

`func (o *CatalystsdwanConfigGroup) HasConfigGroupId() bool`

HasConfigGroupId returns a boolean if a field has been set.

### GetDevices

`func (o *CatalystsdwanConfigGroup) GetDevices() []string`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *CatalystsdwanConfigGroup) GetDevicesOk() (*[]string, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *CatalystsdwanConfigGroup) SetDevices(v []string)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *CatalystsdwanConfigGroup) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### SetDevicesNil

`func (o *CatalystsdwanConfigGroup) SetDevicesNil(b bool)`

 SetDevicesNil sets the value for Devices to be an explicit nil

### UnsetDevices
`func (o *CatalystsdwanConfigGroup) UnsetDevices()`

UnsetDevices ensures that no value is present for Devices, not even an explicit nil
### GetName

`func (o *CatalystsdwanConfigGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CatalystsdwanConfigGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CatalystsdwanConfigGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CatalystsdwanConfigGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumberOfDevices

`func (o *CatalystsdwanConfigGroup) GetNumberOfDevices() int64`

GetNumberOfDevices returns the NumberOfDevices field if non-nil, zero value otherwise.

### GetNumberOfDevicesOk

`func (o *CatalystsdwanConfigGroup) GetNumberOfDevicesOk() (*int64, bool)`

GetNumberOfDevicesOk returns a tuple with the NumberOfDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfDevices

`func (o *CatalystsdwanConfigGroup) SetNumberOfDevices(v int64)`

SetNumberOfDevices sets NumberOfDevices field to given value.

### HasNumberOfDevices

`func (o *CatalystsdwanConfigGroup) HasNumberOfDevices() bool`

HasNumberOfDevices returns a boolean if a field has been set.

### GetSolution

`func (o *CatalystsdwanConfigGroup) GetSolution() string`

GetSolution returns the Solution field if non-nil, zero value otherwise.

### GetSolutionOk

`func (o *CatalystsdwanConfigGroup) GetSolutionOk() (*string, bool)`

GetSolutionOk returns a tuple with the Solution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolution

`func (o *CatalystsdwanConfigGroup) SetSolution(v string)`

SetSolution sets Solution field to given value.

### HasSolution

`func (o *CatalystsdwanConfigGroup) HasSolution() bool`

HasSolution returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


