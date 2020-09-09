# StorageBaseHostGroupAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Short description about the host group. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the host group in storage array. | [optional] [readonly] 
**StorageUtilization** | Pointer to [**StorageBaseCapacity**](storage.BaseCapacity.md) |  | [optional] 

## Methods

### NewStorageBaseHostGroupAllOf

`func NewStorageBaseHostGroupAllOf() *StorageBaseHostGroupAllOf`

NewStorageBaseHostGroupAllOf instantiates a new StorageBaseHostGroupAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageBaseHostGroupAllOfWithDefaults

`func NewStorageBaseHostGroupAllOfWithDefaults() *StorageBaseHostGroupAllOf`

NewStorageBaseHostGroupAllOfWithDefaults instantiates a new StorageBaseHostGroupAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *StorageBaseHostGroupAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StorageBaseHostGroupAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StorageBaseHostGroupAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StorageBaseHostGroupAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *StorageBaseHostGroupAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageBaseHostGroupAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageBaseHostGroupAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageBaseHostGroupAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStorageUtilization

`func (o *StorageBaseHostGroupAllOf) GetStorageUtilization() StorageBaseCapacity`

GetStorageUtilization returns the StorageUtilization field if non-nil, zero value otherwise.

### GetStorageUtilizationOk

`func (o *StorageBaseHostGroupAllOf) GetStorageUtilizationOk() (*StorageBaseCapacity, bool)`

GetStorageUtilizationOk returns a tuple with the StorageUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageUtilization

`func (o *StorageBaseHostGroupAllOf) SetStorageUtilization(v StorageBaseCapacity)`

SetStorageUtilization sets StorageUtilization field to given value.

### HasStorageUtilization

`func (o *StorageBaseHostGroupAllOf) HasStorageUtilization() bool`

HasStorageUtilization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


