# StorageBaseHostGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Short description about the host group. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the host group in storage array. | [optional] [readonly] 
**StorageUtilization** | Pointer to [**StorageBaseCapacity**](storage.BaseCapacity.md) |  | [optional] 

## Methods

### NewStorageBaseHostGroup

`func NewStorageBaseHostGroup() *StorageBaseHostGroup`

NewStorageBaseHostGroup instantiates a new StorageBaseHostGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageBaseHostGroupWithDefaults

`func NewStorageBaseHostGroupWithDefaults() *StorageBaseHostGroup`

NewStorageBaseHostGroupWithDefaults instantiates a new StorageBaseHostGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *StorageBaseHostGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StorageBaseHostGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StorageBaseHostGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StorageBaseHostGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *StorageBaseHostGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageBaseHostGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageBaseHostGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageBaseHostGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStorageUtilization

`func (o *StorageBaseHostGroup) GetStorageUtilization() StorageBaseCapacity`

GetStorageUtilization returns the StorageUtilization field if non-nil, zero value otherwise.

### GetStorageUtilizationOk

`func (o *StorageBaseHostGroup) GetStorageUtilizationOk() (*StorageBaseCapacity, bool)`

GetStorageUtilizationOk returns a tuple with the StorageUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageUtilization

`func (o *StorageBaseHostGroup) SetStorageUtilization(v StorageBaseCapacity)`

SetStorageUtilization sets StorageUtilization field to given value.

### HasStorageUtilization

`func (o *StorageBaseHostGroup) HasStorageUtilization() bool`

HasStorageUtilization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


