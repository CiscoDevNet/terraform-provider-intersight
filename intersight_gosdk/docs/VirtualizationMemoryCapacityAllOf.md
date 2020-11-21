# VirtualizationMemoryCapacityAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.MemoryCapacity"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.MemoryCapacity"]
**Capacity** | Pointer to **int64** | The total memory capacity of the entity in bytes. | [optional] 
**Free** | Pointer to **int64** | Free memory (bytes) that is unused and available for allocation, as a point-in-time snapshot. The available memory capacity is reported for an entity (such as Host or Cluster) when inventory data is collected for that entity. As part of the inventory data, a snapshot of the free and used memory capacity is also reported. | [optional] 
**Used** | Pointer to **int64** | Memory (bytes) that has been already used up, as a point-in-time snapshot. | [optional] 

## Methods

### NewVirtualizationMemoryCapacityAllOf

`func NewVirtualizationMemoryCapacityAllOf(classId string, objectType string, ) *VirtualizationMemoryCapacityAllOf`

NewVirtualizationMemoryCapacityAllOf instantiates a new VirtualizationMemoryCapacityAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationMemoryCapacityAllOfWithDefaults

`func NewVirtualizationMemoryCapacityAllOfWithDefaults() *VirtualizationMemoryCapacityAllOf`

NewVirtualizationMemoryCapacityAllOfWithDefaults instantiates a new VirtualizationMemoryCapacityAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationMemoryCapacityAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationMemoryCapacityAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationMemoryCapacityAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationMemoryCapacityAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationMemoryCapacityAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationMemoryCapacityAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCapacity

`func (o *VirtualizationMemoryCapacityAllOf) GetCapacity() int64`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *VirtualizationMemoryCapacityAllOf) GetCapacityOk() (*int64, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *VirtualizationMemoryCapacityAllOf) SetCapacity(v int64)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *VirtualizationMemoryCapacityAllOf) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetFree

`func (o *VirtualizationMemoryCapacityAllOf) GetFree() int64`

GetFree returns the Free field if non-nil, zero value otherwise.

### GetFreeOk

`func (o *VirtualizationMemoryCapacityAllOf) GetFreeOk() (*int64, bool)`

GetFreeOk returns a tuple with the Free field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree

`func (o *VirtualizationMemoryCapacityAllOf) SetFree(v int64)`

SetFree sets Free field to given value.

### HasFree

`func (o *VirtualizationMemoryCapacityAllOf) HasFree() bool`

HasFree returns a boolean if a field has been set.

### GetUsed

`func (o *VirtualizationMemoryCapacityAllOf) GetUsed() int64`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *VirtualizationMemoryCapacityAllOf) GetUsedOk() (*int64, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *VirtualizationMemoryCapacityAllOf) SetUsed(v int64)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *VirtualizationMemoryCapacityAllOf) HasUsed() bool`

HasUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


