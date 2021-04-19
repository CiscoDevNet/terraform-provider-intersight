# VirtualizationMemoryAllocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.MemoryAllocation"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.MemoryAllocation"]
**Free** | Pointer to **int64** | Free memory on the entity in bytes. | [optional] 
**Reserved** | Pointer to **int64** | Reserved memory on the entity in bytes. These reserved memory can be used for system purposes. | [optional] 
**Total** | Pointer to **int64** | The total memory capacity of the entity in bytes. | [optional] 
**Used** | Pointer to **int64** | Used or allocated memory on the entity represented in bytes. | [optional] 

## Methods

### NewVirtualizationMemoryAllocation

`func NewVirtualizationMemoryAllocation(classId string, objectType string, ) *VirtualizationMemoryAllocation`

NewVirtualizationMemoryAllocation instantiates a new VirtualizationMemoryAllocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationMemoryAllocationWithDefaults

`func NewVirtualizationMemoryAllocationWithDefaults() *VirtualizationMemoryAllocation`

NewVirtualizationMemoryAllocationWithDefaults instantiates a new VirtualizationMemoryAllocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationMemoryAllocation) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationMemoryAllocation) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationMemoryAllocation) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationMemoryAllocation) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationMemoryAllocation) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationMemoryAllocation) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFree

`func (o *VirtualizationMemoryAllocation) GetFree() int64`

GetFree returns the Free field if non-nil, zero value otherwise.

### GetFreeOk

`func (o *VirtualizationMemoryAllocation) GetFreeOk() (*int64, bool)`

GetFreeOk returns a tuple with the Free field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree

`func (o *VirtualizationMemoryAllocation) SetFree(v int64)`

SetFree sets Free field to given value.

### HasFree

`func (o *VirtualizationMemoryAllocation) HasFree() bool`

HasFree returns a boolean if a field has been set.

### GetReserved

`func (o *VirtualizationMemoryAllocation) GetReserved() int64`

GetReserved returns the Reserved field if non-nil, zero value otherwise.

### GetReservedOk

`func (o *VirtualizationMemoryAllocation) GetReservedOk() (*int64, bool)`

GetReservedOk returns a tuple with the Reserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserved

`func (o *VirtualizationMemoryAllocation) SetReserved(v int64)`

SetReserved sets Reserved field to given value.

### HasReserved

`func (o *VirtualizationMemoryAllocation) HasReserved() bool`

HasReserved returns a boolean if a field has been set.

### GetTotal

`func (o *VirtualizationMemoryAllocation) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *VirtualizationMemoryAllocation) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *VirtualizationMemoryAllocation) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *VirtualizationMemoryAllocation) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUsed

`func (o *VirtualizationMemoryAllocation) GetUsed() int64`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *VirtualizationMemoryAllocation) GetUsedOk() (*int64, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *VirtualizationMemoryAllocation) SetUsed(v int64)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *VirtualizationMemoryAllocation) HasUsed() bool`

HasUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


