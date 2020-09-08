# StorageBaseVolumeAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Short description about the volume. | [optional] [readonly] 
**NaaId** | Pointer to **string** | NAA id of volume. It is a significant number to identify corresponding lun path in hypervisor. | [optional] [readonly] 
**Name** | Pointer to **string** | Named entity of the volume. | [optional] [readonly] 
**Size** | Pointer to **int64** | User provisioned volume size. It is the size exposed to host. | [optional] [readonly] 
**StorageUtilization** | Pointer to [**StorageBaseCapacity**](storage.BaseCapacity.md) |  | [optional] 

## Methods

### NewStorageBaseVolumeAllOf

`func NewStorageBaseVolumeAllOf() *StorageBaseVolumeAllOf`

NewStorageBaseVolumeAllOf instantiates a new StorageBaseVolumeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageBaseVolumeAllOfWithDefaults

`func NewStorageBaseVolumeAllOfWithDefaults() *StorageBaseVolumeAllOf`

NewStorageBaseVolumeAllOfWithDefaults instantiates a new StorageBaseVolumeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *StorageBaseVolumeAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StorageBaseVolumeAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StorageBaseVolumeAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StorageBaseVolumeAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetNaaId

`func (o *StorageBaseVolumeAllOf) GetNaaId() string`

GetNaaId returns the NaaId field if non-nil, zero value otherwise.

### GetNaaIdOk

`func (o *StorageBaseVolumeAllOf) GetNaaIdOk() (*string, bool)`

GetNaaIdOk returns a tuple with the NaaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaaId

`func (o *StorageBaseVolumeAllOf) SetNaaId(v string)`

SetNaaId sets NaaId field to given value.

### HasNaaId

`func (o *StorageBaseVolumeAllOf) HasNaaId() bool`

HasNaaId returns a boolean if a field has been set.

### GetName

`func (o *StorageBaseVolumeAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageBaseVolumeAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageBaseVolumeAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageBaseVolumeAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSize

`func (o *StorageBaseVolumeAllOf) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *StorageBaseVolumeAllOf) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *StorageBaseVolumeAllOf) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *StorageBaseVolumeAllOf) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetStorageUtilization

`func (o *StorageBaseVolumeAllOf) GetStorageUtilization() StorageBaseCapacity`

GetStorageUtilization returns the StorageUtilization field if non-nil, zero value otherwise.

### GetStorageUtilizationOk

`func (o *StorageBaseVolumeAllOf) GetStorageUtilizationOk() (*StorageBaseCapacity, bool)`

GetStorageUtilizationOk returns a tuple with the StorageUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageUtilization

`func (o *StorageBaseVolumeAllOf) SetStorageUtilization(v StorageBaseCapacity)`

SetStorageUtilization sets StorageUtilization field to given value.

### HasStorageUtilization

`func (o *StorageBaseVolumeAllOf) HasStorageUtilization() bool`

HasStorageUtilization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


