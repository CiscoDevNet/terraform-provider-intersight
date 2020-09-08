# VirtualizationMemoryCapacity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capacity** | Pointer to **int64** | The total memory capacity of the entity in bytes. | [optional] 
**Free** | Pointer to **int64** | Free memory (bytes) that is unused and available for allocation, as a point-in-time snapshot. The available memory capacity is reported for an entity (such as Host or Cluster) when inventory data is collected for that entity. As part of the inventory data, a snapshot of the free and used memory capacity is also reported. | [optional] 
**Used** | Pointer to **int64** | Memory (bytes) that has been already used up, as a point-in-time snapshot. | [optional] 

## Methods

### NewVirtualizationMemoryCapacity

`func NewVirtualizationMemoryCapacity() *VirtualizationMemoryCapacity`

NewVirtualizationMemoryCapacity instantiates a new VirtualizationMemoryCapacity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationMemoryCapacityWithDefaults

`func NewVirtualizationMemoryCapacityWithDefaults() *VirtualizationMemoryCapacity`

NewVirtualizationMemoryCapacityWithDefaults instantiates a new VirtualizationMemoryCapacity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapacity

`func (o *VirtualizationMemoryCapacity) GetCapacity() int64`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *VirtualizationMemoryCapacity) GetCapacityOk() (*int64, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *VirtualizationMemoryCapacity) SetCapacity(v int64)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *VirtualizationMemoryCapacity) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetFree

`func (o *VirtualizationMemoryCapacity) GetFree() int64`

GetFree returns the Free field if non-nil, zero value otherwise.

### GetFreeOk

`func (o *VirtualizationMemoryCapacity) GetFreeOk() (*int64, bool)`

GetFreeOk returns a tuple with the Free field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree

`func (o *VirtualizationMemoryCapacity) SetFree(v int64)`

SetFree sets Free field to given value.

### HasFree

`func (o *VirtualizationMemoryCapacity) HasFree() bool`

HasFree returns a boolean if a field has been set.

### GetUsed

`func (o *VirtualizationMemoryCapacity) GetUsed() int64`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *VirtualizationMemoryCapacity) GetUsedOk() (*int64, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *VirtualizationMemoryCapacity) SetUsed(v int64)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *VirtualizationMemoryCapacity) HasUsed() bool`

HasUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


