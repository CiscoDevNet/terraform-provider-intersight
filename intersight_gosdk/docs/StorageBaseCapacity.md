# StorageBaseCapacity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **int64** | Total consumable storage capacity represented in bytes. System may reserve some space for internal purposes which is excluded from total capacity. | [optional] [readonly] 
**CapacityUtilization** | Pointer to **float32** | Percentage of used capacity. | [optional] [readonly] 
**Free** | Pointer to **int64** | Unused space available for applications to consume, represented in bytes. | [optional] [readonly] 
**Total** | Pointer to **int64** | Total storage capacity, represented in bytes. It is set by the component manufacturer. | [optional] [readonly] 
**Used** | Pointer to **int64** | Used or consumed storage capacity, represented in bytes. | [optional] [readonly] 

## Methods

### NewStorageBaseCapacity

`func NewStorageBaseCapacity() *StorageBaseCapacity`

NewStorageBaseCapacity instantiates a new StorageBaseCapacity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageBaseCapacityWithDefaults

`func NewStorageBaseCapacityWithDefaults() *StorageBaseCapacity`

NewStorageBaseCapacityWithDefaults instantiates a new StorageBaseCapacity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *StorageBaseCapacity) GetAvailable() int64`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *StorageBaseCapacity) GetAvailableOk() (*int64, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *StorageBaseCapacity) SetAvailable(v int64)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *StorageBaseCapacity) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetCapacityUtilization

`func (o *StorageBaseCapacity) GetCapacityUtilization() float32`

GetCapacityUtilization returns the CapacityUtilization field if non-nil, zero value otherwise.

### GetCapacityUtilizationOk

`func (o *StorageBaseCapacity) GetCapacityUtilizationOk() (*float32, bool)`

GetCapacityUtilizationOk returns a tuple with the CapacityUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityUtilization

`func (o *StorageBaseCapacity) SetCapacityUtilization(v float32)`

SetCapacityUtilization sets CapacityUtilization field to given value.

### HasCapacityUtilization

`func (o *StorageBaseCapacity) HasCapacityUtilization() bool`

HasCapacityUtilization returns a boolean if a field has been set.

### GetFree

`func (o *StorageBaseCapacity) GetFree() int64`

GetFree returns the Free field if non-nil, zero value otherwise.

### GetFreeOk

`func (o *StorageBaseCapacity) GetFreeOk() (*int64, bool)`

GetFreeOk returns a tuple with the Free field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree

`func (o *StorageBaseCapacity) SetFree(v int64)`

SetFree sets Free field to given value.

### HasFree

`func (o *StorageBaseCapacity) HasFree() bool`

HasFree returns a boolean if a field has been set.

### GetTotal

`func (o *StorageBaseCapacity) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *StorageBaseCapacity) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *StorageBaseCapacity) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *StorageBaseCapacity) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUsed

`func (o *StorageBaseCapacity) GetUsed() int64`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *StorageBaseCapacity) GetUsedOk() (*int64, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *StorageBaseCapacity) SetUsed(v int64)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *StorageBaseCapacity) HasUsed() bool`

HasUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


