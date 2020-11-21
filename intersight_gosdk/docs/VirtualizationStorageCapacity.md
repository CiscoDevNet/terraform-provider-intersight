# VirtualizationStorageCapacity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.StorageCapacity"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.StorageCapacity"]
**Capacity** | Pointer to **int64** | The total capacity of the entity (bytes). | [optional] 
**Free** | Pointer to **int64** | Free storage space remaining in the entity (bytes) as a point-in-time snapshot. The available space is reported for an entity (such as Host or Cluster) when inventory data is collected for that entity. As part of the inventory data, a snapshot of the free and used storage capacity is also reported. | [optional] 
**Used** | Pointer to **int64** | Spaced already used by this entity (bytes), as a point-in-time snapshot. | [optional] 

## Methods

### NewVirtualizationStorageCapacity

`func NewVirtualizationStorageCapacity(classId string, objectType string, ) *VirtualizationStorageCapacity`

NewVirtualizationStorageCapacity instantiates a new VirtualizationStorageCapacity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationStorageCapacityWithDefaults

`func NewVirtualizationStorageCapacityWithDefaults() *VirtualizationStorageCapacity`

NewVirtualizationStorageCapacityWithDefaults instantiates a new VirtualizationStorageCapacity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationStorageCapacity) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationStorageCapacity) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationStorageCapacity) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationStorageCapacity) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationStorageCapacity) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationStorageCapacity) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCapacity

`func (o *VirtualizationStorageCapacity) GetCapacity() int64`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *VirtualizationStorageCapacity) GetCapacityOk() (*int64, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *VirtualizationStorageCapacity) SetCapacity(v int64)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *VirtualizationStorageCapacity) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetFree

`func (o *VirtualizationStorageCapacity) GetFree() int64`

GetFree returns the Free field if non-nil, zero value otherwise.

### GetFreeOk

`func (o *VirtualizationStorageCapacity) GetFreeOk() (*int64, bool)`

GetFreeOk returns a tuple with the Free field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree

`func (o *VirtualizationStorageCapacity) SetFree(v int64)`

SetFree sets Free field to given value.

### HasFree

`func (o *VirtualizationStorageCapacity) HasFree() bool`

HasFree returns a boolean if a field has been set.

### GetUsed

`func (o *VirtualizationStorageCapacity) GetUsed() int64`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *VirtualizationStorageCapacity) GetUsedOk() (*int64, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *VirtualizationStorageCapacity) SetUsed(v int64)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *VirtualizationStorageCapacity) HasUsed() bool`

HasUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


